## Context

当前系统架构：
- **套餐类型（Cards 表）**：仅支持 `Days`（天数）和 `Points`（点数）两个维度
- **兑换码（Card 表）**：用于兑换套餐时长/点数
- **用户注册**：`RegMode=0` 需要兑换码，`RegMode=1` 可免费注册但无初始套餐
- **配额管理**：无概念，用户可无限使用服务

业务需求：
- 套餐类型支持配置多维度配额限制（可选）
- 创建项目时默认提供至少两套预设方案
- 预设方案只需配置时长，配额规则由管理员自行定义
- 支持永久套餐（时长为永久）
- 支持免费套餐作为用户保底

技术约束：
- 需兼容现有数据，不破坏现有用户数据
- 配额重置需支持多种周期（每日/每周/每月）
- 高并发下配额扣减需保证原子性
- 配额配置需灵活，支持多维度限制

## Goals / Non-Goals

**Goals:**
- 概念重命名：激活码类型 → 套餐类型；激活码 → 兑换码
- 套餐类型支持配置多维度配额（JSON 格式，可选）
- 配额支持多种周期：每日/每周/每月
- 项目创建时提供至少两套预设方案（可自定义命名和时长）
- 支持永久套餐（`days=0` 或 `days=-1` 表示永久）
- 用户注册时自动获得免费套餐（作为保底）
- SDK 提供配额检查、扣减接口
- 付费套餐到期后自动降级为免费套餐

**Non-Goals:**
- 不支持自定义周期（如每 3 天）
- 不支持配额转赠或共享
- 不修改现有计费逻辑（点数系统独立存在）
- 不强制预设方案的具体配额数值（由管理员自行配置）

## Decisions

### D1: 数据库表重命名策略

**选择**：保持表名不变，通过 UI 和 API 层面重命名概念

**映射关系**：
- `cards` 表 → 套餐类型（Plan Type）
- `card` 表 → 兑换码（Redeem Code）
- 前端 UI：激活码类型 → 套餐类型；激活码 → 兑换码

### D2: 套餐类型字段扩展

**新增字段**：
```sql
ALTER TABLE cards ADD COLUMN quota_rules JSON DEFAULT NULL COMMENT '配额规则JSON（可选）';
ALTER TABLE cards ADD COLUMN is_free_tier TINYINT(1) DEFAULT 0 COMMENT '是否为免费套餐';
ALTER TABLE cards ADD COLUMN priority INT DEFAULT 0 COMMENT '套餐优先级（越大越高）';
```

**关键字段说明**：
- `days`: 天数，`0` 或 `-1` 表示永久有效
- `quota_rules`: 配额规则 JSON，`NULL` 表示无配额限制
- `is_free_tier`: 是否为免费套餐（免费套餐不可删除，作为用户保底）
- `priority`: 优先级，用于升降级判断

### D3: 配额配置格式

**选择**：JSON 对象存储配额规则（可选，不强制配置）

**JSON Schema**：
```json
{
  "quotas": [
    {
      "key": "download",           // 配额标识
      "name": "下载次数",           // 显示名称
      "limit": 20,                 // 限额数量
      "period": "daily",           // 周期：daily/weekly/monthly
      "reset_day": 1,              // 重置日
      "unit": "count"              // 单位：count/bytes/custom
    }
  ]
}
```

**说明**：
- `quota_rules` 为 `NULL` 时，表示该套餐无配额限制
- 管理员可根据业务需求自行配置配额规则
- 系统不预设具体的配额数值

### D4: 用户配额状态存储

**表结构**：
```sql
CREATE TABLE user_quotas (
    id INT PRIMARY KEY AUTO_INCREMENT,
    member_id INT NOT NULL,
    quota_key VARCHAR(64) NOT NULL,
    limit_value BIGINT NOT NULL DEFAULT 0,
    used_value BIGINT NOT NULL DEFAULT 0,
    period VARCHAR(16) NOT NULL,
    reset_time DATETIME NOT NULL,
    expire_time DATETIME,              -- NULL 表示永不过期
    plan_id INT NOT NULL,
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW() ON UPDATE NOW(),
    UNIQUE KEY uk_member_quota (member_id, quota_key),
    INDEX idx_reset (reset_time),
    INDEX idx_expire (expire_time)
);
```

### D5: 预设套餐方案

**选择**：项目创建时提供预设模板，管理员可自定义命名和时长

