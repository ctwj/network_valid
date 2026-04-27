# Go SDK

网络验证系统 Go SDK，提供完整的 API 接口封装。

## 安装

```bash
go get github.com/kerwin/network_valid/sdk/golang
```

## 快速开始

```go
package main

import (
    "fmt"
    "time"

    sdk "github.com/kerwin/network_valid/sdk/golang"
)

func main() {
    // 创建客户端
    client, err := sdk.NewClient(sdk.Config{
        BaseURL:   "https://app.l9.lc",
        AppKey:    "your_appkey",
        SecretKey: "your_secretkey",
        Version:   "1.00",
    })
    if err != nil {
        panic(err)
    }

    // 登录
    userInfo, err := client.Login(sdk.LoginRequest{
        Username: "your_username",
        Password: "your_password",
    })
    if err != nil {
        panic(err)
    }
    fmt.Printf("登录成功! Token: %s\n", userInfo.Client)

    // 启动心跳
    go func() {
        ticker := time.NewTicker(30 * time.Second)
        defer ticker.Stop()
        for range ticker.C {
            if err := client.Heartbeat(); err != nil {
                fmt.Printf("心跳失败: %v\n", err)
                return
            }
        }
    }()

    // 业务逻辑...

    // 退出时下线
    defer client.Logout()
}
```

## 核心 API

### 登录认证

```go
// 登录
userInfo, err := client.Login(sdk.LoginRequest{
    Username: "user@example.com",
    Password: "password123",
})

// 注册
err := client.Register(sdk.RegisterRequest{
    Username: "user@example.com",
    Password: "password123",
    Code:     "VIP12345678",  // 激活码（如需要）
})

// 心跳（保持在线）
err := client.Heartbeat()

// 下线
err := client.Logout()
```

### 配额系统

配额系统支持多维度使用限制，如下载次数、流量、API 调用次数等。

#### 查询配额

```go
// 查询所有配额
quotas, err := client.CheckQuota(nil)

// 查询指定配额
quotas, err := client.CheckQuota([]string{"download", "traffic"})

for _, q := range quotas {
    fmt.Printf("%s: %d/%d (剩余 %d)\n", q.Name, q.Used, q.Limit, q.Remaining)
    fmt.Printf("  周期: %s, 单位: %s\n", q.Period, q.Unit)
    if q.Unlimited {
        fmt.Println("  无限制")
    }
}
```

#### 扣减配额

```go
// 扣减单个配额
result, err := client.DeductQuota("download", 1)
if err != nil {
    panic(err)
}

if result.Success {
    fmt.Printf("扣减成功，剩余: %d\n", result.Remaining)
} else {
    fmt.Printf("扣减失败: %s\n", result.Message)
}
```

#### 批量扣减

```go
// 批量扣减多个配额（事务保证）
results, err := client.DeductQuotaBatch([]sdk.QuotaDeductItem{
    {Key: "download", Amount: 1},
    {Key: "traffic", Amount: 1048576}, // 1MB
})

for _, r := range results {
    if r.Success {
        fmt.Printf("%s: 剩余 %d\n", r.Key, r.Remaining)
    } else {
        fmt.Printf("%s: 失败 - %s\n", r.Key, r.Message)
    }
}
```

### 套餐管理

#### 获取当前套餐

```go
plan, err := client.GetCurrentPlan()
if err != nil {
    panic(err)
}

fmt.Printf("当前套餐: %s\n", plan.PlanName)
fmt.Printf("状态: %s\n", plan.Status)
if plan.ExpireTime != nil {
    fmt.Printf("过期时间: %s\n", *plan.ExpireTime)
} else {
    fmt.Println("永久有效")
}
```

#### 获取排队套餐

```go
plans, err := client.GetQueuedPlans()
for _, p := range plans {
    fmt.Printf("排队套餐: %s, 将在 %s 激活\n", p.PlanName, p.ActivatedAt)
}
```

#### 获取项目套餐列表

```go
plans, err := client.GetPlanList()
for _, p := range plans {
    fmt.Printf("套餐: %v\n", p)
}
```

### 其他功能

```go
// 充值
err := client.Recharge(sdk.RechargeRequest{
    Username: "user@example.com",
    Amount:   30,  // 充值天数
})

// 扣点
err := client.Deduct(sdk.DeductRequest{
    Username: "user@example.com",
    Amount:   10,
})

// 查询在线状态
status, err := client.GetOnlineStatus("user@example.com")

// 获取远程变量
variables, err := client.GetRemoteVariables()
```

## 数据结构

### QuotaInfo 配额信息

| 字段 | 类型 | 说明 |
|------|------|------|
| Key | string | 配额标识 |
| Name | string | 显示名称 |
| Limit | int64 | 限额 |
| Used | int64 | 已用 |
| Remaining | int64 | 剩余 |
| Period | string | 周期：daily/weekly/monthly |
| ResetTime | string | 下次重置时间 |
| Unit | string | 单位：count/bytes/custom |
| Unlimited | bool | 是否无限制 |

### PlanInfo 套餐信息

| 字段 | 类型 | 说明 |
|------|------|------|
| PlanID | int | 套餐类型 ID |
| PlanName | string | 套餐名称 |
| Status | string | 状态：active/expired/queued |
| ExpireTime | *string | 过期时间（NULL 表示永久） |
| Days | float64 | 天数 |
| Priority | int | 优先级 |
| IsFreeTier | int | 是否为免费套餐 |

## 错误处理

```go
userInfo, err := client.Login(req)
if err != nil {
    if apiErr, ok := err.(*sdk.APIError); ok {
        switch apiErr.Errno {
        case 400:
            fmt.Println("用户不存在或密码错误")
        case 10009:
            fmt.Println("版本号不存在")
        case 10011:
            fmt.Println("未绑定登录规则")
        case 10006:
            fmt.Println("项目已停止运营")
        default:
            fmt.Printf("错误: %s (码: %d)\n", apiErr.Errmsg, apiErr.Errno)
        }
    }
    return
}
```

## 配额规则配置

配额规则通过 JSON 配置，详见 [配额规则 JSON Schema](../../docs/配额规则JSON-Schema.md)。

### 示例配置

```json
{
  "quotas": [
    {
      "key": "download",
      "name": "下载次数",
      "limit": 20,
      "period": "daily",
      "unit": "count"
    },
    {
      "key": "traffic",
      "name": "流量",
      "limit": 1073741824,
      "period": "monthly",
      "reset_day": 1,
      "unit": "bytes"
    }
  ]
}
```

## 注意事项

1. **版本号格式**：使用 `X.XX` 格式，如 `1.00` 而非 `1` 或 `1.0`
2. **心跳间隔**：建议每 30 秒发送一次心跳
3. **配额扣减**：扣减前建议先调用 `CheckQuota` 检查剩余配额
4. **批量扣减**：使用 `DeductQuotaBatch` 可保证事务性

## License

MIT
