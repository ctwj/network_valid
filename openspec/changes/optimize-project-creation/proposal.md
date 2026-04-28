## Why

当前项目创建流程需要用户手动配置多个关联项（登录规则、版本号、套餐方案），导致创建后无法立即对接 SDK 使用，增加了配置复杂度和出错风险。用户需要"一步创建即可对接 SDK"的简化体验。

## What Changes

- **运营模式/加密模式/签名算法提示优化**：在创建项目 UI 中提供更友好的提示说明，帮助用户理解各选项对系统和 SDK 的影响
- **套餐方案增强**：
  - 在套餐方案外新增"单月费用"输入框（默认 30 元）
  - 3 组套餐方案的金额基于单月费用自动计算
  - 用户修改单月费用后，套餐内金额同步变更
- **默认登录规则**：项目创建时自动创建一条登录规则
  - 绑定模式：普通登录（mode=1）
  - 注册模式：普通注册（reg_mode=1）
  - 解绑模式：任意解绑（unbind_mode=3）
- **默认版本号管理**：项目创建时自动创建一条版本号记录，初始版本 1.00

## Capabilities

### New Capabilities
- `project-creation-defaults`: 项目创建时自动生成默认配置（登录规则、版本号），实现一步创建即可对接 SDK

### Modified Capabilities
- `sdk-client`: SDK 需适配默认登录规则（普通登录、普通注册、任意解绑）的行为说明
- `sdk-auth`: SDK 需适配默认登录规则下的认证流程说明

## Impact

- **后端**：`backend/models/project.go` - Add 方法需扩展，创建项目时同步创建登录规则和版本号
- **后端**：`backend/models/projectLogin.go` - 需支持默认规则创建
- **后端**：`backend/models/project.go` - 套餐方案金额计算逻辑需支持动态月费基准
- **前端**：`frontend/src/views/project/comp/projectEdit.vue` - 套餐方案 UI 需新增单月费用输入框
- **SDK**：文档需说明默认登录规则下的对接方式