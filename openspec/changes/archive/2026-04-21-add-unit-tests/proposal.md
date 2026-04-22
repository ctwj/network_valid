## Why

当前项目缺少单元测试覆盖，仅有一个示例测试文件。核心业务逻辑（激活码生成、会员认证、加密函数等）未经测试验证，存在潜在 bug 风险。添加单元测试可以验证系统正确性，发现并修复潜在问题。

## What Changes

- 添加纯函数单元测试（models/common.go, models/pager.go, controllers/common/common.go）
- 添加验证逻辑测试（validation/api/valid.go）
- 添加业务逻辑测试（models/keys.go, models/member.go, models/manager.go）
- 运行测试验证，发现并修复 bug
- 添加测试依赖（testify）

## Capabilities

### New Capabilities

- `unit-testing`: 单元测试规范，定义测试覆盖要求和测试策略

### Modified Capabilities

- 无

## Impact

- **后端代码**: 新增测试文件（`*_test.go`）
- **依赖**: 添加 `github.com/stretchr/testify` 测试库
- **可能修复**: 测试过程中发现的 bug
