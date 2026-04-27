## 1. 数据库变更

- [x] 1.1 `cards` 表新增 `quota_rules`（JSON）、`is_free_tier`（TINYINT）、`priority`（INT）字段
- [x] 1.2 创建 `user_quotas` 表（member_id, quota_key, limit_value, used_value, period, reset_time, expire_time, plan_id）
- [x] 1.3 创建 `member_plans` 表（member_id, plan_id, status, expire_time, activated_at, expired_at）
- [x] 1.4 创建 `quota_usage_logs` 表记录使用历史（可选，用于统计分析）
- [x] 1.5 添加系统配置项：`PlanPresets`（JSON，预设套餐方案）

## 2. 后端模型层

- [x] 2.1 创建 `QuotaRule` 结构体（解析 JSON 配额规则，含 key/name/limit/period/reset_day/unit）
- [x] 2.2 创建 `UserQuota` 模型（`backend/models/userQuota.go`）
- [x] 2.3 创建 `MemberPlan` 模型（`backend/models/memberPlan.go`）
- [x] 2.4 `Cards` 模型新增 `QuotaRules`、`IsFreeTier`、`Priority` 字段
- [x] 2.5 实现 `UserQuota` 的 CRUD 方法
- [x] 2.6 实现 `MemberPlan` 的 CRUD 方法
- [x] 2.7 实现配额重置时间计算方法（根据 period 和 reset_day）

## 3. 后端服务层 - 配额服务

- [x] 3.1 创建 `QuotaService`（`backend/services/quotaService.go`）
- [x] 3.2 实现配额规则解析方法 `ParseQuotaRules(json)`
- [x] 3.3 实现配额初始化方法 `InitQuotas(memberId, quotaRules, expireTime, planId)`
- [x] 3.4 实现配额检查方法 `CheckQuotas(memberId, keys)` - 返回指定或所有配额状态
- [x] 3.5 实现配额扣减方法 `DeductQuota(memberId, key, amount)` - 原子操作，防超扣
- [x] 3.6 实现批量扣减方法 `DeductQuotaBatch(memberId, items)` - 事务保证
- [x] 3.7 实现懒重置逻辑 - 检测 reset_time 过期自动重置
- [x] 3.8 实现配额使用记录方法 `LogQuotaUsage(memberId, key, amount)`

## 4. 后端服务层 - 套餐服务

- [x] 4.1 创建 `PlanService`（`backend/services/planService.go`）
- [x] 4.2 实现免费套餐发放方法 `GrantFreeTier(memberId)`
- [x] 4.3 实现套餐兑换方法 `RedeemPlan(memberId, planId)` - 含升级/排队逻辑
- [x] 4.4 实现套餐升级方法 `UpgradePlan(memberId, newPlan)`
- [x] 4.5 实现套餐降级方法 `DowngradeToFreeTier(memberId)`
- [x] 4.6 实现套餐排队方法 `QueuePlan(memberId, planId, afterExpireTime)`
- [x] 4.7 实现获取当前套餐方法 `GetCurrentPlan(memberId)`
- [x] 4.8 实现获取排队套餐方法 `GetQueuedPlans(memberId)`

## 5. 后端服务层 - 预设方案

- [x] 5.1 创建 `PlanPresetService`（`backend/services/planPresetService.go`）
- [x] 5.2 实现获取预设方案方法 `GetPresets()`
- [x] 5.3 实现从预设创建套餐方法 `CreateFromPreset(projectId, presetNames)`
- [x] 5.4 实现预设方案配置方法 `SavePresets(presets)`

## 6. 后端 API 层

- [x] 6.1 新增 `quota.check` 接口 - 查询用户配额（支持 keys 参数）
- [x] 6.2 新增 `quota.deduct` 接口 - 扣减单个配额
- [x] 6.3 新增 `quota.deductBatch` 接口 - 批量扣减配额
- [x] 6.4 修改 `user.register` 接口 - 支持免费套餐发放
- [x] 6.5 修改 `user.recharge` 接口 - 兑换码兑换时调用套餐服务
- [x] 6.6 新增 `plan.list` 接口 - 获取项目套餐列表
- [x] 6.7 新增 `plan.current` 接口 - 获取用户当前套餐
- [x] 6.8 新增 `plan.queued` 接口 - 获取用户排队套餐

## 7. 定时任务

- [x] 7.1 创建配额重置定时任务（每小时执行）
- [x] 7.2 创建套餐到期检查定时任务（每小时执行）
- [x] 7.3 实现批量重置逻辑（更新 reset_time 和 used_value）
- [x] 7.4 实现过期配额清理逻辑（删除 expire_time 已过的记录）

## 8. SDK 更新

- [x] 8.1 新增 `QuotaInfo` 结构体（Key, Name, Limit, Used, Remaining, Period, ResetTime, Unit）
- [x] 8.2 新增 `QuotaDeductItem` 结构体（Key, Amount）
- [x] 8.3 新增 `CheckQuota(keys []string)` 方法 - 查询配额
- [x] 8.4 新增 `DeductQuota(key string, amount int64)` 方法 - 扣减单个配额
- [x] 8.5 新增 `DeductQuotaBatch(items []QuotaDeductItem)` 方法 - 批量扣减
- [x] 8.6 新增 `GetCurrentPlan()` 方法 - 获取当前套餐信息

## 9. SDK 测试

- [x] 9.1 编写 `CheckQuota` 单元测试
- [x] 9.2 编写 `DeductQuota` 单元测试
- [x] 9.3 编写 `DeductQuotaBatch` 单元测试
- [x] 9.4 编写套餐升降级集成测试
- [x] 9.5 编写多维度配额集成测试

## 10. 管理后台（前端 Vue.js）

- [x] 10.1 项目创建向导 - 强制用户登录模式，禁用单码模式
- [x] 10.2 项目创建时自动创建默认套餐（免费套餐 + 标准套餐）
- [ ] 10.3 套餐类型管理页面 - 配额规则可视化编辑器
- [ ] 10.4 套餐类型管理页面 - JSON 高级编辑模式
- [ ] 10.5 套餐类型管理页面 - 优先级配置
- [ ] 10.6 系统配置页面 - 预设套餐方案配置
- [ ] 10.7 用户详情页 - 当前套餐和配额使用情况
- [ ] 10.8 用户详情页 - 排队套餐展示
- [x] 10.9 UI 文案更新：激活码类型 → 套餐类型；激活码 → 兑换码；激活码管理 → 兑换码管理

## 11. 文档更新

- [x] 11.1 更新 `docs/项目配置指南.md` - 新增多维度配额系统说明
- [x] 11.2 编写配额 JSON Schema 文档
- [x] 11.3 编写客户端配额交互流程文档
- [x] 11.4 编写套餐升降级逻辑文档
- [x] 11.5 更新 SDK README - 新增配额接口使用示例
- [x] 11.6 编写迁移指南 - 现有项目如何启用配额系统