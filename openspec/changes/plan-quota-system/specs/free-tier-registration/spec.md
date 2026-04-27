## ADDED Requirements

### Requirement: 用户注册自动获得免费套餐

系统 SHALL 在用户注册时自动发放免费套餐（无需兑换码）。

#### Scenario: 启用免费套餐时注册成功
- **GIVEN** 系统配置 `FreeTierEnabled=true`，`FreeTierQuota=5`，`FreeTierDays=30`
- **WHEN** 用户注册账号（不提供兑换码）
- **THEN** 用户账号创建成功，自动获得 30 天有效期，每日配额 5 次

#### Scenario: 禁用免费套餐时注册失败
- **GIVEN** 系统配置 `FreeTierEnabled=false`，登录规则 `RegMode=0`（需要兑换码）
- **WHEN** 用户尝试注册账号（不提供兑换码）
- **THEN** 系统返回错误"注册需要兑换码"

#### Scenario: 注册时提供兑换码覆盖免费套餐
- **GIVEN** 系统配置 `FreeTierEnabled=true`
- **WHEN** 用户注册账号并提供兑换码（白银套餐，配额 20，天数 120）
- **THEN** 用户获得白银套餐（配额 20，天数 120），不使用免费套餐

### Requirement: 免费套餐参数可配置

系统 SHALL 提供管理配置项控制免费套餐参数。

#### Scenario: 配置免费套餐配额
- **WHEN** 管理员设置系统配置 `FreeTierQuota=10`
- **THEN** 新注册用户获得每日 10 次配额

#### Scenario: 配置免费套餐有效期
- **WHEN** 管理员设置系统配置 `FreeTierDays=60`
- **THEN** 新注册用户获得 60 天有效期

### Requirement: 免费套餐到期处理

系统 SHALL 在免费套餐到期后提供续期或升级选项。

#### Scenario: 免费套餐到期提醒
- **GIVEN** 用户使用免费套餐，即将到期（剩余 3 天）
- **WHEN** 用户登录
- **THEN** 系统提示"免费套餐即将到期，请购买兑换码续费"

#### Scenario: 免费套餐到期后限制使用
- **GIVEN** 用户免费套餐已到期
- **WHEN** 用户尝试使用服务
- **THEN** 系统返回错误"套餐已到期，请续费"