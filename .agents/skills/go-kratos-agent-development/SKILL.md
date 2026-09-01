---
name: go-kratos-agent-development
description: Implement or modify Go services in Kratos-style, go.work multi-module repositories while treating project ADRs, protobuf contracts, key tests, and module ownership as execution boundaries. Use for feature, refactor, or bug-fix coding in these repositories; do not use for general Go questions or architecture discussion without implementation.
---

# Go Kratos Agent Development

Complete coherent changes autonomously inside boundaries already decided by the user and repository. Do not turn the workflow into a sequence of line-by-line approvals.

## Establish the execution contract

Before editing:

1. Read applicable repository instructions.
2. Inspect `go.work`, relevant `go.mod` files, the target module, and nearby implementations.
3. Find accepted ADRs and current domain/project context. Common locations include `docs/adr/`, `docs/architecture/`, `CONTEXT.md`, and module-local documentation; use the repository's actual layout.
4. Identify generated files, protobuf sources, hand-written interfaces, key tests, migrations, and build/generation commands.
5. Briefly state the goal, intended scope, protected seams, relevant decisions, and material assumptions. Raise at most three high-impact unknowns.

Continue without waiting when the work stays within established boundaries. Stop only when a decision listed under **Escalate semantic changes** is genuinely required.

Use this authority order:

1. The user's current request and mandatory repository instructions.
2. Accepted ADRs and explicit domain contracts.
3. Current repository conventions and reference implementations.
4. Defaults in this skill's references.

Do not silently treat a request as overriding a conflicting ADR. Surface the conflict and its consequences.

## Load only relevant guidance

- Read [workspace-and-kratos.md](references/workspace-and-kratos.md) when selecting a module, changing dependencies, or placing code in Kratos layers.
- Read [contract-seams.md](references/contract-seams.md) when protobuf, public interfaces, generated code, or key tests are involved.
- Read [persistence-gorm.md](references/persistence-gorm.md) when changing schemas, GORM models, repositories, query options, or transactions.
- Read [stateful-changes.md](references/stateful-changes.md) when behavior involves states, transitions, retries, idempotency, recovery, concurrency, or consistency.
- Read [wire-and-lifecycle.md](references/wire-and-lifecycle.md) when changing providers, dependency assembly, goroutines, servers, consumers, schedulers, or cleanup.

Do not load every reference by default. Ordinary Go style, naming, and testing knowledge do not need to be restated here; follow the repository and standard Go conventions.

## Implement within the boundary

- Prefer one coherent vertical change over many approval-gated micro-edits.
- Let human-authored contracts constrain the implementation; infer routine names and local details from package, receiver, method, module, and nearby code.
- Preserve user changes and avoid unrelated cleanup.
- Edit source files rather than generated outputs, then use repository-provided generation commands.
- Reuse established abstractions before adding packages, interfaces, configuration, schema fields, or infrastructure.
- Treat comments and shorthand as local vocabulary only when their meaning is established by project context.

## Escalate semantic changes

Stop before making a change that would:

- modify protobuf or another external contract outside the explicitly approved scope;
- weaken, replace, or contradict a key human-authored test;
- change domain states, transitions, invariants, idempotency, retry, recovery, or error semantics;
- expand or move a transaction boundary;
- introduce another writer for owned data or cross an ownership boundary;
- add a new cross-module dependency, extract a shared package, or change `go.work`/module boundaries;
- require an irreversible or data-losing migration;
- contradict an accepted ADR or expose a missing architecture decision;
- start a long-running process implicitly where lifecycle ownership is not already established.

State the current boundary, proposed change, reason, affected components, and the smallest decision needed from the user. Do not draft or accept a new ADR unless requested.

## Complete the change

Run repository generation, formatting, compilation, and targeted tests. Expand verification according to risk; use race or concurrency tests for relevant stateful code. A change is complete only when generated artifacts are current, failures are explained, and the final report identifies:

- behavior implemented;
- verification performed and its result;
- assumptions retained from ADRs or repository context;
- remaining risks and focused human review points.
