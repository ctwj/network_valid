package services

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"

	"verification/models"
)

// PlanService 套餐服务
type PlanService struct{}

// GrantFreeTier 发放免费套餐
func (s *PlanService) GrantFreeTier(memberId int, projectId int) error {
	o := orm.NewOrm()

	// 查找项目的免费套餐
	var card models.Cards
	err := o.QueryTable("cards").
		Filter("project_id", projectId).
		Filter("is_free_tier", 1).
		One(&card)
	if err != nil {
		logs.Error("查找免费套餐失败:", err)
		return err
	}

	// 创建用户套餐关联（永久有效）
	memberPlan := &models.MemberPlan{
		MemberId:    memberId,
		PlanId:      card.ID,
		Status:      "active",
		ExpireTime:  nil, // NULL 表示永久
		ActivatedAt: time.Now(),
	}
	_, err = o.Insert(memberPlan)
	if err != nil {
		logs.Error("创建用户套餐关联失败:", err)
		return err
	}

	// 初始化配额（如果有）
	if card.QuotaRules != "" {
		quotaRules, err := models.ParseQuotaRules(card.QuotaRules)
		if err == nil && quotaRules != nil {
			quotaService := &QuotaService{}
			err = quotaService.InitQuotas(memberId, quotaRules, nil, card.ID)
			if err != nil {
				logs.Error("初始化配额失败:", err)
			}
		}
	}

	return nil
}

// RedeemPlan 兑换套餐（含升级/排队逻辑）
func (s *PlanService) RedeemPlan(memberId int, planId int) error {
	o := orm.NewOrm()

	// 获取新套餐
	newPlan := &models.Cards{ID: planId}
	err := o.Read(newPlan)
	if err != nil {
		return err
	}

	// 获取当前套餐
	currentPlan, hasCurrent := s.GetCurrentPlan(memberId)

	if !hasCurrent {
		// 没有当前套餐，直接激活
		return s.ActivatePlan(memberId, newPlan)
	}

	// 获取当前套餐信息
	currentCard := &models.Cards{ID: currentPlan.PlanId}
	o.Read(currentCard)

	// 比较优先级
	if newPlan.Priority > currentCard.Priority {
		// 高优先级，立即升级
		return s.UpgradePlan(memberId, newPlan, currentPlan)
	} else if newPlan.Priority == currentCard.Priority {
		// 同优先级，时长叠加
		return s.ExtendPlan(memberId, newPlan, currentPlan)
	} else {
		// 低优先级，进入排队
		return s.QueuePlan(memberId, newPlan, currentPlan)
	}
}

// ActivatePlan 激活套餐
func (s *PlanService) ActivatePlan(memberId int, plan *models.Cards) error {
	o := orm.NewOrm()

	// 计算过期时间
	var expireTime *time.Time
	if plan.Days > 0 {
		t := time.Now().AddDate(0, 0, int(plan.Days))
		expireTime = &t
	}

	// 创建用户套餐关联
	memberPlan := &models.MemberPlan{
		MemberId:    memberId,
		PlanId:      plan.ID,
		Status:      "active",
		ExpireTime:  expireTime,
		ActivatedAt: time.Now(),
	}
	_, err := o.Insert(memberPlan)
	if err != nil {
		return err
	}

	// 初始化配额
	if plan.QuotaRules != "" {
		quotaRules, err := models.ParseQuotaRules(plan.QuotaRules)
		if err == nil && quotaRules != nil {
			quotaService := &QuotaService{}
			err = quotaService.InitQuotas(memberId, quotaRules, expireTime, plan.ID)
			if err != nil {
				logs.Error("初始化配额失败:", err)
			}
		}
	}

	return nil
}

