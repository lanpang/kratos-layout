# Non-Goals

本文记录明确不做、暂不做或拒绝反复讨论的方向。它不是 backlog；如果未来要重启某个方向，需要新增 feature request 或 ADR 说明触发条件变化。

## 条目格式

- **方向**：一句话描述不做什么。
- **状态**：active / superseded。
- **原因**：为什么当前不做。
- **重启条件**：什么变化发生后可以重新评估。
- **相关文档**：roadmap / ADR / API / bug / research。

## 当前非目标

- **方向**：不维护 `docs/specs/` 作为第二套规格系统。
- **状态**：active。
- **原因**：需求、计划、架构、协议、验证已经分别由 `feature-requests.md`、`roadmap.md`、`adr/`、`api/README.md`、`smoke.md` 承载；再引入 specs 目录会产生重复事实来源。
- **重启条件**：团队决定引入独立规格治理工具，并先通过 ADR 统一目录和迁移规则。
- **相关文档**：[README.md](README.md)。
