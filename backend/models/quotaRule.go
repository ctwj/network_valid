package models

import (
	"encoding/json"
	"time"
)

// QuotaRule 配额规则（JSON 解析用）
type QuotaRule struct {
	Key      string `json:"key"`       // 配额标识
	Name     string `json:"name"`      // 显示名称
	Limit    int64  `json:"limit"`     // 限额数量
	Period   string `json:"period"`    // 周期：daily/weekly/monthly
	ResetDay int    `json:"reset_day"` // 重置日：daily=每日，weekly=周几(1-7)，monthly=几号(1-31)
	Unit     string `json:"unit"`      // 单位：count/bytes/custom
}

// QuotaRules 配额规则集合
type QuotaRules struct {
	Quotas []QuotaRule `json:"quotas"`
}

// QuotaDeductItem 配额扣减项
type QuotaDeductItem struct {
	Key    string `json:"key"`
	Amount int64  `json:"amount"`
}

// PlanPreset 预设套餐方案
type PlanPreset struct {
	Name       string      `json:"name"`        // 套餐名称
	Days       int         `json:"days"`        // 天数，0 表示永久
	IsFreeTier bool        `json:"is_free_tier"` // 是否为免费套餐
	Priority   int         `json:"priority"`    // 优先级
	QuotaRules *QuotaRules `json:"quota_rules"` // 配额规则（可选）
}

// PlanPresets 预设套餐方案集合
type PlanPresets struct {
	Presets []PlanPreset `json:"presets"`
}

// ParseQuotaRules 解析配额规则 JSON
func ParseQuotaRules(jsonStr string) (*QuotaRules, error) {
	if jsonStr == "" {
		return nil, nil
	}
	var rules QuotaRules
	err := json.Unmarshal([]byte(jsonStr), &rules)
	if err != nil {
		return nil, err
	}
	return &rules, nil
}

// ToJSON 将配额规则转换为 JSON
func (q *QuotaRules) ToJSON() (string, error) {
	if q == nil {
		return "", nil
	}
	data, err := json.Marshal(q)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// CalculateNextResetTime 计算下次重置时间
func CalculateNextResetTime(period string, resetDay int) time.Time {
	now := time.Now()
	loc := now.Location()

	switch period {
	case "daily":
		// 次日 00:00
		next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, loc)
		return next

	case "weekly":
		// 下个 resetDay（1=周一，7=周日）
		currentWeekday := int(now.Weekday())
		if currentWeekday == 0 {
			currentWeekday = 7 // 将周日从 0 转为 7
		}
		daysUntilReset := resetDay - currentWeekday
		if daysUntilReset <= 0 {
			daysUntilReset += 7
		}
		next := time.Date(now.Year(), now.Month(), now.Day()+daysUntilReset, 0, 0, 0, 0, loc)
		return next

	case "monthly":
		// 下月 resetDay 号
		year := now.Year()
		month := now.Month()

		// 如果当前日期已过本月的 resetDay，则计算下个月
		if now.Day() >= resetDay {
			month++
			if month > 12 {
				month = 1
				year++
			}
		}

		// 处理月份天数不足的情况（如 2 月 30 号）
		day := resetDay
		daysInMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, loc).Day()
		if day > daysInMonth {
			day = daysInMonth
		}

		next := time.Date(year, month, day, 0, 0, 0, 0, loc)
		return next

	default:
		// 默认次日
		return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, loc)
	}
}

// IsExpired 检查配额是否需要重置
func IsExpired(resetTime time.Time) bool {
	return time.Now().After(resetTime) || time.Now().Equal(resetTime)
}

// IsPlanExpired 检查套餐是否过期
func IsPlanExpired(expireTime *time.Time) bool {
	if expireTime == nil {
		return false // NULL 表示永不过期
	}
	return time.Now().After(*expireTime)
}
