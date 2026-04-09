package com.exchangedemo.notification.service;

import com.exchangedemo.notification.config.NotificationProperties;
import com.exchangedemo.notification.model.EventEnvelope;
import com.exchangedemo.notification.model.NotificationChannel;
import com.exchangedemo.notification.model.UserNotification;
import java.util.Optional;
import org.springframework.stereotype.Component;

@Component
public class NotificationRouter {

    private final NotificationProperties properties;

    public NotificationRouter(NotificationProperties properties) {
        this.properties = properties;
    }

    public Optional<UserNotification> route(EventEnvelope envelope) {
        if (envelope.userId() == null) {
            return Optional.empty();
        }

        NotificationChannel channel = switch (envelope.eventType()) {
            case "OrderAccepted", "OrderRejected", "OrderCanceled" -> NotificationChannel.USER_ORDERS;
            case "TradeExecuted" -> NotificationChannel.USER_TRADES;
            case "FundsReserved", "FundsReleased", "LedgerPosted" -> NotificationChannel.USER_BALANCES;
            default -> null;
        };

        if (channel == null) {
            return Optional.empty();
        }

        String suffix = switch (channel) {
            case USER_ORDERS -> "orders";
            case USER_TRADES -> "trades";
            case USER_BALANCES -> "balances";
        };

        return Optional.of(new UserNotification(
                envelope.userId(),
                channel,
                properties.websocketPrefix() + "/" + envelope.userId() + "/" + suffix,
                envelope.eventType(),
                envelope.orderId(),
                envelope.symbol(),
                envelope.occurredAt(),
                envelope.payload()
        ));
    }
}
