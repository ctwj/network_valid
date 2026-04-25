## Why

当前会员管理系统已成功运行并提供完整的 API 服务，但缺少 Golang SDK，导致 Go 开发者需要自行实现签名算法、HTTP 请求封装等底层逻辑。创建官方 SDK 可以降低接入成本，提高开发效率，确保签名和加密实现的正确性。

## What Changes

- 新增 `sdk/` 目录，包含完整的 Golang SDK 实现
- 实现 MD5 签名生成机制
- 实现机器码自动获取逻辑（支持用户自定义或使用默认逻辑）
- 封装所有 18 个 API 接口，提供类型安全的 Go 方法
- 提供统一的错误处理和响应解析
- 支持自定义 HTTP Client 配置

## Capabilities

### New Capabilities

- `sdk-client`: SDK 客户端核心功能，包括配置、签名生成、机器码获取、HTTP 请求封装
- `sdk-auth`: 用户认证相关接口（注册、登录、心跳、解绑、下线）
- `sdk-account`: 账户操作接口（充值、扣点、拉黑、查询在线、找回账号）
- `sdk-captcha`: 验证码相关接口（获取验证码、发送验证码邮件）
- `sdk-common`: 基础接口（时间戳、权限、软件信息、会员标签、远程变量）

### Modified Capabilities

无（这是新增 SDK，不修改现有功能）

## Impact

- 新增 `sdk/` 目录，不影响现有代码
- Go 开发者可通过 `import` 方式使用 SDK
- 需要提供使用示例和文档
