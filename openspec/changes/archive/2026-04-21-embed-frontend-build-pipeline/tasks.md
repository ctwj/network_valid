## 1. 配置文件模板化

- [x] 1.1 将 `backend/config.conf` 重命名为 `backend/config.conf.example`
- [x] 1.2 更新 `.gitignore` 添加 `backend/config.conf`
- [x] 1.3 更新 `.gitignore` 添加 `backend/static/`

## 2. Go embed 静态文件集成

- [x] 2.1 在 `backend/main.go` 添加 `embed` 包导入和 `//go:embed static` 指令
- [x] 2.2 创建静态文件服务 handler，从 embed.FS 读取文件
- [x] 2.3 实现 SPA fallback 路由，未匹配路由返回 `index.html`
- [x] 2.4 配置 Beego 路由，确保 API 路由优先于静态文件路由

## 3. GitHub Action 构建工作流

- [x] 3.1 创建 `.github/workflows/build.yml` 工作流文件
- [x] 3.2 配置 `workflow_dispatch` 触发器，添加版本号输入参数
- [x] 3.3 添加前端构建步骤（安装依赖、构建、复制到 static 目录）
- [x] 3.4 添加 Go 环境设置步骤
- [x] 3.5 添加跨平台编译步骤（Linux amd64, Windows amd64）
- [x] 3.6 添加 GitHub Release 创建和产物上传步骤

## 4. 文档更新

- [x] 4.1 更新 README 添加配置文件设置说明
- [x] 4.2 更新 README 添加构建和部署说明
