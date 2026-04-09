# Implementation Plan

## Working style

- Discuss the next feature slice before implementation.
- Implement in small vertical slices.
- Confirm completion after each feature.
- Update docs and index files in the same pass as code changes.

## Proposed sequence

1. Repo scaffold, Go module, config shape, Docker Compose, migrations baseline, docs baseline.
   Include `depin` bootstrap conventions from the start so service wiring stays consistent.
2. Core domain models: order, trade, account, ledger, events.
3. Single-symbol matching engine with price-time priority and unit tests.
4. REST order placement, cancel, and query flow with PostgreSQL order persistence.
5. Funds reservation and release with append-only ledger entries.
6. Event-driven settlement and idempotent consumers.
7. Redis-backed market data projection and public WebSocket.
8. Private user WebSocket updates.
9. Symbol sharding, bounded queue, overload policies.
10. Replay tool.
11. Observability, load tests, and benchmark docs.

## Phase checkpoint expectations

### Phase 1

- One symbol is enough.
- LIMIT orders first.
- End-to-end place, cancel, query, and matching correctness matter more than infra completeness.

### Phase 2

- Available and frozen balance transitions must be correct.
- Duplicate event delivery must not double-post ledger entries.
- Current progress: basic reserve-on-place and release-on-cancel flow is implemented with in-memory balance and ledger stores.

### Phase 3+

- Market data and private notifications stay async.
- Public and private fan-out remain isolated.

## Discussion rule

When a slice has non-obvious tradeoffs, pause for confirmation before implementation. Otherwise proceed and report back with verification and follow-up options.
