# Service Boundaries

## Write ownership

- Matching Engine: order book state, match sequencing, trade generation.
- Ledger Service: available balance, frozen balance, ledger entries.
- Order Projection: order and trade read models for query APIs.
- Market Data Service: Redis snapshots, top of book, ticker, depth cache.
- Notification Service: private feed fan-out preparation. Implemented as a Java Spring Boot module for Kafka/WebSocket integration.

## Interaction model

- REST Gateway accepts external order and query traffic.
- Order Service validates, enforces idempotency, checks balances, reserves funds, stores pending order records, and publishes commands.
- Matching Engine consumes symbol-ordered commands and emits durable domain events.
- Ledger Service settles trades and releases funds from matching outcomes.
- Projection consumers update query tables.
- WebSocket Gateway serves isolated public and private channels.

## Important design guardrails

- Never let matching write balances directly.
- Never let ledger mutate order book state.
- Never let Redis become authoritative state.
- Never let public market-data pressure block matching.
