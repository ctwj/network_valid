# 软中心 (Soft-Center)

一个前后端一体化的管理系统，前端使用 Vue 3 + Vite + Naive UI，后端使用 Go + Beego。

## 快速开始

### 1. 配置文件设置

首次运行前，需要创建配置文件：

```bash
cd backend
cp config.conf.example config.conf
```

编辑 `config.conf` 文件，配置数据库和缓存信息：

```ini
[app]
cache    = file    # 缓存类型: file 或 redis
key      = your-secret-key

[sql]
type     = sqlite  # 数据库类型: sqlite 或 mysql
ip       = 127.0.0.1
port     = 3306
user     = your_user
pwd      = your_password
db       = your_database
rebuild  = false
```

### 2. 本地开发

#### 前端开发

```bash
cd frontend
pnpm install
pnpm dev
```

#### 后端开发

```bash
cd backend
go mod tidy
go run main.go
```

### 3. 构建生产版本

#### 手动构建

```bash
# 1. 构建前端
cd frontend
pnpm install
pnpm build

# 2. 复制前端构建产物到后端
cp -r dist ../backend/static

# 3. 构建后端
cd ../backend
go build -o soft-center .
```

#### GitHub Actions 自动构建

1. 进入 GitHub 仓库的 Actions 页面
2. 选择 "Build and Release" 工作流
3. 点击 "Run workflow"
4. 输入版本号（如 `v1.0.0`）
5. 点击 "Run workflow"

构建完成后，会在 Releases 页面生成对应版本的二进制文件：
- `soft-center-linux-amd64` - Linux 版本
- `soft-center-windows-amd64.exe` - Windows 版本

## 部署

### 单文件部署

构建后的二进制文件已包含前端静态资源，只需：

1. 下载对应平台的二进制文件
2. 在同目录创建 `config.conf` 配置文件
3. 运行程序

```bash
./soft-center
```

默认监听端口为 9960，访问 `http://localhost:9960` 即可。

### 默认账号

首次运行会自动创建管理员账号：
- 用户名：`admin`
- 密码：`112233`

**请在首次登录后立即修改密码！**

## 目录结构

```
.
├── backend/                 # 后端代码
│   ├── config.conf.example  # 配置文件模板
│   ├── main.go              # 入口文件
│   ├── controllers/         # 控制器
│   ├── models/              # 数据模型
│   ├── routers/             # 路由配置
│   └── static/              # 前端构建产物（不纳入版本控制）
├── frontend/                # 前端代码
│   ├── src/                 # 源代码
│   └── dist/                # 构建产物
└── .github/
    └── workflows/
        └── build.yml        # GitHub Actions 构建配置
```

## 技术栈

### 前端
- Vue 3
- Vite
- Naive UI
- TypeScript

### 后端
- Go 1.21+
- Beego 2.0
- SQLite / MySQL

## License

MIT
