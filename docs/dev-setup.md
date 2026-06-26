# Development Setup

本文记录本地开发、生成代码、运行和调试的入口。命令以仓库根目录为工作目录。

## Requirements

- Go 版本以 `go.mod` 为准。
- `buf` 用于 protobuf 生成。
- `wire` 用于依赖注入生成。
- PostgreSQL 用于默认 GORM 示例配置。

安装常用工具：

```bash
make init
```

## Common Commands

```bash
make proto
make wire
make build
make run
make test
make tidy
```

如果使用 Task：

```bash
task proto
task build
```

## Configuration

默认配置目录是：

```text
configs/
```

默认配置文件 `configs/config.yaml` 包含：

- HTTP server: `0.0.0.0:8000`
- gRPC server: `0.0.0.0:9000`
- PostgreSQL DSN 示例
- zap + lumberjack 日志配置

本地运行前确认 PostgreSQL 连接信息可用，或按项目实际存储方案调整 `internal/data` 和配置。

## Code Generation

API 和配置生成：

```bash
make proto
```

Wire 生成：

```bash
make wire
```

不要手改以下生成文件：

- `api/**/*.pb.go`
- `api/**/*_grpc.pb.go`
- `api/**/*_http.pb.go`
- `internal/conf/conf.pb.go`
- `cmd/**/wire_gen.go`
- OpenAPI 生成产物

## Troubleshooting

- `make proto` 失败：先确认 `buf dep update` 已执行，生成插件已安装。
- `make wire` 失败：检查 provider set 是否包含新构造函数，接口实现是否能被 Wire 推导。
- `make run` 启动失败：检查 `configs/config.yaml`、PostgreSQL、端口占用和日志目录权限。
- HTTP 请求 400：检查 proto validate 规则和请求路径参数。
