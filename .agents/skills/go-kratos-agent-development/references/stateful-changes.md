# Stateful changes

Read this reference for state machines, consistency, retries, idempotency, recovery, or concurrent mutation.

## Reconstruct semantics before coding

Use ADRs, domain types, protobuf contracts, callers, persistence code, and key tests to identify:

- states and terminal states;
- commands or events that cause transitions;
- guards and rejected transitions;
- side effects and their ordering;
- transaction and durability boundaries;
- idempotency keys and duplicate behavior;
- concurrency ownership and conflict handling;
- retry, timeout, crash, and recovery behavior.

Write down only the parts relevant to the requested change. Do not invent missing transitions or recovery policy to make the implementation convenient.

## Implementation boundary

The Agent may implement mechanics once semantics are determined: transition lookup, validation, persistence, error propagation, observability, and tests. Escalate when the change requires choosing new states, transitions, invariants, consistency level, conflict winner, retry policy, or recovery outcome.

Keep state transitions explicit in domain or use-case code. Avoid distributing one transition across GORM callbacks, constructors, background goroutines, and transport handlers.

## Verification

Cover the relevant state/event matrix, including rejected and repeated operations. Add focused tests for:

- duplicate delivery or replay;
- concurrent attempts when possible;
- failure between durable writes and external side effects;
- cancellation and shutdown for background processing;
- restart or recovery when the design promises it.

Run race detection or repository-specific concurrency checks when shared in-memory state or goroutines are involved.
