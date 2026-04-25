package sdk

import "fmt"

// APIError API 错误
type APIError struct {
	Errno     int    // 错误码
	Errmsg    string // 错误信息
	UID       string // 请求ID
	Timestamp int64  // 时间戳
}

// Error 实现 error 接口
func (e *APIError) Error() string {
	return fmt.Sprintf("API error %d: %s (uid: %s)", e.Errno, e.Errmsg, e.UID)
}

// IsSuccess 判断是否成功
func (e *APIError) IsSuccess() bool {
	return e.Errno == 0
}

// 错误码常量
const (
	ErrSuccess         = 0    // 请求成功
	ErrBadRequest      = 400  // 请求参数错误
	ErrUnauthorized    = 401  // 未授权/签名验证失败
	ErrForbidden       = 403  // 权限不足
	ErrNotFound        = 404  // 资源不存在
	ErrInternalServer  = 500  // 服务器内部错误

	// 业务错误码
	ErrUserNotFound     = 1001 // 用户不存在
	ErrPasswordWrong    = 1002 // 密码错误
	ErrAccountDisabled  = 1003 // 账号已被禁用
	ErrAccountBanned    = 1004 // 账号已被拉黑
	ErrEmailUsed        = 1005 // 邮箱已被使用
	ErrCaptchaWrong     = 1006 // 验证码错误
	ErrCaptchaExpired   = 1007 // 验证码已过期
	ErrUserOnline       = 1008 // 用户已在线
	ErrBalanceInsufficient = 1009 // 余额不足
	ErrDeviceLimitExceeded = 1010 // 设备绑定数量超限
)

// ErrMsg 错误码对应的错误信息
var ErrMsg = map[int]string{
	ErrSuccess:         "请求成功",
	ErrBadRequest:      "请求参数错误",
	ErrUnauthorized:    "未授权/签名验证失败",
	ErrForbidden:       "权限不足",
	ErrNotFound:        "资源不存在",
	ErrInternalServer:  "服务器内部错误",
	ErrUserNotFound:     "用户不存在",
	ErrPasswordWrong:    "密码错误",
	ErrAccountDisabled:  "账号已被禁用",
	ErrAccountBanned:    "账号已被拉黑",
	ErrEmailUsed:        "邮箱已被使用",
	ErrCaptchaWrong:     "验证码错误",
	ErrCaptchaExpired:   "验证码已过期",
	ErrUserOnline:       "用户已在线",
	ErrBalanceInsufficient: "余额不足",
	ErrDeviceLimitExceeded: "设备绑定数量超限",
}

// GetErrMsg 获取错误码对应的错误信息
func GetErrMsg(errno int) string {
	if msg, ok := ErrMsg[errno]; ok {
		return msg
	}
	return "未知错误"
}

// NewAPIError 创建 API 错误
func NewAPIError(errno int, errmsg, uid string, timestamp int64) *APIError {
	return &APIError{
		Errno:     errno,
		Errmsg:    errmsg,
		UID:       uid,
		Timestamp: timestamp,
	}
}
