package com.exchangedemo.notification.service;

import com.exchangedemo.notification.config.NotificationProperties;
import com.exchangedemo.notification.model.EventEnvelope;
import com.exchangedemo.notification.model.NotificationChannel;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.time.Instant;
import java.util.Optional;
import java.util.UUID;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

class NotificationRouterTest {

    private final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    void routesTradeExecutedToUserTradesChannel() {
        NotificationRouter router = new NotificationRouter(new NotificationProperties(
                "exchange.private.events",
                "/topic/private"
        ));

        EventEnvelope envelope = new EventEnvelope(
                UUID.randomUUID(),
                "TradeExecuted",
                Instant.parse("2026-04-09T00:00:00Z"),
                UUID.randomUUID(),
                UUID.randomUUID(),
                "BTC/USDT",
                UUID.randomUUID(),
                UUID.fromString("11111111-1111-1111-1111-111111111111"),
                "shard-1",
                1,
                objectMapper.createObjectNode().put("tradeId", UUID.randomUUID().toString())
        );

        Optional<com.exchangedemo.notification.model.UserNotification> routed = router.route(envelope);

        assertTrue(routed.isPresent());
        assertEquals(NotificationChannel.USER_TRADES, routed.get().channel());
        assertEquals("/topic/private/11111111-1111-1111-1111-111111111111/trades", routed.get().topic());
    }
}
