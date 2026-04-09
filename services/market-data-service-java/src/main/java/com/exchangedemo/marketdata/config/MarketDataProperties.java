package com.exchangedemo.marketdata.config;

import org.springframework.boot.context.properties.ConfigurationProperties;

@ConfigurationProperties(prefix = "exchange.market-data")
public record MarketDataProperties(
        String streamPrefix,
        String redisChannel,
        boolean redisEnabled
) {
}
