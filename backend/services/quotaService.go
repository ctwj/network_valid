package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"

	"verification/models"
)

// QuotaService 配额服务
type QuotaService struct{}

// ParseQuotaRules 解析配额规则 JSON
func (s *QuotaService) ParseQuotaRules(jsonStr string) (*models.QuotaRules, error) {
	return models.ParseQuotaRules(jsonStr)
}

// InitQuotas 初始化用户配额
func (s *QuotaService) InitQuotas(memberId int, quotaRules *models.QuotaRules, expireTime *time.Time, planId int) error {
	if quotaRules == nil || len(quotaRules.Quotas) == 0 {
		return nil // 无配额规则，不创建配额记录
	}

	o := orm.NewOrm()

	for _, rule := range quotaRules.Quotas {
		resetTime := models.CalculateNextResetTime(rule.Period, rule.ResetDay)

		quota := &models.UserQuota{
			MemberId:   memberId,
			QuotaKey:   rule.Key,
			LimitValue: rule.Limit,
			UsedValue:  0,
			Period:     rule.Period,
			ResetDay:   rule.ResetDay,
			ResetTime:  resetTime,
			ExpireTime: expireTime,
			PlanId:     planId,
		}

		_, err := o.Insert(quota)
		if err != nil {
			logs.Error("初始化配额失败:", err, "memberId:", memberId, "quotaKey:", rule.Key)
			return err
		}
	}

	return nil
}

// QuotaInfo 配额信息（返回给客户端）
type QuotaInfo struct {
	Key       string    `json:"key"`
	Name      string    `json:"name"`
	Limit     int64     `json:"limit"`
	Used      int64     `json:"used"`
	Remaining int64     `json:"remaining"`
	Period    string    `json:"period"`
	ResetTime time.Time `json:"reset_time"`
	Unit      string    `json:"unit"`
	Unlimited bool      `json:"unlimited"` // 是否无限制
}

// CheckQuotas 检查用户配额状态
func (s *QuotaService) CheckQuotas(memberId int, keys []string) ([]QuotaInfo, error) {
	o := orm.NewOrm()

	// 获取用户所有配额
	var quotas []*models.UserQuota
	if len(keys) > 0 {
		_, err := o.QueryTable("user_quotas").
			Filter("member_id", memberId).
			Filter("quota_key__in", keys).
			All(&quotas)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := o.QueryTable("user_quotas").
			Filter("member_id", memberId).
			All(&quotas)
		if err != nil {
			return nil, err
		}
	}

	// 如果没有配额记录，返回无限制
	if len(quotas) == 0 {
		return []QuotaInfo{{Unlimited: true}}, nil
	}

	// 转换为 QuotaInfo
	result := make([]QuotaInfo, 0, len(quotas))
	for _, quota := range quotas {
		// 检查是否需要重置（懒重置）
		s.checkAndReset(quota)

		remaining := quota.LimitValue - quota.UsedValue
		if remaining < 0 {
			remaining = 0
		}

		info := QuotaInfo{
			Key:       quota.QuotaKey,
			Limit:     quota.LimitValue,
			Used:      quota.UsedValue,
			Remaining: remaining,
			Period:    quota.Period,
			ResetTime: quota.ResetTime,
		}

		result = append(result, info)
	}

	return result, nil
}

// checkAndReset 检查并重置配额（懒重置）
func (s *QuotaService) checkAndReset(quota *models.UserQuota) {
	if models.IsExpired(quota.ResetTime) {
		quota.UsedValue = 0
		quota.ResetTime = models.CalculateNextResetTime(quota.Period, quota.ResetDay)
		o := orm.NewOrm()
		_, err := o.Update(quota)
		if err != nil {
			logs.Error("重置配额失败:", err)
		}
	}
}

// DeductResult 扣减结果
type DeductResult struct {
	Key       string `json:"key"`
	Success   bool   `json:"success"`
	Used      int64  `json:"used"`
	Remaining int64  `json:"remaining"`
	Message   string `json:"message,omitempty"`
}

