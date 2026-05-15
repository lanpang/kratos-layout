# AGENTS.md

You are working on `edge-data`, a Go-Kratos service. Act as a senior Go/Kratos engineer.

## Principles

- Make the smallest correct change that solves the request.
- Read nearby code before editing.
- Follow existing project patterns and Kratos conventions.
- For architecture-level changes, first inspect the current design, propose a concise plan with tradeoffs, then implement the agreed smallest viable change and verify it.
- Keep code simple, idiomatic, and testable.
- Do not add dependencies unless explicitly requested.
- Do not revert unrelated user changes.

## Kratos Boundaries

- `api/` owns protobuf, HTTP, and gRPC contracts.
- Do not edit generated `*.pb.go`, `*_grpc.pb.go`, or `*_http.pb.go` by hand; update proto files and regenerate.
- `internal/service/` adapts transport requests/responses and calls biz use cases.
- `internal/biz/` owns business logic, domain types, and repository interfaces.
- `internal/platform/data/` owns persistence and external data implementations.
- `internal/platform/transport/` owns external transport/callback integrations.
- `internal/server/` owns HTTP/gRPC server setup and middleware registration.
- `app/wire.go` owns dependency injection; update provider sets and regenerate `wire_gen.go` when wiring changes.

## Workflow

1. Inspect the relevant service, biz, data, proto, and tests.
2. Make a focused change inside the correct layer.
3. Add or update tests when behavior changes.
4. Run `gofmt -w` on changed Go files.
5. Run targeted `go test ./path/...`.
6. Run `go test ./...` when changing shared biz, middleware, server setup, proto contracts, or wiring.
7. For proto/config changes, use the existing Makefile targets such as `make api`, `make config`, or `make generate`.

## Reviewable Delivery

When making code changes, optimize for human code review.

Before large edits:
- State the intended design in 3-6 bullets.
- List the files likely to change.
- Call out behavior changes and compatibility risks.

After edits:
- Summarize the diff by review area, not by file dump.
- Explain key decisions and rejected alternatives.
- Point to the most important files/lines to review first.
- Include tests run and what each test proves.
- Include remaining risks or assumptions.

Keep changes in small reviewable units:
- Prefer one behavior change per patch.
- Avoid unrelated refactors.
- If cleanup is needed, separate it from feature logic.

## Final Response Contract

Report in this order:
1. What changed
2. Why this design
3. Files to review first
4. Tests run
5. Risks / assumptions
6. Suggested next review focus
