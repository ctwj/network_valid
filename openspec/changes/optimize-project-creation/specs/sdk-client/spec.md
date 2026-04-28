## ADDED Requirements

### Requirement: 默认登录规则对接说明

SDK 文档 SHALL 说明项目默认登录规则的配置：
- 绑定模式：普通登录（mode=1）- 用户登录不强制绑定设备
- 注册模式：普通注册（reg_mode=1）- 用户可直接注册无需激活码
- 解绑模式：任意解绑（unbind_mode=3）- 用户可随时解绑设备

#### Scenario: 使用默认规则的项目对接
- **WHEN** 用户使用新创建的项目对接 SDK
- **THEN** 无需额外配置登录规则即可正常使用注册、登录、解绑功能
