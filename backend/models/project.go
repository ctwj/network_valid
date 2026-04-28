package models

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
	"verification/controllers/common"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

// PlanSchemeConfig 预设方案套餐配置
type PlanSchemeConfig struct {
	Name       string  `json:"name"`        // 套餐名称
	Days       int     `json:"days"`        // 天数，0=永久
	Price      float64 `json:"price"`       // 价格
	Priority   int     `json:"priority"`    // 优先级
	IsFreeTier bool    `json:"is_free_tier"` // 是否为免费套餐
	QuotaLimit int64   `json:"quota_limit"` // 每日下载次数限制
	Savings    string  `json:"savings"`     // 节省提示，如"比月卡省10元"
}

// PlanScheme 预设方案
type PlanScheme struct {
	Name        string             `json:"name"`        // 方案名称
	Description string             `json:"description"` // 方案描述
	Plans       []PlanSchemeConfig `json:"plans"`       // 套餐列表
}

// GetDefaultPlanSchemes 返回默认预设方案（销售引导设计）
func GetDefaultPlanSchemes() []PlanScheme {
	return []PlanScheme{
		{
			Name:        "入门引导",
			Description: "适合个人轻度使用，价格亲民",
			Plans: []PlanSchemeConfig{
				{Name: "免费套餐", Days: 0, Price: 0, Priority: 0, IsFreeTier: true, QuotaLimit: 5, Savings: ""},
				{Name: "月卡套餐", Days: 30, Price: 30, Priority: 10, QuotaLimit: 20, Savings: ""},
				{Name: "季卡套餐", Days: 90, Price: 80, Priority: 20, QuotaLimit: 20, Savings: "比月卡省10元"},
			},
		},
		{
			Name:        "标准推荐",
			Description: "适合日常使用，性价比最高",
			Plans: []PlanSchemeConfig{
				{Name: "免费套餐", Days: 0, Price: 0, Priority: 0, IsFreeTier: true, QuotaLimit: 5, Savings: ""},
				{Name: "月卡套餐", Days: 30, Price: 49, Priority: 10, QuotaLimit: 50, Savings: ""},
				{Name: "季卡套餐", Days: 90, Price: 129, Priority: 20, QuotaLimit: 50, Savings: "比月卡省18元"},
				{Name: "年卡套餐", Days: 365, Price: 399, Priority: 30, QuotaLimit: 50, Savings: "比月卡省189元"},
			},
		},
		{
			Name:        "高级专业",
			Description: "适合重度用户，包含永久套餐",
			Plans: []PlanSchemeConfig{
				{Name: "免费套餐", Days: 0, Price: 0, Priority: 0, IsFreeTier: true, QuotaLimit: 5, Savings: ""},
				{Name: "月卡套餐", Days: 30, Price: 99, Priority: 10, QuotaLimit: 100, Savings: ""},
				{Name: "季卡套餐", Days: 90, Price: 249, Priority: 20, QuotaLimit: 100, Savings: "比月卡省48元"},
				{Name: "年卡套餐", Days: 365, Price: 699, Priority: 30, QuotaLimit: 100, Savings: "比月卡省489元"},
				{Name: "永久套餐", Days: 0, Price: 1999, Priority: 40, QuotaLimit: 100, Savings: "相当于3年年卡"},
			},
		},
	}
}

