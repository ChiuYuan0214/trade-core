package com.exchangedemo.marketdata.service;

import com.exchangedemo.marketdata.model.MarketDataSnapshot;
import java.util.Collection;
import java.util.Comparator;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;
import org.springframework.stereotype.Component;

@Component
public class MarketDataSnapshotStore {

    private final ConcurrentMap<String, MarketDataSnapshot> snapshots = new ConcurrentHashMap<>();

    public MarketDataSnapshot save(MarketDataSnapshot snapshot) {
        snapshots.put(snapshot.symbol(), snapshot);
        return snapshot;
    }

    public Optional<MarketDataSnapshot> get(String symbol) {
        return Optional.ofNullable(snapshots.get(symbol));
    }

    public Collection<MarketDataSnapshot> list() {
        return snapshots.values().stream()
                .sorted(Comparator.comparing(MarketDataSnapshot::symbol))
                .toList();
    }
}
