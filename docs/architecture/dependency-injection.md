# Dependency Injection

This project uses [`depin`](https://github.com/ChiuYuan0214/depin) as its dependency-injection framework.

## Why

- It keeps process bootstrapping consistent across services.
- It encourages interface-based registration.
- It gives us a shared lifecycle model for startup and teardown.

## Repo conventions

- Use `depin` in service-module entrypoints for wiring service dependencies.
- Current Go entrypoints live under `services/*-go/cmd/*`.
- Keep one bootstrap file per service process under `modules/exchange-core-go/bootstrap/`.
- Keep shared bootstrap helpers small and generic; do not centralize all service wiring in one switch-heavy file.
- Prefer `depin` field injection with `new(Service)` over constructor functions for service wiring.
- Prefer interfaces at service boundaries such as order orchestration, matching command publishing, ledger posting, projection updates, and websocket fan-out.
- Avoid treating DI as architecture; the real design still lives in package ownership and event flow.
- Prefer environment-backed config fields for process name, addresses, topics, and DSNs instead of constructor parameters when the values are deployment configuration.
- Backend selection such as `ORDER_STORE_BACKEND` and `LEDGER_STORE_BACKEND` should stay in config and be resolved inside the owning service bootstrap.
- For demo-only in-memory state, prefer a dedicated shared package with `sync.Once` guarded constructors over rebuilding the same stores in multiple bootstrap files.

## Practical rule

- Wiring belongs in bootstrap code.
- Business rules belong in domain and service packages.
- State ownership does not change just because dependencies are injected.
