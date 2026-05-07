# Kratos Layout Codex Agent Guide

## 角色与目标

你是这个仓库的 Go 技术伙伴，熟悉 `go-kratos`、分层架构、`protobuf`、`wire`、`buf`、`protovalidate`。

这个目录下的项目默认按 `go-kratos` 项目处理。只要需求涉及以下任一主题，就必须优先按 Kratos 项目思维处理，而不是按普通 Go 项目随手实现：

- `kratos` / `go-kratos`
- `protobuf` API 或配置定义
- HTTP / gRPC Service
- `buf` / `buf.gen.yaml` / `buf.validate`
- `wire`
- middleware、日志、鉴权、文件上传、自定义 HTTP 路由
- 目录中出现 `api`、`cmd`、`internal/service`、`internal/biz`、`internal/data`、`internal/server`、`internal/conf`

贴合当前仓库事实，不要为了“标准目录”强行大重构。Kratos 常见目录结构只作为参考，本仓库实际结构优先。

## 随仓库携带的 Skills

本 layout 自带 Codex skills，位于项目相对路径：

- `./.codex/skills/go-kratos`
- `./.codex/skills/cc-skills-golang`
- `./.codex/skills/agent-rules-books`

引用 skill 时只使用 skill 名称或上述相对路径，不要写本机用户主目录下的绝对路径。

### Go Style Decisions（最高优先级）

所有 Go / Golang 代码风格必须优先遵循 Google Go Style Decisions，其次遵循 Uber Go Style Guide：

- Google Go Style Decisions: https://google.github.io/styleguide/go/decisions
- Uber Go Style Guide: https://github.com/uber-go/guide/blob/master/style.md

Go 代码风格优先级：

1. Google Go Style Decisions（最高优先级）
2. Uber Go Style Guide
3. `cc-skills-golang` 中的具体子技能
4. `agent-rules-books` 中的 Clean Code / Refactoring / 软件设计书本规则
5. 本 `AGENTS.md` / `go-kratos` skill 的 Kratos 分层、架构与生成代码规则
6. 当前仓库既有风格和通用 Go 惯用写法

如果规则冲突，优先按更高优先级执行；除非系统 / 开发者 / 用户指令明确要求例外。

### Go Skills 选择

所有 Go 代码相关任务都要按需使用 `cc-skills-golang` 的具体子技能；不要只笼统引用整个技能包。

- 写 Go 代码、命名、注释、格式化：`cc-skills-golang:golang-code-style`、`cc-skills-golang:golang-naming`
- 写测试：`cc-skills-golang:golang-testing`、`cc-skills-golang:golang-stretchr-testify`
- 处理 error / wrapping / sentinel error：`cc-skills-golang:golang-error-handling`
- 处理 `context.Context`：`cc-skills-golang:golang-context`
- 处理数据库 / GORM / SQL：`cc-skills-golang:golang-database`
- 处理 gRPC / protobuf：`cc-skills-golang:golang-grpc`
- 处理并发、goroutine、channel、锁、worker pool：`cc-skills-golang:golang-concurrency`
- 处理安全、鉴权、输入校验、敏感信息：`cc-skills-golang:golang-security`
- 做性能分析、优化、benchmark、pprof：`cc-skills-golang:golang-performance` 或 `cc-skills-golang:golang-benchmark`
- 处理依赖、Go 版本升级、现代化改造：`cc-skills-golang:golang-dependency-management`、`cc-skills-golang:golang-modernize`

涉及 Kratos、protobuf API、配置、分层、wire、buf 时，优先同时使用 `go-kratos` skill。

### 书本原则 Skills

涉及 cleanup、refactor、命名、模块边界、复杂度控制、遗留代码接缝、设计评审时，使用 `agent-rules-books` skill。该 skill 随仓库携带，来源于 `agent-rules-books`，覆盖 Clean Code、Refactoring、A Philosophy of Software Design、Working Effectively with Legacy Code、Clean Architecture、Code Complete、DDD、PoEAA、DDIA、Release It!、The Pragmatic Programmer、Refactoring.Guru 等规则集；按任务加载 `references/<book>/<book>.mini.md`，不要引用本机绝对路径。

---

## 仓库事实

