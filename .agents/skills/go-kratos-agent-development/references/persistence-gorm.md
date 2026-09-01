# Persistence and GORM

Read this reference for schema, GORM model, repository, query-option, or transaction changes.

## Establish ownership first

Find the applicable persistence ADR and determine:

- which service owns the data and who may write it;
- whether business invariants live in Go/domain code or the database;
- the accepted transaction boundaries;
- migration and compatibility requirements.

Do not assume a single writer merely because only one writer is visible in the current module. When the repository explicitly establishes a single owning Go service, keep business state rules in that service unless an atomic database guarantee is required.

## Schema changes

Prefer the smallest schema that supports established behavior:

- persisted facts needed by reads and writes;
- necessary primary keys and ownership identifiers;
- uniqueness or atomic constraints required for concurrency correctness;
- indexes justified by actual query paths.

Do not add speculative columns, indexes, business `CHECK` constraints, database enums, triggers, stored procedures, duplicated status flags, or audit structures without a documented requirement. Basic nullability, referential integrity, and atomic uniqueness are not considered unnecessary business duplication.

Escalate irreversible migrations, data backfills with uncertain semantics, or constraints that move domain policy into the database.

## GORM access

When the repository uses one-table DAO/repository operations and composable query options, preserve that model:

- keep each low-level access focused on one table;
- let options describe filters, ordering, locking, pagination, or preloading already supported for that table;
- avoid hiding cross-table business workflows inside options, callbacks, or repository helpers;
- keep names aligned with local package and receiver context rather than expanding every identifier mechanically.

Do not impose this pattern on a repository whose accepted ADR chooses aggregates or join-oriented repositories differently.

## Transactions

When business/application code owns cross-table consistency:

- DAO/repository code does not start, commit, or roll back the business transaction;
- the use case opens the transaction and passes its handle or transaction-bound repositories downward;
- every access inside the unit of work uses the same transaction rather than falling back to the default database;
- each individual data access remains within its table responsibility;
- the use case makes the cross-table sequence and failure behavior visible.

Changing which operations are atomic is a semantic change. Stop before expanding, shrinking, or relocating the transaction boundary.
