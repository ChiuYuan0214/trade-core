# Table Design

This file groups relational schema by ownership domain so later work can load only the needed section.

## Migration note

- Phase 1 currently applies all `.sql` files in lexical order from the repo `migrations/` directory during postgres-backed startup.
- This is intentionally simple for the demo and can later evolve into a richer migration table/versioning workflow.

## Identity

### `users`

- Purpose: minimal user registry for demo ownership.
- Primary key: `user_id`
- Key fields: `created_at`

## Accounts and ledger

### `accounts`

- Owner: Ledger/Account domain
- Purpose: current available and frozen balances by user and asset.
- Current write owner: `ledger-service`
- Primary key: `(user_id, asset)`
- Important fields:
  - `available_balance`
  - `frozen_balance`
  - `updated_at`

### `ledger_entries`

- Owner: Ledger/Account domain
- Purpose: append-only durable balance mutations.
- Current write owner: `ledger-service`
- Primary key: `entry_id`
- Important fields:
  - `user_id`
  - `asset`
  - `delta_available`
  - `delta_frozen`
  - `reference_type`
  - `reference_id`
  - `event_id` unique
  - `created_at`

## Orders and trades

### `orders`

- Owner: order write path plus query projection
- Purpose: durable order record and status tracking.
- Primary key: `order_id`
- Phase 1 note: this becomes the first durable table used by the order application flow.
- Important fields:
  - `client_order_id`
  - `user_id`
  - `symbol`
  - `side`
  - `type`
  - `price`
  - `quantity`
  - `filled_quantity`
  - `status`
  - `rejection_reason`
  - `created_at`
  - `updated_at`

### `trades`

- Owner: trade projection
- Purpose: durable execution history for query and audit.
- Primary key: `trade_id`
- Important fields:
  - `symbol`
  - `maker_order_id`
  - `taker_order_id`
  - `maker_user_id`
  - `taker_user_id`
  - `price`
  - `quantity`
  - `maker_fee`
  - `taker_fee`
  - `executed_at`

## Infrastructure support

### `processed_events`

- Owner: idempotent consumers
- Purpose: dedupe marker for at-least-once delivery handling.
- Primary key: `(consumer_name, event_id)`
- Important fields:
  - `processed_at`

### `order_events`

- Owner: optional event-store mirror
- Purpose: searchable mirror of order-related events and payloads.
- Primary key: `event_id`
- Important fields:
  - `order_id`
  - `event_type`
  - `payload`
  - `occurred_at`
