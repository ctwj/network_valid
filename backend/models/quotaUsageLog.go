package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
)

// QuotaUsageLog 配额使用日志
type QuotaUsageLog struct {
	ID        int64     `orm:"column(id);pk;auto" json:"id"`
	MemberId  int       `orm:"column(member_id);index;description(用户ID)" json:"member_id"`
	QuotaKey  string    `orm:"column(quota_key);size(64);index;description(配额标识)" json:"quota_key"`
	Amount    int64     `orm:"column(amount);description(扣减数量)" json:"amount"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime);index" json:"created_at"`
}

func init() {
	orm.RegisterModel(new(QuotaUsageLog))
}

// TableName 设置表名
func (q *QuotaUsageLog) TableName() string {
	return "quota_usage_logs"
}

// Create 创建使用日志
func (q *QuotaUsageLog) Create() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(q)
}

// GetByMember 获取用户使用日志
func (q *QuotaUsageLog) GetByMember(memberId int, limit int) ([]*QuotaUsageLog, error) {
	o := orm.NewOrm()
	var logs []*QuotaUsageLog
	_, err := o.QueryTable(q.TableName()).
		Filter("member_id", memberId).
		OrderBy("-created_at").
		Limit(limit).
		All(&logs)
	return logs, err
}

// GetByMemberAndKey 获取用户指定配额的使用日志
func (q *QuotaUsageLog) GetByMemberAndKey(memberId int, quotaKey string, limit int) ([]*QuotaUsageLog, error) {
	o := orm.NewOrm()
	var logs []*QuotaUsageLog
	_, err := o.QueryTable(q.TableName()).
		Filter("member_id", memberId).
		Filter("quota_key", quotaKey).
		OrderBy("-created_at").
		Limit(limit).
		All(&logs)
	return logs, err
}
