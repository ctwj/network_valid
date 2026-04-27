package services

import (
	"encoding/json"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"

	"verification/models"
)

// PlanPresetService 预设套餐方案服务
type PlanPresetService struct{}

// GetPresets 获取预设套餐方案
func (s *PlanPresetService) GetPresets(projectId int) (*models.PlanPresets, error) {
	o := orm.NewOrm()

	project := &models.Project{ID: projectId}
	err := o.Read(project)
	if err != nil {
		return nil, err
	}

	// 如果项目没有配置预设方案，返回默认方案
	if project.PlanPresets == "" {
		return s.GetDefaultPresets(), nil
	}

	var presets models.PlanPresets
	err = json.Unmarshal([]byte(project.PlanPresets), &presets)
	if err != nil {
		logs.Error("解析预设方案失败:", err)
		return s.GetDefaultPresets(), nil
	}

	return &presets, nil
}

// GetDefaultPresets 获取默认预设方案
func (s *PlanPresetService) GetDefaultPresets() *models.PlanPresets {
	return &models.PlanPresets{
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
}

// SavePresets 保存预设方案配置
func (s *PlanPresetService) SavePresets(projectId int, presets *models.PlanPresets) error {
	o := orm.NewOrm()

	data, err := json.Marshal(presets)
	if err != nil {
		return err
	}

	project := &models.Project{ID: projectId}
	err = o.Read(project)
	if err != nil {
		return err
	}

	project.PlanPresets = string(data)
	_, err = o.Update(project, "PlanPresets")
	return err
}

// CreateFromPreset 从预设方案创建套餐类型
func (s *PlanPresetService) CreateFromPreset(projectId int, managerId int, presetNames []string) error {
	presets, err := s.GetPresets(projectId)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	for _, preset := range presets.Presets {
		// 检查是否在选中列表中
		selected := false
		for _, name := range presetNames {
			if name == preset.Name {
				selected = true
				break
			}
		}
		if !selected {
			continue
		}

		// 转换配额规则为 JSON
		var quotaRulesJSON string
		if preset.QuotaRules != nil {
			data, _ := json.Marshal(preset.QuotaRules)
			quotaRulesJSON = string(data)
		}

		// 创建套餐类型
		card := &models.Cards{
			ProjectId:  projectId,
			ManagerId:  managerId,
			Title:      preset.Name,
			Days:       float64(preset.Days),
			QuotaRules: quotaRulesJSON,
			IsFreeTier: func() int { if preset.IsFreeTier { return 1 }; return 0 }(),
			Priority:   preset.Priority,
		}

		_, err := o.Insert(card)
		if err != nil {
			logs.Error("创建套餐类型失败:", err, "preset:", preset.Name)
			continue
		}
	}

	return nil
}

// CreateDefaultPlans 创建默认套餐（项目创建时调用）
func (s *PlanPresetService) CreateDefaultPlans(projectId int, managerId int) error {
	presets := s.GetDefaultPresets()

	o := orm.NewOrm()

	for _, preset := range presets.Presets {
		// 转换配额规则为 JSON
		var quotaRulesJSON string
		if preset.QuotaRules != nil {
			data, _ := json.Marshal(preset.QuotaRules)
			quotaRulesJSON = string(data)
		}

		// 创建套餐类型
		card := &models.Cards{
			ProjectId:  projectId,
			ManagerId:  managerId,
			Title:      preset.Name,
			Days:       float64(preset.Days),
			QuotaRules: quotaRulesJSON,
			IsFreeTier: func() int { if preset.IsFreeTier { return 1 }; return 0 }(),
			Priority:   preset.Priority,
		}

		_, err := o.Insert(card)
		if err != nil {
			logs.Error("创建默认套餐失败:", err, "preset:", preset.Name)
			continue
		}
	}

	return nil
}
