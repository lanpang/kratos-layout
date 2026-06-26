# Architecture Decision Records

ADR 记录长期架构取舍。它回答“为什么这么做”，不代替 roadmap 或实现说明。

## When To Add An ADR

需要 ADR 的情况：

- 改变 Kratos 分层职责或依赖方向。
- 改变 API 协议、兼容性策略或外部集成边界。
- 改变持久化技术、迁移方式、事务边界。
- 引入消息队列、后台 worker、scheduler、缓存一致性模型。
- 引入或移除关键依赖。
- 安全、认证、鉴权、租户隔离、恢复策略发生长期变化。

不一定需要 ADR 的情况：

- 小 bugfix。
- 不改变边界的内部重构。
- 普通字段或文案调整。

## File Naming

```text
NNNN-short-title.md
```

示例：

```text
0001-use-postgresql-gorm.md
```

## Index

| ADR | 状态 | 标题 |
|---|---|---|
| _none_ | | |
