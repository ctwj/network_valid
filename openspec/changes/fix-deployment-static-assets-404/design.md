## Context

当前系统使用 Go 的 `embed` 包将前端静态文件嵌入到二进制文件中。`main.go` 中的 `staticFileHandler` 函数负责服务这些静态文件并处理 SPA 路由回退。

**问题分析**：
1. 当请求 `/assets/index.b4b296db.css` 时，`staticFiles.Open("static" + path)` 正确打开了 `static/assets/index.b4b296db.css`
2. 但随后使用 `http.FileServer(http.FS(staticFiles))` 时，它从文件系统根目录开始服务，而不是从 `static/` 子目录
3. 这导致文件路径不匹配，返回 404
4. 由于文件未正确返回，MIME 类型也未能正确设置

**当前代码问题**：
```go
http.FileServer(http.FS(staticFiles)).ServeHTTP(ctx.ResponseWriter, ctx.Request)
```
这行代码创建了一个从根目录开始的文件服务器，但请求路径是 `/assets/...`，而文件实际在 `static/assets/...`。

## Goals / Non-Goals

**Goals:**
- 修复静态文件服务，确保所有静态资源能正确返回
- 确保静态文件返回正确的 MIME 类型
- 保持 SPA 路由回退功能不变
- 保持 API 路由不受影响

**Non-Goals:**
- 不修改前端构建配置
- 不修改 Nginx 配置（问题在应用层，不在反向代理层）
- 不添加新的静态文件处理功能

## Decisions

### 决策 1：使用 `http.StripPrefix` 修复路径问题

**选择**：使用 `http.StripPrefix` 配合 `http.FS` 来正确处理嵌入文件系统的路径。

**理由**：
- `http.StripPrefix` 可以移除请求路径的前缀，使文件服务器能正确匹配嵌入文件系统中的路径
- 这是 Go 标准库推荐的处理嵌入文件系统的方式
- 替代方案（如手动读取文件并设置 Content-Type）需要更多代码且容易出错

**实现方式**：
```go
// 创建一个从 static/ 目录服务的文件服务器
fs := http.FileServer(http.FS(staticFiles))
http.StripPrefix("/", fs).ServeHTTP(ctx.ResponseWriter, ctx.Request)
```

但更好的方式是使用 `fs.FS` 子目录：
```go
subFS, _ := fs.Sub(staticFiles, "static")
http.FileServer(http.FS(subFS)).ServeHTTP(ctx.ResponseWriter, ctx.Request)
```

### 决策 2：保持 SPA 回退逻辑

**选择**：在文件未找到时返回 `index.html`，保持现有逻辑不变。

**理由**：
- Vue Router 使用 history 模式，需要服务器返回 `index.html` 来处理前端路由
- 这是 SPA 应用的标准做法

## Risks / Trade-offs

**风险 1**：修改可能影响其他静态文件路径
- **缓解**：在本地测试所有静态文件路径，包括 CSS、JS、图片、字体等

**风险 2**：MIME 类型可能仍然不正确
- **缓解**：Go 的 `http.FileServer` 会自动根据文件扩展名设置 MIME 类型，使用正确的文件服务器实现即可解决

**权衡**：使用 `fs.Sub` 创建子文件系统比手动处理路径更简洁，但需要导入 `io/fs` 包
