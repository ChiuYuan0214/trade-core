You are helping build a portfolio-grade centralized exchange core ("exchange-demo") for learning and interview use.

Important context:
- This is NOT a production exchange and must NOT include real fiat integration, real KYC/AML, real custody, or real blockchain deposit/withdrawal handling.
- The goal is to build an exchange-core demo that shows correct architecture, matching, accounting, real-time streaming, and spike handling.
- The implementation must optimize for correctness, clean architecture, replayability, and demonstrable engineering tradeoffs.
- Language: Go
- Primary storage: PostgreSQL
- Cache / hot distribution: Redis
- Durable event log / replay backbone: Kafka preferred
- Local development: Docker Compose
- Internal APIs: gRPC where useful
- External APIs: REST + WebSocket
- Observability: OpenTelemetry + Prometheus metrics
- Load testing: k6

==================================================
1. PROJECT GOAL
==================================================

Build a demo centralized exchange core with these capabilities:

1. User account balances
2. Available / frozen balances
3. Limit orders and market orders
4. Place order / cancel order / query order
5. In-memory matching engine
6. Real-time order book
7. Trade execution events
8. Append-only ledger
9. Public market data WebSocket
10. Private user order/trade update WebSocket
11. Replay capability from durable log
12. Bounded queues / backpressure handling for spikes

This project is intended for:
- portfolio use
- system design interview discussion
- demonstrating event-driven architecture
- demonstrating consistency and replay concepts

==================================================
2. NON-GOALS
==================================================

Do NOT implement:
- margin trading
- liquidation
- perpetual futures
- options
- real fiat rails
- real on-chain deposits/withdrawals
- compliance / KYC / AML workflows
- multi-region active-active matching
- HFT-grade microsecond optimization
- multi-tenant exchange support

==================================================
3. BUSINESS SCOPE
==================================================

Initial supported symbols:
- BTC/USDT
- ETH/USDT
- SOL/USDT

Supported order types:
- LIMIT
- MARKET

Supported order sides:
- BUY
- SELL

Order lifecycle statuses:
- PENDING_ACCEPT
- OPEN
- PARTIALLY_FILLED
- FILLED
- CANCELED
- REJECTED

==================================================
4. HIGH-LEVEL ARCHITECTURE
==================================================

Use this architecture:

[Client / Web / App]
    |
    | REST / WebSocket
    v
[Ingress / Reverse Proxy]
    |------------------------------|
    |                              |
    v                              v
[REST API Gateway]            [WebSocket Gateway]
    |                              |
    v                              |
[Order Service]                    |
    |                              |
    | validate / idempotency       |
    | reserve funds                |
    | create pending order         |
    v                              |
[Command Router]                   |
    |                              |
    | route by symbol shard        |
    v                              |
[Kafka command topic(s)]           |
    |                              |
    v                              |
[Matching Engine Shard(s)]---------|
    |
    | emits events:
    | OrderAccepted
    | OrderRejected
    | TradeExecuted
    | OrderCanceled
    | BookUpdated
    v
[Kafka event topic(s)]
    |--------------|-------------------|---------------------|
    v              v                   v                     v
[Ledger Service] [Order Projection] [Market Data Service] [Notification Service]
    |              |                   |                     |
    v              v                   v                     v
[PostgreSQL]    [PostgreSQL]        [Redis]            [WebSocket Gateway]

Important design rule:
- Order book state is owned only by Matching Engine
- Account/balance state is owned only by Ledger/Account domain
- Market data cache is owned only by Market Data Service
- Query views are projections, not the core state owner

==================================================
5. CRITICAL DOMAIN RULES
==================================================

Use these domain definitions:

- symbol: trading pair, e.g. BTC/USDT
- base asset: the asset being traded, e.g. BTC in BTC/USDT
- quote asset: the asset used to price the base asset, e.g. USDT in BTC/USDT
- qty: quantity of base asset
- price: amount of quote asset per 1 base asset

Examples:
- BUY BTC/USDT qty=0.5 price=60000
  means buy 0.5 BTC at 60000 USDT per BTC
- required quote amount before fee = 30000 USDT

==================================================
6. STATE OWNERSHIP RULES
==================================================

These are strict rules:

1. Matching Engine is the only writer of:
   - in-memory order book
   - execution sequencing
   - matching results

2. Ledger Service is the only writer of:
   - available_balance
   - frozen_balance
   - ledger entries

3. Market Data Service is the only writer of:
   - Redis order book snapshots
   - top of book
   - ticker
   - depth cache

4. Order Projection / Query Model is the only source for:
   - API order queries
   - order history read model
   - trade history read model

No service may directly mutate another service’s owned state.

