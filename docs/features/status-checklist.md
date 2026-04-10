# Status Checklist

Use this file as the quick truth source for what is implemented, partially implemented, and still pending.

## Completed

- Go workspace layout with separate service modules under `services/*-go`
- Shared Go core module under `modules/exchange-core-go`
- `rest-gateway -> order-service` gRPC boundary
- `order-service -> matching-engine` gRPC boundary
- `order-service -> ledger-service` gRPC boundary
- Single-symbol in-memory matching with price-time priority
- Order placement, cancel, and order query flow
- Balance query flow
- Reserve, release, and settlement application logic
- `orders` PostgreSQL persistence path
- `ledger-service` `memory|postgres` backend selection
- SQL migrations baseline
- Java `notification-service-java`
- Java `market-data-service-java`
- Java 21 virtual-thread configuration in both Java services
- Browser demo page for Java notification service
- Browser demo page for Java market-data service
- Local Java demo start/stop scripts
- Split docs, architecture docs, schema docs, and method index
- Portfolio-facing `README.md`

## Partially completed

- Ledger durability:
  - `ledger-service` has postgres stores and tests, but full end-to-end durable flow has not been exercised in a complete local infra stack recently
- Market-data integration:
  - Java market-data service is implemented and demoable, but the upstream Go projection/publishing path is not yet fully wired as the primary flow
- Notification integration:
  - Java notification service is implemented and demoable, and Go can publish private events, but Kafka-first production-style flow is not the main default path yet
- Matching engine structure:
  - The matching rules work for the current slice, but the data structure is still demo-oriented rather than a more production-like price-level FIFO implementation
- Service decomposition:
  - Core Go boundaries now exist, but some surrounding services like websocket gateway and replay tool are still thin placeholders

## Not completed yet

- Kafka as the true durable command/event backbone across the main flow
- Full projection consumers and durable read models for orders, trades, and market data
- Dedicated Go WebSocket gateway implementation
- Replay tool implementation beyond placeholder structure
- Symbol sharding coordination beyond the current demo slice
- Bounded queue and overload policy implementation
- Full observability layer:
  - metrics
  - tracing
  - benchmark docs
- Load testing / k6 / benchmarking workflow
- Complete integration-test matrix across Go services plus Kafka/Redis/PostgreSQL
- Production-style matching data structure upgrades

## Best current portfolio story

The strongest completed story today is:

1. Go owns the execution-critical service boundaries.
2. Java owns two portfolio-quality boundary services:
   - private notifications
   - public market-data delivery
3. Both Java services are:
   - Java 21
   - Spring Boot
   - virtual-thread based
   - testable
   - packageable
   - locally demoable from built-in browser pages

## Recommended next priorities

If the goal is portfolio polish first:

1. Add screenshots or short demo captures to `README.md`
2. Keep tightening the Java demo experience
3. Add a concise architecture decision write-up for why Go vs Java was split this way

If the goal is system completeness first:

1. Make Kafka the main event backbone
2. Finish read-side projections
3. Implement replay tooling and observability
4. Upgrade matching internals and system-level testing