type Project struct {
	ID          int       `orm:"column(id)" json:"id"`
	ManagerId   int       `orm:"index;default(0)" json:"manager_id"`
	Name        string    `orm:"size(40)" json:"name" valid:"MaxSize(40)"`
	AppKey      string    `orm:"index;size(32)" json:"app_key" valid:"MaxSize(32)"`
	SecretKey   string    `orm:"size(32)" json:"secret_key" valid:"MaxSize(32)"`
	Type        int       `orm:"default(0);description(0单码,1账号)" json:"type" valid:"Max(1);MaxSize(1)"`
	StatusType  int       `orm:"default(0);description(0收费,1停止,2免费)" json:"status_type"`
	PublicKey   string    `orm:"type(text);size(4000);description(RSA公钥);null" json:"public_key" valid:"MaxSize(4000)"`
	PrivateKey  string    `orm:"type(text);size(4000);description(RSA私钥);null" json:"private_key" valid:"MaxSize(4000)"`
	KeyA        string    `orm:"size(16);description(AES加密密匙A);null" json:"key_a" valid:"MaxSize(16)"`
	KeyB        string    `orm:"size(16);description(AES加密密匙B);null" json:"key_b" valid:"MaxSize(16)"`
	Sign        int       `orm:"default(0);description(0 MD5 1 SHA1 2 SHA224 3 SHA256 4 SHA384 5 SHA512)" json:"sign" valid:"MaxSize(1)"`
	Encrypt     int       `orm:"default(0);description(0开放签名API,1AES)" json:"encrypt" valid:"MaxSize(1)"`
	LoginType   int       `orm:"default(0)" json:"login_type" valid:"MaxSize(1)"`
	GiftType    int       `orm:"default(0)" json:"gift_type" valid:"MaxSize(1)"`
	Notice      string    `orm:"type(text);description(公告);size(9000);null" json:"notice" valid:"MaxSize(9000)"`
	Api         string    `orm:"type(text);description(公告);size(5000);null" json:"api" valid:"MaxSize(5000)"`
	PlanPresets string    `orm:"type(text);null;description(预设套餐方案JSON)" json:"plan_presets"`
	CreateTime  time.Time `orm:"auto_now_add;type(datetime);index" json:"create_time"`
}

func (p *Project) GetProjectList(pageSize int64, page int64) (status bool, pager Pager) {
	var data []Project
	o := orm.NewOrm()
	count, err := o.QueryTable(&p).Count()
	if err != nil {
		logs.Error(err)
		return false, Pager{}
	}
	totalPage, offset, currentPage := PageCount(count, pageSize, page)
	_, err = o.QueryTable(&p).Limit(pageSize, offset).All(&data)
	return true, Pager{
		Count:       count,
		CurrentPage: currentPage,
		Data:        data,
		PageSize:    pageSize,
		TotalPages:  totalPage,
	}
}

func (p *Project) ProjectList() (status bool, pager Pager) {
	var data []Project
	o := orm.NewOrm()
	_, err := o.QueryTable(&p).All(&data, "ID", "Name")
	if err != nil {
		return false, Pager{}
	}
	return true, Pager{
		Count:       0,
		CurrentPage: 0,
		Data:        data,
		PageSize:    0,
		TotalPages:  0,
	}
}

func (p *Project) GetAgentProjectList(list []int) (status bool, pager Pager) {
	var data []Project
	o := orm.NewOrm()
	_, err := o.QueryTable(&p).Filter("ID__in", list).All(&data, "ID", "Name")
	if err != nil {
		return false, Pager{}
	}
	return true, Pager{
		Count:       0,
		CurrentPage: 0,
		Data:        data,
		PageSize:    0,
		TotalPages:  0,
	}
}

func (c *Project) Add(name string, projectType int, statusType int, encrypt int, notice string, api string, managerId int, sign int, schemeName string, monthlyPrice float64) (projectId int64) {
	status, publicKey, privateKey := GetRsaKey()
	if status == false {
		return 0
	}
	// 强制使用用户登录模式（禁用单码模式）
	projectType = 1
	p := Project{
		ManagerId:  managerId,
		Name:       name,
		AppKey:     RandStr(32),
		SecretKey:  RandStr(32),
		Type:       projectType,
		StatusType: statusType,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		KeyA:       RandStr(16),
		KeyB:       RandStr(16),
		Encrypt:    encrypt,
		LoginType:  0,
		GiftType:   0,
		Notice:     notice,
		Api:        api,
		Sign:       sign,
	}
	o := orm.NewOrm()
	id, err := o.Insert(&p)
	if err != nil {
		return 0
	}
	projectId = id

	// 创建默认登录规则
	loginRule := &ProjectLogin{
		ManagerId:    managerId,
		Title:        "默认规则",
		Mode:         1, // 普通登录
		RegMode:      1, // 普通注册
		UnbindMode:   3, // 任意解绑
		UnbindTimes:  3,
		UnbindDate:   1,
		NumberMore:   10,
		NumberWeaken: 1,
	}
	loginId, err := o.Insert(loginRule)
	if err == nil && loginId > 0 {
		// 绑定登录规则到项目
		p.ID = int(id)
		p.LoginType = int(loginId)
		o.Update(&p, "LoginType")
	}

	// 创建默认版本号
	version := &ProjectVersion{
		ManagerId:    managerId,
		ProjectId:    int(id),
		Version:      1.00,
		IsMustUpdate: 0,
		IsActive:     0,
	}
	_, err = o.Insert(version)
	if err != nil {
		logs.Error("创建默认版本号失败:", err)
	}

	// 根据选择的方案创建套餐
	c.createPlansFromScheme(int(id), managerId, schemeName, monthlyPrice)

	_, ac := common.GetCacheAC()
	data, err := json.Marshal(&p)
	_ = ac.Put(common.GetStringMd5(p.AppKey), string(data), 365*24*60*60*time.Second)
	_ = ac.Put(p.AppKey, string(data), 365*24*60*60*time.Second)
	return id
}

