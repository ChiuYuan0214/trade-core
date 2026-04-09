# Java Notification Demo

Use this flow to demo the Java/Spring Boot notification service as part of the portfolio project.

## Start the Java service

```bash
docker compose -f deployments/docker-compose.yml up -d kafka notification-service-java
```

## Verify the Java build locally

```bash
./scripts/java/test_notification_service.sh
```

The helper script resolves Java 21 and Maven explicitly so the module does not depend on whichever shell PATH happens to be active.

## Regenerate visible gRPC sources

```bash
cd services/notification-service-java
mvn generate-sources
```

Generated protobuf and gRPC classes will appear under `src/generated`, including `PrivateNotificationServiceGrpc.java`.

## Send a demo private event

```bash
./scripts/seed/send_private_event_demo.sh
```

## Connect Go order flow to Java notifications

Run the Go gateway with the Java gRPC address configured:

```bash
NOTIFICATION_GRPC_ADDR=127.0.0.1:10084 \
PROCESS_NAME=rest-gateway \
HTTP_ADDR=127.0.0.1:18080 \
ORDER_STORE_BACKEND=memory \
go run ./services/rest-gateway-go/cmd/rest-gateway
```

Once that is running, order placement, cancel, reserve/release, and trade settlement will emit private events from Go to the Java notification service over gRPC.

For a local gRPC-only demo without Kafka noise, start the Java service with:

```bash
EXCHANGE_NOTIFICATION_KAFKA_ENABLED=false \
java -jar services/notification-service-java/target/notification-service-java-0.1.0-SNAPSHOT.jar
```

## What this demonstrates

- A non-core exchange service can be implemented in Java/Spring Boot while the execution core stays in Go.
- The Java service can consume private events from Kafka or accept them from Go over gRPC, then route them to user-specific WebSocket topics.
- The architecture is polyglot by service boundary, not mixed arbitrarily inside the core matching path.

## Build note

- The module has been verified with Java 21 plus Maven using the helper script above.
