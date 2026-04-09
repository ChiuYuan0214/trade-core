# Java Notification Service

This repo implements `notification-service` as a Java Spring Boot service to make the project stronger as a polyglot portfolio piece.

## Why this service is the best Java candidate

- It is event-driven and boundary-oriented.
- It does not own the core matching state.
- It fits well with Spring Boot strengths: Kafka consumption, WebSocket fan-out, actuator endpoints, and operational visibility.
- It is easy to discuss as an example of mixed-language service architecture without compromising the single-writer matching core.

## Current module

- Module path: `services/notification-service-java`
- Stack: Spring Boot, Spring Kafka, Spring WebSocket, Actuator
- Responsibility: consume private domain events and fan them out to user-specific channels

## Current interfaces

- Kafka topic consumption via `PrivateEventConsumer`
- WebSocket/STOMP endpoint at `/ws/private`
- Demo ingest endpoint at `POST /internal/notifications/events`

## Routing model

- `OrderAccepted`, `OrderRejected`, `OrderCanceled` -> user orders channel
- `TradeExecuted` -> user trades channel
- `FundsReserved`, `FundsReleased`, `LedgerPosted` -> user balances channel

## Portfolio talking point

The exchange core remains Go-first for low-level ownership domains, while Java/Spring Boot is used for a highly integrative edge service where framework productivity and ecosystem support are strong advantages.
