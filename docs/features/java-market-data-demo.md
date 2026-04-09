# Java Market Data Demo

Use this flow to demo the Java/Spring Boot market-data service as part of the portfolio project.

## Verify the Java build locally

```bash
./scripts/java/test_market_data_service.sh
```

## Start the Java service

```bash
docker compose -f deployments/docker-compose.yml up -d redis market-data-service-java
```

## Push a demo snapshot

```bash
curl -sS -X POST http://127.0.0.1:8095/internal/market-data/snapshots \
  -H 'Content-Type: application/json' \
  -d '{
    "symbol":"BTC/USDT",
    "bestBid":"60000",
    "bestAsk":"60010",
    "bidSize":"1.5",
    "askSize":"0.8",
    "lastPrice":"60005",
    "sequence":42,
    "occurredAt":"2026-04-10T00:00:00Z",
    "source":"demo"
  }'
```

## Read the snapshot back

```bash
curl -sS http://127.0.0.1:8095/api/v1/market-data/BTC%2FUSDT
```

## Stream updates publicly

```bash
curl -N http://127.0.0.1:8095/api/v1/market-data/stream?symbol=BTC/USDT
```

## What this demonstrates

- A second non-core exchange service can live in Java without weakening the Go-first execution core.
- The Java service uses Java 21 virtual threads for public request and stream handling.
- The service is ready to sit behind Redis pub/sub for public market-data fan-out.