- 模块名：`github.com/lanpang/kratos-layout`（通过 `kratos new` 创建项目时会替换为新项目模块名）
- 当前主要目录：
  - `proto/`：protobuf 源文件
    - `proto/helloworld/v1/`：API proto
    - `proto/conf/`：配置 proto
  - `api/`：API 生成代码、HTTP/gRPC binding、OpenAPI 产物
  - `cmd/server/`：服务启动入口和 wire 入口
  - `internal/service/`：Kratos service 适配层，只做协议转换、请求校验和调用 biz
  - `internal/biz/`：用例编排、仓储接口、业务规则
  - `internal/data/`：数据访问实现
  - `internal/server/`：HTTP / gRPC server 启动与注册
  - `internal/conf/`：配置 Go 生成代码，不放 proto 源
  - `internal/pkg/`：项目内部通用辅助包
  - `configs/`：本地示例配置
  - `docs/`：OpenAPI / 文档生成产物，不手改
- proto 源与 Go 产物分离：
  - API proto 源在 `proto/helloworld/v1/`，Go 代码输出到 `api/helloworld/v1/`
  - 配置 proto 源在 `proto/conf/`，Go 代码输出到 `internal/conf/`
- 当前 proto 生成入口：`make proto`
  - API: `buf generate`
  - 配置: `buf generate --template internal/conf/buf.gen.yaml`
- wire 入口：`cmd/server/wire.go`，生成产物是 `cmd/server/wire_gen.go`
- 默认构建入口：`make build`
- 默认测试入口：`go test ./...`

如果后续出现规则冲突，优先级如下：

1. 当前目录代码事实
2. 本 `AGENTS.md`
3. Kratos 既有生成代码和目录约定
4. 通用 Go 最佳实践

---

## 核心原则

### 需求理解

- 先问“这个问题值不值得解决”，再想解决方案
- 优先用 User Story + Acceptance Criteria 描述需求，而不是直接跳到实现
- 如果用户直接给实现方案，也先检查这是业务问题、建模问题，还是基础设施问题

推荐理解模板：

```text
User Story:
作为 <角色>，我希望 <能力>，从而 <业务价值>

Acceptance Criteria:
1. ...
2. ...
3. ...
```

### 设计与分层

- 用 interface 替代继承，用组合替代 class 层级
- 每个 struct 尽量只做一件事，interface 优先保持在 1 到 3 个方法
- `biz` 层不依赖 `data` 具体类型，只依赖 interface
- 先满足可测试性、可维护性，再考虑性能与抽象优雅度
- 不过度设计；只在有 2 个以上真实用例时才提取抽象
- 外部系统数据结构不得泄漏进 `biz`
- 协议层 DTO / protobuf message 不要贯穿所有层；在边界处做转换

### 重构与遗留代码

- 小步修改，每次修改尽量保持测试绿
- 命名即文档
- 没有测试保护时，先补 characterization test 再重构
- 对同一个大功能模块，默认优先收敛在一个文件内维护；只有职责边界清晰、单文件明显影响理解或测试时才拆分
- 不要把一个功能模块机械地拆成很多小碎文件

### 零值语义

- 默认把零值当成“未设置 / 不存在 / 未传”
- 尽量不要让零值承担真实业务语义
- 如果业务必须区分“未传”和“显式传零值”，不要依赖普通 proto3 标量猜语义，优先使用：
  - `optional`
  - message wrapper
  - `FieldMask`
  - 指针字段对应的 biz 参数对象
  - 明确枚举

---

## 快速决策树

- 新增 API / 改 proto：看“Proto / OpenAPI 规则”
- 改 `service` / `biz` / `data`：看“分层职责”和“依赖注入”
- 改配置 / 启动入口：看“配置与启动”
- 改 HTTP 返回、自定义路由、文件上传下载：看“HTTP 定制”
- 改鉴权 / JWT / Casbin / 幂等：看“安全与鉴权”
- 改中间件 / 日志 / recovery / validate：看“中间件与日志”
- 排查奇怪行为：看“排错清单”

---

## 工作方式

### 默认改动策略

- 优先小步修改，避免一口气重构多层
- 新增逻辑时先找接缝，优先通过 biz interface 注入依赖
- 不在 biz 层拼 SQL
- 不在 service 层写条件分支业务
- 不为单一实现过早抽象
- 能复用现有 `ProviderSet` / `wire` 装配方式时，不额外造新的注入风格
- 不引入没有明确收益的新依赖

### 强执行顺序

收到改动请求时，默认按下面顺序执行：

1. 判断改动属于 `proto / api / service / biz / data / middleware / server / conf` 哪一层
2. 先找已有实现和接缝，再决定是否新增抽象
3. 设计最小可行改动，不做顺手大重构
4. 修改对应层代码
5. 补最小必要测试
6. 如果涉及 proto / wire / 生成代码，执行对应生成命令并验证

### 默认不做

- 不新建无用 interface
- 不把已有目录整体搬家
- 不引入新的 framework 式封装
- 不把简单查询过度抽象成复杂 builder
- 不为了“看起来像 DDD”做无收益的大搬家

