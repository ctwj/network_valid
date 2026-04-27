package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

// UserQuota 用户配额状态
type UserQuota struct {
	ID         int64     `orm:"column(id);pk;auto" json:"id"`
	MemberId   int       `orm:"column(member_id);index;description(用户ID)" json:"member_id"`
	QuotaKey   string    `orm:"column(quota_key);size(64);description(配额标识)" json:"quota_key"`
	LimitValue int64     `orm:"column(limit_value);default(0);description(当前限额)" json:"limit_value"`
	UsedValue  int64     `orm:"column(used_value);default(0);description(已使用量)" json:"used_value"`
	Period     string    `orm:"column(period);size(16);description(周期:daily/weekly/monthly)" json:"period"`
	ResetDay   int       `orm:"column(reset_day);default(1);description(重置日)" json:"reset_day"`
	ResetTime  time.Time `orm:"column(reset_time);type(datetime);description(下次重置时间)" json:"reset_time"`
	ExpireTime *time.Time `orm:"column(expire_time);type(datetime);null;description(过期时间,NULL表示永不过期)" json:"expire_time"`
	PlanId     int       `orm:"column(plan_id);index;description(关联套餐类型ID)" json:"plan_id"`
	CreatedAt  time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt  time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
}

func init() {
	orm.RegisterModel(new(UserQuota))
}

// TableName 设置表名
func (u *UserQuota) TableName() string {
	return "user_quotas"
}

// GetByMemberAndKey 根据用户ID和配额标识获取配额
func (u *UserQuota) GetByMemberAndKey(memberId int, quotaKey string) (bool, *UserQuota) {
	o := orm.NewOrm()
	quota := UserQuota{MemberId: memberId, QuotaKey: quotaKey}
	err := o.Read(&quota, "MemberId", "QuotaKey")
	if err != nil {
		return false, &quota
	}
	return true, &quota
}

// GetByMember 获取用户所有配额
func (u *UserQuota) GetByMember(memberId int) ([]*UserQuota, error) {
	o := orm.NewOrm()
	var quotas []*UserQuota
	_, err := o.QueryTable(u.TableName()).Filter("member_id", memberId).All(&quotas)
	return quotas, err
}

// Create 创建配额记录
func (u *UserQuota) Create() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(u)
}

// Update 更新配额记录
func (u *UserQuota) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(u)
	return err
}

// Delete 删除配额记录
func (u *UserQuota) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(u)
	return err
}

// DeleteByMember 删除用户所有配额
func (u *UserQuota) DeleteByMember(memberId int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(u.TableName()).Filter("member_id", memberId).Delete()
	return err
}

// DeleteByMemberAndPlan 删除用户指定套餐的配额
func (u *UserQuota) DeleteByMemberAndPlan(memberId int, planId int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(u.TableName()).Filter("member_id", memberId).Filter("plan_id", planId).Delete()
	return err
}
