# System Overview

This project is a portfolio-grade exchange core demo focused on correctness, replayability, and explainable service boundaries.

## Core principles

- Matching engine owns live in-memory order book state.
- Ledger/account domain owns available and frozen balances through append-only ledger entries.
- Query models own read APIs and projections, not core business truth.
- Kafka is the durable ordered backbone for commands and events.
- PostgreSQL is the durable business record.
- Redis is cache and fan-out infrastructure only.

## Initial architecture shape

Planned services:
- REST Gateway
- WebSocket Gateway
- Order Service
- Matching Engine
- Ledger Service
- Market Data Service
- Notification Service
- Replay Tool

## Delivery strategy

The implementation will begin with strict package boundaries inside one repo and incrementally add infra complexity by phase. Early slices may run a smaller subset of processes, but code boundaries should already reflect the intended service ownership.
