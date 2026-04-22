# 会员管理系统 API 对接文档

**版本**: v1.0
**更新日期**: 2026-04-22
**项目名称**: 会员管理系统

---

## 目录

- [1. 概述](#1-概述)
- [2. 认证机制](#2-认证机制)
- [3. 公共参数](#3-公共参数)
- [4. 基础接口](#4-基础接口)
- [5. 验证码接口](#5-验证码接口)
- [6. 用户管理接口](#6-用户管理接口)
- [7. 账户操作接口](#7-账户操作接口)
- [8. 错误码对照表](#8-错误码对照表)

---

## 1. 概述

### 1.1 系统简介

会员管理系统是一套完整的用户认证和权限管理平台，为第三方软件提供以下核心功能：

- **用户认证**：用户注册、登录、心跳检测
- **权限验证**：获取用户权限列表、会员标签
- **账户管理**：充值、扣点、解绑、拉黑
- **安全机制**：MD5签名、RSA加密

### 1.2 通信协议

| 项目 | 说明 |
|------|------|
| 协议 | HTTP/HTTPS |
| 请求方法 | POST |
| 数据格式 | JSON |
| 编码 | UTF-8 |

### 1.3 接口地址

请向管理员获取API服务器地址。

---

## 2. 认证机制

### 2.1 签名算法

所有接口请求都需要携带签名参数，签名生成流程如下：

**步骤1：获取时间戳**

调用时间戳接口获取服务器时间戳（秒级）。

**步骤2：组合签名字符串**

```
签名字符串 = appkey + secretkey + version + timestamp + mac
```

| 参数 | 说明 |
|------|------|
| appkey | 应用标识，由管理员分配 |
| secretkey | 应用密钥，由管理员分配 |
| version | 软件版本号 |
| timestamp | 秒级时间戳 |
| mac | 客户端机器码 |

**步骤3：MD5加密**

对签名字符串进行MD5加密，生成32位小写签名。

**步骤4：请求携带签名**

将生成的签名放入请求参数中。

### 2.2 签名生成示例（JavaScript）

```javascript
const crypto = require('crypto');

function generateSign(appkey, secretkey, version, timestamp, mac) {
    const signStr = appkey + secretkey + version + timestamp + mac;
    const sign = crypto.createHash('md5').update(signStr).digest('hex');
    return sign;
}

// 使用示例
const appkey = 'your_appkey';
const secretkey = 'your_secretkey';
const version = '1.0.0';
const timestamp = '1658038920';
const mac = 'client_machine_code';

const sign = generateSign(appkey, secretkey, version, timestamp, mac);
console.log('签名:', sign);
```

### 2.3 RSA签名机制

服务器响应中包含 `signal` 字段，为RSA加密签名，用于验证响应数据的完整性和来源可信度。

---

## 3. 公共参数

### 3.1 公共请求参数

以下参数在所有接口中通用：

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| appkey | String | 是 | 应用标识 |
| timestamp | String | 是 | 秒级时间戳 |
| sign | String | 是 | MD5签名 |
| version | String | 是 | 软件版本号 |
| mac | String | 是 | 客户端机器码 |

### 3.2 统一响应格式

所有接口返回统一的JSON格式：

```json
{
    "errno": 0,
    "data": {},
    "errmsg": "请求成功",
    "uid": "1c9516b5-c732-47a0-adb2-00a61f4ae8ae",
    "timestamp": 1658038920639,
    "sign": "12d681c19a0c784c096c44ead62e0a88",
    "signal": "J8sqpBlzY32CKPABv0MIBpNtwi/evCyoRuVjgJimoE6o78xHQSIKVpoxu6E4eA+J..."
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| errno | Number | 状态码，0表示成功，非0表示失败 |
| data | Object/Array | 返回数据 |
| errmsg | String | 提示信息 |
| uid | String | 唯一请求标识UUID |
| timestamp | Number | 服务器毫秒级时间戳 |
| sign | String | 服务器返回的MD5签名 |
| signal | String | 服务器返回的RSA签名 |

---

## 4. 基础接口

### 4.1 时间戳接口

获取服务器当前时间戳，用于签名计算。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 时间戳接口 |
| 请求方法 | POST |
| Content-Type | application/x-www-form-urlencoded |

**请求参数**

无特殊参数，携带公共参数即可。

**成功响应示例**

```json
{
    "errno": 0,
    "data": 1658038920,
    "errmsg": "请求成功",
    "uid": "1c9516b5-c732-47a0-adb2-00a61f4ae8ae",
    "timestamp": 1658038920639,
    "sign": "12d681c19a0c784c096c44ead62e0a88",
    "signal": "J8sqpBlzY32CKPABv0MIBpNtwi/evCyoRuVjgJimoE6o78xHQSIKVpoxu6E4eA+J8a54gPV2AeZr3FHX/FoP3bJO/R3hIvEF7r6mL+FchW+5ThFMEEOCHcCbD0762+QnB+VjFTzZmD5IDsDFeyLyWYe0b2jUY09WMJJez/D9khU="
}
```

**响应字段说明**

| 字段 | 类型 | 说明 |
|------|------|------|
| data | Number | 服务器当前时间戳（秒级） |

---

### 4.2 获取权限

获取当前用户的权限列表。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 获取权限 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |

**成功响应示例**

```json
{
    "errno": 0,
    "data": [
        {
            "id": 1,
            "pid": 0,
            "project_id": 1,
            "path": "func",
            "name": "操作权限",
            "free": 0,
            "show": 0,
            "description": "操作权限",
            "children": [
                {
                    "id": 2,
                    "pid": 1,
                    "path": "func.taobao",
                    "name": "淘宝权限",
                    "children": [...]
                }
            ]
        }
    ],
    "errmsg": "获取成功",
    "uid": "cab3f7a7-52a8-429c-90ae-27652092cf59",
    "timestamp": 1702921490481,
    "sign": "a954a972a6f06bc9b305456efe764b09"
}
```

**失败响应示例**

```json
{
    "errno": 400,
    "errmsg": "权限获取失败",
    "uid": "a98b9311-abac-4354-bb36-aa62a9f6c926",
    "timestamp": 1658038702771,
    "sign": "850f62ba5984330a7b2923630e90a8bf"
}
```

---

### 4.3 获取软件信息

获取软件版本和配置信息。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 获取软件信息 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |

---

### 4.4 获取会员标签

获取会员标签列表。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 获取会员标签 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |

---

### 4.5 获取远程变量

获取服务器端配置的远程变量。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 获取远程变量 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |

---

## 5. 验证码接口

### 5.1 验证码接口

获取图形验证码。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 验证码接口 |
| 请求方法 | POST |
| Content-Type | application/x-www-form-urlencoded |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |

---

### 5.2 注册验证码邮件

发送注册验证码邮件。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 注册验证码邮件 |
| 请求方法 | POST |
| Content-Type | application/x-www-form-urlencoded |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| email | String | 是 | 注册邮箱地址 |

---

### 5.3 找回账号验证码邮件

发送找回账号验证码邮件。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 找回账号验证码邮件 |
| 请求方法 | POST |
| Content-Type | application/x-www-form-urlencoded |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| email | String | 是 | 注册邮箱地址 |

---

## 6. 用户管理接口

### 6.1 用户注册

用户注册新账号。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 用户注册 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| username | String | 是 | 用户名/邮箱 |
| password | String | 是 | 密码 |
| code | String | 否 | 邮件验证码（如开启验证码注册） |
| captcha | String | 否 | 验证码密码（如开启验证码注册） |

**说明**：如果后台登录规则开启了邮件验证码，则需要传递code（邮件验证码）、captcha（验证码密码），否则将无法注册成功。

---

### 6.2 用户登录

用户登录验证。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 用户登录 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| username | String | 是 | 用户名/邮箱 |
| password | String | 是 | 密码 |

---

### 6.3 用户解绑

解绑用户设备绑定。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 用户解绑 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| username | String | 是 | 用户名/邮箱 |

---

### 6.4 用户下线

强制用户下线。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 用户下线 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| username | String | 是 | 用户名/邮箱 |

---

### 6.5 用户心跳

保持用户在线状态心跳检测。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 用户心跳 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| username | String | 是 | 用户名/邮箱 |

**说明**：客户端应定期（如每30秒）调用此接口保持在线状态。

---

## 7. 账户操作接口

### 7.1 用户充值

为用户账户充值。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 用户充值 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| username | String | 是 | 用户名/邮箱 |
| amount | Number | 是 | 充值金额/天数 |

---

### 7.2 账号扣点

扣除用户账户点数。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 账号扣点 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| username | String | 是 | 用户名/邮箱 |
| amount | Number | 是 | 扣除点数 |

---

### 7.3 账号拉黑

将用户账号加入黑名单。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 账号拉黑 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| username | String | 是 | 用户名/邮箱 |
| reason | String | 否 | 拉黑原因 |

---

### 7.4 查询用户在线

查询用户当前在线状态。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 查询用户在线 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| username | String | 是 | 用户名/邮箱 |

---

### 7.5 找回账号

通过邮箱找回账号。

**接口信息**

| 项目 | 说明 |
|------|------|
| 接口名称 | 找回账号 |
| 请求方法 | POST |
| Content-Type | multipart/form-data |

**请求参数**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| 公共参数 | - | 是 | appkey, timestamp, sign, version, mac |
| email | String | 是 | 注册邮箱地址 |
| code | String | 是 | 邮件验证码 |
| captcha | String | 是 | 验证码密码 |

---

## 8. 错误码对照表

### 8.1 常见错误码

| errno | 说明 | 处理建议 |
|-------|------|----------|
| 0 | 请求成功 | 正常处理返回数据 |
| 400 | 请求参数错误 | 检查请求参数是否完整正确 |
| 401 | 未授权/签名验证失败 | 检查签名算法和参数是否正确 |
| 403 | 权限不足 | 检查用户权限或联系管理员 |
| 404 | 资源不存在 | 检查请求地址是否正确 |
| 500 | 服务器内部错误 | 联系技术支持 |

### 8.2 业务错误码

| errno | 说明 | 处理建议 |
|-------|------|----------|
| 1001 | 用户不存在 | 检查用户名是否正确 |
| 1002 | 密码错误 | 检查密码是否正确 |
| 1003 | 账号已被禁用 | 联系管理员 |
| 1004 | 账号已被拉黑 | 联系管理员 |
| 1005 | 邮箱已被使用 | 更换邮箱地址 |
| 1006 | 验证码错误 | 重新获取验证码 |
| 1007 | 验证码已过期 | 重新获取验证码 |
| 1008 | 用户已在线 | 无需重复登录 |
| 1009 | 余额不足 | 请先充值 |
| 1010 | 设备绑定数量超限 | 解绑其他设备 |

### 8.3 错误处理建议

1. **签名验证失败**：检查时间戳是否在有效范围内，确认签名算法正确
2. **参数错误**：对照接口文档检查必填参数是否完整
3. **权限不足**：确认用户是否有对应功能权限
4. **网络超时**：建议设置合理的超时时间，并实现重试机制

---

## 附录

### A. 联系方式

如有技术问题，请联系管理员获取支持。

### B. 更新日志

| 版本 | 日期 | 更新内容 |
|------|------|----------|
| v1.0 | 2026-04-22 | 初始版本，包含18个接口文档 |

---

*本文档由会员管理系统自动生成*
