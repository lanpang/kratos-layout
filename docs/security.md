# Security

本文记录信任边界、安全假设和已知缺口。安全策略变化时同步更新 ADR、API 文档和 runbook。

## Trust Boundaries

| 边界 | 当前状态 | 需要关注 |
|---|---|---|
| HTTP/gRPC inbound | 模板默认未启用认证 | 项目落地后明确认证、鉴权、租户隔离和审计策略 |
| Config files | `configs/` 是示例配置 | 不提交真实密码、token、证书、客户数据 |
| Database | 默认 PostgreSQL DSN 示例 | 使用最小权限账号，生产密码走 secret 管理 |
| Logs | zap + lumberjack | 避免输出 token、密码、个人信息、请求大 payload |
| Generated OpenAPI | 来自 proto annotations | 对外发布前检查 server URL、示例值和内部字段 |

## Rules

- 不要提交真实凭据、私有 token、生产 DSN、客户数据或 live smoke 凭据。
- 新增认证、鉴权、租户上下文或外部回调时，更新本文和 [api/README.md](api/README.md)。
- 对外部输入使用 proto validate、service 层校验和 biz 层业务校验分层处理。
- 文件路径、URL、回调地址、SQL 条件、消息 topic 等边界必须显式校验。

## Open Items

- [ ] 为真实项目选择认证方式，例如 JWT、mTLS、API key 或网关前置认证。
- [ ] 定义权限模型和错误返回规则。
- [ ] 定义日志脱敏字段清单。
