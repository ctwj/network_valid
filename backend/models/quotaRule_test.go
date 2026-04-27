package models

import (
	"testing"
	"time"
)

func TestParseQuotaRules(t *testing.T) {
	// 测试空字符串
	rules, err := ParseQuotaRules("")
	if err != nil {
		t.Errorf("expected no error for empty string, got: %v", err)
	}
	if rules != nil {
		t.Errorf("expected nil for empty string, got: %v", rules)
	}

	// 测试有效 JSON
	jsonStr := `{"quotas": [{"key": "download", "name": "下载次数", "limit": 20, "period": "daily", "unit": "count"}]}`
	rules, err = ParseQuotaRules(jsonStr)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if rules == nil {
		t.Fatal("expected rules, got nil")
	}
	if len(rules.Quotas) != 1 {
		t.Errorf("expected 1 quota, got %d", len(rules.Quotas))
	}
	if rules.Quotas[0].Key != "download" {
		t.Errorf("expected key 'download', got '%s'", rules.Quotas[0].Key)
	}
	if rules.Quotas[0].Limit != 20 {
		t.Errorf("expected limit 20, got %d", rules.Quotas[0].Limit)
	}

	// 测试无效 JSON
	_, err = ParseQuotaRules("invalid json")
	if err == nil {
		t.Error("expected error for invalid json")
	}
}

func TestQuotaRulesToJSON(t *testing.T) {
	rules := &QuotaRules{
		Quotas: []QuotaRule{
			{Key: "download", Name: "下载次数", Limit: 20, Period: "daily", Unit: "count"},
		},
	}

	jsonStr, err := rules.ToJSON()
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if jsonStr == "" {
		t.Error("expected non-empty json string")
	}

	// 测试 nil
	var nilRules *QuotaRules
	jsonStr, err = nilRules.ToJSON()
	if err != nil {
		t.Errorf("expected no error for nil, got: %v", err)
	}
	if jsonStr != "" {
		t.Errorf("expected empty string for nil, got: %s", jsonStr)
	}
}

func TestCalculateNextResetTime_Daily(t *testing.T) {
	// 测试每日重置
	now := time.Now()
	next := CalculateNextResetTime("daily", 0)

	// 应该是明天 00:00
	expected := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	if !next.Equal(expected) {
		t.Errorf("daily reset: expected %v, got %v", expected, next)
	}
}

func TestCalculateNextResetTime_Weekly(t *testing.T) {
	// 测试每周重置
	now := time.Now()
	currentWeekday := int(now.Weekday())
	if currentWeekday == 0 {
		currentWeekday = 7
	}

	// 测试重置日为周一 (1)
	next := CalculateNextResetTime("weekly", 1)

	// 计算预期天数
	daysUntilMonday := 1 - currentWeekday
	if daysUntilMonday <= 0 {
		daysUntilMonday += 7
	}
	expected := time.Date(now.Year(), now.Month(), now.Day()+daysUntilMonday, 0, 0, 0, 0, now.Location())

	if !next.Equal(expected) {
		t.Errorf("weekly reset (Monday): expected %v, got %v", expected, next)
	}

	// 验证结果是周一
	if next.Weekday() != time.Monday {
		t.Errorf("expected Monday, got %v", next.Weekday())
	}
}

func TestCalculateNextResetTime_Monthly(t *testing.T) {
	// 测试每月重置
	now := time.Now()

	// 测试重置日为 1 号
	next := CalculateNextResetTime("monthly", 1)

	// 应该是下月 1 号或本月 1 号（如果当前日期 < 1）
	var expectedMonth time.Month
	expectedYear := now.Year()

	if now.Day() >= 1 {
		expectedMonth = now.Month() + 1
		if expectedMonth > 12 {
			expectedMonth = 1
			expectedYear++
		}
	} else {
		expectedMonth = now.Month()
	}

	expected := time.Date(expectedYear, expectedMonth, 1, 0, 0, 0, 0, now.Location())

	if !next.Equal(expected) {
		t.Errorf("monthly reset (1st): expected %v, got %v", expected, next)
	}

	// 验证日期是 1 号
	if next.Day() != 1 {
		t.Errorf("expected day 1, got %d", next.Day())
	}
}

func TestCalculateNextResetTime_Monthly_EdgeCase(t *testing.T) {
	// 测试月份天数不足的情况（如 2 月 30 号）
	// 这种情况应该使用该月最后一天

	// 这里只验证函数不会 panic
	next := CalculateNextResetTime("monthly", 30)
	if next.IsZero() {
		t.Error("expected non-zero time for monthly reset with day 30")
	}
}

func TestCalculateNextResetTime_Default(t *testing.T) {
	// 测试未知周期（默认为每日）
	now := time.Now()
	next := CalculateNextResetTime("unknown", 0)

	expected := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	if !next.Equal(expected) {
		t.Errorf("default reset: expected %v, got %v", expected, next)
	}
}

func TestIsExpired(t *testing.T) {
	// 测试已过期
	past := time.Now().Add(-1 * time.Hour)
	if !IsExpired(past) {
		t.Error("expected expired for past time")
	}

	// 测试未过期
	future := time.Now().Add(1 * time.Hour)
	if IsExpired(future) {
		t.Error("expected not expired for future time")
	}
}

func TestIsPlanExpired(t *testing.T) {
	// 测试 nil（永不过期）
	if IsPlanExpired(nil) {
		t.Error("expected not expired for nil expire time")
	}

	// 测试已过期
	past := time.Now().Add(-1 * time.Hour)
	if !IsPlanExpired(&past) {
		t.Error("expected expired for past time")
	}

	// 测试未过期
	future := time.Now().Add(24 * time.Hour)
	if IsPlanExpired(&future) {
		t.Error("expected not expired for future time")
	}
}

func TestQuotaRule(t *testing.T) {
	rule := QuotaRule{
		Key:      "download",
		Name:     "下载次数",
		Limit:    20,
		Period:   "daily",
		ResetDay: 1,
		Unit:     "count",
	}

	if rule.Key != "download" {
		t.Errorf("expected key 'download', got '%s'", rule.Key)
	}
	if rule.Limit != 20 {
		t.Errorf("expected limit 20, got %d", rule.Limit)
	}
}

func TestQuotaDeductItem(t *testing.T) {
	item := QuotaDeductItem{
		Key:    "download",
		Amount: 5,
	}

	if item.Key != "download" {
		t.Errorf("expected key 'download', got '%s'", item.Key)
	}
	if item.Amount != 5 {
		t.Errorf("expected amount 5, got %d", item.Amount)
	}
}

func TestPlanPreset(t *testing.T) {
	preset := PlanPreset{
		Name:       "免费套餐",
		Days:       0,
		IsFreeTier: true,
		Priority:   0,
		QuotaRules: &QuotaRules{
			Quotas: []QuotaRule{
				{Key: "download", Name: "下载次数", Limit: 5, Period: "daily", Unit: "count"},
			},
		},
	}

	if preset.Name != "免费套餐" {
		t.Errorf("expected name '免费套餐', got '%s'", preset.Name)
	}
	if !preset.IsFreeTier {
		t.Error("expected IsFreeTier to be true")
	}
	if preset.QuotaRules == nil {
		t.Error("expected QuotaRules not nil")
	}
}