==================================================
7. CORE CONSISTENCY MODEL
==================================================

Adopt this consistency model:

- Matching engine memory is the real-time execution state for a symbol shard
- Kafka is the durable ordered event log backbone
- PostgreSQL is the durable business record / query / ledger store
- Redis is cache and distribution only, never the source of truth

Key guarantees:
- at-least-once delivery is acceptable
- consumers must be idempotent
- every command/event must have a unique id
- every ledger mutation must be append-only
- system must support replay from Kafka and/or durable DB projections

==================================================
8. CRASH SAFETY REQUIREMENTS
==================================================

Never use "memory only then maybe emit later" as the only flow.

Required crash-safe behavior:
- Commands must be durably written to Kafka before considered accepted into processing flow
- Matching engine must consume commands in symbol order
- Matching results must be published as durable events
- On restart, engine must be able to rebuild state from:
  a. open order snapshot + durable event log
  or
  b. periodic snapshot + delta replay

Implement a replay tool that can reconstruct:
- open orders
- balances (if needed via ledger projection)
- recent order book state

==================================================
9. SHARDING MODEL
==================================================

Implement symbol-based sharding.

Initial shard strategy example:
- shard-1: BTC/USDT
- shard-2: ETH/USDT
- shard-3: SOL/USDT

For demo simplicity:
- one matching engine process can host multiple shards
- but architecture must clearly support "one hot symbol = one dedicated shard"

Important:
- do NOT implement multi-writer matching for the same symbol
- same symbol must be processed in strict order by one active shard owner
- for future extensibility, design with active + standby in mind

==================================================
10. SPIKE / BURST HANDLING REQUIREMENTS
==================================================

Design for burst tolerance.

Targets:
- API burst: 2000 RPS
- hot symbol order intake: 1000 orders/sec
- WebSocket concurrent clients: 30000-50000 simulated target for design discussion
- public market data fan-out must be isolated from private user feeds

Required protections:
1. bounded per-shard queue
2. reject or degrade when queue is full
3. optional policy: allow cancel orders while rejecting new orders under extreme pressure
4. public and private WS channels must be isolated
5. market data pipeline must be async and must not block matching
6. support degraded mode:
   - lower depth levels
   - reduce update frequency
   - disable non-critical analytics

==================================================
11. ORDER FLOW (BUY LIMIT ORDER)
==================================================

Implement this flow:

1. client sends POST /api/v1/orders
2. REST gateway authenticates request
3. Order Service validates:
   - symbol exists
   - side/type valid
   - qty/price valid
   - client_order_id / idempotency key
4. Order Service checks available balance via account projection
5. Order Service creates reservation request:
   - for BUY: freeze quote asset
   - for SELL: freeze base asset
6. Ledger/Account domain reserves funds
7. Order Service inserts pending order record into PostgreSQL
8. Order Service publishes PlaceOrder command to symbol shard topic
9. Matching Engine consumes command in order
10. Matching Engine either:
   - rejects order
   - accepts and places in book
   - immediately matches partially/fully
11. Matching Engine emits durable events
12. Ledger Service consumes trade/cancel/reject events and adjusts balances
13. Order Projection updates order/trade query tables
14. Market Data Service updates Redis snapshots/deltas
15. Notification Service pushes private updates to user
16. WebSocket Gateway broadcasts public market data

If order placement fails after reservation:
- release frozen funds
- mark order REJECTED

==================================================
12. BALANCE / FREEZE RULES
==================================================

Implement these rules:

BUY order:
- reserve quote asset = price * qty (+ optional fee buffer)
- decrease available_quote
- increase frozen_quote

SELL order:
- reserve base asset = qty
- decrease available_base
- increase frozen_base

On partial fill:
- decrement frozen accordingly
- increment received asset accordingly
- charge fee according to configured fee schedule

On cancel:
- release remaining frozen amount back to available

Do NOT directly overwrite balances without ledger entries.

==================================================
13. MATCHING ENGINE DESIGN
==================================================

Implement an in-memory order book with price-time priority.

Preferred design:
- one event loop goroutine per symbol shard
- avoid concurrent mutation of a single symbol book
- keep data structures simple and explainable

Suggested structures:
- buy side:
  - max-heap or sorted structure by price desc then time asc
- sell side:
  - min-heap or sorted structure by price asc then time asc
- order lookup:
  - map[orderID]*Order
- price level optimization optional but nice-to-have

Required matching behaviors:
- limit buy crosses best ask => execute
- limit sell crosses best bid => execute
- market order consumes opposite book until qty filled or liquidity exhausted
- support partial fills
- preserve maker/taker concepts
- emit trades with maker/taker IDs

