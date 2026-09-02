# Error handling

Read this reference when adding or changing errors, sentinel errors, layer mappings, logging, cleanup, or asynchronous failure propagation.

Start by inspecting the repository's existing error conventions and Go version. Preserve an established coherent contract. When the repository has no stronger rule, use standard Go error chains rather than inventing a parallel error framework.

## Propagate with identity

- Add useful operation context with `%w`, for example `fmt.Errorf("load workflow %s: %w", wfID, err)`.
- Use `errors.Is` and `errors.As` for decisions. Never compare error strings.
- Wrap at meaningful ownership or operation boundaries, not mechanically in every function.
- Return `err` unchanged when the current layer cannot add useful context.
- Do not use `%v` where callers must retain the underlying error identity.
- Keep error text concise, lowercase, and free of trailing punctuation unless it is a user-facing message.
- Do not include secrets or large payloads in error text.

Wrapping is for causal context. It does not authorize changing the caller-visible error category.

## Sentinel and typed errors

Use a sentinel only when callers need a stable category for control flow, such as not found, conflict, invalid transition, or already completed.

- Define the sentinel once in the package that owns the semantic decision, for example `var ErrWorkflowNotFound = errors.New("workflow not found")`.
- Preserve it through `%w` and test it with `errors.Is`.
- Do not create a new sentinel for every message or operation.
- Use a typed error when callers require structured data beyond a category; support `errors.As`, and implement `Is` only when category matching is intentional.
- Do not add a shared error package merely to centralize unrelated sentinels.

## Layer responsibilities

Keep one error chain while translating only at real abstraction boundaries:

- Data/repository code translates recognized driver or GORM conditions into repository/domain sentinels and wraps unexpected infrastructure failures with operation context. For example, map `gorm.ErrRecordNotFound` when the repository contract distinguishes absence.
- Biz/domain code returns domain errors and preserves their identity. It must not depend on HTTP, gRPC, or Kratos transport status codes.
- Transport/service code maps stable domain categories to the external protobuf, Kratos, HTTP, or gRPC contract. Unexpected internal causes are logged but not leaked to clients.
- Process boundaries log an error once with request or operation context. Avoid logging an error in a lower layer and then returning the same error for every upper layer to log again.

Before introducing a new translation, identify which layer owns the new semantic category. Changing an existing category or public mapping is an error-semantics change and must be escalated.

## Transactions, cleanup, and goroutines

- Return transaction callback errors without destroying their identity; add transaction context only where it helps diagnosis.
- Preserve the primary operation failure when rollback or cleanup also fails. Use the repository's convention; use `errors.Join` only when the module's Go version supports it and both failures are actionable.
- Long-running goroutines must propagate fatal errors through the owning `Run`, error group, or application lifecycle. Do not only log and continue when the process owner needs to react.
- Cancellation is normally propagated as `context.Canceled` or `context.DeadlineExceeded`; do not replace it with an unrelated sentinel.
- Reserve panic for violated programmer invariants or impossible initialization states already treated that way by the repository, not ordinary runtime failure.

## Verification

- Test stable categories with `errors.Is` and structured errors with `errors.As`.
- Test transport mappings at the transport boundary.
- Assert full error strings only when the exact text is itself a documented contract; otherwise verify category and the useful context fragment separately.
- Include failure-path tests for transaction rollback, cleanup, and asynchronous error propagation when changed.
