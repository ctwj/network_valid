package api

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/adapter/orm"

	"verification/models"
	"verification/validation/api"
	"verification/services"
)

// QuotaController 配额相关接口
type QuotaController struct {
	IndexController
}

// Check 检查配额
// @router / [post,get]
func (p *QuotaController) Check() {
	// 获取用户信息
	param := api.ClientParam{}
	if p.IsEncrypt == false {
		if p.IsPost == true && p.ParamContentType == ApplicationJSON {
			_ = p.Ctx.BindJSON(&param)
		} else {
			param.User = p.GetString("user", "")
		}
	}
	if p.IsEncrypt == true {
		_ = json.Unmarshal([]byte(p.decryptJsonString), &param)
	}

	// 获取用户
	u := models.Member{}
	if p.Project.Type == 0 {
		u.Name = param.User
	} else {
		u.Name = param.User
	}
	status, _ := u.CheckMember()
	if !status {
		p.CallErrorJson("用户不存在", nil)
		return
	}

	// 获取要查询的配额标识（可选）
	keysParam := p.GetString("keys", "")
	var keys []string
	if keysParam != "" {
		keys = strings.Split(keysParam, ",")
	}

	// 查询配额
	quotaService := &services.QuotaService{}
	quotas, err := quotaService.CheckQuotas(u.ID, keys)
	if err != nil {
		p.CallErrorJson("查询配额失败", err.Error())
		return
	}

	p.CallJson("查询成功", map[string]interface{}{
		"quotas": quotas,
	})
}

// Deduct 扣减配额
// @router / [post,get]
func (p *QuotaController) Deduct() {
	// 获取用户信息
	param := api.ClientParam{}
	if p.IsEncrypt == false {
		if p.IsPost == true && p.ParamContentType == ApplicationJSON {
			_ = p.Ctx.BindJSON(&param)
		} else {
			param.User = p.GetString("user", "")
		}
	}
	if p.IsEncrypt == true {
		_ = json.Unmarshal([]byte(p.decryptJsonString), &param)
	}

	// 获取用户
	u := models.Member{}
	u.Name = param.User
	status, _ := u.CheckMember()
	if !status {
		p.CallErrorJson("用户不存在", nil)
		return
	}

	// 获取参数
	key := p.GetString("key", "")
	amountStr := p.GetString("amount", "1")
	amount, _ := strconv.ParseInt(amountStr, 10, 64)

	if key == "" {
		p.CallErrorJson("配额标识不能为空", nil)
		return
	}

	// 扣减配额
	quotaService := &services.QuotaService{}
	result, err := quotaService.DeductQuota(u.ID, key, amount)
	if err != nil {
		p.CallErrorJson("扣减配额失败", err.Error())
		return
	}

	if !result.Success {
		p.CallErrorJson(result.Message, result)
		return
	}

	p.CallJson("扣减成功", result)
}

// DeductBatch 批量扣减配额
// @router / [post,get]
func (p *QuotaController) DeductBatch() {
	// 获取用户信息
	param := api.ClientParam{}
	if p.IsEncrypt == false {
		if p.IsPost == true && p.ParamContentType == ApplicationJSON {
			_ = p.Ctx.BindJSON(&param)
		} else {
			param.User = p.GetString("user", "")
		}
	}
	if p.IsEncrypt == true {
		_ = json.Unmarshal([]byte(p.decryptJsonString), &param)
	}

	// 获取用户
	u := models.Member{}
	u.Name = param.User
	status, _ := u.CheckMember()
	if !status {
		p.CallErrorJson("用户不存在", nil)
		return
	}

	// 获取参数（JSON 格式的 items）
	itemsParam := p.GetString("items", "")
	if itemsParam == "" {
		p.CallErrorJson("扣减项不能为空", nil)
		return
	}

	var items []models.QuotaDeductItem
	err := json.Unmarshal([]byte(itemsParam), &items)
	if err != nil {
		p.CallErrorJson("参数格式错误", err.Error())
		return
	}

	// 批量扣减
	quotaService := &services.QuotaService{}
	results, err := quotaService.DeductQuotaBatch(u.ID, items)
	if err != nil {
		p.CallErrorJson("批量扣减失败", err.Error())
		return
	}

	p.CallJson("扣减成功", map[string]interface{}{
		"results": results,
	})
}

// PlanController 套餐相关接口
type PlanController struct {
	IndexController
}

// List 获取项目套餐列表
// @router / [post,get]
func (p *PlanController) List() {
	projectId := p.Project.ID

	// 获取套餐列表
	o := orm.NewOrm()
	var cards []models.Cards
	_, err := o.QueryTable("cards").
		Filter("project_id", projectId).
		All(&cards)
	if err != nil {
		p.CallErrorJson("获取套餐列表失败", err.Error())
		return
	}

	p.CallJson("获取成功", map[string]interface{}{
		"plans": cards,
	})
}

// Current 获取当前套餐
// @router / [post,get]
func (p *PlanController) Current() {
	// 获取用户信息
	param := api.ClientParam{}
	if p.IsEncrypt == false {
		if p.IsPost == true && p.ParamContentType == ApplicationJSON {
			_ = p.Ctx.BindJSON(&param)
		} else {
			param.User = p.GetString("user", "")
		}
	}
	if p.IsEncrypt == true {
		_ = json.Unmarshal([]byte(p.decryptJsonString), &param)
	}

	// 获取用户
	u := models.Member{}
	u.Name = param.User
	status, _ := u.CheckMember()
	if !status {
		p.CallErrorJson("用户不存在", nil)
		return
	}

	planService := &services.PlanService{}
	plan, hasPlan := planService.GetCurrentPlan(u.ID)

	if !hasPlan {
		p.CallJson("暂无套餐", map[string]interface{}{
			"plan": nil,
		})
		return
	}

	// 获取套餐详情
	o := orm.NewOrm()
	card := &models.Cards{ID: plan.PlanId}
	o.Read(card)

	p.CallJson("获取成功", map[string]interface{}{
		"plan": map[string]interface{}{
			"id":           plan.ID,
			"plan_id":      plan.PlanId,
			"plan_name":    card.Title,
			"status":       plan.Status,
			"expire_time":  plan.ExpireTime,
			"activated_at": plan.ActivatedAt,
			"days":         card.Days,
			"priority":     card.Priority,
			"is_free_tier": card.IsFreeTier,
		},
	})
}

// Queued 获取排队套餐
// @router / [post,get]
func (p *PlanController) Queued() {
	// 获取用户信息
	param := api.ClientParam{}
	if p.IsEncrypt == false {
		if p.IsPost == true && p.ParamContentType == ApplicationJSON {
			_ = p.Ctx.BindJSON(&param)
		} else {
			param.User = p.GetString("user", "")
		}
	}
	if p.IsEncrypt == true {
		_ = json.Unmarshal([]byte(p.decryptJsonString), &param)
	}

	// 获取用户
	u := models.Member{}
	u.Name = param.User
	status, _ := u.CheckMember()
	if !status {
		p.CallErrorJson("用户不存在", nil)
		return
	}

	planService := &services.PlanService{}
	plans, err := planService.GetQueuedPlans(u.ID)
	if err != nil {
		p.CallErrorJson("获取排队套餐失败", err.Error())
		return
	}

	p.CallJson("获取成功", map[string]interface{}{
		"plans": plans,
	})
}