Do NOT use PostgreSQL as the live matching loop state machine.

==================================================
14. EVENT TYPES
==================================================

Define event schemas for at least:

Commands:
- PlaceOrderCommand
- CancelOrderCommand

Domain events:
- FundsReserved
- FundsReservationFailed
- OrderAccepted
- OrderRejected
- OrderCanceled
- TradeExecuted
- OrderBookUpdated
- FundsReleased
- LedgerPosted

Each event must contain:
- event_id
- event_type
- occurred_at
- correlation_id
- causation_id
- symbol (if relevant)
- order_id (if relevant)
- user_id (if relevant)
- shard_id (if relevant)
- version (optional)

==================================================
15. DATABASE SCHEMA
==================================================

Use PostgreSQL with migrations.

Required tables:

users
- user_id (uuid pk)
- created_at

accounts
- user_id
- asset
- available_balance decimal(36,18)
- frozen_balance decimal(36,18)
- updated_at
- PK(user_id, asset)

orders
- order_id uuid pk
- client_order_id varchar nullable
- user_id uuid not null
- symbol varchar not null
- side varchar not null
- type varchar not null
- price decimal(36,18) nullable
- quantity decimal(36,18) not null
- filled_quantity decimal(36,18) not null default 0
- status varchar not null
- rejection_reason varchar nullable
- created_at timestamptz not null
- updated_at timestamptz not null

trades
- trade_id uuid pk
- symbol varchar not null
- maker_order_id uuid not null
- taker_order_id uuid not null
- maker_user_id uuid not null
- taker_user_id uuid not null
- price decimal(36,18) not null
- quantity decimal(36,18) not null
- maker_fee decimal(36,18) not null
- taker_fee decimal(36,18) not null
- executed_at timestamptz not null

ledger_entries
- entry_id uuid pk
- user_id uuid nullable
- asset varchar not null
- delta_available decimal(36,18) not null
- delta_frozen decimal(36,18) not null
- reference_type varchar not null
- reference_id uuid not null
- event_id uuid not null unique
- created_at timestamptz not null

processed_events
- consumer_name varchar not null
- event_id uuid not null
- processed_at timestamptz not null
- PK(consumer_name, event_id)

order_events (optional event store mirror)
- event_id uuid pk
- order_id uuid
- event_type varchar
- payload jsonb
- occurred_at timestamptz

==================================================
16. REDIS USAGE
==================================================

Use Redis only for:
- public order book snapshots
- recent top-of-book
- ticker cache
- rate limiting
- optional ws session metadata

Do NOT treat Redis as:
- the source of truth for balances
- the durable replay backbone
- the authoritative order state

==================================================
17. WEBSOCKET DESIGN
==================================================

Implement separate channels:

Public channels:
- orderbook.{symbol}
- trades.{symbol}
- ticker.{symbol}

Private channels:
- user.orders
- user.trades
- user.balances

