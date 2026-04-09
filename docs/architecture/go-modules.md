# Go Modules

The Go codebase now uses a workspace-style layout instead of treating the repo root as one large Go module.

## Current layout

```text
go.work
modules/
  exchange-core-go/
services/
  ledger-service-go/
  matching-engine-go/
  order-service-go/
  replay-tool-go/
  rest-gateway-go/
  ws-gateway-go/
  notification-service-java/
  market-data-service-java/
```

## Module intent

- `modules/exchange-core-go`: shared Go domain, application, transport, postgres, matching, config, and generated protobuf code.
- `services/rest-gateway-go`: thin HTTP edge binary that calls `order-service` over gRPC.
- `services/order-service-go`: internal Go service that owns the current order, reservation, and settlement application flow.
- `services/ledger-service-go`: balance, ledger, and settlement owner process. It now supports `memory` and `postgres` backends behind the same gRPC contract.
- `services/matching-engine-go`: matching owner process for the current in-memory single-node book slice.
- `services/ws-gateway-go`: current placeholder Go module for the future WebSocket gateway process.
- `services/replay-tool-go`: current placeholder Go module for the future replay tool process.

## Design rule

- Each deployable Go service should have its own `go.mod`.
- Shared Go code must live in a reusable module, not under a root-level `internal/` package.
- Java services stay in the same monorepo, but outside the Go module boundary.

## Current note

- The root-level Go module has been retired.
- New Go work should target the workspace modules only.
