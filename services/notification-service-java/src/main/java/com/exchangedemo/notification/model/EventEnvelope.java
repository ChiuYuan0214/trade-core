package com.exchangedemo.notification.model;

import com.fasterxml.jackson.databind.JsonNode;
import java.time.Instant;
import java.util.UUID;

public record EventEnvelope(
        UUID eventId,
        String eventType,
        Instant occurredAt,
        UUID correlationId,
        UUID causationId,
        String symbol,
        UUID orderId,
        UUID userId,
        String shardId,
        Integer version,
        JsonNode payload
) {
}
