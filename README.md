# kratos-layout

`kratos-layout` 是一个基于 [go-kratos](https://github.com/go-kratos/kratos) 官方 layout 扩展的 Go 微服务项目模板。

当前仓库地址：

```bash
git@github.com:lanpang/kratos-layout.git
```

## 特性

- 基于 go-kratos v2 的分层项目结构
- 使用 [buf](https://github.com/bufbuild/buf) 管理 protobuf 生成
- 使用 Wire 生成依赖注入代码
- 内置 zap + lumberjack 日志配置
- 提供 PostgreSQL / GORM 配置示例
- 提供 `Makefile` 和 `Taskfile.yml` 常用开发命令

## 环境准备

安装项目常用工具：

```bash
make init
```

也可以按需单独安装：

```bash
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/google/wire/cmd/wire@latest
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

如果需要使用 Task：

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

## 使用模板创建项目

推荐使用 SSH 地址：

```bash
kratos new <your-app-name> -r git@github.com:lanpang/kratos-layout.git
```

也可以直接克隆当前模板仓库：

```bash
git clone git@github.com:lanpang/kratos-layout.git
cd kratos-layout
```

## 常用命令

```bash
# 生成 protobuf、gRPC、HTTP、校验和文档代码
make proto

# 生成 Wire 依赖注入代码
make wire

# 构建服务
make build

# 运行服务，默认读取 ./configs
make run

# 运行测试
make test

# 整理 Go module
make tidy
```

如果使用 Task：

```bash
task proto
task build
```

## protobuf 生成说明

首次生成或更新 buf 依赖：

```bash
buf dep update
```

清理未使用的 buf 依赖：

```bash
buf dep prune
```

生成 API 和配置 protobuf 代码（proto 源文件在 `proto/`，Go 生成代码分别输出到 `api/` 和 `internal/conf/`）：

```bash
make proto
```

等价的 buf 命令：

```bash
buf generate
buf generate --template internal/conf/buf.gen.yaml
```

> 注意：buf BSR 对未认证请求有频率限制。如果频繁使用 remote plugins 生成代码，建议登录 buf 或配置本地插件。

## 配置

默认配置文件位于：

```text
configs/config.yaml
```

其中包含 HTTP、gRPC、数据库和日志配置示例。运行前请根据本地环境调整数据库连接信息。

## 文档工作流

`docs/README.md` 是工程记忆入口。新需求先写入 `docs/feature-requests.md`，决定要做后进入 `docs/roadmap.md`；长期架构取舍写 `docs/adr/`，API 契约索引写 `docs/api/README.md`，关键端到端验证写 `docs/smoke.md`。

`make proto` 可能在 `docs/` 下生成 OpenAPI 产物。生成产物不要手改，协议源头始终是 `proto/**/*.proto`。

## 目录结构

```text
api/        生成后的 API 代码
cmd/        服务启动入口
configs/    配置文件
internal/   业务实现代码
proto/      protobuf 源文件
docs/       工程记忆文档与 OpenAPI 生成产物
```
