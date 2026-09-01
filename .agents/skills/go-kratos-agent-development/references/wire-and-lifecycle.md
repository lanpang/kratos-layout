# Wire and lifecycle

Read this reference for dependency providers, assembly, goroutines, servers, consumers, schedulers, and cleanup.

## Wire repositories

If the target module already uses Wire:

- edit provider constructors, provider sets, and `wire.go` inputs;
- never hand-edit `wire_gen.go`;
- regenerate with the repository's pinned command and inspect the diff;
- preserve provider cleanup and error propagation;
- do not introduce a second assembly style for the same application without a documented reason.

If the repository uses manual assembly, follow it unless the task explicitly includes an assembly change. The existence or archival status of Wire alone is not a reason to migrate.

## Construction and runtime

Default lifecycle semantics are:

- `New` constructs a value and validates construction dependencies;
- `Open` acquires an external resource when the repository uses that distinction;
- `Start` begins asynchronous work;
- `Run` owns blocking execution and returns runtime failure;
- `Stop(ctx)` performs bounded graceful shutdown;
- `Close` releases final resources.

Constructors and Wire providers should not secretly launch long-running goroutines. A consumer, server, scheduler, cluster member, watcher, or recovery loop should become live through an explicit application-owned lifecycle call.

Use provider cleanup for failed construction and final resource release. Use the application lifecycle for ordered startup, readiness, graceful stop, and runtime error handling. Keep shutdown idempotent where the surrounding framework may invoke cleanup through more than one path.

If existing code starts work in a constructor, do not copy the pattern automatically. Determine who observes startup failure, who owns cancellation, and who waits for termination. Escalate if correcting it changes public lifecycle behavior beyond the requested scope.
