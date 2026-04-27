package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

// MemberPlan 用户套餐关联
type MemberPlan struct {
	ID          int64      `orm:"column(id);pk;auto" json:"id"`
	MemberId    int        `orm:"column(member_id);index;description(用户ID)" json:"member_id"`
	PlanId      int        `orm:"column(plan_id);index;description(套餐类型ID)" json:"plan_id"`
	Status      string     `orm:"column(status);size(16);default(active);description(状态:active/expired/queued)" json:"status"`
	ExpireTime  *time.Time `orm:"column(expire_time);type(datetime);null;description(过期时间,NULL表示永久)" json:"expire_time"`
	ActivatedAt time.Time  `orm:"column(activated_at);type(datetime);null;description(激活时间)" json:"activated_at"`
	ExpiredAt   *time.Time `orm:"column(expired_at);type(datetime);null;description(失效时间)" json:"expired_at"`
	CreatedAt   time.Time  `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
}

func init() {
	orm.RegisterModel(new(MemberPlan))
}

// TableName 设置表名
func (m *MemberPlan) TableName() string {
	return "member_plans"
}

// GetActiveByMember 获取用户当前激活的套餐
func (m *MemberPlan) GetActiveByMember(memberId int) (bool, *MemberPlan) {
	o := orm.NewOrm()
	plan := MemberPlan{}
	err := o.QueryTable(m.TableName()).
		Filter("member_id", memberId).
		Filter("status", "active").
		One(&plan)
	if err != nil {
		return false, &plan
	}
	return true, &plan
}

// GetQueuedByMember 获取用户排队中的套餐（按优先级排序）
func (m *MemberPlan) GetQueuedByMember(memberId int) ([]*MemberPlan, error) {
	o := orm.NewOrm()
	var plans []*MemberPlan
	_, err := o.QueryTable(m.TableName()).
		Filter("member_id", memberId).
		Filter("status", "queued").
		OrderBy("-plan_id").
		All(&plans)
	return plans, err
}

// GetAllByMember 获取用户所有套餐关联记录
func (m *MemberPlan) GetAllByMember(memberId int) ([]*MemberPlan, error) {
	o := orm.NewOrm()
	var plans []*MemberPlan
	_, err := o.QueryTable(m.TableName()).
		Filter("member_id", memberId).
		OrderBy("-created_at").
		All(&plans)
	return plans, err
}

// Create 创建用户套餐关联
func (m *MemberPlan) Create() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// Update 更新用户套餐关联
func (m *MemberPlan) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(m)
	return err
}

// Delete 删除用户套餐关联
func (m *MemberPlan) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(m)
	return err
}

// GetExpiredPlans 获取已过期但状态仍为active的套餐（用于定时任务）
func (m *MemberPlan) GetExpiredPlans() ([]*MemberPlan, error) {
	o := orm.NewOrm()
	var plans []*MemberPlan
	now := time.Now()
	_, err := o.QueryTable(m.TableName()).
		Filter("status", "active").
		Filter("expire_time__lt", now).
		Filter("expire_time__isnull", false).
		All(&plans)
	return plans, err
}

// Activate 激活套餐
func (m *MemberPlan) Activate() error {
	m.Status = "active"
	m.ActivatedAt = time.Now()
	return m.Update()
}

// Expire 使套餐失效
func (m *MemberPlan) Expire() error {
	m.Status = "expired"
	now := time.Now()
	m.ExpiredAt = &now
	return m.Update()
}

// Queue 使套餐进入排队状态
func (m *MemberPlan) Queue() error {
	m.Status = "queued"
	return m.Update()
}