// DeductQuota 扣减单个配额（原子操作）
func (s *QuotaService) DeductQuota(memberId int, key string, amount int64) (*DeductResult, error) {
	o := orm.NewOrm()

	// 查询配额
	quota := &models.UserQuota{MemberId: memberId, QuotaKey: key}
	err := o.Read(quota, "MemberId", "QuotaKey")
	if err != nil {
		// 配额不存在，表示无限制
		return &DeductResult{
			Key:     key,
			Success: true,
			Message: "无配额限制",
		}, nil
	}

	// 检查是否需要重置
	s.checkAndReset(quota)

	// 检查是否过期
	if quota.ExpireTime != nil && time.Now().After(*quota.ExpireTime) {
		return &DeductResult{
			Key:     key,
			Success: false,
			Message: "套餐已过期",
		}, nil
	}

	// 检查配额是否足够
	if quota.UsedValue + amount > quota.LimitValue {
		return &DeductResult{
			Key:     key,
			Success: false,
			Message: "配额不足",
		}, nil
	}

	// 原子更新（使用 SQL 直接更新，防止并发超扣）
	sql := `UPDATE user_quotas SET used_value = used_value + ? WHERE member_id = ? AND quota_key = ? AND used_value + ? <= limit_value AND (expire_time IS NULL OR expire_time > NOW())`
	res, err := o.Raw(sql, amount, memberId, key, amount).Exec()
	if err != nil {
		return nil, err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return &DeductResult{
			Key:     key,
			Success: false,
			Message: "配额不足或并发冲突",
		}, nil
	}

	// 记录使用日志
	s.LogQuotaUsage(memberId, key, amount)

	// 查询更新后的值
	o.Read(quota, "MemberId", "QuotaKey")
	remaining := quota.LimitValue - quota.UsedValue

	return &DeductResult{
		Key:       key,
		Success:   true,
		Used:      quota.UsedValue,
		Remaining: remaining,
	}, nil
}

// DeductQuotaBatch 批量扣减配额（事务保证）
func (s *QuotaService) DeductQuotaBatch(memberId int, items []models.QuotaDeductItem) ([]DeductResult, error) {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return nil, err
	}

	results := make([]DeductResult, 0, len(items))
	allSuccess := true

	for _, item := range items {
		result, err := s.DeductQuota(memberId, item.Key, item.Amount)
		if err != nil {
			o.Rollback()
			return nil, err
		}
		results = append(results, *result)
		if !result.Success {
			allSuccess = false
		}
	}

	if !allSuccess {
		o.Rollback()
		return results, errors.New("部分配额扣减失败")
	}

	o.Commit()
	return results, nil
}

// LogQuotaUsage 记录配额使用日志
func (s *QuotaService) LogQuotaUsage(memberId int, key string, amount int64) {
	log := &models.QuotaUsageLog{
		MemberId: memberId,
		QuotaKey: key,
		Amount:   amount,
	}
	o := orm.NewOrm()
	_, err := o.Insert(log)
	if err != nil {
		logs.Error("记录配额使用日志失败:", err)
	}
}

// ResetExpiredQuotas 批量重置过期配额（定时任务）
func (s *QuotaService) ResetExpiredQuotas() (int64, error) {
	o := orm.NewOrm()
	now := time.Now()

	// 查询需要重置的配额
	var quotas []*models.UserQuota
	_, err := o.QueryTable("user_quotas").
		Filter("reset_time__lt", now).
		All(&quotas)
	if err != nil {
		return 0, err
	}

	count := int64(0)
	for _, quota := range quotas {
		quota.UsedValue = 0
		quota.ResetTime = models.CalculateNextResetTime(quota.Period, quota.ResetDay)
		_, err := o.Update(quota)
		if err == nil {
			count++
		}
	}

	logs.Info("重置配额数量:", count)
	return count, nil
}

// CleanExpiredQuotas 清理过期配额（定时任务）
func (s *QuotaService) CleanExpiredQuotas() (int64, error) {
	o := orm.NewOrm()
	now := time.Now()

	// 删除已过期的配额（expire_time 不为 NULL 且已过期）
	count, err := o.QueryTable("user_quotas").
		Filter("expire_time__isnull", false).
		Filter("expire_time__lt", now).
		Delete()
	if err != nil {
		return 0, err
	}

	logs.Info("清理过期配额数量:", count)
	return count, nil
}

// GetDefaultPresets 获取默认预设套餐方案
func (s *QuotaService) GetDefaultPresets() string {
	presets := models.PlanPresets{
		Presets: []models.PlanPreset{
			{
				Name:       "免费套餐",
				Days:       0,
				IsFreeTier: true,
				Priority:   0,
				QuotaRules: nil,
			},
			{
				Name:       "标准套餐",
				Days:       30,
				IsFreeTier: false,
				Priority:   10,
				QuotaRules: nil,
			},
		},
	}
	data, _ := json.Marshal(presets)
	return string(data)
}