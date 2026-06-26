# Runbook

本文记录运行排障流程。它面向已经部署或本地运行中的服务，重点是如何定位和恢复。

## Service Does Not Start

检查配置加载：

```bash
make run
```

常见原因：

- `configs/config.yaml` 路径不对。
- `proto/conf/conf.proto` 的校验规则和 YAML 示例不一致。
- HTTP 或 gRPC 端口被占用。
- PostgreSQL DSN 不可连接。
- 日志目录无写入权限。

## Code Generation Fails

重新安装工具：

```bash
make init
```

更新 buf 依赖：

```bash
buf dep update
```

重新生成：

```bash
make proto
make wire
```

## API Returns Validation Error

定位步骤：

1. 查看对应 `proto/**/*.proto` 的 `buf.validate` 规则。
2. 确认 HTTP path 参数和 query/body 字段映射。
3. 检查 `internal/service` 是否在调用 biz 前执行请求校验。
4. 如果是业务规则失败，检查 `internal/biz` 的错误返回和 Kratos error reason。

## Add A Runbook Entry

```markdown
## Symptom

- 影响:
- 快速判断:
- 诊断命令:
- 恢复动作:
- 回滚动作:
- 相关日志:
- 相关文档:
```
