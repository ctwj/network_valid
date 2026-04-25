## ADDED Requirements

### Requirement: 用户充值接口

SDK SHALL 提供用户充值方法，支持参数：
- username: 用户名/邮箱（必填）
- amount: 充值金额/天数（必填）

使用 multipart/form-data 格式提交。

#### Scenario: 充值成功
- **WHEN** 传入用户名和充值金额
- **THEN** 为用户账户充值

### Requirement: 账号扣点接口

SDK SHALL 提供账号扣点方法，支持参数：
- username: 用户名/邮箱（必填）
- amount: 扣除点数（必填）

使用 multipart/form-data 格式提交。

#### Scenario: 扣点成功
- **WHEN** 传入用户名和扣除点数
- **THEN** 从用户账户扣除点数

#### Scenario: 余额不足
- **WHEN** 用户余额小于扣除点数
- **THEN** 返回错误码 1009

### Requirement: 账号拉黑接口

SDK SHALL 提供账号拉黑方法，支持参数：
- username: 用户名/邮箱（必填）
- reason: 拉黑原因（可选）

使用 multipart/form-data 格式提交。

#### Scenario: 拉黑账号
- **WHEN** 传入用户名和拉黑原因
- **THEN** 将账号加入黑名单

### Requirement: 查询用户在线接口

SDK SHALL 提供查询用户在线状态方法。

#### Scenario: 查询在线状态
- **WHEN** 传入用户名
- **THEN** 返回用户当前在线状态

### Requirement: 找回账号接口

SDK SHALL 提供找回账号方法，支持参数：
- email: 注册邮箱地址（必填）
- code: 邮件验证码（必填）
- captcha: 验证码密码（必填）

使用 multipart/form-data 格式提交。

#### Scenario: 找回账号
- **WHEN** 传入邮箱和验证码
- **THEN** 返回账号信息或重置密码
