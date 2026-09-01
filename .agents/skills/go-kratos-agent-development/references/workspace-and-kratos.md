# Workspace and Kratos boundaries

Read this reference when locating work or changing dependencies in a Kratos-style multi-module repository.

## Discover the actual structure

1. Read `go.work` to list participating modules; do not assume the repository root is a Go module.
2. Read the target module's `go.mod`, local instructions, build files, and nearby entry points.
3. Trace imports and provider assembly to determine ownership before selecting a package.
4. Use the repository's existing Kratos interpretation. Names such as `service`, `biz`, `data`, `server`, and `conf` are evidence, not a universal layer specification.

Typical responsibilities are:

- transport/service code adapts protobuf or HTTP requests and responses;
- biz/application code owns use-case orchestration and transaction scope;
- domain code owns business state and invariants when the repository has an explicit domain layer;
- data code implements repositories and persistence details;
- command or app code owns process assembly and runtime lifecycle.

Prefer the local pattern when it is coherent and does not conflict with an ADR.

## Boundary rules

- Keep the change in the owning module.
- Do not move code to a shared module merely to make an import convenient.
- Do not add cross-module imports, change `go.work`, or create a new module without an explicit decision.
- Before adding an interface or package, verify that it hides a meaningful policy or volatile implementation rather than only forwarding calls.
- When a use case spans modules or services, identify the ownership and consistency contract instead of composing internal data access across boundaries.

If module ownership is ambiguous, report the candidate owners and the dependency direction each choice would create.
