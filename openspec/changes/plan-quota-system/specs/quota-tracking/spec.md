## ADDED Requirements

### Requirement: 配额检查接口

系统 SHALL 提供配额检查接口，供客户端查询当前配额状态。

#### Scenario: 查询所有配额
- **WHEN** 客户端调用 `quota.check`（不传 keys 参数）
- **THEN** 系统返回用户所有配额列表，每项包含：
  - `key`: 配额标识
  - `name`: 配额名称
  - `limit`: 限额
  - `used`: 已用
  - `remaining`: 剩余
  - `period`: 周期
  - `reset_time`: 下次重置时间
  - `unit`: 单位

#### Scenario: 查询指定配额
- **WHEN** 客户端调用 `quota.check`，传入 `keys=["download", "traffic"]`
- **THEN** 系统返回指定的配额信息

#### Scenario: 无配额限制时返回特殊标识
- **GIVEN** 用户套餐无配额限制（`quota_rules=NULL`）
- **WHEN** 客户端查询配额
- **THEN** 系统返回 `unlimited=true`

### Requirement: 配额扣减接口

系统 SHALL 提供配额扣减接口，供客户端上报使用量。

#### Scenario: 正常扣减配额
- **GIVEN** 用户 `download` 配额已用 3 次，限额 5 次
- **WHEN** 客户端调用 `quota.deduct`，`key="download"`，`amount=1`
- **THEN** 用户该配额已用变为 4 次，返回成功

#### Scenario: 配额不足时拒绝扣减
- **GIVEN** 用户 `download` 配额已用 5 次，限额 5 次
- **WHEN** 客户端调用 `quota.deduct`，`key="download"`，`amount=1`
- **THEN** 系统返回错误"配额不足"，已用次数不变

#### Scenario: 并发扣减不超扣
- **GIVEN** 用户 `download` 配额已用 4 次，限额 5 次
- **WHEN** 同时发起 3 个扣减请求
- **THEN** 只有 1 个请求成功（已用变为 5），其余 2 个返回"配额不足"

#### Scenario: 流量类配额扣减
- **GIVEN** 用户 `traffic` 配额已用 500MB，限额 1GB
- **WHEN** 客户端调用 `quota.deduct`，`key="traffic"`，`amount=1048576`（1MB）
- **THEN** 用户该配额已用变为 501MB

### Requirement: 批量扣减接口

系统 SHALL 提供批量扣减接口，用于一次操作涉及多个配额的场景。

#### Scenario: 批量扣减多个配额
- **GIVEN** 用户有 `download` 和 `traffic` 两个配额
- **WHEN** 客户端调用 `quota.deductBatch`：
  ```json
  {"items": [{"key": "download", "amount": 1}, {"key": "traffic", "amount": 1048576}]}
  ```
- **THEN** 两个配额同时扣减成功，返回各自结果

#### Scenario: 批量扣减部分失败
- **GIVEN** 用户 `download` 配额已用完，`traffic` 配额充足
- **WHEN** 客户端调用 `quota.deductBatch` 同时扣减两个配额
- **THEN** 整体返回失败，两个配额都不扣减（事务回滚）

### Requirement: 客户端交互流程

系统 SHALL 定义清晰的客户端与服务端配额交互流程。

#### Scenario: 下载文件流程
- **WHEN** 客户端需要下载文件
- **THEN** 推荐流程：
  1. 调用 `quota.check` 检查流量是否足够
  2. 执行下载操作
  3. 下载完成后调用 `quota.deduct` 上报实际使用量

#### Scenario: API 调用流程
- **WHEN** 客户端需要调用受限 API
- **THEN** 推荐流程：
  1. 调用 `quota.check` 检查配额
  2. 调用 `quota.deduct` 预扣配额
  3. 执行 API 调用
  4. 若调用失败，可忽略（配额已扣）或请求回滚

### Requirement: SDK 配额接口

SDK SHALL 提供配额检查和扣减方法。

#### Scenario: SDK 检查配额
- **WHEN** 调用 `client.CheckQuota(keys []string)`
- **THEN** SDK 返回 `map[string]QuotaInfo`，key 为配额标识

#### Scenario: SDK 扣减单个配额
- **WHEN** 调用 `client.DeductQuota(key string, amount int64)`
- **THEN** SDK 返回扣减结果（成功/失败，剩余量）

#### Scenario: SDK 批量扣减
- **WHEN** 调用 `client.DeductQuotaBatch(items []QuotaDeductItem)`
- **THEN** SDK 返回批量扣减结果

### Requirement: 配额使用记录

系统 SHALL 记录配额使用历史，便于统计分析。

#### Scenario: 记录每次扣减
- **WHEN** 用户扣减配额成功
- **THEN** 系统记录：用户ID、配额标识、扣减时间、扣减数量

#### Scenario: 查询使用历史
- **WHEN** 管理员查询用户配额使用历史
- **THEN** 系统返回该用户近 30 天的配额使用记录