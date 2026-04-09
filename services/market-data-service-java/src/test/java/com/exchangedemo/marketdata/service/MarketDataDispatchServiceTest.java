package com.exchangedemo.marketdata.service;

import com.exchangedemo.marketdata.model.MarketDataSnapshot;
import java.math.BigDecimal;
import java.time.Instant;
import java.util.ArrayList;
import java.util.List;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

class MarketDataDispatchServiceTest {

    @Test
    void ingestStoresSnapshotAndPublishesIt() {
        MarketDataSnapshotStore store = new MarketDataSnapshotStore();
        RecordingPublisher publisher = new RecordingPublisher();
        MarketDataDispatchService service = new MarketDataDispatchService(store, publisher);

        MarketDataSnapshot snapshot = new MarketDataSnapshot(
                "BTC/USDT",
                new BigDecimal("60000"),
                new BigDecimal("60010"),
                new BigDecimal("1.5"),
                new BigDecimal("0.8"),
                new BigDecimal("60005"),
                42L,
                Instant.parse("2026-04-10T00:00:00Z"),
                "demo"
        );

        service.ingest(snapshot);

        assertTrue(service.get("BTC/USDT").isPresent());
        assertEquals(1, publisher.snapshots.size());
        assertEquals("BTC/USDT", publisher.snapshots.getFirst().symbol());
    }

    private static final class RecordingPublisher implements MarketDataPublisher {
        private final List<MarketDataSnapshot> snapshots = new ArrayList<>();

        @Override
        public void publish(MarketDataSnapshot snapshot) {
            snapshots.add(snapshot);
        }
    }
}
