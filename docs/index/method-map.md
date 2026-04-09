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
| REST Gateway | `POST /api/v1/orders` handler | `services/rest-gateway-go/cmd/rest-gateway` + `modules/exchange-core-go/api` | Accept external order placement requests. |
| REST Gateway | `DELETE /api/v1/orders/{order_id}` handler | `services/rest-gateway-go/cmd/rest-gateway` + `modules/exchange-core-go/api` | Accept cancel requests. |
| Bootstrap | `depin.Set` registrations | `services/*-go/cmd/*` + `modules/exchange-core-go/bootstrap` | Register per-process dependencies and lifecycle-managed services. |
| Bootstrap | `RunRESTGateway` | `modules/exchange-core-go/bootstrap/rest_gateway.go` | Wire and run the REST gateway process. |
| Bootstrap | `RunOrderService` | `modules/exchange-core-go/bootstrap/order_service.go` | Wire and run the order-service process. |
| Bootstrap | `RunMatchingEngine` | `modules/exchange-core-go/bootstrap/matching_engine.go` | Wire and run the matching-engine process. |
| Bootstrap | `RunLedgerService` | `modules/exchange-core-go/bootstrap/generic.go` | Wire and run the ledger-service process. |
| Bootstrap | `RunWSGateway` | `modules/exchange-core-go/bootstrap/generic.go` | Run the current placeholder ws-gateway process. |
| Bootstrap | `RunReplayTool` | `modules/exchange-core-go/bootstrap/generic.go` | Run the current placeholder replay-tool process. |
| App | `NewStdLogger` | `modules/exchange-core-go/app/logger.go` | Build the process logger registered through `depin`. |
| App | `NewProcessService` | `modules/exchange-core-go/app/service.go` | Build the minimal process service used by each command binary. |
| Config | `Load` | `modules/exchange-core-go/config/config.go` | Resolve process-level settings from env vars and defaults. |
| Proto Tooling | `generate.sh` | `scripts/proto/generate.sh` | Regenerate Go protobuf and gRPC bindings from the shared `proto/` contracts. |
| Order App | `PlaceOrder` | `modules/exchange-core-go/app/order_application.go` | Create a pending order, route it to the matching shard, and persist resulting order state. |
| Order App | `CancelOrder` | `modules/exchange-core-go/app/order_application.go` | Cancel a resting order through the owning shard and persist the updated status. |
| Order App | `GetOrder` | `modules/exchange-core-go/app/order_application.go` | Read the current order view from the active store. |
| Account App | `Reserve` | `modules/exchange-core-go/app/account_application.go` | Reserve funds through the active ledger backend. |
| Account App | `Release` | `modules/exchange-core-go/app/account_application.go` | Release funds through the active ledger backend. |
| Account App | `ApplyTrade` | `modules/exchange-core-go/app/account_application.go` | Settle a trade through the active ledger backend. |
| Private Events | `Publish` | `modules/exchange-core-go/app/private_event_publisher.go` | Push private order/trade/balance events from Go services to the Java notification gRPC endpoint. |
| Order Domain | `ValidateCreateInput` | `modules/exchange-core-go/domain/order/order.go` | Validate order creation inputs against symbol, side, type, price, and quantity rules. |
| Order Domain | `NewPending` | `modules/exchange-core-go/domain/order/order.go` | Create a pending order aggregate from validated input. |
| Account Domain | `Reserve` | `modules/exchange-core-go/domain/account/balance.go` | Move funds from available to frozen. |
| Account Domain | `Release` | `modules/exchange-core-go/domain/account/balance.go` | Return funds from frozen to available. |
| Account Domain | `ReservationAsset` | `modules/exchange-core-go/domain/account/reservation.go` | Resolve which asset should be reserved for an order. |
| Account Domain | `ReservationAmount` | `modules/exchange-core-go/domain/account/reservation.go` | Compute the amount to reserve for a limit order. |
| Matching Book | `Apply` | `modules/exchange-core-go/matching/book/book.go` | Match an incoming order, create trades, and rest the remainder when appropriate. |
| Matching Book | `Cancel` | `modules/exchange-core-go/matching/book/book.go` | Remove a resting order from the book. |
| Matching Book | `Snapshot` | `modules/exchange-core-go/matching/book/book.go` | Build aggregated bid/ask depth levels for read models. |
| Matching Book | `lookup map[order_id]side` | `modules/exchange-core-go/matching/book/book.go` | Support fast cancel lookup; it is not the primary matching index. |
| Matching Engine | `NewShard` | `modules/exchange-core-go/matching/engine/shard.go` | Create the single-writer symbol shard wrapper around the in-memory book. |
| Matching Engine | `Apply` | `modules/exchange-core-go/matching/engine/shard.go` | Serialize order application and advance shard sequence. |
| Matching Engine | `Cancel` | `modules/exchange-core-go/matching/engine/shard.go` | Serialize cancel processing and advance shard sequence. |
| API | `Handler` | `modules/exchange-core-go/api/http_server.go` | Register the phase-1 order HTTP endpoints on a standard mux. |
| API | `handlePlaceOrder` | `modules/exchange-core-go/api/http_server.go` | Accept `POST /api/v1/orders` requests and translate them into order application input. |
| API | `handleCancelOrder` | `modules/exchange-core-go/api/http_server.go` | Accept `DELETE /api/v1/orders/{order_id}` requests. |
| API | `handleGetOrder` | `modules/exchange-core-go/api/http_server.go` | Accept `GET /api/v1/orders/{order_id}` requests. |
| API | `handleGetBalance` | `modules/exchange-core-go/api/http_server.go` | Return current available and frozen balance for a user asset. |
| API | `PlaceOrder` | `modules/exchange-core-go/api/order_grpc_server.go` | Serve internal gRPC order placement inside `order-service`. |
| API | `CancelOrder` | `modules/exchange-core-go/api/order_grpc_server.go` | Serve internal gRPC order cancellation inside `order-service`. |
| API | `GetOrder` | `modules/exchange-core-go/api/order_grpc_server.go` | Serve internal gRPC order lookup inside `order-service`. |
| API | `GetBalance` | `modules/exchange-core-go/api/order_grpc_server.go` | Serve internal gRPC balance lookup inside `order-service`. |
| API | `PlaceOrder` | `modules/exchange-core-go/api/order_grpc_client.go` | Call `order-service` from `rest-gateway` over internal gRPC. |
| API | `CancelOrder` | `modules/exchange-core-go/api/order_grpc_client.go` | Call `order-service` cancel flow from `rest-gateway`. |
| API | `GetOrder` | `modules/exchange-core-go/api/order_grpc_client.go` | Call `order-service` order query from `rest-gateway`. |
| API | `GetBalance` | `modules/exchange-core-go/api/order_grpc_client.go` | Call `order-service` balance query from `rest-gateway`. |
| API | `ApplyOrder` | `modules/exchange-core-go/api/matching_grpc_server.go` | Serve internal gRPC matching requests inside `matching-engine`. |
| API | `CancelOrder` | `modules/exchange-core-go/api/matching_grpc_server.go` | Serve internal gRPC cancel requests inside `matching-engine`. |
| API | `Apply` | `modules/exchange-core-go/api/matching_grpc_client.go` | Call `matching-engine` from `order-service` for order-book mutation. |
| API | `Cancel` | `modules/exchange-core-go/api/matching_grpc_client.go` | Call `matching-engine` from `order-service` for cancel mutation. |
| API | `ReserveFunds` | `modules/exchange-core-go/api/ledger_grpc_server.go` | Serve internal gRPC fund reservation inside `ledger-service`. |
| API | `ReleaseFunds` | `modules/exchange-core-go/api/ledger_grpc_server.go` | Serve internal gRPC fund release inside `ledger-service`. |
| API | `ApplyTrade` | `modules/exchange-core-go/api/ledger_grpc_server.go` | Serve internal gRPC trade settlement inside `ledger-service`. |
| API | `GetBalance` | `modules/exchange-core-go/api/ledger_grpc_server.go` | Serve internal gRPC balance lookup inside `ledger-service`. |
| API | `Reserve` | `modules/exchange-core-go/api/ledger_grpc_client.go` | Call `ledger-service` from `order-service` to reserve balances. |
| API | `Release` | `modules/exchange-core-go/api/ledger_grpc_client.go` | Call `ledger-service` from `order-service` to release balances. |
| API | `ApplyTrade` | `modules/exchange-core-go/api/ledger_grpc_client.go` | Call `ledger-service` from `order-service` to settle trades. |
| API | `GetBalance` | `modules/exchange-core-go/api/ledger_grpc_client.go` | Call `ledger-service` from `order-service` to read balances. |
| Shared In-Memory | `SharedOrderState` | `modules/exchange-core-go/shared/inmemory/order_state.go` | Create or reuse the process-local demo order stores. |
| Shared In-Memory | `SharedMatchingState` | `modules/exchange-core-go/shared/inmemory/matching_state.go` | Create or reuse the process-local demo shard router. |
| Shared In-Memory | `SharedLedgerState` | `modules/exchange-core-go/shared/inmemory/ledger_state.go` | Create or reuse the process-local demo balance and ledger stores. |
| Java Notification | `consume` | `services/notification-service-java/.../PrivateEventConsumer.java` | Consume private Kafka events and hand them to the notification dispatcher. |
| Java Notification | `dispatch` | `services/notification-service-java/.../NotificationDispatchService.java` | Route a domain event into a user-facing private notification. |
| Java Notification | `PrivateNotificationServiceImplBase` | `services/notification-service-java/src/generated/java/.../PrivateNotificationServiceGrpc.java` | Generated gRPC server base class from `private_event.proto`. |
| Java Notification | `publish` | `services/notification-service-java/.../GrpcNotificationService.java` | Accept private events over gRPC and hand them to the notification dispatcher. |
| Java Notification | `notificationVirtualThreadExecutor` | `services/notification-service-java/.../VirtualThreadConfig.java` | Provide the shared Java 21 virtual-thread executor for HTTP and gRPC request handling. |
| Java Notification | `ingest` | `services/notification-service-java/.../NotificationIngestController.java` | Accept demo/private events over HTTP for local portfolio demos without Kafka. |
| Java Market Data | `ingest` | `services/market-data-service-java/.../MarketDataController.java` | Accept public market-data snapshots over HTTP for local demos or adapter intake. |
| Java Market Data | `getSnapshot` | `services/market-data-service-java/.../MarketDataController.java` | Return the latest public snapshot for one symbol. |
| Java Market Data | `stream` | `services/market-data-service-java/.../MarketDataController.java` | Open the public SSE stream for a symbol. |
| Java Market Data | `ingest` | `services/market-data-service-java/.../MarketDataDispatchService.java` | Store a public snapshot and publish it to SSE subscribers. |
| Java Market Data | `subscribe` | `services/market-data-service-java/.../SseMarketDataPublisher.java` | Register a public symbol subscriber and keep the SSE emitter alive. |
| Java Market Data | `notificationVirtualThreadExecutor` | `services/market-data-service-java/.../VirtualThreadConfig.java` | Provide the shared Java 21 virtual-thread executor for HTTP request handling. |
| Postgres | `Run` | `modules/exchange-core-go/postgres/db.go` | Open and ping the PostgreSQL connection pool. |
| Postgres | `Run` | `modules/exchange-core-go/postgres/migrations.go` | Apply repo SQL migrations during postgres-backed startup. |
| Postgres | `Run` | `modules/exchange-core-go/postgres/demo_ledger_seed.go` | Seed demo users and account balances when `ledger-service` starts with the postgres backend. |
| Postgres | `Save` | `modules/exchange-core-go/postgres/order_store.go` | Upsert durable order state into the `orders` table. |
| Postgres | `Get` | `modules/exchange-core-go/postgres/order_store.go` | Read durable order state from the `orders` table. |
| Postgres | `SaveBalance` | `modules/exchange-core-go/postgres/balance_store.go` | Upsert durable account balance state into the `accounts` table. |
| Postgres | `GetBalance` | `modules/exchange-core-go/postgres/balance_store.go` | Read durable account balance state from the `accounts` table. |
| Postgres | `AppendEntries` | `modules/exchange-core-go/postgres/ledger_store.go` | Append ledger entries transactionally into the `ledger_entries` table with event-id dedupe. |
| Order Service | `PlaceOrder` | `services/order-service-go` + `modules/exchange-core-go` | Validate request, reserve funds, create pending order, publish command. |
| Order Service | `CancelOrder` | `services/order-service-go` + `modules/exchange-core-go` | Validate cancel intent and publish cancel command. |
| Matching Engine | `RunShard` | `modules/exchange-core-go/matching/engine` | Consume symbol commands in order. |
| Matching Engine | `Match` | `modules/exchange-core-go/matching/book` | Execute price-time priority matching. |
| Ledger Service | `ReserveFunds` | `modules/exchange-core-go/api/ledger_grpc_server.go` + `modules/exchange-core-go/app` | Move available balance to frozen balance via ledger posting. |
| Ledger Service | `ReleaseFunds` | `modules/exchange-core-go/api/ledger_grpc_server.go` + `modules/exchange-core-go/app` | Return frozen balance to available balance through the ledger owner service. |
| Ledger Service | `ApplyTrade` | `modules/exchange-core-go/api/ledger_grpc_server.go` + `modules/exchange-core-go/app` | Settle executed trades. |
| Ledger Service | `GetBalance` | `modules/exchange-core-go/api/ledger_grpc_server.go` + `modules/exchange-core-go/app` | Serve the current account balance view from the ledger owner service. |
| Replay Tool | `ReplaySymbol` | `services/replay-tool-go` | Rebuild recent state from durable records. |