**预设方案结构**：
```json
{
  "presets": [
    {
      "name": "免费套餐",           // 可自定义命名
      "days": 0,                  // 0 表示永久
      "is_free_tier": true,
      "priority": 0,
      "quota_rules": null         // 无配额限制，或由管理员配置
    },
    {
      "name": "标准套餐",           // 可自定义命名
      "days": 30,                 // 可自定义时长
      "priority": 10,
      "quota_rules": null         // 由管理员配置
    }
  ]
}
```

**项目创建流程**：
```
1. 管理员创建项目
2. 系统展示至少两套预设方案（免费套餐 + 至少一套付费套餐）
3. 管理员可修改方案名称、时长
4. 管理员可配置配额规则（可选）
5. 系统自动创建对应的套餐类型记录
```

**预设方案要求**：
- 必须包含一套免费套餐（`is_free_tier=true`，`days=0` 永久）
- 至少包含一套付费套餐（时长可自定义）
- 管理员可新增、修改、删除预设方案（免费套餐不可删除）

### D6: 配额扣减交互流程

**选择**：客户端主动上报 + 服务端校验扣减

**交互流程**：
```
┌─────────────┐                    ┌─────────────┐
│   Client    │                    │   Server    │
└──────┬──────┘                    └──────┬──────┘
       │                                  │
       │ 1. CheckQuota(download)          │
       │ ────────────────────────────────>│
       │<──────────────────────────────── │ 返回 {limit, used, remaining}
       │                                  │
       │ 2. 执行业务操作                    │
       │                                  │
       │ 3. DeductQuota(download, 1)      │
       │ ────────────────────────────────>│
       │<──────────────────────────────── │ 返回成功/失败
       │                                  │
```

**API 设计**：

1. **quota.check** - 检查配额
2. **quota.deduct** - 扣减配额
3. **quota.deductBatch** - 批量扣减

### D7: 套餐升降级

**选择**：基于优先级的自动升降级

**规则**：
- 每个套餐有 `priority` 字段，数值越大优先级越高
- 用户兑换高优先级套餐 → 立即升级
- 用户兑换同优先级套餐 → 时长叠加
- 用户兑换低优先级套餐 → 进入排队
- 付费套餐到期 → 自动降级为免费套餐

**用户套餐关联表**：
```sql
CREATE TABLE member_plans (
    id INT PRIMARY KEY AUTO_INCREMENT,
    member_id INT NOT NULL,
    plan_id INT NOT NULL,
    status VARCHAR(16) NOT NULL,        -- active/expired/queued
    expire_time DATETIME,               -- NULL 表示永久
    activated_at DATETIME,
    expired_at DATETIME,
    created_at DATETIME DEFAULT NOW(),
    INDEX idx_member_status (member_id, status),
    INDEX idx_expire (expire_time)
);
```

### D8: 配额扣减并发控制

**选择**：数据库行锁 + 原子更新

```sql
UPDATE user_quotas
SET used_value = used_value + ?
WHERE member_id = ? AND quota_key = ?
  AND used_value + ? <= limit_value
  AND (expire_time IS NULL OR expire_time > NOW());
```

### D9: 配额重置机制

**选择**：懒重置 + 定时任务双保险

- **懒重置**：每次操作前检查 `reset_time`，若已过期则重置
- **定时任务**：每小时扫描即将重置的配额，批量处理

## Risks / Trade-offs

| 风险 | 缓解措施 |
|------|----------|
| JSON 配额规则配置复杂 | 提供前端可视化配置界面 |
| 高并发下数据库行锁竞争 | 若 QPS > 1000，考虑引入 Redis |
| 免费套餐被滥用 | 添加注册频率限制 |
| 配额表数据量增长 | 定期清理过期数据 |

## Migration Plan

### 阶段一：数据库变更
1. `cards` 表新增字段
2. 创建 `user_quotas` 表
3. 创建 `member_plans` 表
4. 添加预设方案配置

### 阶段二：后端实现
1. 配额服务
2. 套餐服务（含升降级）
3. 预设方案服务
4. API 接口

### 阶段三：SDK 更新
1. 配额检查、扣减方法

### 阶段四：管理后台
1. 项目创建向导
2. 套餐类型管理
3. 配额配置界面

## Open Questions

1. **永久套餐标识**：`days=0` 还是 `days=-1` 表示永久？
   - 建议：`days=0` 表示永久，与现有逻辑兼容

2. **配额预警**：是否需要配额即将用完的通知？
   - 建议：支持，可配置阈值