# Service Boundaries

## Write ownership

- Matching Engine: order book state, match sequencing, trade generation.
- Ledger Service: available balance, frozen balance, ledger entries.
- Order Projection: order and trade read models for query APIs.
- Market Data Service: Redis snapshots, top of book, ticker, depth cache. Implemented as a Java Spring Boot module for public read-side fan-out and Redis integration.
- Notification Service: private feed fan-out preparation. Implemented as a Java Spring Boot module for Kafka/WebSocket integration.

## Interaction model

- REST Gateway accepts external order and query traffic.
- REST Gateway forwards order and balance requests to `order-service` over internal gRPC.
- Order Service validates requests, stores pending order records, and orchestrates write-path calls across downstream services.
- Order Service forwards book mutations to `matching-engine` over internal gRPC.
- Order Service forwards reserve, release, settlement, and balance reads to `ledger-service` over internal gRPC.
- Matching Engine owns the in-memory order books, match sequencing, and trade generation for the current demo slice.
- Ledger Service owns available and frozen balances, append-only ledger posting, and trade settlement for the current demo slice.
- Projection consumers update query tables.
- WebSocket Gateway serves isolated public and private channels.

## Current internal RPC edge

- Current Go-to-Go RPC boundary: `exchange.order.v1.OrderService`
- Current methods:
  - `PlaceOrder`
- `CancelOrder`
- `GetOrder`
- `GetBalance`
- Current reason: keep `rest-gateway` transport-only while moving business state and reservation logic behind a service boundary

- Current Go-to-Go RPC boundary: `exchange.matching.v1.MatchingEngineService`
- Current methods:
  - `ApplyOrder`
  - `CancelOrder`
- Current reason: keep order-book mutation inside `matching-engine` so `order-service` no longer hosts local shard state

- Current Go-to-Go RPC boundary: `exchange.ledger.v1.LedgerService`
- Current methods:
  - `ReserveFunds`
  - `ReleaseFunds`
  - `ApplyTrade`
  - `GetBalance`
- Current reason: keep balance mutation and ledger posting inside `ledger-service` so `order-service` does not own account state in-process

## Current Go module placement

- `rest-gateway` binary lives under `services/rest-gateway-go`.
- `order-service` binary lives under `services/order-service-go`.
- `matching-engine` binary lives under `services/matching-engine-go`.
- `ledger-service` binary lives under `services/ledger-service-go`.
- Shared Go contracts, transport code, and application logic live under `modules/exchange-core-go`.

## Important design guardrails

- Never let matching write balances directly.
- Never let ledger mutate order book state.
- Never let Redis become authoritative state.
- Never let public market-data pressure block matching.
