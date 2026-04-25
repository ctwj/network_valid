## ADDED Requirements

### Requirement: 获取验证码接口

SDK SHALL 提供获取图形验证码方法。

使用 application/x-www-form-urlencoded 格式提交。

#### Scenario: 获取验证码
- **WHEN** 调用验证码接口
- **THEN** 返回验证码图片或标识

### Requirement: 注册验证码邮件接口

SDK SHALL 提供发送注册验证码邮件方法，支持参数：
- email: 注册邮箱地址（必填）

使用 application/x-www-form-urlencoded 格式提交。

#### Scenario: 发送注册验证码
- **WHEN** 传入邮箱地址
- **THEN** 向该邮箱发送注册验证码

### Requirement: 找回账号验证码邮件接口

SDK SHALL 提供发送找回账号验证码邮件方法，支持参数：
- email: 注册邮箱地址（必填）

使用 application/x-www-form-urlencoded 格式提交。

#### Scenario: 发送找回验证码
- **WHEN** 传入邮箱地址
- **THEN** 向该邮箱发送找回账号验证码
