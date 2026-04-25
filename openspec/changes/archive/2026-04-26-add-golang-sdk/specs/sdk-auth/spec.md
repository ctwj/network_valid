## ADDED Requirements

### Requirement: 用户注册接口

SDK SHALL 提供用户注册方法，支持参数：
- username: 用户名/邮箱（必填）
- password: 密码（必填）
- code: 邮件验证码（可选）
- captcha: 验证码密码（可选）

使用 multipart/form-data 格式提交。

#### Scenario: 普通注册
- **WHEN** 传入用户名和密码
- **THEN** 调用注册接口并返回结果

#### Scenario: 验证码注册
- **WHEN** 传入用户名、密码、验证码
- **THEN** 调用注册接口并返回结果

### Requirement: 用户登录接口

SDK SHALL 提供用户登录方法，支持参数：
- username: 用户名/邮箱（必填）
- password: 密码（必填）

使用 multipart/form-data 格式提交。

#### Scenario: 登录成功
- **WHEN** 传入正确的用户名和密码
- **THEN** 返回登录成功响应

#### Scenario: 登录失败
- **WHEN** 传入错误的用户名或密码
- **THEN** 返回错误码 1001 或 1002

### Requirement: 用户心跳接口

SDK SHALL 提供用户心跳方法，保持用户在线状态。

#### Scenario: 发送心跳
- **WHEN** 调用心跳接口
- **THEN** 返回心跳状态

### Requirement: 用户解绑接口

SDK SHALL 提供用户解绑方法，解绑设备绑定。

#### Scenario: 解绑设备
- **WHEN** 传入用户名
- **THEN** 解绑该用户的设备绑定

### Requirement: 用户下线接口

SDK SHALL 提供强制用户下线方法。

#### Scenario: 强制下线
- **WHEN** 传入用户名
- **THEN** 强制该用户下线
