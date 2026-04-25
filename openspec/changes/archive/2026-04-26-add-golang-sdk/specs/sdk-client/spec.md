## ADDED Requirements

### Requirement: SDK 客户端配置

SDK SHALL 提供客户端配置结构体，包含以下字段：
- BaseURL: API 服务器地址（必填）
- AppKey: 应用标识（必填）
- SecretKey: 应用密钥（必填）
- Version: 软件版本号（必填）
- MachineCode: 客户端机器码（可选，不传则自动获取）

#### Scenario: 创建默认客户端
- **WHEN** 使用必填参数创建客户端，不传递 MachineCode
- **THEN** 客户端成功初始化，自动获取机器码，使用默认 HTTP Client（30秒超时）

#### Scenario: 创建自定义机器码的客户端
- **WHEN** 传入自定义 MachineCode
- **THEN** 客户端使用传入的机器码

#### Scenario: 创建自定义 HTTP Client 的客户端
- **WHEN** 传入自定义 HTTPClient 配置
- **THEN** 客户端使用自定义的 HTTP Client

### Requirement: 机器码自动获取

SDK SHALL 提供默认机器码获取功能：
- 基于本机 MAC 地址和主机名生成
- 生成算法：MD5(MAC地址 + 主机名) 前 16 位
- 用户可调用 DefaultMachineCode() 函数获取默认机器码

#### Scenario: 自动获取机器码
- **WHEN** 创建客户端时未传递 MachineCode
- **THEN** SDK 自动调用 DefaultMachineCode() 获取机器码

#### Scenario: 手动获取机器码
- **WHEN** 用户调用 DefaultMachineCode() 函数
- **THEN** 返回基于本机信息生成的机器码字符串

### Requirement: MD5 签名生成

SDK SHALL 正确生成 MD5 签名，签名算法为：
```
sign = MD5(appkey + secretkey + version + timestamp + mac)
```

签名结果为 32 位小写字符串。

#### Scenario: 生成有效签名
- **WHEN** 调用签名生成函数
- **THEN** 返回正确的 32 位小写 MD5 字符串

### Requirement: 公共参数自动填充

SDK SHALL 自动为每个请求填充公共参数：
- appkey
- timestamp（秒级时间戳）
- sign（MD5签名）
- version
- mac

#### Scenario: 请求自动携带公共参数
- **WHEN** 调用任意 API 方法
- **THEN** 请求自动包含所有公共参数

### Requirement: 统一响应解析

SDK SHALL 解析统一响应格式，包含字段：
- errno: 状态码
- data: 返回数据
- errmsg: 提示信息
- uid: 请求标识
- timestamp: 服务器时间戳
- sign: 服务器签名
- signal: RSA签名

#### Scenario: 成功响应解析
- **WHEN** 服务器返回 errno=0
- **THEN** 解析 data 字段并返回

#### Scenario: 错误响应处理
- **WHEN** 服务器返回 errno!=0
- **THEN** 返回包含错误码和错误信息的 APIError
