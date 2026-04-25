## ADDED Requirements

### Requirement: 时间戳接口

SDK SHALL 提供获取服务器时间戳方法。

使用 application/x-www-form-urlencoded 格式提交。

#### Scenario: 获取时间戳
- **WHEN** 调用时间戳接口
- **THEN** 返回服务器当前时间戳（秒级）

### Requirement: 获取权限接口

SDK SHALL 提供获取用户权限列表方法。

使用 multipart/form-data 格式提交。

#### Scenario: 获取权限成功
- **WHEN** 用户已登录
- **THEN** 返回用户权限列表（树形结构）

#### Scenario: 获取权限失败
- **WHEN** 用户未登录或无权限
- **THEN** 返回错误码 400

### Requirement: 获取软件信息接口

SDK SHALL 提供获取软件版本和配置信息方法。

使用 multipart/form-data 格式提交。

#### Scenario: 获取软件信息
- **WHEN** 调用软件信息接口
- **THEN** 返回软件版本和配置信息

### Requirement: 获取会员标签接口

SDK SHALL 提供获取会员标签列表方法。

使用 multipart/form-data 格式提交。

#### Scenario: 获取会员标签
- **WHEN** 调用会员标签接口
- **THEN** 返回会员标签列表

### Requirement: 获取远程变量接口

SDK SHALL 提供获取服务器端远程变量方法。

使用 multipart/form-data 格式提交。

#### Scenario: 获取远程变量
- **WHEN** 调用远程变量接口
- **THEN** 返回服务器配置的远程变量
