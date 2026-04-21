## Why

当前项目需要单独部署前端和后端，增加了部署复杂度。通过将前端静态文件 embed 到 Go 二进制文件中，可以实现单一可执行文件部署，简化分发和运维流程。同时需要完善 CI/CD 流水线支持自动化构建和版本管理，并保护敏感配置信息不被提交到版本控制。

## What Changes

- 使用 Go 1.16+ 的 `embed` 功能将前端构建产物嵌入到后端二进制文件中
- 后端静态文件目录（`backend/static`）添加到 `.gitignore`
- 创建 GitHub Action 工作流，支持手动触发构建并指定版本号
- 配置文件（`backend/config.conf`）添加到 `.gitignore`
- 原配置文件重命名为 `config.conf.example` 作为模板

## Capabilities

### New Capabilities

- `embed-static-files`: 将前端构建产物嵌入 Go 二进制文件，实现单一可执行文件部署
- `github-actions-build`: GitHub Action 工作流，支持手动触发构建、指定版本号、跨平台编译
- `config-template`: 配置文件模板机制，保护敏感配置信息

### Modified Capabilities

- 无

## Impact

- **后端代码**: `main.go` 需要添加 embed 指令和静态文件服务路由
- **构建流程**: 需要先构建前端，再构建后端
- **部署方式**: 从前后端分离部署变为单一二进制文件部署
- **版本控制**: `.gitignore` 需要更新，配置文件需要重命名
- **CI/CD**: 新增 GitHub Action 工作流文件
