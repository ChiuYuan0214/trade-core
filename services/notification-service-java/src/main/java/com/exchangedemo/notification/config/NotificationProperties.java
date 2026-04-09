package com.exchangedemo.notification.config;

import org.springframework.boot.context.properties.ConfigurationProperties;

@ConfigurationProperties(prefix = "exchange.notification")
public record NotificationProperties(
        String eventTopic,
        String websocketPrefix
) {
}