Important routing note:
- WebSocket upgrade terminates at WebSocket Gateway (or ingress that proxies WS)
- REST gateway is not responsible for holding WS sessions
- ingress routes:
  - /api/* -> REST gateway
  - /ws/public/* -> WebSocket gateway
  - /ws/private/* -> WebSocket gateway

Implement:
- auth for private WS
- subscribe/unsubscribe
- heartbeats/ping-pong
- backpressure handling
- slow consumer disconnect policy
- snapshot + delta model for order book sync

==================================================
18. API ENDPOINTS
==================================================

REST endpoints:

POST /api/v1/orders
- place order

DELETE /api/v1/orders/{order_id}
- cancel order

GET /api/v1/orders/{order_id}
- query order

GET /api/v1/orders?symbol=BTC/USDT&status=OPEN
- list orders

GET /api/v1/trades?symbol=BTC/USDT
- list trades

GET /api/v1/accounts/balances
- show balances

GET /api/v1/market/orderbook?symbol=BTC/USDT&depth=20
- read market depth

GET /api/v1/market/ticker?symbol=BTC/USDT
- read ticker

For order placement request body:
{
  "client_order_id": "string-optional-but-recommended",
  "symbol": "BTC/USDT",
  "side": "BUY",
  "type": "LIMIT",
  "price": "60000",
  "quantity": "0.5"
}

==================================================
19. IDEMPOTENCY RULES
==================================================

Implement idempotency for order placement and event consumption.

Requirements:
- client_order_id or explicit Idempotency-Key header
- duplicate order submissions with same key must not create double orders
- all consumers must dedupe by event_id using processed_events table or equivalent

==================================================
20. OBSERVABILITY
==================================================

Implement:
- structured logs
- trace propagation
- Prometheus metrics
- OpenTelemetry traces

Minimum metrics:
- API request count/latency
- command publish latency
- shard queue depth
- matching latency
- trades/sec
- rejected orders count
- WS connections
- WS fan-out latency
- projection lag
- consumer lag

==================================================
21. FAILURE SCENARIOS TO HANDLE
==================================================

Implement and document behavior for:

1. client timeout after command accepted
   - order query must reveal eventual status
   - idempotent retry must be safe

2. command published but order later rejected
   - release reserved funds
   - mark order rejected

3. duplicate event delivery
   - consumer dedup must prevent double ledger posting

4. WebSocket client reconnect
   - fetch snapshot then continue delta stream

5. Redis unavailable
   - public market data degraded
   - matching and ledger continue

6. matching engine restart
   - replay state from durable source

7. queue overload
   - reject new orders or degrade appropriately
   - optionally prioritize cancel commands

==================================================
22. LOAD TEST / BENCHMARK REQUIREMENTS
==================================================

Create k6 scenarios:

Scenario A: normal mix
- place order 300 RPS
- cancel order 100 RPS
- query order 300 RPS

Scenario B: hot symbol spike
- BTC/USDT order burst to 2000 RPS
- measure p50/p95/p99
- measure queue depth and reject rate

Scenario C: websocket fan-out
- simulate large public subscribers
- validate snapshot + delta broadcast stability

Produce benchmark report in docs/benchmarks.md

==================================================
23. CODEBASE STRUCTURE
==================================================

Use this repo layout:

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
    docker-compose.yml
  scripts/
    seed/
    loadtest/
  docs/
    architecture.md
    order-lifecycle.md
    failure-scenarios.md
    benchmarks.md
    resume-bullets.md

==================================================
24. DEVELOPMENT PHASES
==================================================

Implement in phases:

Phase 1:
- basic REST API
- single-symbol in-memory matching engine
- PostgreSQL orders/accounts
- place/cancel/query order
- limit order matching

Phase 2:
- ledger entries
- frozen/available balances
- event-driven settlement
- idempotent consumers

Phase 3:
- Redis market data cache
- public WebSocket
- private WebSocket
- snapshot + delta

Phase 4:
- symbol sharding
- bounded queues
- overload policies
- replay tool

Phase 5:
- metrics, tracing
- k6 load tests
- benchmark docs
- polished README and architecture docs

==================================================
25. ACCEPTANCE CRITERIA
==================================================

The project is acceptable only if:

1. A user can place a LIMIT BUY/SELL and it is correctly matched by price-time priority
2. Partial fills work correctly
3. Balances move through available/frozen correctly
4. Canceling an unfilled order releases remaining frozen balance
5. Duplicate event delivery does not double-post ledger entries
6. Market data is published asynchronously and does not block matching
7. Public and private WebSocket channels are separated
8. Replay tool can rebuild a symbol’s recent state
9. Documentation clearly explains tradeoffs and ownership of state
10. System can be discussed convincingly in a backend/system design interview

==================================================
26. IMPORTANT IMPLEMENTATION TRADEOFFS TO EXPLAIN IN DOCS
==================================================

Document these tradeoffs explicitly:

- Why matching state stays in memory instead of DB
- Why one symbol should have one active matching owner
- Why Redis is cache, not truth
- Why append-only ledger is safer than balance overwrite
- Why at-least-once + idempotency is acceptable here
- Why market data is eventually consistent relative to matching core
- Why public/private WS separation helps under spike traffic
- Why hot symbols typically scale up first before exotic scale-out approaches

==================================================
27. README REQUIREMENTS
==================================================

The README must include:
- project goal
- non-goals
- architecture diagram
- order lifecycle
- balance lifecycle
- API examples
- websocket examples
- how to run locally
- how to seed demo balances
- how to replay events
- how to run load tests
- known limitations
- interview talking points

==================================================
28. CODING STYLE REQUIREMENTS
==================================================

Use:
- idiomatic Go
- clear package boundaries
- small interfaces where needed
- explicit error handling
- context.Context for request flow
- table-driven tests for domain logic
- comments only where helpful
- deterministic unit tests for matching logic

Must include tests for:
- price-time priority
- partial fills
- cancel behavior
- fund reservation/release
- idempotent event consumption

==================================================
29. FINAL OUTPUT EXPECTATION
==================================================

Generate:
1. full project skeleton
2. Docker Compose
3. initial PostgreSQL migrations
4. minimal working matching engine
5. REST API for orders and balances
6. Kafka topics config
7. Redis market data publisher
8. WebSocket gateway
9. replay tool
10. tests
11. docs

Start with a minimal but runnable version, then iterate by phases.
When uncertain, favor correctness and clear architecture over premature optimization.