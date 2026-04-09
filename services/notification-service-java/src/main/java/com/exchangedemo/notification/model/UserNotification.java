package com.exchangedemo.notification.model;

import java.time.Instant;
import java.util.UUID;

public record UserNotification(
        UUID userId,
        NotificationChannel channel,
        String topic,
        String eventType,
        UUID orderId,
        String symbol,
        Instant occurredAt,
        Object payload
) {
}
