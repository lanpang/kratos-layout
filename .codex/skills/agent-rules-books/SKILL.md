---
name: agent-rules-books
description: Use book-derived AI coding rules from the bundled agent-rules-books reference set. Apply when work involves Clean Code, Refactoring, Clean Architecture, DDD, legacy-code seams, complexity reduction, enterprise patterns, data-intensive design, production reliability, or general pragmatic engineering judgment.
---

# Agent Rules Books Skill

This skill wraps the bundled `agent-rules-books` rule set copied into this repository. Use project-relative paths only.

## Bundled References

Rule files live under:

- `references/clean-code/`
- `references/refactoring/`
- `references/a-philosophy-of-software-design/`
- `references/working-effectively-with-legacy-code/`
- `references/clean-architecture/`
- `references/domain-driven-design/`
- `references/domain-driven-design-distilled/`
- `references/implementing-domain-driven-design/`
- `references/patterns-of-enterprise-application-architecture/`
- `references/code-complete/`
- `references/the-pragmatic-programmer/`
- `references/designing-data-intensive-applications/`
- `references/release-it/`
- `references/refactoring-guru/`
- `references/_compatibility/`

Each rule set usually has:

- `*.nano.md` for tiny always-on reminders
- `*.mini.md` for most task-level use
- `*.md` for full reference/audit use

## How To Choose

- Everyday code quality: `clean-code`, `code-complete`
- Refactoring: `refactoring`, `refactoring-guru`, `a-philosophy-of-software-design`
- Legacy code: `working-effectively-with-legacy-code`, then `refactoring`
- Architecture boundaries: `clean-architecture`, `patterns-of-enterprise-application-architecture`
- Domain modeling: `domain-driven-design-distilled`, `domain-driven-design`, `implementing-domain-driven-design`
- Data/consistency/distributed systems: `designing-data-intensive-applications`
- Production resilience: `release-it`
- General engineering practice: `the-pragmatic-programmer`

Default to `mini` for a task. Use `nano` for always-on reminders and `full` only for deep audits or focused sessions.

## Workflow

1. Identify the task pressure: naming, complexity, boundary, refactor, legacy seam, domain model, data flow, or reliability.
2. Load only the smallest relevant reference file from `references/<book>/<book>.mini.md`.
3. If two books conflict, check `references/_compatibility/` when available; otherwise prefer repo `AGENTS.md` and current code facts.
4. Make the smallest behavior-preserving change that improves the relevant pressure.
5. Verify with targeted tests, `go test ./...`, and build/generation commands when applicable.

## Guardrails

- Do not quote large portions of the bundled references in final answers.
- Do not treat book rules as stronger than user instructions or repository facts.
- Do not use local absolute paths; all references must be project-relative.
- Do not do broad cleanup when the user asked for a narrow fix.