// createPlansFromScheme 根据预设方案创建套餐
func (c *Project) createPlansFromScheme(projectId int, managerId int, schemeName string, monthlyPrice float64) {
	schemes := GetDefaultPlanSchemes()

	// 查找选择的方案，默认使用"标准推荐"
	var selectedScheme *PlanScheme
	for i := range schemes {
		if schemes[i].Name == schemeName {
			selectedScheme = &schemes[i]
			break
		}
	}
	if selectedScheme == nil {
		// 默认使用标准推荐
		for i := range schemes {
			if schemes[i].Name == "标准推荐" {
				selectedScheme = &schemes[i]
				break
			}
		}
	}
	if selectedScheme == nil {
		selectedScheme = &schemes[0]
	}

	// 如果没有指定月费，使用默认值 30
	if monthlyPrice <= 0 {
		monthlyPrice = 30
	}

	o := orm.NewOrm()
	for _, plan := range selectedScheme.Plans {
		isFreeTier := 0
		if plan.IsFreeTier {
			isFreeTier = 1
		}

		// 基于月费动态计算价格
		price := calculatePlanPrice(plan, monthlyPrice, selectedScheme.Name)

		// 生成配额规则 JSON
		quotaRules := fmt.Sprintf(`{"quotas":[{"key":"download","name":"下载次数","limit":%d,"period":"daily","unit":"count"}]}`, plan.QuotaLimit)

		card := &Cards{
			ProjectId:  projectId,
			Title:      plan.Name,
			Price:      price,
			Days:       float64(plan.Days),
			Points:     0,
			Tag:        plan.Name + "用户",
			IsFreeTier: isFreeTier,
			Priority:   plan.Priority,
			QuotaRules: quotaRules,
			ManagerId:  managerId,
		}
		_, err := o.Insert(card)
		if err != nil {
			logs.Error("创建套餐失败:", err, "plan:", plan.Name)
		}
	}

	logs.Info("项目 %d 根据方案 %s 创建套餐成功，月费基准: %.2f", projectId, selectedScheme.Name, monthlyPrice)
}

// calculatePlanPrice 根据月费基准计算套餐价格
func calculatePlanPrice(plan PlanSchemeConfig, monthlyPrice float64, schemeName string) float64 {
	// 免费套餐
	if plan.IsFreeTier {
		return 0
	}

	// 永久套餐（天数=0 且非免费）
	if plan.Days == 0 {
		switch schemeName {
		case "高级专业":
			return roundPrice(monthlyPrice * 66.6) // 约 3 年年卡
		}
		return roundPrice(monthlyPrice * 66.6)
	}

	// 根据天数和方案类型计算价格
	switch schemeName {
	case "入门引导":
		switch plan.Days {
		case 30:
			return roundPrice(monthlyPrice)
		case 90:
			return roundPrice(monthlyPrice * 2.67) // 约省 10%
		}
	case "标准推荐":
		switch plan.Days {
		case 30:
			return roundPrice(monthlyPrice)
		case 90:
			return roundPrice(monthlyPrice * 2.63) // 约省 12%
		case 365:
			return roundPrice(monthlyPrice * 8.14) // 约省 32%
		}
	case "高级专业":
		switch plan.Days {
		case 30:
			return roundPrice(monthlyPrice * 3.3)
		case 90:
			return roundPrice(monthlyPrice * 8.3)
		case 365:
			return roundPrice(monthlyPrice * 23.3)
		}
	}

	// 默认：按天数比例计算
	return roundPrice(monthlyPrice * float64(plan.Days) / 30)
}

// roundPrice 价格保留两位小数
func roundPrice(price float64) float64 {
	return math.Round(price*100) / 100
}

