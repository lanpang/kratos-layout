# Roadmap

本文记录已经决定要做的阶段计划、backlog 和完成状态。长期架构见 [architecture.md](architecture.md)，协议入口见 [api/README.md](api/README.md)，架构取舍见 [adr/](adr/README.md)。

每个条目的完成标志：

- 代码实现完成。
- 相关测试通过。
- roadmap、API、ADR、smoke、runbook 等相关文档已更新。
- 必要时标记兼容性风险或 `BREAKING`。

## 一览

| 阶段 | 主题 | 状态 | 说明 |
|---|---|---|---|
| MVP-0 | 项目文档模板 | done | 建立 docs 工程记忆系统；不使用 `docs/specs/` 作为第二套规格目录 |

## 条目模板

```markdown
### MVP-N 主题

- 状态: planned / in-progress / done / deferred
- 相关文档: [API](api/README.md)、[ADR-NNNN](adr/NNNN-title.md)、[smoke](smoke.md)
- 验证: `go test ./...`

#### 目标

- ...

#### 范围

- ...

#### 非目标

- ...

#### 完成标准

- ...

#### 测试计划

- ...
```

## Backlog

### 待办

- [ ] 替换 demo `helloworld` API 为项目真实 API。
- [ ] 根据项目实际持久化方案更新 [architecture.md](architecture.md#persistence)。
- [ ] 补第一条真实端到端 smoke 路径。
