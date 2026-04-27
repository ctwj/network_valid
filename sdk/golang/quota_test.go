package sdk

import (
	"testing"
)

// TestQuotaInfo 测试 QuotaInfo 结构体
func TestQuotaInfo(t *testing.T) {
	quota := QuotaInfo{
		Key:       "download",
		Name:      "下载次数",
		Limit:     20,
		Used:      5,
		Remaining: 15,
		Period:    "daily",
		Unit:      "count",
	}

	if quota.Key != "download" {
		t.Errorf("expected key 'download', got '%s'", quota.Key)
	}
	if quota.Remaining != 15 {
		t.Errorf("expected remaining 15, got %d", quota.Remaining)
	}
}

// TestQuotaDeductItem 测试 QuotaDeductItem 结构体
func TestQuotaDeductItem(t *testing.T) {
	item := QuotaDeductItem{
		Key:    "download",
		Amount: 1,
	}

	if item.Key != "download" {
		t.Errorf("expected key 'download', got '%s'", item.Key)
	}
	if item.Amount != 1 {
		t.Errorf("expected amount 1, got %d", item.Amount)
	}
}

// TestPlanInfo 测试 PlanInfo 结构体
func TestPlanInfo(t *testing.T) {
	plan := PlanInfo{
		PlanID:     1,
		PlanName:   "标准套餐",
		Status:     "active",
		Days:       30,
		Priority:   10,
		IsFreeTier: 0,
	}

	if plan.PlanName != "标准套餐" {
		t.Errorf("expected plan name '标准套餐', got '%s'", plan.PlanName)
	}
	if plan.Status != "active" {
		t.Errorf("expected status 'active', got '%s'", plan.Status)
	}
}

// TestCheckQuota 测试配额查询（需要真实服务器）
// 此测试需要配置真实的服务器地址和凭证
func TestCheckQuota(t *testing.T) {
	// 跳过集成测试，除非设置了环境变量
	t.Skip("Integration test - requires real server")
}

// TestDeductQuota 测试配额扣减（需要真实服务器）
func TestDeductQuota(t *testing.T) {
	t.Skip("Integration test - requires real server")
}

// TestDeductQuotaBatch 测试批量扣减（需要真实服务器）
func TestDeductQuotaBatch(t *testing.T) {
	t.Skip("Integration test - requires real server")
}

// TestGetCurrentPlan 测试获取当前套餐（需要真实服务器）
func TestGetCurrentPlan(t *testing.T) {
	t.Skip("Integration test - requires real server")
}