// UpgradePlan 升级套餐
func (s *PlanService) UpgradePlan(memberId int, newPlan *models.Cards, currentPlan *models.MemberPlan) error {
	// 使当前套餐失效
	currentPlan.Expire()

	// 删除当前套餐的配额
	quota := &models.UserQuota{}
	quota.DeleteByMemberAndPlan(memberId, currentPlan.PlanId)

	// 激活新套餐
	return s.ActivatePlan(memberId, newPlan)
}

// ExtendPlan 延长套餐（同优先级叠加时长）
func (s *PlanService) ExtendPlan(memberId int, newPlan *models.Cards, currentPlan *models.MemberPlan) error {
	// 计算新的过期时间
	if currentPlan.ExpireTime == nil {
		// 当前套餐永久，不需要延长
		return nil
	}
	newExpireTime := currentPlan.ExpireTime.AddDate(0, 0, int(newPlan.Days))
	currentPlan.ExpireTime = &newExpireTime

	return currentPlan.Update()
}

// DowngradeToFreeTier 降级为免费套餐
func (s *PlanService) DowngradeToFreeTier(memberId int, projectId int) error {
	// 获取当前套餐
	currentPlan, hasCurrent := s.GetCurrentPlan(memberId)
	if hasCurrent {
		// 使当前套餐失效
		currentPlan.Expire()

		// 删除当前套餐的配额
		quota := &models.UserQuota{}
		quota.DeleteByMemberAndPlan(memberId, currentPlan.PlanId)
	}

	// 发放免费套餐
	return s.GrantFreeTier(memberId, projectId)
}

// QueuePlan 套餐排队
func (s *PlanService) QueuePlan(memberId int, newPlan *models.Cards, currentPlan *models.MemberPlan) error {
	o := orm.NewOrm()

	// 计算排队激活时间（当前套餐过期后）
	var activateAfter *time.Time
	if currentPlan.ExpireTime != nil {
		activateAfter = currentPlan.ExpireTime
	}

	// 创建排队状态的套餐关联
	memberPlan := &models.MemberPlan{
		MemberId:   memberId,
		PlanId:     newPlan.ID,
		Status:     "queued",
		ExpireTime: activateAfter,
	}
	_, err := o.Insert(memberPlan)
	return err
}

// GetCurrentPlan 获取当前套餐
func (s *PlanService) GetCurrentPlan(memberId int) (*models.MemberPlan, bool) {
	plan := &models.MemberPlan{}
	hasPlan, result := plan.GetActiveByMember(memberId)
	return result, hasPlan
}

// GetQueuedPlans 获取排队套餐
func (s *PlanService) GetQueuedPlans(memberId int) ([]*models.MemberPlan, error) {
	plan := &models.MemberPlan{}
	return plan.GetQueuedByMember(memberId)
}

// CheckExpiredPlans 检查过期套餐（定时任务）
func (s *PlanService) CheckExpiredPlans() (int64, error) {
	o := orm.NewOrm()

	// 查找已过期但状态仍为 active 的套餐
	var expiredPlans []*models.MemberPlan
	now := time.Now()
	_, err := o.QueryTable("member_plans").
		Filter("status", "active").
		Filter("expire_time__isnull", false).
		Filter("expire_time__lt", now).
		All(&expiredPlans)
	if err != nil {
		return 0, err
	}

	count := int64(0)
	for _, plan := range expiredPlans {
		// 获取用户的项目ID
		member := &models.Member{ID: plan.MemberId}
		err := o.Read(member)
		if err != nil {
			continue
		}

		// 检查是否有排队套餐
		queuedPlans, _ := s.GetQueuedPlans(plan.MemberId)
		if len(queuedPlans) > 0 {
			// 激活排队套餐
			nextPlan := queuedPlans[0]
			nextPlan.Activate()
			count++
		} else {
			// 降级为免费套餐
			err = s.DowngradeToFreeTier(plan.MemberId, member.ProjectId)
			if err == nil {
				count++
			}
		}
	}

	logs.Info("处理过期套餐数量:", count)
	return count, nil
}
