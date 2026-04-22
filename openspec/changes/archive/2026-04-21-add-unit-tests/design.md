## Context

当前项目使用 Go + Beego 框架，仅有一个示例测试文件 `tests/default_test.go`，使用 GoConvey 测试框架。核心业务逻辑未经测试验证。

## Goals / Non-Goals

**Goals:**
- 为纯函数添加单元测试（无外部依赖）
- 为验证逻辑添加单元测试
- 为核心业务逻辑添加单元测试
- 运行测试发现并修复 bug

**Non-Goals:**
- 不添加集成测试（需要数据库/缓存环境）
- 不添加 API 端到端测试
- 不追求 100% 覆盖率

## Decisions

### 1. 测试框架选择

**选择**: 使用 `testify` + Go 标准库 `testing`
**原因**:
- testify 提供丰富的断言和 mock 功能
- 与现有 GoConvey 兼容
- Go 社区广泛使用

### 2. 测试优先级

**优先级排序**:
1. **纯函数测试**：无依赖，易于测试，快速验证
   - `models/common.go`: RandStr, In, GetRsaKey
   - `models/pager.go`: PageCount
   - `controllers/common/common.go`: GetToken, GetStringMd5, VerifyEmailFormat, AesEncrypt/AesDecrypt

2. **验证逻辑测试**：输入验证，边界条件
   - `validation/api/valid.go`: 邮箱、用户名、密码格式验证

3. **业务逻辑测试**：需要 mock 数据库
   - `models/keys.go`: 激活码生成逻辑
   - `models/member.go`: 会员验证逻辑
   - `models/manager.go`: 管理员登录逻辑

### 3. 测试文件结构

```
backend/
  models/
    common_test.go
    pager_test.go
    keys_test.go
    member_test.go
    manager_test.go
  controllers/
    common/
      common_test.go
  validation/
    api/
      valid_test.go
```

## Risks / Trade-offs

- **数据库依赖测试复杂** → 优先测试纯函数，业务逻辑使用 mock
- **测试发现 bug** → 记录并修复，更新任务列表
- **测试覆盖不完整** → 后续迭代补充
