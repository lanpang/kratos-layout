# Project Docs

本目录是 Kratos 项目的工程记忆模板。目标是让人和 Codex 在改代码前先知道：系统现在怎么设计、为什么这么设计、哪些需求已经决定、哪些问题修过、怎么验证。

> 约定：不要创建 `docs/specs/` 作为第二套规格系统。新诉求先进 `feature-requests.md`，决定要做后进入 `roadmap.md`；长期架构取舍写 ADR。

## 文档分工

| 文档 | 用途 | 什么时候更新 |
|---|---|---|
| [architecture.md](architecture.md) | 当前长期架构总览 | 改模块职责、数据流、外部依赖、部署拓扑时 |
| [adr/](adr/README.md) | 架构决策记录 | 涉及服务边界、协议、持久化、一致性、恢复、依赖选型时 |
| [roadmap.md](roadmap.md) | 已决定的阶段计划、backlog 和行为变更摘要 | 功能进入计划、完成、延期或拆分时 |
| [feature-requests.md](feature-requests.md) | 还未决定的需求入口 | 用户、外部系统、运维或开发者提出新诉求时 |
| [bugs.md](bugs.md) | bug 索引和修复记录 | 发现 bug、修复 bug、保留 workaround 时 |
| [non-goals.md](non-goals.md) | 明确不做的方向 | 拒绝一个会反复出现的方向时 |
| [glossary.md](glossary.md) | 术语表 | 出现容易混淆的业务或技术词时 |
| [api/README.md](api/README.md) | HTTP、gRPC、OpenAPI、消息和回调契约索引 | proto、路由、错误码、消息 topic 或回调入口变化时 |
| [dev-setup.md](dev-setup.md) | 本地开发和调试 | 启动方式、依赖、常见坑变化时 |
| [smoke.md](smoke.md) | 端到端冒烟验证 | 跨层行为、外部契约或部署入口变化时 |
| [security.md](security.md) | 信任边界和安全缺口 | 认证、鉴权、外部回调、存储、密钥边界变化时 |
| [arch-guards.md](arch-guards.md) | 架构不变量和守护测试 | 某类错误不应再发生，且能测试、lint 或断言时 |
| [runbook.md](runbook.md) | 运行排障流程 | 新增常见故障、恢复动作、诊断命令时 |
| [benchmarks/](benchmarks/README.md) | 性能基线 | 热路径、批量处理、调度或存储性能变化时 |
| [postmortems/](postmortems/README.md) | 严重事故复盘 | 用户可见故障、状态错误、检测盲点暴露时 |
| [research/](research/README.md) | 调研记录 | 方案还没定，需要外部或内部证据比较时 |

## 生成产物

`make proto` 会根据 `buf.gen.yaml` 生成 protobuf、gRPC、HTTP、错误码和 OpenAPI 产物。生成产物不是事实来源，不要手改：

- API 源头是 `proto/**/*.proto`。
- Go API 生成物在 `api/`。
- 配置生成物在 `internal/conf/`。
- OpenAPI 生成物可能输出到 `docs/` 下的 proto 包路径中。

## 复用流程

1. 新诉求先进 `feature-requests.md`，尽量保留原话。
2. 决定要做后，在 `roadmap.md` 记录目标行为、影响范围、验收标准和测试计划。
3. 如果涉及协议，同步更新 `api/README.md`，并以 `proto/**/*.proto` 为契约源头。
4. 如果涉及长期架构取舍，同步新增或更新 ADR。
5. 实现时按 roadmap、API、ADR、smoke、runbook 等对应文档改代码；实现偏离设计时先改文档再继续。
6. 完成后更新 `roadmap.md`、`bugs.md`、`api/README.md`、`smoke.md`、`glossary.md` 等受影响文档。
7. 对跨层或高风险改动，补 `arch-guards.md` 里的不变量和对应测试。

## 轻重判断

- 小 bugfix：`bugs.md` + 测试即可。
- 普通功能：`roadmap.md` + 测试 + 必要时 API、smoke、runbook。
- 架构级变更：`docs/adr/` + `architecture.md` 或 roadmap + smoke 或 arch guard。
- 严重线上或集成事故：`bugs.md` + `postmortems/` + arch guard 或 runbook。
