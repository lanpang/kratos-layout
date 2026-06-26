# Smoke Tests

本文记录关键端到端冒烟路径。冒烟测试不是完整回归测试，目标是在真实启动条件下快速证明核心入口可用。

## Preconditions

- 已运行 `make proto` 和 `make wire`。
- `configs/config.yaml` 指向可用的本地依赖。
- 默认 PostgreSQL 示例可连接，或 data 层已替换为项目实际可用实现。

## Local Demo HTTP Smoke

启动服务：

```bash
make run
```

请求 demo greeting：

```bash
curl -i http://127.0.0.1:8000/helloworld/alice
```

期望：

- HTTP status 为 200。
- 响应包含 `Hello alice`。
- 服务日志无 panic 或配置加载错误。

请求 demo search：

```bash
curl -i http://127.0.0.1:8000/search/kratos
```

期望：

- HTTP status 为 200。
- 响应包含 Google search redirect URL 和 302 status code 字段。

## Add A Smoke Path

新增跨层功能时，按以下格式追加：

```markdown
### Scenario Name

- 目的:
- 前置条件:
- 命令:
- 期望结果:
- 日志或外部系统证据:
- 最近验证日期:
```
