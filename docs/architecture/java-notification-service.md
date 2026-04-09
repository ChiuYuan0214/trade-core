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
- Concurrency model: Java 21 virtual threads for application tasks, HTTP fallback handling, and gRPC request execution
- Responsibility: consume private domain events and fan them out to user-specific channels

## Current interfaces

- Kafka topic consumption via `PrivateEventConsumer`
- WebSocket/STOMP endpoint at `/ws/private`
- gRPC service endpoint via `PrivateNotificationService.Publish`
- HTTP service endpoint at `POST /internal/notifications/events` kept only as a local fallback path

## Generated gRPC sources

- Java protobuf and gRPC sources are generated into `services/notification-service-java/src/generated`.
- The generated gRPC base class is `com.exchangedemo.notification.proto.PrivateNotificationServiceGrpc`.
- The concrete server implementation is `GrpcNotificationService`, which extends `PrivateNotificationServiceGrpc.PrivateNotificationServiceImplBase`.
- Regenerate them with `mvn generate-sources` inside the Java service module.

## Routing model

- `OrderAccepted`, `OrderRejected`, `OrderCanceled` -> user orders channel
- `TradeExecuted` -> user trades channel
- `FundsReserved`, `FundsReleased`, `LedgerPosted` -> user balances channel

## Portfolio talking point

The exchange core remains Go-first for low-level ownership domains, while Java/Spring Boot is used for a highly integrative edge service where framework productivity and ecosystem support are strong advantages.

The Go and Java processes stay independent. They communicate over service boundaries such as gRPC or Kafka; the project is polyglot by deployment boundary, not by mixing runtimes inside one component.

The Java service also demonstrates a modern Java 21 runtime style: request-oriented work runs on virtual threads so blocking-style service code stays straightforward without forcing the design into callback-heavy flow control.
