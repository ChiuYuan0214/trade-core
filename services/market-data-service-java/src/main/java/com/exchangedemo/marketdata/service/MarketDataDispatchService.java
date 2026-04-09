package com.exchangedemo.marketdata.service;

import com.exchangedemo.marketdata.model.MarketDataSnapshot;
import java.util.Collection;
import java.util.Optional;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

@Service
public class MarketDataDispatchService {

    private static final Logger logger = LoggerFactory.getLogger(MarketDataDispatchService.class);

    private final MarketDataSnapshotStore snapshotStore;
    private final MarketDataPublisher publisher;

    public MarketDataDispatchService(MarketDataSnapshotStore snapshotStore, MarketDataPublisher publisher) {
        this.snapshotStore = snapshotStore;
        this.publisher = publisher;
    }

    public MarketDataSnapshot ingest(MarketDataSnapshot snapshot) {
        MarketDataSnapshot stored = snapshotStore.save(snapshot);
        publisher.publish(stored);
        logger.info("published market data symbol={} sequence={} bid={} ask={}",
                stored.symbol(), stored.sequence(), stored.bestBid(), stored.bestAsk());
        return stored;
    }

    public Optional<MarketDataSnapshot> get(String symbol) {
        return snapshotStore.get(symbol);
    }

    public Collection<MarketDataSnapshot> list() {
        return snapshotStore.list();
    }
}
