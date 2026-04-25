package sdk

import "encoding/json"

// Config SDK 客户端配置
type Config struct {
	BaseURL     string // API 服务器地址（必填）
	AppKey      string // 应用标识（必填）
	SecretKey   string // 应用密钥（必填）
	Version     string // 软件版本号（必填）
	MachineCode string // 客户端机器码（可选，不传则自动获取）
	HTTPClient  HTTPClient // 自定义 HTTP Client（可选）
}

// HTTPClient HTTP 客户端接口
type HTTPClient interface {
	Do(req interface{}) (resp interface{}, err error)
}

// Response 统一响应结构
type Response struct {
	Errno     int             `json:"errno"`
	Data      json.RawMessage `json:"data"`
	Errmsg    string          `json:"errmsg"`
	UID       string          `json:"uid"`
	Timestamp int64           `json:"timestamp"`
	Sign      string          `json:"sign"`
	Signal    string          `json:"signal"`
}

// Permission 权限项
type Permission struct {
	ID          int          `json:"id"`
	PID         int          `json:"pid"`
	ProjectID   int          `json:"project_id"`
	Path        string       `json:"path"`
	Name        string       `json:"name"`
	Free        int          `json:"free"`
	Show        int          `json:"show"`
	Description string       `json:"description"`
	Children    []Permission `json:"children,omitempty"`
}

// SoftwareInfo 软件信息
type SoftwareInfo struct {
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// MemberTag 会员标签
type MemberTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// RemoteVariable 远程变量
type RemoteVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// UserInfo 用户信息
type UserInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Balance  int    `json:"balance"`
	VIPLevel int    `json:"vip_level"`
}

// OnlineStatus 用户在线状态
type OnlineStatus struct {
	Username string `json:"username"`
	Online   bool   `json:"online"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string // 用户名/邮箱（必填）
	Password string // 密码（必填）
	Code     string // 邮件验证码（可选）
	Captcha  string // 验证码密码（可选）
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string // 用户名/邮箱（必填）
	Password string // 密码（必填）
}

// RechargeRequest 充值请求
type RechargeRequest struct {
	Username string // 用户名/邮箱（必填）
	Amount   int    // 充值金额/天数（必填）
}

// DeductRequest 扣点请求
type DeductRequest struct {
	Username string // 用户名/邮箱（必填）
	Amount   int    // 扣除点数（必填）
}

// BanRequest 拉黑请求
type BanRequest struct {
	Username string // 用户名/邮箱（必填）
	Reason   string // 拉黑原因（可选）
}

// RecoverRequest 找回账号请求
type RecoverRequest struct {
	Email    string // 注册邮箱地址（必填）
	Code     string // 邵件验证码（必填）
	Captcha  string // 验证码密码（必填）
}

// SendCodeRequest 发送验证码请求
type SendCodeRequest struct {
	Email string // 邮箱地址（必填）
}