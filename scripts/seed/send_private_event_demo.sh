#!/usr/bin/env bash
set -euo pipefail

BASE_URL="${1:-http://127.0.0.1:8094}"

curl -sS -X POST "${BASE_URL}/internal/notifications/events" \
  -H 'Content-Type: application/json' \
  -d '{
    "eventId": "11111111-aaaa-bbbb-cccc-111111111111",
    "eventType": "TradeExecuted",
    "occurredAt": "2026-04-09T12:00:00Z",
    "correlationId": "22222222-aaaa-bbbb-cccc-222222222222",
    "causationId": "33333333-aaaa-bbbb-cccc-333333333333",
    "symbol": "BTC/USDT",
    "orderId": "44444444-aaaa-bbbb-cccc-444444444444",
    "userId": "11111111-1111-1111-1111-111111111111",
    "shardId": "shard-1",
    "version": 1,
    "payload": {
      "tradeId": "55555555-aaaa-bbbb-cccc-555555555555",
      "price": "60000",
      "quantity": "0.25"
    }
  }'

printf '\n'
