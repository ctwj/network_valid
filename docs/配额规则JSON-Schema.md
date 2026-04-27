# 配额规则 JSON Schema

本文档定义配额规则的 JSON Schema 规范。

## Schema 定义

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "quotas": {
      "type": "array",
      "items": {
        "type": "object",
        "required": ["key", "limit", "period"],
        "properties": {
          "key": {
            "type": "string",
            "description": "配额标识，用于 API 调用",
            "pattern": "^[a-zA-Z0-9_]+$",
            "maxLength": 64
          },
          "name": {
            "type": "string",
            "description": "显示名称",
            "maxLength": 128
          },
          "limit": {
            "type": "integer",
            "description": "限额数量（>=0，0表示无限制）",
            "minimum": 0
          },
          "period": {
            "type": "string",
            "description": "重置周期",
            "enum": ["daily", "weekly", "monthly"]
          },
          "reset_day": {
            "type": "integer",
            "description": "重置日：daily=每日(忽略)，weekly=周几(1-7)，monthly=几号(1-31)",
            "minimum": 1,
            "maximum": 31
          },
          "unit": {
            "type": "string",
            "description": "单位",
            "enum": ["count", "bytes", "custom"]
          }
        }
      }
    }
  },
  "required": ["quotas"]
}
```

## 字段说明

### key（必填）
- 类型：字符串
- 说明：配额唯一标识，用于 API 调用
- 规则：仅允许字母、数字、下划线，最长 64 字符
- 示例：`download`、`api_call`、`traffic`

### name（可选）
- 类型：字符串
- 说明：配额显示名称，用于前端展示
- 示例：`下载次数`、`API调用`、`流量`

### limit（必填）
- 类型：整数
- 说明：配额限额
- 特殊值：`0` 表示无限制

### period（必填）
- 类型：字符串
- 可选值：
  - `daily`：每日重置
  - `weekly`：每周重置
  - `monthly`：每月重置

### reset_day（可选）
- 类型：整数
- 说明：重置日
- 不同周期的含义：
  - `daily`：忽略此字段
  - `weekly`：1=周一，7=周日
  - `monthly`：1-31 表示几号

### unit（可选）
- 类型：字符串
- 可选值：
  - `count`：次数
  - `bytes`：字节（用于流量等）
  - `custom`：自定义单位

## 示例

### 下载次数限制（每日）
```json
{
  "quotas": [
    {
      "key": "download",
      "name": "下载次数",
      "limit": 20,
      "period": "daily",
      "unit": "count"
    }
  ]
}
```

### 流量限制（每月）
```json
{
  "quotas": [
    {
      "key": "traffic",
      "name": "流量",
      "limit": 1073741824,
      "period": "monthly",
      "reset_day": 1,
      "unit": "bytes"
    }
  ]
}
```

### 多维度配额
```json
{
  "quotas": [
    {
      "key": "download",
      "name": "下载次数",
      "limit": 20,
      "period": "daily",
      "unit": "count"
    },
    {
      "key": "api_call",
      "name": "API调用",
      "limit": 100,
      "period": "weekly",
      "reset_day": 1,
      "unit": "count"
    },
    {
      "key": "traffic",
      "name": "流量",
      "limit": 1073741824,
      "period": "monthly",
      "reset_day": 1,
      "unit": "bytes"
    }
  ]
}
```

### 无限制
```json
{
  "quotas": []
}
```

或直接设置 `quota_rules` 为 `NULL`。
