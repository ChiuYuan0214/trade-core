# Repo Structure

Current top-level layout:

```text
exchange-demo/
  go.work
  modules/
    exchange-core-go/
      api/
      app/
      bootstrap/
      config/
      domain/
      events/
      gen/
      matching/
      postgres/
      shared/
  migrations/
  deployments/
  scripts/
    java/
    proto/
    seed/
    loadtest/
  services/
    ledger-service-go/
      cmd/ledger-service/
    matching-engine-go/
      cmd/matching-engine/
    order-service-go/
      cmd/order-service/
    replay-tool-go/
      cmd/replay-tool/
    rest-gateway-go/
      cmd/rest-gateway/
    ws-gateway-go/
      cmd/ws-gateway/
    notification-service-java/
    market-data-service-java/
  docs/
```

## Folder intent

- `go.work`: Connects the active Go modules during local development.
- `modules/exchange-core-go/`: Shared Go code used by multiple Go services.
- `modules/exchange-core-go/bootstrap/`: Common `depin` bootstrap flow and per-process registration helpers.
- `modules/exchange-core-go/shared/`: Shared demo runtime helpers such as process-local in-memory singleton state.
- `modules/exchange-core-go/domain/`: Core business types and rules.
- `modules/exchange-core-go/matching/`: In-memory book, shard loop, and queueing helpers.
- `modules/exchange-core-go/events/`: Command and event envelopes plus shared metadata.
- `modules/exchange-core-go/postgres/`: PostgreSQL adapters and migration runner.
- `modules/exchange-core-go/api/`: Shared HTTP and gRPC transport implementations.
- `services/*-go/`: Thin Go service modules with their own `go.mod` and binary entrypoints.
- Current Go service modules: `rest-gateway-go`, `order-service-go`, `ledger-service-go`, `matching-engine-go`, `ws-gateway-go`, and `replay-tool-go`.
- `migrations/`: Durable schema evolution.
- `deployments/`: Docker Compose and deployment scaffolding.
- `scripts/`: Seeding and load testing.
- `services/`: Deployable service modules, including Go and Java services.
- `docs/`: Modular architecture, feature, schema, and index references.

## File-size guidance

- Prefer one concept per file.
- Split large services into request handling, orchestration, repository, and tests.
- Keep indexes and schemas in separate docs so later work can load only the relevant slice.
