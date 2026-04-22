## 1. 权限系统文档

- [x] 1.1 创建 `权限.md` 文档，包含权限模型说明（User、Role、RoleItem 字段和关系）
- [x] 1.2 添加权限校验流程说明（请求处理流程、Prepare 方法逻辑）
- [x] 1.3 添加当前问题列表和优化建议

## 2. 权限校验测试

- [x] 2.1 创建 `controllers/admin/base_test.go`，测试开发者权限校验（Pid=0）
- [x] 2.2 测试代理商权限校验（Pid>0，基于 RoleItem）
- [x] 2.3 测试未登录用户拦截

## 3. 角色模型测试

- [x] 3.1 创建 `models/role_test.go`，测试 Role 模型 CRUD
- [x] 3.2 创建 `models/roleItem_test.go`，测试 RoleItem 模型 CRUD

## 4. 验证

- [x] 4.1 运行所有测试确保通过
