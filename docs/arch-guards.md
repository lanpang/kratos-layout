# Architecture Guards

本文记录项目必须守住的架构不变量。ADR 记录“为什么做这个选择”，本文记录“为了让这个选择长期成立，每次改代码必须守住什么”。

## 条目格式

```markdown
### ARCH-NNN: 规则名

- **范畴**：dependency / protocol / state-machine / security / runtime / persistence
- **规则**：一句话说清不变量
- **动机**：违反会怎样，引用 ADR / roadmap / API / bug
- **实现**：测试、lint、运行时断言或人工 review 点
- **绕过**：允许的例外和注释格式
- **添加日期**：YYYY-MM-DD
```

## 绕过约定

允许显式绕过，但必须写理由：

```go
// arch:ARCH-NNN-ok reason
```

## 当前规则

### ARCH-001: service 层不得直接访问 data 实现

- **范畴**：dependency
- **规则**：`internal/service` 只调用 biz usecase，不直接依赖 GORM、SQL client、Redis client、外部 SDK 或 `internal/data` 具体类型。
- **动机**：保持 Kratos 分层，避免协议适配层吞掉业务状态和测试边界。
- **实现**：当前为 review 规则；后续可加 import lint。
- **绕过**：仅 server 注册、自定义 HTTP handler adapter 等边界代码可例外。
- **添加日期**：YYYY-MM-DD

### ARCH-002: generated files 不得手改

- **范畴**：protocol
- **规则**：`api/**/*.pb.go`、`api/**/*_grpc.pb.go`、`api/**/*_http.pb.go`、`internal/conf/conf.pb.go`、`cmd/**/wire_gen.go` 和 OpenAPI 生成产物不得手工编辑。
- **动机**：生成物会被 `make proto` 或 `make wire` 覆盖，手改会造成不可复现行为。
- **实现**：review 检查；必要时在 CI 增加 regenerate diff check。
- **绕过**：无。需要改源文件或生成配置。
- **添加日期**：YYYY-MM-DD

### ARCH-003: biz 层不得依赖 data 实现

- **范畴**：dependency / persistence
- **规则**：`internal/biz` 定义 repo interface 和业务模型，不 import `internal/data`、数据库 driver、ORM model 或外部 SDK。
- **动机**：保持业务规则可测试，避免持久化细节泄漏到领域层。
- **实现**：review 检查；后续可加 import lint。
- **绕过**：需要 ADR 说明新的分层边界。
- **添加日期**：YYYY-MM-DD

## 计划规则

- [ ] API 路径、字段、错误码变更必须同步更新 `docs/api/README.md` 和 smoke。
- [ ] 配置字段变更必须同步更新 `configs/` 示例和 `docs/dev-setup.md`。
- [ ] 外部回调失败不能静默吞掉，必须记录状态或返回可观测错误。

## 新增规则流程

1. 从 bug、review 或事故里识别“这类错误不该再发生”。
2. 判断能否用单元测试、AST/import lint、集成测试或运行时断言守住。
3. 同一变更里补规则说明和对应测试。
4. 确认绕过机制，避免规则变成不可维护的硬编码。
