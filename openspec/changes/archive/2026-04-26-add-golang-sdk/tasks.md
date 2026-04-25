## 1. SDK 基础结构

- [x] 1.1 创建 sdk/ 目录结构
- [x] 1.2 实现 types.go - 定义请求/响应类型结构体
- [x] 1.3 实现 errors.go - 定义 APIError 和错误码常量
- [x] 1.4 实现 sign.go - 实现 MD5 签名生成函数
- [x] 1.5 实现 machinecode.go - 实现默认机器码获取逻辑（基于 MAC 地址 + 主机名）

## 2. 核心客户端

- [x] 2.1 实现 client.go - 定义 Config 结构体和 Client 结构体
- [x] 2.2 实现 NewClient 构造函数（支持 MachineCode 可选，不传则自动获取）
- [x] 2.3 实现公共参数自动填充方法
- [x] 2.4 实现 HTTP 请求封装方法（支持 multipart 和 urlencoded）
- [x] 2.5 实现统一响应解析方法

## 3. 基础接口实现 (common.go)

- [x] 3.1 实现 GetTimestamp - 获取服务器时间戳
- [x] 3.2 实现 GetPermissions - 获取用户权限列表
- [x] 3.3 实现 GetSoftwareInfo - 获取软件信息
- [x] 3.4 实现 GetMemberTags - 获取会员标签
- [x] 3.5 实现 GetRemoteVariables - 获取远程变量

## 4. 验证码接口实现 (captcha.go)

- [x] 4.1 实现 GetCaptcha - 获取图形验证码
- [x] 4.2 实现 SendRegisterCode - 发送注册验证码邮件
- [x] 4.3 实现 SendRecoverCode - 发送找回账号验证码邮件

## 5. 用户认证接口实现 (auth.go)

- [x] 5.1 实现 Register - 用户注册
- [x] 5.2 实现 Login - 用户登录
- [x] 5.3 实现 Heartbeat - 用户心跳
- [x] 5.4 实现 Unbind - 用户解绑
- [x] 5.5 实现 Logout - 用户下线

## 6. 账户操作接口实现 (account.go)

- [x] 6.1 实现 Recharge - 用户充值
- [x] 6.2 实现 Deduct - 账号扣点
- [x] 6.3 实现 Ban - 账号拉黑
- [x] 6.4 实现 IsOnline - 查询用户在线状态
- [x] 6.5 实现 Recover - 找回账号

## 7. 测试与文档

- [x] 7.1 编写 example_test.go - 使用示例
- [x] 7.2 编写 README.md - SDK 使用文档
