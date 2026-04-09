# Method Map

Use this file as the quick lookup table for key methods, handlers, consumers, and their purpose.

## Current status

The codebase has not been implemented yet. Populate this file incrementally as code is added.

## Registration rules

- Add one line for each business-significant public handler, service method, consumer, publisher, or replay entrypoint.
- Keep each entry short: component, method, location, purpose.
- Prefer links to files once they exist.

## Planned entries

| Component | Method / Handler | Planned Location | Purpose |
| --- | --- | --- | --- |
| REST Gateway | `POST /api/v1/orders` handler | `cmd/rest-gateway` + `internal/api` | Accept external order placement requests. |
| REST Gateway | `DELETE /api/v1/orders/{order_id}` handler | `cmd/rest-gateway` + `internal/api` | Accept cancel requests. |
| Bootstrap | `depin.Set` registrations | `cmd/*` | Register per-process dependencies and lifecycle-managed services. |
| Bootstrap | `RunProcess` | `internal/bootstrap/run.go` | Build the minimal process graph and start shared bootstrap flow. |
| App | `NewStdLogger` | `internal/app/logger.go` | Build the process logger registered through `depin`. |
| App | `NewProcessService` | `internal/app/service.go` | Build the minimal process service used by each command binary. |
| Config | `Load` | `internal/config/config.go` | Resolve process-level settings from env vars and defaults. |
| Order App | `PlaceOrder` | `internal/app/order_application.go` | Create a pending order, route it to the matching shard, and persist resulting order state. |
| Order App | `CancelOrder` | `internal/app/order_application.go` | Cancel a resting order through the owning shard and persist the updated status. |
| Order App | `GetOrder` | `internal/app/order_application.go` | Read the current order view from the active store. |
| Account App | `Reserve` | `internal/app/account_application.go` | Move balances from available to frozen and append ledger entries. |
| Account App | `Release` | `internal/app/account_application.go` | Release frozen balances and append ledger entries. |
| Account App | `ApplyTrade` | `internal/app/account_application.go` | Settle a trade by reducing frozen inventory and crediting received assets. |
| Order Domain | `ValidateCreateInput` | `internal/domain/order/order.go` | Validate order creation inputs against symbol, side, type, price, and quantity rules. |
| Order Domain | `NewPending` | `internal/domain/order/order.go` | Create a pending order aggregate from validated input. |
| Account Domain | `Reserve` | `internal/domain/account/balance.go` | Move funds from available to frozen. |
| Account Domain | `Release` | `internal/domain/account/balance.go` | Return funds from frozen to available. |
| Account Domain | `ReservationAsset` | `internal/domain/account/reservation.go` | Resolve which asset should be reserved for an order. |
| Account Domain | `ReservationAmount` | `internal/domain/account/reservation.go` | Compute the amount to reserve for a limit order. |
| Matching Book | `Apply` | `internal/matching/book/book.go` | Match an incoming order, create trades, and rest the remainder when appropriate. |
| Matching Book | `Cancel` | `internal/matching/book/book.go` | Remove a resting order from the book. |
| Matching Book | `Snapshot` | `internal/matching/book/book.go` | Build aggregated bid/ask depth levels for read models. |
| Matching Book | `lookup map[order_id]side` | `internal/matching/book/book.go` | Support fast cancel lookup; it is not the primary matching index. |
| Matching Engine | `NewShard` | `internal/matching/engine/shard.go` | Create the single-writer symbol shard wrapper around the in-memory book. |
| Matching Engine | `Apply` | `internal/matching/engine/shard.go` | Serialize order application and advance shard sequence. |
| Matching Engine | `Cancel` | `internal/matching/engine/shard.go` | Serialize cancel processing and advance shard sequence. |
| API | `Handler` | `internal/api/http_server.go` | Register the phase-1 order HTTP endpoints on a standard mux. |
| API | `handlePlaceOrder` | `internal/api/http_server.go` | Accept `POST /api/v1/orders` requests and translate them into order application input. |
| API | `handleCancelOrder` | `internal/api/http_server.go` | Accept `DELETE /api/v1/orders/{order_id}` requests. |
| API | `handleGetOrder` | `internal/api/http_server.go` | Accept `GET /api/v1/orders/{order_id}` requests. |
| API | `handleGetBalance` | `internal/api/http_server.go` | Return current available and frozen balance for a user asset. |
| Java Notification | `consume` | `services/notification-service-java/.../PrivateEventConsumer.java` | Consume private Kafka events and hand them to the notification dispatcher. |
| Java Notification | `dispatch` | `services/notification-service-java/.../NotificationDispatchService.java` | Route a domain event into a user-facing private notification. |
| Java Notification | `ingest` | `services/notification-service-java/.../NotificationIngestController.java` | Accept demo/private events over HTTP for local portfolio demos without Kafka. |
| Postgres | `Run` | `internal/postgres/db.go` | Open and ping the PostgreSQL connection pool. |
| Postgres | `Run` | `internal/postgres/migrations.go` | Apply repo SQL migrations during postgres-backed startup. |
| Postgres | `Save` | `internal/postgres/order_store.go` | Upsert durable order state into the `orders` table. |
| Postgres | `Get` | `internal/postgres/order_store.go` | Read durable order state from the `orders` table. |
| Order Service | `PlaceOrder` | `internal/...` | Validate request, reserve funds, create pending order, publish command. |
| Order Service | `CancelOrder` | `internal/...` | Validate cancel intent and publish cancel command. |
| Matching Engine | `RunShard` | `internal/matching/engine` | Consume symbol commands in order. |
| Matching Engine | `Match` | `internal/matching/book` | Execute price-time priority matching. |
| Ledger Service | `ReserveFunds` | `internal/domain/account` or service layer | Move available balance to frozen balance via ledger posting. |
| Ledger Service | `ApplyTrade` | `internal/domain/ledger` or service layer | Settle executed trades. |
| Replay Tool | `ReplaySymbol` | `cmd/replay-tool` | Rebuild recent state from durable records. |
