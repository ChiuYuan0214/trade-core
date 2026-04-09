# Matching Engine Design

This document explains the current Phase 1 matching structure and the intended upgrade path.

## Current Phase 1 structure

- One in-memory book per symbol shard.
- One active shard wrapper serializes mutations for that symbol.
- A single process may host multiple shards.
- Matching uses price-time priority.
- Cancel lookup uses `order_id`.

## Why multiple shards can exist in one process

- The critical rule is one active writer per symbol shard, not one machine per shard.
- In early phases and many real systems, one process can host multiple low-traffic symbols.
- Hot symbols should still be able to move to a dedicated shard owner later.

## Why the current book is not keyed by UUID for matching

- UUID lookup is only for locating a resting order quickly during cancel or state updates.
- Matching itself is driven by price ordering on the buy and sell sides.
- The current implementation uses sorted resting slices because they are easy to reason about and good for a demo/interview codebase.

## Planned upgrade path

If throughput or cancel efficiency becomes a bottleneck, evolve the book into:

- `price -> price level`
- FIFO queue inside each price level for time priority
- `order_id -> node/level reference` for fast cancel/update

That upgrade preserves the same business rules while improving lookup and mutation cost.
