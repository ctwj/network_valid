package sdk

import (
	"encoding/json"
	"fmt"
)

// CheckQuota 查询配额状态
// keys: 要查询的配额标识列表，为空则查询所有
func (c *Client) CheckQuota(keys []string) ([]QuotaInfo, error) {
	params := map[string]string{
		"action": "quota.check",
	}
	if len(keys) > 0 {
		// 将 keys 数组转为逗号分隔的字符串
		keysStr := ""
		for i, k := range keys {
			if i > 0 {
				keysStr += ","
			}
			keysStr += k
		}
		params["keys"] = keysStr
	}

	resp, err := c.doMultipartRequest(endpointAPI, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Quotas []QuotaInfo `json:"quotas"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	return result.Quotas, nil
}

// DeductQuota 扣减单个配额
// key: 配额标识
// amount: 扣减数量
func (c *Client) DeductQuota(key string, amount int64) (*QuotaDeductResult, error) {
	params := map[string]string{
		"action": "quota.deduct",
		"key":    key,
	}

	// amount 转为字符串
	params["amount"] = fmt.Sprintf("%d", amount)

	resp, err := c.doMultipartRequest(endpointAPI, params)
	if err != nil {
		return nil, err
	}

	var result QuotaDeductResult
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	return &result, nil
}

// DeductQuotaBatch 批量扣减配额
// items: 扣减项列表
func (c *Client) DeductQuotaBatch(items []QuotaDeductItem) ([]QuotaDeductResult, error) {
	// 将 items 转为 JSON 字符串
	itemsJSON, err := json.Marshal(items)
	if err != nil {
		return nil, fmt.Errorf("marshal items failed: %w", err)
	}

	params := map[string]string{
		"action": "quota.deductBatch",
		"items":  string(itemsJSON),
	}

	resp, err := c.doMultipartRequest(endpointAPI, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Results []QuotaDeductResult `json:"results"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	return result.Results, nil
}

// GetCurrentPlan 获取当前套餐信息
func (c *Client) GetCurrentPlan() (*PlanInfo, error) {
	params := map[string]string{
		"action": "plan.current",
	}

	resp, err := c.doMultipartRequest(endpointAPI, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Plan *PlanInfo `json:"plan"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	return result.Plan, nil
}

// GetPlanList 获取项目套餐列表
func (c *Client) GetPlanList() ([]map[string]interface{}, error) {
	params := map[string]string{
		"action": "plan.list",
	}

	resp, err := c.doMultipartRequest(endpointAPI, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Plans []map[string]interface{} `json:"plans"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	return result.Plans, nil
}

// GetQueuedPlans 获取排队套餐列表
func (c *Client) GetQueuedPlans() ([]PlanInfo, error) {
	params := map[string]string{
		"action": "plan.queued",
	}

	resp, err := c.doMultipartRequest(endpointAPI, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Plans []PlanInfo `json:"plans"`
	}
	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("parse response failed: %w", err)
	}

	return result.Plans, nil
}
