# Trading System Docs Index

Use this file as the entry point for the repo.

## Core Navigation

- [architecture/system-overview.md](/Users/adam/trading_system/docs/architecture/system-overview.md): High-level architecture, state ownership, and execution model.
- [architecture/repo-structure.md](/Users/adam/trading_system/docs/architecture/repo-structure.md): Planned folder structure and why each area exists.
- [architecture/go-modules.md](/Users/adam/trading_system/docs/architecture/go-modules.md): Current Go workspace layout and how shared Go code is separated from service modules.
- [architecture/service-boundaries.md](/Users/adam/trading_system/docs/architecture/service-boundaries.md): Service responsibilities, write ownership, and integration edges.
- [architecture/dependency-injection.md](/Users/adam/trading_system/docs/architecture/dependency-injection.md): Why the project uses `depin` and how service wiring should be organized.
- [architecture/java-notification-service.md](/Users/adam/trading_system/docs/architecture/java-notification-service.md): Why `notification-service` is implemented in Java/Spring Boot and how it fits the polyglot design.
- [architecture/java-market-data-service.md](/Users/adam/trading_system/docs/architecture/java-market-data-service.md): Why `market-data-service` is implemented in Java/Spring Boot and how it fits the public read-side design.
- [architecture/matching-engine-design.md](/Users/adam/trading_system/docs/architecture/matching-engine-design.md): Current matching-book structure, shard ownership model, and upgrade path.
- [features/implementation-plan.md](/Users/adam/trading_system/docs/features/implementation-plan.md): Feature order, milestones, and discussion checkpoints.
- [features/java-notification-demo.md](/Users/adam/trading_system/docs/features/java-notification-demo.md): How to demo the Java Spring Boot notification service in the portfolio project.
- [features/java-market-data-demo.md](/Users/adam/trading_system/docs/features/java-market-data-demo.md): How to demo the Java Spring Boot market-data service in the portfolio project.
- [schema/tables.md](/Users/adam/trading_system/docs/schema/tables.md): PostgreSQL table design split by domain ownership.
- [schema/events.md](/Users/adam/trading_system/docs/schema/events.md): Command and domain event contracts.
- [index/method-map.md](/Users/adam/trading_system/docs/index/method-map.md): Lookup table for key methods, handlers, consumers, and their purpose.

## Usage Rules

- Read this file first, then load only the targeted child document.
- Update the nearest child doc when architecture, schema, feature sequence, or method ownership changes.
- Keep detailed rationale in the child files, not in this index.
- The reusable Codex skill lives under `~/.codex/skills`; this repo keeps the durable project conventions in `docs/` so they remain visible and portable.