---

## 分层职责

### `proto/`

- 放 protobuf 源文件
- API proto 放 `proto/helloworld/v1/` 或后续业务包目录
- 配置 proto 放 `proto/conf/`
- API 契约优先于实现
- 不把数据库模型、ORM 标签、外部系统字段直接暴露进 proto
- 对外契约要明确区分“未传”“零值”“默认策略”
- `go_package` 使用项目相对路径，例如 `./api/...`、`./internal/conf;conf`，避免 `kratos new` 替换 module path 后破坏 protobuf raw descriptor

### `api/`

- 放由 proto 生成的 Go API、HTTP/gRPC binding、错误码 helper、OpenAPI 产物
- 默认不手改 `*.pb.go`、`*_grpc.pb.go`、`*_http.pb.go`、`docs/openapi.yaml`
- 修改 API 时先改 `proto/`，再运行 `make proto`

### `internal/service/`

- Kratos service 适配层
- 负责请求参数校验、协议对象和 biz 参数转换、调用 biz、组装响应
- 不写核心业务规则
- 不直接访问 DB / cache / MQ / 外部 HTTP

### `internal/biz/`

- 放 Use Case / Application Service
- 定义 Repository / Gateway interface
- 编排业务流程和业务规则
- 依赖抽象，不依赖 `data` 具体实现
- 不直接处理 HTTP / gRPC / transport 细节
- 参数校验和业务前置条件校验优先在这里返回明确业务错误

### `internal/data/`

- 实现 `biz` 定义的接口
- 处理 DB、缓存、外部 API、对象存储等基础设施访问
- 做 model / DTO 到 biz 对象的转换
- 不泄漏基础设施细节到 `biz`
- 不把 ORM model 直接返回给 `biz`

### `internal/server/`

- 只负责 transport server 创建、middleware 编排、service 注册
- 不写业务逻辑
- HTTP / gRPC 注册应该保持可读、集中、可测试

### `internal/conf/`

- 放配置 Go 生成代码和配置读取相关代码
- proto 源文件在 `proto/conf/`
- 修改配置结构时先改 `proto/conf/conf.proto`，再运行 `make proto`

---

## Proto / OpenAPI 规则

- 修改 proto 后必须运行 `make proto`
- API proto 使用 `proto3`
- package 建议带版本，例如 `helloworld.v1`
- 字段编号一旦发布，不复用、不改语义
- 删除字段必须 reserved 字段号和字段名
- HTTP 注解放在 rpc 上，路径保持资源语义
- OpenAPI 注解只描述契约，不写实现细节
- 错误码 enum 使用稳定、业务可读的 reason
- `buf.validate` 用于基础输入约束；复杂业务校验放 biz 层
- 生成产物不手改

---

## 配置与启动

- 启动入口在 `cmd/server/main.go`
- wire 入口在 `cmd/server/wire.go`
- 修改依赖注入后运行 `make wire`
- 修改配置结构后运行 `make proto`，再检查 `configs/config.yaml`
- 示例配置可以用于本地开发，但不要把真实密码、token、生产地址写入仓库

---

## 测试与验证

默认验证顺序：

1. `make proto`（涉及 proto / conf 时）
2. `make wire`（涉及依赖注入时）
3. `go test ./...`
4. `go build -o bin/server ./cmd/server`
5. 必要时启动 smoke test：`./bin/server -conf configs/config.yaml`

如果本地缺少数据库、Redis 或其他外部依赖，说明验证缺口，不要把外部依赖连接失败误判为编译失败。

---

## 排错清单

- protobuf 启动 panic：优先检查 `go_package` 是否被 `kratos new` 替换进生成产物 raw descriptor；本 layout 使用相对 `go_package` 避免该问题
- 找不到生成类型：先运行 `make proto`，再检查 `proto/` 和输出目录是否一致
- wire 注入失败：检查 ProviderSet、构造函数参数、`cmd/server/wire.go`
- 配置读取失败：检查 `configs/config.yaml` 与 `proto/conf/conf.proto` 是否一致
- service 无路由：检查 `internal/server/http.go` / `grpc.go` 是否注册了对应 service
- validate 不生效：检查 proto 上的 `buf.validate` 规则和 runtime `protovalidate.Validator` 是否注入

---

## 提交前检查

- 没有手改生成代码，或生成代码变更能由 `make proto` / `make wire` 复现
- `go test ./...` 通过，或清楚说明无法运行的原因
- `go build` 通过，或清楚说明阻塞原因
- 没有新增绝对路径依赖，尤其是 skill 路径、临时目录、用户主目录
- 没有提交真实密钥、token、生产密码
