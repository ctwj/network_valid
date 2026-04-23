## Why

部署后前端静态资源（CSS、JS、manifest.webmanifest）返回 404 错误，CSS 文件 MIME 类型错误（返回 `text/plain` 而非 `text/css`），导致前端页面无法正常加载。

**根本原因**：`main.go` 中的 `staticFileHandler` 函数在服务静态文件时存在路径问题。当使用 `http.FileServer(http.FS(staticFiles))` 时，它尝试从嵌入文件系统的根目录服务文件，而不是从 `static/` 子目录服务，导致文件无法找到。

## What Changes

- 修复 `staticFileHandler` 函数中静态文件服务的路径问题
- 确保所有静态资源（CSS、JS、图片、字体、manifest.webmanifest）能正确返回
- 确保静态文件返回正确的 MIME 类型（`text/css`、`application/javascript` 等）
- 保持 SPA 路由回退到 `index.html` 的功能不变

## Capabilities

### New Capabilities

无新增能力。

### Modified Capabilities

- `embed-static-files`: 修复静态文件服务逻辑，确保嵌入的静态资源能正确返回并带有正确的 MIME 类型。

## Impact

- **代码变更**：`backend/main.go` 中的 `staticFileHandler` 函数
- **API 影响**：无 API 变更
- **依赖影响**：无依赖变更
- **系统影响**：修复后前端页面将能正常加载静态资源
