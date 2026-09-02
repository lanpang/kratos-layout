# Trust boundaries and zero-value invariants

Read this reference when adding nil checks, validation, defaulting, constructors, pointer fields, optional values, or external adapters.

## Discover the project invariant

Read accepted ADRs, project context, constructors, domain types, adapters, and tests. Determine whether the project establishes facts such as:

- required internal pointers and dependencies are initialized before use;
- constructors or factories return fully usable objects;
- each relevant Go zero value has an intentional domain meaning;
- external values are validated or normalized before entering internal code.

When these facts are explicitly established, rely on them. Do not infer a repository-wide invariant from a few examples, and do not copy this personal preference into a project whose documented contract differs.

## Trusted internal code

Inside a boundary that guarantees initialized pointers and defined zero values:

- do not add repeated `nil` checks for required fields, arguments, or results;
- do not silently allocate a missing dependency, return an empty result, or skip work to recover from an impossible internal `nil`;
- do not turn `0`, `""`, `false`, or another defined zero value into “missing”, a default, or an error;
- do not add pointer fields merely to distinguish absence when the domain already defines the zero value;
- preserve zero values through mappings and persistence unless the contract explicitly transforms them.

If internal code can violate the invariant, fix the constructor, factory, mapper, or producer that owns it. A consumer-side guard can hide the defect and create a partially valid object. Validate required injected dependencies once at construction when necessary, not again in every method.

A nil map that must be written, a nil channel that must run, or another unusable zero value must be initialized by the owning constructor or initializer. “Zero values are defined” does not make every raw Go zero value operational by itself.

## External and untrusted boundaries

Treat decoded requests, configuration, database or RPC responses, files, network data, plugin callbacks, and third-party library results as untrusted according to their contracts.

- Check `nil`, presence, range, and shape when the external contract permits invalid or absent values.
- Follow documented library guarantees; do not add a nil check to every external call when the API guarantees a usable value on success.
- Translate and normalize external representations into the internal invariant at the adapter boundary.
- After successful normalization, pass the trusted internal type inward and avoid repeating the same checks in biz/domain code.
- Keep optionality explicit. If absence is meaningful, model it in the boundary or domain contract rather than guessing from a zero value.

Interfaces are trust boundaries according to ownership, not syntax alone. An interface implemented entirely inside the owning module may share its invariant; an extension point implemented by other modules or users requires boundary validation.

## Semantic changes and tests

Changing whether a pointer may be nil, whether zero means absent, or where normalization occurs changes the contract. Escalate before changing that meaning.

- Test constructors and factories for their promised postconditions.
- Test external adapters with missing and malformed values allowed by the external format.
- Test the normalized internal behavior without adding cases for states the internal contract forbids.
- When an impossible internal nil appears in production, repair and test the producing boundary rather than distributing defensive branches among consumers.