func (c *Project) Update(id int, name string, projectType int, statusType int, encrypt int, notice string, api string, updateRsa int, updateKey int, updateAppKey int, updateSecretKey int, sign int) bool {
	p := Project{
		ID: id,
	}
	_, ac := common.GetCacheAC()
	o := orm.NewOrm()
	err := o.Read(&p)
	if err != nil {
		logs.Error("查询项目失败")
		return false
	}
	p.Name = name
	p.Type = projectType
	p.StatusType = statusType
	p.Encrypt = encrypt
	p.Sign = sign
	if notice != "" {
		p.Notice = notice
	}
	if api != "" {
		p.Api = api
	}
	if updateRsa > 0 {
		status, publicKey, privateKey := GetRsaKey()
		if status == false {
			return false
		}
		p.PublicKey = publicKey
		p.PrivateKey = privateKey
	}
	if updateKey > 0 {
		p.KeyA = RandStr(16)
		p.KeyB = RandStr(16)
	}
	if updateAppKey > 0 {
		_ = ac.Delete(common.GetStringMd5(p.AppKey))
		_ = ac.Delete(p.AppKey)
		p.AppKey = RandStr(32)
	}
	if updateSecretKey > 0 {
		p.SecretKey = RandStr(32)
	}

	row, err := o.Update(&p)
	if err != nil {
		return false
	}
	if row > 0 {

		data, _ := json.Marshal(&p)
		_ = ac.Put(common.GetStringMd5(p.AppKey), string(data), 3*365*24*60*60*time.Second)
		_ = ac.Put(p.AppKey, string(data), 3*365*24*60*60*time.Second)
		return true
	}
	return false
}

func (c *Project) Delete(id int) bool {
	p := Project{
		ID: id,
	}
	o := orm.NewOrm()
	err := o.Read(&p)
	if err != nil {
		logs.Error("查询项目失败")
		return false
	}
	rows, err := o.Delete(&p)
	if rows > 0 {
		_, err = o.Raw("delete from cards where project_id = ?", p.ID).Exec()
		_, err = o.Raw("delete from keys where project_id = ?", p.ID).Exec()
		_, err = o.Raw("delete from project_version where project_id = ?", p.ID).Exec()
		_, err = o.Raw("delete from order where project_id = ?", p.ID).Exec()
		_, err = o.Raw("delete from member where project_id = ?", p.ID).Exec()
		_, err = o.Raw("delete from member_login where project_id = ?", p.ID).Exec()
		_, err = o.Raw("delete from manager_cards where project_id = ?", p.ID).Exec()
		_, ac := common.GetCacheAC()
		_ = ac.Delete(fmt.Sprintf(cacheIndex[0], id))
		_ = ac.Delete(common.GetStringMd5(p.AppKey))
		_ = ac.Delete(p.AppKey)
		return true
	}
	return false
}

func (c *Project) Bind(id int, projectLoginId int) (status bool, msg string) {
	o := orm.NewOrm()
	if projectLoginId > 0 {

		projectLogin := ProjectLogin{ID: projectLoginId}
		err := o.Read(&projectLogin)
		if err != nil {
			return false, "登录规则不存在"
		}
	}
	project := Project{ID: id}
	err := o.Read(&project)
	if err != nil {
		return false, "项目不存在"
	}
	if project.LoginType > 0 {
		project.LoginType = 0
		msg = "解绑登录规则成功"
	} else {
		project.LoginType = projectLoginId
		msg = "绑定登录规则成功"
	}

	row, err := o.Update(&project)
	if row > 0 {
		_, ac := common.GetCacheAC()
		data, _ := json.Marshal(&project)
		_ = ac.Put(fmt.Sprintf(cacheIndex[0], id), string(data), 3*365*24*60*60*time.Second)
		_ = ac.Put(common.GetStringMd5(project.AppKey), string(data), 3*365*24*60*60*time.Second)
		_ = ac.Put(project.AppKey, string(data), 3*365*24*60*60*time.Second)
		return true, msg
	}
	return false, "操作失败"
}

func (p *Project) GetCount() int64 {
	managerId := p.ManagerId
	o := orm.NewOrm()
	qs := o.QueryTable("Project")
	count, err := qs.Filter("ManagerId", managerId).Count()
	if err != nil {
		logs.Error("查询项目总数错误：", err)
		return 0
	}
	return count
}

func (p *Project) AllProject() []Project {
	data := make([]Project, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable(&p).All(&data)
	if err != nil {
		return data
	}
	return data
}
