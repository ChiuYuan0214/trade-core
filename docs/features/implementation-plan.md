# Implementation Plan

## Working style

- Discuss the next feature slice before implementation.
- Implement in small vertical slices.
- Confirm completion after each feature.
- Update docs and index files in the same pass as code changes.

## Proposed sequence

1. Repo scaffold, Go workspace modules, config shape, Docker Compose, migrations baseline, docs baseline.
   Include `depin` bootstrap conventions from the start so service wiring stays consistent.
2. Core domain models: order, trade, account, ledger, events.
3. Single-symbol matching engine with price-time priority and unit tests.
4. REST order placement, cancel, and query flow with PostgreSQL order persistence.
   Current progress: `rest-gateway` now proxies order and balance requests to `order-service` over gRPC, so the HTTP edge is no longer coupled directly to in-process trading logic.
   Current structure: `rest-gateway` and `order-service` now have separate Go modules, while shared code lives in `modules/exchange-core-go`.
   Current progress: `order-service` now forwards order-book mutations to `matching-engine` over gRPC, so matching state is no longer hosted inside the order service process.
5. Funds reservation and release with append-only ledger entries.
   Current progress: `order-service` now forwards reserve/release and balance reads to `ledger-service` over gRPC, so account state is no longer owned in-process.
6. Event-driven settlement and idempotent consumers.
   Current progress: trade settlement now executes through `ledger-service` over gRPC rather than local application-layer balance mutation.
7. Redis-backed market data projection and public WebSocket.
   Current direction: the public market-data edge is being shaped as a Java Spring Boot service with SSE and optional Redis pub/sub intake.
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
- Current progress: reserve-on-place, release-on-cancel, settlement, and balance reads all route through `ledger-service`, and the service now supports `memory` or `postgres` storage behind the same gRPC edge.

### Phase 3+

- Market data and private notifications stay async.
- Public and private fan-out remain isolated.

## Discussion rule

When a slice has non-obvious tradeoffs, pause for confirmation before implementation. Otherwise proceed and report back with verification and follow-up options.
