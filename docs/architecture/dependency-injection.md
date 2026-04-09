# Dependency Injection

This project uses [`depin`](https://github.com/ChiuYuan0214/depin) as its dependency-injection framework.

## Why

- It keeps process bootstrapping consistent across services.
- It encourages interface-based registration.
- It gives us a shared lifecycle model for startup and teardown.

## Repo conventions

- Use `depin` in `cmd/*` entrypoints for wiring service dependencies.
- Prefer `depin` field injection with `new(Service)` over constructor functions for service wiring.
- Prefer interfaces at service boundaries such as order orchestration, matching command publishing, ledger posting, projection updates, and websocket fan-out.
- Avoid treating DI as architecture; the real design still lives in package ownership and event flow.
- Prefer environment-backed config fields for process name, addresses, topics, and DSNs instead of constructor parameters when the values are deployment configuration.

## Practical rule

- Wiring belongs in bootstrap code.
- Business rules belong in domain and service packages.
- State ownership does not change just because dependencies are injected.
