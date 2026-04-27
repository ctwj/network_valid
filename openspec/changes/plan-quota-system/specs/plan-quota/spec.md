## ADDED Requirements

### Requirement: 套餐类型支持多维度配额配置

系统 SHALL 允许管理员为套餐类型配置多个配额限制（JSON 格式，可选）。

#### Scenario: 创建带配额的套餐类型
- **WHEN** 管理员创建套餐类型，配置配额规则：
  ```json
  {
    "quotas": [
      {"key": "download", "name": "下载次数", "limit": 20, "period": "daily", "unit": "count"}
    ]
  }
  ```
- **THEN** 系统保存套餐类型，`quota_rules` 字段存储该 JSON

#### Scenario: 配额规则为空表示无限制
- **WHEN** 管理员创建套餐类型，未配置配额规则
- **THEN** 系统 `quota_rules` 字段为 NULL，表示该套餐无配额限制

#### Scenario: 配额规则可选
- **GIVEN** 管理员创建套餐类型
- **WHEN** 管理员不配置配额规则
- **THEN** 套餐类型创建成功，用户使用不受配额限制

### Requirement: 配额支持多种周期

系统 SHALL 支持每日、每周、每月三种配额重置周期。

#### Scenario: 每日配额自动重置
- **GIVEN** 用户配额 `period=daily`
- **WHEN** 次日 00:00 到达
- **THEN** 用户该配额的 `used_value` 重置为 0

#### Scenario: 每周配额按指定日重置
- **GIVEN** 用户配额 `period=weekly`，`reset_day=1`（周一）
- **WHEN** 下周一 00:00 到达
- **THEN** 用户该配额的 `used_value` 重置为 0

#### Scenario: 每月配额按指定日期重置
- **GIVEN** 用户配额 `period=monthly`，`reset_day=1`（每月1号）
- **WHEN** 下月1号 00:00 到达
- **THEN** 用户该配额的 `used_value` 重置为 0

### Requirement: 套餐优先级与升降级

系统 SHALL 支持套餐优先级，实现自动升降级。

#### Scenario: 购买高优先级套餐自动升级
- **GIVEN** 用户当前使用套餐 A（`priority=0`）
- **WHEN** 用户兑换套餐 B（`priority=10`）
- **THEN** 用户立即升级为套餐 B，配额按套餐 B 规则更新

#### Scenario: 付费套餐到期自动降级
- **GIVEN** 用户使用付费套餐，即将到期，无排队套餐
- **WHEN** 套餐到期
- **THEN** 用户自动降级为免费套餐

#### Scenario: 套餐排队等待激活
- **GIVEN** 用户使用套餐 A（`priority=20`），剩余 100 天
- **WHEN** 用户兑换套餐 B（`priority=10`）
- **THEN** 套餐 B 进入排队状态，等待套餐 A 到期后激活

### Requirement: 兑换码兑换时同步配额

系统 SHALL 在用户使用兑换码兑换套餐时，根据套餐的配额规则创建用户配额记录。

#### Scenario: 使用带配额的兑换码兑换
- **GIVEN** 套餐类型配置了配额规则
- **WHEN** 用户使用该套餐类型的兑换码兑换
- **THEN** 系统根据配额规则创建用户配额记录

#### Scenario: 使用无配额限制的兑换码兑换
- **GIVEN** 套餐类型 `quota_rules=NULL`
- **WHEN** 用户使用该套餐类型的兑换码兑换
- **THEN** 用户获得对应时长，不创建配额记录，使用不受限

### Requirement: 管理后台配额配置界面

系统 SHALL 在套餐类型管理页面提供配额规则配置界面。

#### Scenario: 可视化配置配额规则
- **WHEN** 管理员编辑套餐类型
- **THEN** 页面显示配额规则配置区域，支持添加/删除配额项

#### Scenario: JSON 高级编辑模式
- **WHEN** 管理员切换到高级模式
- **THEN** 页面显示 JSON 编辑器，支持直接编辑配额规则 JSON