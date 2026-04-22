## Why

当前系统的权限管理缺乏文档说明，开发者(Pid=0)与代理商(Pid>0)的权限区别不清晰，导致维护困难。权限核心逻辑缺少测试覆盖，存在潜在风险。需要建立完整的权限系统文档，并为核心权限校验逻辑添加单元测试。

## What Changes

- 新增 `权限.md` 文档，记录权限系统的设计说明、实现细节和优化建议
- 为 `controllers/admin/base.go` 中的权限校验逻辑添加单元测试
- 为 `models/role.go` 和 `models/roleItem.go` 添加单元测试

## Capabilities

### New Capabilities

- `permission-system`: 权限系统文档，包含权限模型说明、校验流程、角色定义、优化建议

### Modified Capabilities

无需求变更，仅添加文档和测试

## Impact

- **文档**: 新增 `权限.md` 权限系统说明文档
- **测试**: 新增 `controllers/admin/base_test.go`、`models/role_test.go`、`models/roleItem_test.go`
- **代码**: 无功能代码变更
