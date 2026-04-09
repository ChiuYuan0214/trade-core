# Java Notification Demo

Use this flow to demo the Java/Spring Boot notification service as part of the portfolio project.

## Start the Java service

```bash
docker compose -f deployments/docker-compose.yml up -d kafka notification-service-java
```

## Send a demo private event

```bash
./scripts/seed/send_private_event_demo.sh
```

## What this demonstrates

- A non-core exchange service can be implemented in Java/Spring Boot while the execution core stays in Go.
- The Java service can consume or ingest private events and route them to user-specific WebSocket topics.
- The architecture is polyglot by service boundary, not mixed arbitrarily inside the core matching path.

## Current caveat

- In this environment, Maven resolution to Spring Boot artifacts may fail intermittently even though the module source is valid. The repo still includes the full module, Dockerfile, and runtime wiring for portfolio presentation and future local builds.
