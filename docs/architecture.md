# Architecture

本文记录当前服务的长期架构。它描述“系统现在如何组织”，不记录每个阶段的任务清单；阶段计划见 [roadmap.md](roadmap.md)，架构取舍见 [adr/](adr/README.md)。

## Project Shape

这是一个 Go-Kratos 服务模板。默认启动入口在 `cmd/server`，配置从 `configs/` 加载，Wire 负责依赖注入，HTTP 和 gRPC 服务由 Kratos server 注册。

核心目录：

| 目录 | 职责 |
|---|---|
| `proto/` | API 和配置的 protobuf 源文件 |
| `api/` | 从 `proto/` 生成的 Go API、gRPC、HTTP、错误码代码 |
| `cmd/server/` | 进程入口、配置加载、logger、Wire graph、Kratos app 组装 |
| `configs/` | 本地和部署示例配置 |
| `internal/server/` | HTTP/gRPC server 构造、中间件、路由注册 |
| `internal/service/` | 协议适配层：校验请求、调用 biz、转换响应 |
| `internal/biz/` | 业务用例、领域模型、repo interface、状态机和领域规则 |
| `internal/data/` | 持久化和外部数据源实现 |
| `internal/conf/` | 从配置 proto 生成的 Go 配置类型 |
| `internal/pkg/` | 项目内共享工具，如日志、middleware、ID、时间等 |

## Dependency Direction

保持依赖方向清晰：

```text
server -> service -> biz <- data
cmd/server -> server/service/biz/data/conf
```

- `service` 只做协议适配，不承载领域决策。
- `biz` 拥有业务规则和 repo interface，不依赖 `internal/data` 的具体实现。
- `data` 实现 `biz` 定义的 interface，并负责把数据库、缓存、外部 SDK 的错误转换成业务层可理解的错误。
- `server` 只组装 inbound transport，不跳过 service/biz 直接操作 data。
- `cmd/server` 只负责进程生命周期、配置、logger、依赖注入和 Kratos app。

## Request Flow

HTTP/gRPC 请求默认路径：

```text
client
  -> internal/server
  -> generated api handler
  -> internal/service
  -> internal/biz
  -> internal/data or external gateways
```

协议字段校验优先放在 protobuf validate 规则和 `service` 层；跨对象、跨状态、跨资源的判断放在 `biz` 层。

## Persistence

模板默认提供 PostgreSQL + GORM 示例。项目落地后如果切换为 sqlc、Ent、MongoDB、Redis 或其他存储模型，需要更新本文和新增 ADR，说明：

- 存储边界和 repo interface 是否变化。
- migration 如何管理。
- 事务边界归属 data 还是 biz usecase。
- 错误转换、超时、连接池和重试策略。

## API Contracts

`proto/**/*.proto` 是 HTTP/gRPC/OpenAPI 的事实来源。修改 API 时需要：

1. 更新 proto 源文件。
2. 运行 `make proto` 生成 `api/`、`internal/conf/` 和 OpenAPI 产物。
3. 更新 [api/README.md](api/README.md) 的契约索引。
4. 如有破坏性变更，在 roadmap/API 文档中标记 `BREAKING`。

## Configuration

配置入口是 `configs/` 和 `proto/conf/conf.proto`。新增配置项时需要同时更新：

- `proto/conf/conf.proto`。
- `configs/config.yaml` 示例。
- `docs/dev-setup.md` 或 `docs/runbook.md` 中的启动说明。
- 需要时补配置扫描或启动测试。

## Update Rules

当以下内容变化时必须更新本文：

- 新增或删除核心层级、包职责、外部依赖。
- 改变请求流、状态流、数据流或部署拓扑。
- 改变持久化技术、迁移方式、事务边界。
- 引入消息队列、后台 worker、scheduler、controller 等长期运行组件。
