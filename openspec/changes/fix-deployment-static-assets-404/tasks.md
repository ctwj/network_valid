## 1. 修复静态文件服务

- [x] 1.1 修改 `backend/main.go` 中的 `staticFileHandler` 函数，使用 `fs.Sub` 创建子文件系统
- [x] 1.2 确保静态文件从 `static/` 目录正确服务
- [x] 1.3 确保 MIME 类型正确设置（CSS 返回 `text/css`，JS 返回 `application/javascript`）

## 2. 测试验证

- [x] 2.1 本地构建并测试静态文件服务
- [x] 2.2 验证 CSS 文件能正确加载
- [x] 2.3 验证 JS 文件能正确加载
- [x] 2.4 验证 SPA 路由回退功能正常
- [x] 2.5 验证 API 路由不受影响

## 3. 部署验证

- [ ] 3.1 推送代码到远程仓库
- [ ] 3.2 等待 GitHub Action 构建完成
- [ ] 3.3 验证部署后前端页面正常加载