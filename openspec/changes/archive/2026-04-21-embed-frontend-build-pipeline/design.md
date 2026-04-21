## Context

当前项目是一个前后端分离的 Web 应用：
- **前端**: Vue 3 + Vite + Naive UI，构建产物输出到 `frontend/dist`
- **后端**: Go + Beego 框架，静态文件存放在 `backend/static`
- **配置**: 使用 `backend/config.conf` 存储数据库、缓存等配置

当前部署需要分别部署前端静态文件和后端服务，增加了运维复杂度。

## Goals / Non-Goals

**Goals:**
- 实现单一二进制文件部署，前端静态文件嵌入后端可执行文件
- 支持 GitHub Action 手动触发构建，可指定版本号
- 支持跨平台编译（Linux/Windows）
- 保护敏感配置信息不被提交到版本控制

**Non-Goals:**
- 不改变现有的前端构建配置
- 不改变后端 API 结构
- 不实现热更新或动态加载静态文件

## Decisions

### 1. 使用 Go embed 包嵌入静态文件

**选择**: 使用 `embed` 包
**原因**:
- Go 1.16+ 内置支持，无需外部依赖
- 编译时嵌入，运行时无额外 IO 开销
- 支持嵌入整个目录

**实现方式**:
```go
//go:embed static
var staticFiles embed.FS
```

### 2. 静态文件服务路由

**选择**: 使用 Beego 的 `SetStaticPath` 配合 `embed.FS`
**原因**:
- 与现有 Beego 框架集成
- 支持前端 SPA 路由（所有未匹配路由返回 index.html）

**实现方式**:
- 根路径 `/` 指向嵌入的静态文件
- API 路由 `/admin/*` 和 `/api/*` 保持不变
- 未匹配路由返回 `index.html` 支持 Vue Router history 模式

### 3. GitHub Action 工作流设计

**选择**: 使用 `workflow_dispatch` 事件触发
**原因**:
- 支持手动触发
- 支持输入参数（版本号）
- 可配置构建目标平台

**工作流步骤**:
1. 检出代码
2. 设置 Node.js 环境
3. 安装前端依赖并构建
4. 复制前端构建产物到后端 static 目录
5. 设置 Go 环境
6. 跨平台编译后端
7. 创建 Release 并上传产物

### 4. 配置文件管理

**选择**: 使用 `.example` 后缀作为模板
**原因**:
- 常见的配置管理模式
- 用户复制模板并修改
- 便于文档说明

## Risks / Trade-offs

- **二进制文件体积增大** → 前端静态文件通常 1-5MB，对现代服务器影响可忽略
- **更新前端需要重新编译后端** → 通过 CI/CD 自动化构建解决
- **embed 不支持热更新** → 符合单一二进制部署的目标，非问题
