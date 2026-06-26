# API Contracts

本文是 HTTP、gRPC、OpenAPI、消息和回调契约的索引。当前模板的 API 源头是 `proto/helloworld/v1/*.proto`，项目落地后应替换为真实业务 proto。

## Sources Of Truth

| 契约 | 源头 | 生成或实现 |
|---|---|---|
| gRPC service | `proto/**/*.proto` | `api/**/*_grpc.pb.go` |
| HTTP route | `google.api.http` annotations in `proto/**/*.proto` | `api/**/*_http.pb.go` |
| Error reason | `proto/**/error_reason.proto` | `api/**/*error*.pb.go` 和 Kratos error helpers |
| OpenAPI | proto + `gnostic.openapi.v3` annotations | `make proto` 输出到 `docs/` 下的生成文件 |
| Config schema | `proto/conf/conf.proto` | `internal/conf/conf.pb.go` |

生成文件不要手改。修改契约时先改 proto，再运行 `make proto`。

## Current Demo API

| Method | Path | RPC | 说明 |
|---|---|---|---|
| `GET` | `/helloworld/{name}` | `helloworld.v1.GreeterService/SayHello` | demo greeting endpoint |
| `GET` | `/search/{keyword}` | `helloworld.v1.GreeterService/LuckySearch` | demo redirect payload endpoint |

## Change Rules

- 新增、删除、重命名 RPC 或字段时，更新 proto、本文、roadmap 和 smoke。
- 改 HTTP path、method、status、错误码或字段含义时，标记兼容性影响。
- 外部消费者已经依赖的字段或 route 发生不兼容变化时，在文档中标记 `BREAKING`。
- 如果引入 MQTT、Kafka、Webhook 或第三方 callback，在本文增加对应小节并说明 topic、payload、重试和幂等规则。

## Change Entry Template

```markdown
### Contract Name

- 状态: proposed / active / deprecated / removed
- 源头: `proto/...`
- 入口: `GET /...` or `rpc ...`
- 消费者: frontend / external system / internal worker
- 兼容性: compatible / BREAKING
- 验证: `go test ./...` and smoke path

#### Request

#### Response

#### Errors

#### Notes
```
