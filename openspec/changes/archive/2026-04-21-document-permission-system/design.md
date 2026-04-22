## Context

当前权限系统基于 Beego 框架实现，核心逻辑位于 `controllers/admin/base.go` 的 `Prepare()` 方法中。权限模型包含：
- **User 模型**: Pid 字段区分开发者(Pid=0)和代理商(Pid>0)
- **Role 模型**: 定义角色名称和关联权限项
- **RoleItem 模型**: 细粒度权限控制，包含 Path、Index、Value 字段

当前问题：
1. 权限标签过于简单（仅"开发者"/"代理商"），无法体现角色差异
2. RoleList 在代码中硬编码，扩展性差
3. 密码明文存储，存在安全风险
4. 无权限继承机制

## Goals / Non-Goals

**Goals:**
- 创建完整的权限系统文档，说明设计原理和实现细节
- 记录当前存在的问题和优化建议
- 为核心权限校验逻辑添加单元测试

**Non-Goals:**
- 不修改现有权限功能代码
- 不实现优化建议（仅记录供后续参考）
- 不重构权限模型

## Decisions

1. **文档位置**: 在项目根目录创建 `权限.md`，便于快速查阅
2. **测试范围**: 覆盖 `base.go` 的权限校验逻辑、`role.go` 和 `roleItem.go` 的 CRUD 操作
3. **测试框架**: 使用 testify/assert，与现有测试保持一致

## Risks / Trade-offs

- 测试可能发现现有 bug → 记录并修复
- 文档可能随代码演进过时 → 建议在 CLAUDE.md 中添加文档更新提醒
