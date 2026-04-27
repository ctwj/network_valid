## ADDED Requirements

### Requirement: 免费套餐作为保底

系统 SHALL 确保用户始终有一个免费套餐作为保底。

#### Scenario: 免费套餐 expire_time 为 NULL
- **WHEN** 用户获得免费套餐（`is_free_tier=1`）
- **THEN** 用户配额记录 `expire_time=NULL`，表示永不过期

#### Scenario: 免费套餐不受到期检查影响
- **GIVEN** 用户使用免费套餐
- **WHEN** 套餐到期检查定时任务执行
- **THEN** 免费套餐用户被跳过，不触发降级

### Requirement: 套餐优先级

系统 SHALL 为每个套餐类型分配优先级，用于升降级判断。

#### Scenario: 套餐优先级配置
- **WHEN** 管理员创建套餐类型
- **THEN** 可配置 `priority` 字段，数值越大优先级越高

#### Scenario: 免费套餐优先级最低
- **WHEN** 系统创建默认免费套餐
- **THEN** `priority=0`，确保所有付费套餐优先级高于免费套餐

### Requirement: 套餐升级

系统 SHALL 在用户兑换高优先级套餐时自动升级。

#### Scenario: 兑换高优先级套餐立即升级
- **GIVEN** 用户当前使用套餐 A（`priority=0`）
- **WHEN** 用户兑换套餐 B（`priority=10`）
- **THEN** 用户立即升级为套餐 B

#### Scenario: 兑换同优先级套餐叠加时长
- **GIVEN** 用户当前使用套餐 A（`priority=10`），剩余 50 天
- **WHEN** 用户兑换另一个套餐 A（`priority=10`，有效期 120 天）
- **THEN** 用户套餐有效期延长为 170 天

### Requirement: 套餐降级

系统 SHALL 在付费套餐到期时自动降级为免费套餐。

#### Scenario: 付费套餐到期自动降级
- **GIVEN** 用户使用付费套餐，即将到期，无排队套餐
- **WHEN** 套餐到期
- **THEN** 用户自动降级为免费套餐

#### Scenario: 付费套餐到期有排队套餐则激活
- **GIVEN** 用户使用套餐 A，套餐 B 在排队中
- **WHEN** 套餐 A 到期
- **THEN** 套餐 B 自动激活

### Requirement: 套餐排队

系统 SHALL 支持低优先级套餐排队等待激活。

#### Scenario: 兑换低优先级套餐进入排队
- **GIVEN** 用户使用套餐 A（`priority=20`），剩余 100 天
- **WHEN** 用户兑换套餐 B（`priority=10`）
- **THEN** 套餐 B 进入排队状态，等待套餐 A 到期

#### Scenario: 查看排队套餐
- **WHEN** 用户查询当前套餐状态
- **THEN** 系统返回当前激活套餐和排队中的套餐列表

### Requirement: 套餐到期检查定时任务

系统 SHALL 定时检查即将到期的套餐并处理降级。

#### Scenario: 定时任务检查到期套餐
- **WHEN** 套餐到期检查定时任务执行
- **THEN** 系统查询所有已到期的活跃套餐，执行降级逻辑