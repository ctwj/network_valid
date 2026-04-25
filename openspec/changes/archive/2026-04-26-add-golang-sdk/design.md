## Context

会员管理系统提供基于 HTTP 的 REST API，包含用户认证、权限管理、账户操作等功能。API 采用 MD5 签名验证机制，所有请求需要携带公共参数（appkey、timestamp、sign、version、mac）。

当前项目使用 Golang，需要创建一个独立的 SDK 模块，方便 Go 开发者快速接入系统。

## Goals / Non-Goals

**Goals:**

- 提供类型安全的 Go SDK，封装所有 18 个 API 接口
- 实现正确的 MD5 签名生成逻辑
- 实现机器码自动获取逻辑（默认获取，支持用户覆盖）
- 支持自定义 HTTP Client（超时、代理等）
- 提供清晰的错误处理机制
- 支持 multipart/form-data 和 application/x-www-form-urlencoded 两种 Content-Type

**Non-Goals:**

- 不实现 RSA 签名验证（服务端响应验证由使用者自行处理）
- 不提供异步/并发封装（保持同步简单模型）
- 不实现连接池管理（由 HTTP Client 自行配置）

## Decisions

### 1. SDK 目录结构

采用扁平化结构，按功能模块分文件：

```
sdk/
├── client.go       # 核心客户端和配置
├── sign.go         # 签名生成
├── machinecode.go  # 机器码获取
├── common.go       # 基础接口（时间戳、权限等）
├── auth.go         # 用户认证接口
├── account.go      # 账户操作接口
├── captcha.go      # 验证码接口
├── types.go        # 请求/响应类型定义
├── errors.go       # 错误定义
└── example_test.go # 使用示例
```

**理由**: 按功能分文件便于维护，同时避免过度嵌套。

### 2. 机器码获取设计

SDK 提供默认机器码获取逻辑，同时支持用户自定义：

```go
// 默认机器码获取：基于 MAC 地址 + 主机名生成
func DefaultMachineCode() string

// 配置时可选传递机器码
client, err := sdk.NewClient(sdk.Config{
    BaseURL:   "https://api.example.com",
    AppKey:    "your_appkey",
    SecretKey: "your_secretkey",
    Version:   "1.0.0",
    // MachineCode 不传则自动调用 DefaultMachineCode()
})
```

**默认机器码生成算法**：
1. 获取本机第一个非回环网络接口的 MAC 地址
2. 获取主机名
3. 拼接后进行 MD5 哈希，取前 16 位

**理由**: 用户可选择便捷（自动获取）或灵活（自定义）两种方式。

### 3. 客户端设计

采用配置模式创建客户端：

```go
client, err := sdk.NewClient(sdk.Config{
    BaseURL:    "https://api.example.com",
    AppKey:     "your_appkey",
    SecretKey:  "your_secretkey",
    Version:    "1.0.0",
    MachineCode: "client_machine_code",
    HTTPClient: &http.Client{Timeout: 30 * time.Second}, // 可选
})
```

**理由**: 配置结构体清晰，支持可选参数，便于扩展。

### 4. 错误处理

定义业务错误码映射：

```go
type APIError struct {
    Errno     int    // 错误码
    Errmsg    string // 错误信息
    UID       string // 请求ID
    Timestamp int64  // 时间戳
}
```

**理由**: 保留完整错误上下文，便于调试和日志记录。

### 5. 响应解析

统一响应结构：

```go
type Response struct {
    Errno     int             `json:"errno"`
    Data      json.RawMessage `json:"data"` // 延迟解析
    Errmsg    string          `json:"errmsg"`
    UID       string          `json:"uid"`
    Timestamp int64           `json:"timestamp"`
    Sign      string          `json:"sign"`
    Signal    string          `json:"signal"`
}
```

**理由**: Data 使用 RawMessage 延迟解析，适配不同接口返回不同数据类型。

## Risks / Trade-offs

| 风险 | 缓解措施 |
|------|----------|
| 时间戳不同步导致签名失败 | SDK 自动获取服务器时间戳，提供时间戳同步方法 |
| API 变更导致 SDK 不兼容 | 版本化 SDK，提供变更日志和迁移指南 |
| 网络超时处理 | 默认 30 秒超时，支持自定义 HTTP Client |
