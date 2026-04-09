# Event Design

Keep command contracts and domain events separate from relational tables.

## Shared event metadata

Every command and event should carry:

- `event_id`
- `event_type`
- `occurred_at`
- `correlation_id`
- `causation_id`
- `symbol` when relevant
- `order_id` when relevant
- `user_id` when relevant
- `shard_id` when relevant
- `version` optional

## Commands

### `PlaceOrderCommand`

- Purpose: submit a validated order request to the symbol shard.
- Expected payload areas:
  - order identity
  - user identity
  - symbol, side, type
  - price and quantity
  - idempotency metadata

### `CancelOrderCommand`

- Purpose: request cancellation for an existing order on the owning shard.

## Domain events

### `FundsReserved`

- Emitted when funds are frozen successfully before order routing completes.
- Current implementation note: the reserve path is now owned by `ledger-service`, which applies the balance mutation and append-only ledger write for the current demo slice. Durable event publication is still pending.

### `FundsReservationFailed`

- Emitted when reservation cannot be completed.

### `OrderAccepted`

- Emitted when a command becomes an active order in the book or enters immediate execution flow.

### `OrderRejected`

- Emitted when validation or matching rules reject the order.

### `OrderCanceled`

- Emitted when the order is canceled and any remainder should be released.

### `TradeExecuted`

- Emitted for every match result and should include maker/taker identities and execution values.

### `OrderBookUpdated`

- Emitted for market data consumers to derive snapshots and deltas asynchronously.

### `FundsReleased`

- Emitted when frozen balances are returned to available balances.
- Current implementation note: cancel-triggered release now executes through `ledger-service`, with `memory` and `postgres` backends supported behind the same write owner.

### `LedgerPosted`

- Emitted when append-only ledger entries are durably written.
- Current implementation note: the current write owner is the `ledger-service` process.
