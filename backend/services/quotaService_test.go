package services

import (
	"encoding/json"
	"testing"

	"verification/models"
)

// TestQuotaService_ParseQuotaRules 测试配额规则解析
func TestQuotaService_ParseQuotaRules(t *testing.T) {
	s := &QuotaService{}

	// 测试空字符串
	rules, err := s.ParseQuotaRules("")
	if err != nil {
		t.Errorf("expected no error for empty string, got: %v", err)
	}
	if rules != nil {
		t.Errorf("expected nil for empty string, got: %v", rules)
	}

	// 测试有效 JSON
	jsonStr := `{"quotas": [{"key": "download", "name": "下载次数", "limit": 20, "period": "daily", "unit": "count"}]}`
	rules, err = s.ParseQuotaRules(jsonStr)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if rules == nil {
		t.Fatal("expected rules, got nil")
	}
	if len(rules.Quotas) != 1 {
		t.Errorf("expected 1 quota, got %d", len(rules.Quotas))
	}
}

// TestQuotaService_GetDefaultPresets 测试获取默认预设
func TestQuotaService_GetDefaultPresets(t *testing.T) {
	s := &QuotaService{}

	presets := s.GetDefaultPresets()
	if presets == "" {
		t.Error("expected non-empty presets")
	}

	// 验证 JSON 格式
	var planPresets models.PlanPresets
	err := json.Unmarshal([]byte(presets), &planPresets)
	if err != nil {
		t.Errorf("failed to parse presets: %v", err)
	}

	if len(planPresets.Presets) < 2 {
		t.Errorf("expected at least 2 presets, got %d", len(planPresets.Presets))
	}

	// 验证免费套餐
	freeTier := planPresets.Presets[0]
	if !freeTier.IsFreeTier {
		t.Error("first preset should be free tier")
	}
	if freeTier.Priority != 0 {
		t.Errorf("free tier priority should be 0, got %d", freeTier.Priority)
	}
}

// TestDeductResult 测试扣减结果结构体
func TestDeductResult(t *testing.T) {
	result := DeductResult{
		Key:       "download",
		Success:   true,
		Used:      5,
		Remaining: 15,
		Message:   "",
	}

	if result.Key != "download" {
		t.Errorf("expected key 'download', got '%s'", result.Key)
	}
	if !result.Success {
		t.Error("expected success to be true")
	}
	if result.Remaining != 15 {
		t.Errorf("expected remaining 15, got %d", result.Remaining)
	}
}

// TestQuotaInfo 测试配额信息结构体
func TestQuotaInfo(t *testing.T) {
	info := QuotaInfo{
		Key:       "download",
		Name:      "下载次数",
		Limit:     20,
		Used:      5,
		Remaining: 15,
		Period:    "daily",
		Unit:      "count",
		Unlimited: false,
	}

	if info.Key != "download" {
		t.Errorf("expected key 'download', got '%s'", info.Key)
	}
	if info.Unlimited {
		t.Error("expected unlimited to be false")
	}
}

// 以下测试需要数据库连接，在集成测试中运行

func TestQuotaService_InitQuotas(t *testing.T) {
	t.Skip("Requires database connection")
}

func TestQuotaService_CheckQuotas(t *testing.T) {
	t.Skip("Requires database connection")
}

func TestQuotaService_DeductQuota(t *testing.T) {
	t.Skip("Requires database connection")
}

func TestQuotaService_DeductQuotaBatch(t *testing.T) {
	t.Skip("Requires database connection")
}

func TestQuotaService_ResetExpiredQuotas(t *testing.T) {
	t.Skip("Requires database connection")
}
