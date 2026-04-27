## Why

当前系统缺少免费试用机制，无法有效区分不同套餐等级的价值差异。现有套餐类型（原"激活码类型"）仅通过"天数/点数"区分，无法实现按使用量限制的功能。这导致：
1. 新用户无法免费体验产品，增加了获客门槛
2. 不同价位套餐缺乏功能差异化，影响付费转化
3. 无法实现"免费试用 + 付费升级"的商业模式
4. 缺少客户端与服务端的配额同步交互机制

## What Changes

- **概念重命名**：激活码类型 → 套餐类型；激活码 → 兑换码
- **多维度配额系统**：套餐类型支持配置多个配额限制（JSON 格式，可选）
- **多周期支持**：配额支持每日/每周/每月重置
- **永久套餐**：支持时长为永久的套餐（`days=0`）
- **预设方案**：项目创建时提供至少两套预设方案（免费套餐 + 至少一套付费套餐），管理员可自定义命名和时长
- **免费套餐保底**：用户注册时自动获得免费套餐，付费套餐到期后自动降级
- **配额交互接口**：SDK 新增配额检查、扣减接口

## Capabilities

### New Capabilities

- `plan-quota`: 套餐配额配置与管理，支持多维度、多周期限制（JSON 配置，可选）
- `free-tier-registration`: 免费套餐自动发放，作为用户保底套餐
- `quota-tracking`: 用户配额追踪与扣减，支持客户端主动上报
- `plan-lifecycle`: 套餐生命周期管理，支持升降级和到期自动降级
- `plan-presets`: 预设套餐方案，项目创建时快速配置

### Modified Capabilities

- `plan-type`: 套餐类型模型新增 `quota_rules`（JSON）、`is_free_tier`、`priority` 字段，支持永久套餐
- `user-registration`: 注册逻辑修改，支持自动发放免费套餐
- `redeem-code`: 兑换码管理（原激活码），用于兑换套餐

## Impact

- **数据库**：
  - `cards` 表新增 `quota_rules`、`is_free_tier`、`priority` 字段
  - 新增 `user_quotas` 表存储用户配额状态
  - 新增 `member_plans` 表存储用户套餐关联
- **后端 API**：新增 `quota.check`、`quota.deduct`、`quota.deductBatch` 接口
- **SDK**：新增 `CheckQuota()`、`DeductQuota()`、`DeductQuotaBatch()` 方法
- **定时任务**：新增配额重置任务、套餐到期检查任务
- **管理后台**：项目创建向导（选择预设方案）、套餐类型管理页面、配额配置界面