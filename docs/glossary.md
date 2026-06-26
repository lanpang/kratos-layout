# Glossary

本文记录项目内容易混淆的业务和技术术语。新增术语时优先保持和代码包名、proto 字段、ADR 中的叫法一致。

| 术语 | 含义 | 备注 |
|---|---|---|
| Service layer | `internal/service`，协议适配层 | 校验和转换请求，调用 biz usecase |
| Biz layer | `internal/biz`，业务层 | 业务规则、领域模型、repo interface |
| Data layer | `internal/data`，数据实现层 | 实现 biz repo interface，封装数据库和外部数据源 |
| Server layer | `internal/server`，入站 transport 组装层 | HTTP/gRPC server、中间件、路由注册 |
| API source | `proto/**/*.proto` | HTTP/gRPC/OpenAPI 的事实来源 |
| Generated API | `api/` | `make proto` 生成，不手改 |
| Config source | `proto/conf/conf.proto` 和 `configs/` | 配置字段源头和示例值 |
| ADR | Architecture Decision Record | 记录长期架构取舍和后果 |
| Smoke test | 冒烟验证 | 证明关键部署入口和跨层路径可用 |
