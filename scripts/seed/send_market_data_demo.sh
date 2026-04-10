#!/usr/bin/env bash
set -euo pipefail

BASE_URL="${1:-http://127.0.0.1:8095}"
SYMBOL="${2:-BTC/USDT}"

curl -sS -X POST "${BASE_URL}/internal/market-data/snapshots" \
  -H 'Content-Type: application/json' \
  -d "{
    \"symbol\":\"${SYMBOL}\",
    \"bestBid\":\"60000\",
    \"bestAsk\":\"60010\",
    \"bidSize\":\"1.5\",
    \"askSize\":\"0.8\",
    \"lastPrice\":\"60005\",
    \"sequence\":42,
    \"occurredAt\":\"2026-04-10T00:00:00Z\",
    \"source\":\"seed-script\"
  }"

printf '\n'
