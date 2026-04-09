CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS accounts (
    user_id UUID NOT NULL,
    asset VARCHAR(32) NOT NULL,
    available_balance NUMERIC(36, 18) NOT NULL,
    frozen_balance NUMERIC(36, 18) NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (user_id, asset)
);

CREATE TABLE IF NOT EXISTS orders (
    order_id UUID PRIMARY KEY,
    client_order_id VARCHAR(128),
    user_id UUID NOT NULL,
    symbol VARCHAR(32) NOT NULL,
    side VARCHAR(16) NOT NULL,
    type VARCHAR(16) NOT NULL,
    price NUMERIC(36, 18),
    quantity NUMERIC(36, 18) NOT NULL,
    filled_quantity NUMERIC(36, 18) NOT NULL DEFAULT 0,
    status VARCHAR(32) NOT NULL,
    rejection_reason TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id_created_at
    ON orders (user_id, created_at DESC);

CREATE INDEX IF NOT EXISTS idx_orders_symbol_status_created_at
    ON orders (symbol, status, created_at DESC);

CREATE TABLE IF NOT EXISTS trades (
    trade_id UUID PRIMARY KEY,
    symbol VARCHAR(32) NOT NULL,
    maker_order_id UUID NOT NULL,
    taker_order_id UUID NOT NULL,
    maker_user_id UUID NOT NULL,
    taker_user_id UUID NOT NULL,
    price NUMERIC(36, 18) NOT NULL,
    quantity NUMERIC(36, 18) NOT NULL,
    maker_fee NUMERIC(36, 18) NOT NULL,
    taker_fee NUMERIC(36, 18) NOT NULL,
    executed_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS ledger_entries (
    entry_id UUID PRIMARY KEY,
    user_id UUID,
    asset VARCHAR(32) NOT NULL,
    delta_available NUMERIC(36, 18) NOT NULL,
    delta_frozen NUMERIC(36, 18) NOT NULL,
    reference_type VARCHAR(64) NOT NULL,
    reference_id UUID NOT NULL,
    event_id UUID NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS processed_events (
    consumer_name VARCHAR(128) NOT NULL,
    event_id UUID NOT NULL,
    processed_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (consumer_name, event_id)
);

CREATE TABLE IF NOT EXISTS order_events (
    event_id UUID PRIMARY KEY,
    order_id UUID,
    event_type VARCHAR(64) NOT NULL,
    payload JSONB NOT NULL,
    occurred_at TIMESTAMPTZ NOT NULL
);
