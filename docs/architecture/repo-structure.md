# Repo Structure

Planned top-level layout:

```text
exchange-demo/
  cmd/
    rest-gateway/
    ws-gateway/
    order-service/
    matching-engine/
    ledger-service/
    market-data-service/
    notification-service/
    replay-tool/
  internal/
    domain/
      order/
      trade/
      account/
      ledger/
      marketdata/
    matching/
      book/
      engine/
      queue/
    events/
    kafka/
    postgres/
    redis/
    ws/
    api/
    config/
    observability/
  migrations/
  deployments/
  scripts/
    seed/
    loadtest/
  services/
    notification-service-java/
  docs/
```

## Folder intent

- `cmd/`: Thin process entrypoints per service.
- `cmd/` should primarily compose config and `depin` registrations for each process.
- `internal/domain/`: Core business types and rules.
- `internal/matching/`: In-memory book, shard loop, queueing, and replay helpers.
- `internal/events/`: Command and event envelopes plus shared metadata.
- `internal/postgres/`, `internal/kafka/`, `internal/redis/`: Infrastructure adapters.
- `internal/api/`: REST and possibly gRPC transport-layer contracts and handlers.
- `internal/ws/`: Subscription, session, and fan-out logic.
- `migrations/`: Durable schema evolution.
- `deployments/`: Docker Compose and deployment scaffolding.
- `scripts/`: Seeding and load testing.
- `services/`: Non-Go service modules; currently includes the Java Spring Boot notification service.
- `docs/`: Modular architecture, feature, schema, and index references.

## File-size guidance

- Prefer one concept per file.
- Split large services into request handling, orchestration, repository, and tests.
- Keep indexes and schemas in separate docs so later work can load only the relevant slice.
