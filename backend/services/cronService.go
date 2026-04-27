package services

import (
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// CronService 定时任务服务
type CronService struct{}

// RunQuotaReset 配额重置任务（每小时执行）
func (s *CronService) RunQuotaReset() {
	logs.Info("开始执行配额重置任务...")
	quotaService := &QuotaService{}
	count, err := quotaService.ResetExpiredQuotas()
	if err != nil {
		logs.Error("配额重置任务失败:", err)
		return
	}
	logs.Info("配额重置任务完成，重置数量:", count)
}

// RunPlanExpiryCheck 套餐到期检查任务（每小时执行）
func (s *CronService) RunPlanExpiryCheck() {
	logs.Info("开始执行套餐到期检查任务...")
	planService := &PlanService{}
	count, err := planService.CheckExpiredPlans()
	if err != nil {
		logs.Error("套餐到期检查任务失败:", err)
		return
	}
	logs.Info("套餐到期检查任务完成，处理数量:", count)
}

// RunQuotaCleanup 过期配额清理任务（每天执行）
func (s *CronService) RunQuotaCleanup() {
	logs.Info("开始执行过期配额清理任务...")
	quotaService := &QuotaService{}
	count, err := quotaService.CleanExpiredQuotas()
	if err != nil {
		logs.Error("过期配额清理任务失败:", err)
		return
	}
	logs.Info("过期配额清理任务完成，清理数量:", count)
}

// StartCronJobs 启动定时任务
func StartCronJobs() {
	cronService := &CronService{}

	// 每小时执行配额重置
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		for range ticker.C {
			cronService.RunQuotaReset()
		}
	}()

	// 每小时执行套餐到期检查
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		for range ticker.C {
			cronService.RunPlanExpiryCheck()
		}
	}()

	// 每天凌晨执行过期配额清理
	go func() {
		for {
			now := time.Now()
			next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
			time.Sleep(next.Sub(now))
			cronService.RunQuotaCleanup()
		}
	}()

	logs.Info("定时任务已启动")
}
