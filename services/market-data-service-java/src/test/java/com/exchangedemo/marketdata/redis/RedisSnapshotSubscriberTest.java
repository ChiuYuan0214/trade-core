package com.exchangedemo.marketdata.redis;

import com.exchangedemo.marketdata.model.MarketDataSnapshot;
import com.exchangedemo.marketdata.service.MarketDataDispatchService;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.math.BigDecimal;
import java.nio.charset.StandardCharsets;
import java.time.Instant;
import org.junit.jupiter.api.Test;
import org.springframework.data.redis.connection.DefaultMessage;
import org.mockito.ArgumentCaptor;
import org.springframework.data.redis.connection.Message;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.verifyNoMoreInteractions;

class RedisSnapshotSubscriberTest {

    private final ObjectMapper objectMapper = new ObjectMapper().findAndRegisterModules();

    @Test
    void onMessageDeserializesSnapshotAndDispatchesIt() throws Exception {
        MarketDataDispatchService dispatchService = mock(MarketDataDispatchService.class);
        RedisSnapshotSubscriber subscriber = new RedisSnapshotSubscriber(objectMapper, dispatchService);

        MarketDataSnapshot snapshot = new MarketDataSnapshot(
                "BTC/USDT",
                new BigDecimal("60000"),
                new BigDecimal("60010"),
                new BigDecimal("1.5"),
                new BigDecimal("0.8"),
                new BigDecimal("60005"),
                42L,
                Instant.parse("2026-04-10T00:00:00Z"),
                "redis"
        );

        Message message = new DefaultMessage(
                "exchange.market-data.snapshots".getBytes(StandardCharsets.UTF_8),
                objectMapper.writeValueAsBytes(snapshot)
        );

        subscriber.onMessage(message, null);

        ArgumentCaptor<MarketDataSnapshot> captor = ArgumentCaptor.forClass(MarketDataSnapshot.class);
        verify(dispatchService, times(1)).ingest(captor.capture());
        assertEquals("BTC/USDT", captor.getValue().symbol());
        assertEquals(42L, captor.getValue().sequence());
    }

    @Test
    void onMessageIgnoresInvalidPayload() {
        MarketDataDispatchService dispatchService = mock(MarketDataDispatchService.class);
        RedisSnapshotSubscriber subscriber = new RedisSnapshotSubscriber(objectMapper, dispatchService);
        Message message = new DefaultMessage(
                "exchange.market-data.snapshots".getBytes(StandardCharsets.UTF_8),
                "not-json".getBytes(StandardCharsets.UTF_8)
        );

        subscriber.onMessage(message, null);

        verifyNoMoreInteractions(dispatchService);
    }
}
