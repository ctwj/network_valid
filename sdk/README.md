# 会员管理系统 SDK

本目录包含会员管理系统的各语言 SDK 实现。

## 目录结构

```
sdk/
├── golang/           # Golang SDK
│   ├── client.go     # 核心客户端
│   ├── types.go      # 类型定义
│   ├── errors.go     # 错误处理
│   ├── sign.go       # 签名生成
│   ├── machinecode.go # 机器码获取
│   ├── common.go     # 基础接口
│   ├── auth.go       # 用户认证
│   ├── account.go    # 账户操作
│   ├── captcha.go    # 验证码接口
│   ├── go.mod        # Go 模块定义
│   └── example_test.go # 使用示例
└── README.md         # 本文件
```

---

## Golang SDK

### 安装

```bash
go get github.com/kerwin/network_valid/sdk/golang
```

### 快速开始

#### 创建客户端

```go
package main

import (
    "fmt"
    gosdk "github.com/kerwin/network_valid/sdk/golang"
)

func main() {
    // 方式1: 自动获取机器码
    client, err := gosdk.NewClient(gosdk.Config{
        BaseURL:   "https://api.example.com",
        AppKey:    "your_appkey",
        SecretKey: "your_secretkey",
        Version:   "1.0.0",
        // MachineCode 不传，SDK 自动获取
    })
    if err != nil {
        panic(err)
    }
    
    // 方式2: 自定义机器码
    client, err = gosdk.NewClient(gosdk.Config{
        BaseURL:     "https://api.example.com",
        AppKey:      "your_appkey",
        SecretKey:   "your_secretkey",
        Version:     "1.0.0",
        MachineCode: "my_custom_code",
    })
}
```

### API 接口

#### 用户认证

```go
// 用户登录
userInfo, err := client.Login(gosdk.LoginRequest{
    Username: "user@example.com",
    Password: "password123",
})

// 用户注册
err := client.Register(gosdk.RegisterRequest{
    Username: "newuser@example.com",
    Password: "password123",
    Code:     "123456",      // 邮件验证码（可选）
    Captcha:  "abc123",      // 图形验证码（可选）
})

// 用户心跳（建议每 30 秒调用一次）
err := client.Heartbeat("user@example.com")

// 用户解绑
err := client.Unbind("user@example.com")

// 用户下线
err := client.Logout("user@example.com")
```

#### 账户操作

```go
// 用户充值
err := client.Recharge(gosdk.RechargeRequest{
    Username: "user@example.com",
    Amount:   100,
})

// 账号扣点
err := client.Deduct(gosdk.DeductRequest{
    Username: "user@example.com",
    Amount:   10,
})

// 账号拉黑
err := client.Ban(gosdk.BanRequest{
    Username: "user@example.com",
    Reason:   "违规操作",
})

// 查询在线状态
status, err := client.IsOnline("user@example.com")

// 找回账号
err := client.Recover(gosdk.RecoverRequest{
    Email:    "user@example.com",
    Code:     "123456",
    Captcha:  "abc123",
})
```

#### 验证码

```go
// 获取图形验证码
captchaData, err := client.GetCaptcha()

// 发送注册验证码
err := client.SendRegisterCode("user@example.com")

// 发送找回账号验证码
err := client.SendRecoverCode("user@example.com")
```

#### 基础接口

```go
// 获取服务器时间戳
timestamp, err := client.GetTimestamp()

// 获取用户权限
permissions, err := client.GetPermissions()

// 获取软件信息
info, err := client.GetSoftwareInfo()

// 获取会员标签
tags, err := client.GetMemberTags()

// 获取远程变量
variables, err := client.GetRemoteVariables()
value, err := client.GetRemoteVariable("some_key")
```

### 错误处理

```go
userInfo, err := client.Login(gosdk.LoginRequest{
    Username: "user@example.com",
    Password: "wrong_password",
})
if err != nil {
    if apiErr, ok := err.(*gosdk.APIError); ok {
        fmt.Printf("错误码: %d\n", apiErr.Errno)
        fmt.Printf("错误信息: %s\n", apiErr.Errmsg)
        fmt.Printf("请求ID: %s\n", apiErr.UID)
    }
    return
}
```

#### 常见错误码

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 400 | 请求参数错误 |
| 401 | 签名验证失败 |
| 1001 | 用户不存在 |
| 1002 | 密码错误 |
| 1003 | 账号已被禁用 |
| 1004 | 账号已被拉黑 |
| 1009 | 余额不足 |

### 机器码

SDK 提供默认机器码获取功能：

```go
// 获取默认机器码
machineCode := gosdk.DefaultMachineCode()
fmt.Printf("机器码: %s\n", machineCode)
```

**默认机器码生成算法**：
1. 获取本机第一个非回环网络接口的 MAC 地址
2. 获取主机名
3. 拼接后进行 MD5 哈希，取前 16 位

### 自定义 HTTP Client

```go
import (
    "net/http"
    "time"
)

client, err := gosdk.NewClient(gosdk.Config{
    BaseURL:   "https://api.example.com",
    AppKey:    "your_appkey",
    SecretKey: "your_secretkey",
    Version:   "1.0.0",
    HTTPClient: &http.Client{
        Timeout: 60 * time.Second,
    },
})
```

---

## 其他语言 SDK

如需其他语言 SDK，请参考 API 对接文档自行实现，或联系管理员获取支持。

## License

MIT
