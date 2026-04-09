# Java Market Data Service

This repo implements `market-data-service` as a Java Spring Boot service to strengthen the public read-side part of the portfolio.

## Why this service is the best second Java candidate

- It is read-heavy and fan-out oriented.
- It is operationally adjacent to Redis and public subscriptions.
- It benefits from Spring Boot productivity for HTTP, SSE, Redis listeners, and Actuator.
- It complements the Java notification service without moving matching ownership out of Go.

## Current module

- Module path: `services/market-data-service-java`
- Stack: Spring Boot, Spring Data Redis, Actuator
- Concurrency model: Java 21 virtual threads for HTTP request handling and public stream fan-out
- Responsibility: hold the latest public market-data snapshot per symbol and publish updates to SSE subscribers

## Current interfaces

- SSE endpoint at `GET /api/v1/market-data/stream?symbol=...`
- Snapshot query at `GET /api/v1/market-data/{symbol}`
- Local ingest endpoint at `POST /internal/market-data/snapshots`
- Optional Redis pub/sub subscriber when `exchange.market-data.redis-enabled=true`

## Portfolio talking point

Go still owns the execution-critical write path. Java is used here for a public, read-optimized service where framework ergonomics, Redis integration, and Java 21 virtual-thread execution make the service easy to explain and scale independently.
