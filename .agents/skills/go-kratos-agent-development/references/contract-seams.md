# Contract seams

Read this reference when protobuf, public interfaces, generated code, or key tests are involved.

## Protobuf

Treat existing `.proto` sources as human-owned contract skeletons unless the current request explicitly authorizes a specific contract change or repository documentation says otherwise.

- Implement behind the established contract without opportunistically changing messages, field numbers, RPCs, or error semantics.
- Never edit generated protobuf code directly.
- Use the repository's generation command and inspect the resulting diff.
- Escalate when implementation requires a contract change not already approved.

## Go interfaces

Treat an established interface as semantic guidance, not just a list of method signatures. Read its comments, callers, tests, ADRs, and error handling before implementing it.

- Do not widen an interface to simplify one implementation.
- Do not introduce an interface for every concrete type; add one only at a real policy, ownership, or test seam.
- Preserve cancellation, idempotency, ordering, and error semantics visible to callers.

## Key tests

Existing tests that encode business examples, state transitions, consistency, or compatibility are protected seams.

- Implement until they pass; do not rewrite expectations to fit the implementation.
- Add routine coverage inside the established semantics.
- Escalate before weakening or replacing a key test, or when two protected tests express incompatible behavior.

Human-owned does not mean the Agent cannot touch these files when explicitly asked. It means changes require an intentional scope rather than being an implementation side effect.
