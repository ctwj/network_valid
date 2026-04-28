## 1. 后端 - 默认登录规则创建

- [x] 1.1 在 `backend/models/project.go` 的 `Add` 方法中，项目创建后自动创建默认登录规则
- [x] 1.2 默认登录规则字段配置：Title="默认规则"，Mode=1，RegMode=1，UnbindMode=3
- [x] 1.3 将创建的登录规则 ID 自动绑定到新项目的 LoginType 字段

## 2. 后端 - 默认版本号创建

- [x] 2.1 在 `backend/models/project.go` 的 `Add` 方法中，项目创建后自动创建初始版本号
- [x] 2.2 默认版本号字段配置：Version=1.00，IsActive=0，IsMustUpdate=0

## 3. 后端 - 套餐方案动态定价

- [x] 3.1 在 `backend/models/project.go` 中新增 `monthlyPrice` 参数到 `Add` 方法
- [x] 3.2 修改 `createPlansFromScheme` 方法，支持基于月费动态计算套餐金额
- [x] 3.3 在 `backend/controllers/admin/project.go` 的 `CreateProject` 方法中接收 `monthly_price` 参数

## 4. 前端 - 套餐方案 UI 增强

- [x] 4.1 在 `frontend/src/views/project/comp/projectEdit.vue` 套餐方案 Tab 中新增"单月费用"输入框
- [x] 4.2 实现单月费用变更时套餐金额自动计算逻辑
- [x] 4.3 显示计算后的套餐价格，保留两位小数

## 5. 前端 - 配置提示优化

- [x] 5.1 在运营模式 Tab 的 Alert 中增加 SDK 调用影响说明
- [x] 5.2 在加密模式 Tab 的 Alert 中增加 SDK 密钥配置说明
- [x] 5.3 在签名算法 Tab 的 Alert 中增加 SDK 签名算法使用说明

## 6. 测试验证

- [x] 6.1 测试创建项目后自动生成登录规则并绑定
- [x] 6.2 测试创建项目后自动生成初始版本号
- [x] 6.3 测试修改单月费用后套餐金额正确计算
- [x] 6.4 测试前端提示信息正确显示