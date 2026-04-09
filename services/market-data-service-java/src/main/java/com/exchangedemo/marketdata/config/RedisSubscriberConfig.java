package com.exchangedemo.marketdata.config;

import com.exchangedemo.marketdata.redis.RedisSnapshotSubscriber;
import com.exchangedemo.marketdata.service.MarketDataDispatchService;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.listener.PatternTopic;
import org.springframework.data.redis.listener.RedisMessageListenerContainer;

@Configuration
@ConditionalOnProperty(prefix = "exchange.market-data", name = "redis-enabled", havingValue = "true")
public class RedisSubscriberConfig {

    @Bean
    public RedisSnapshotSubscriber redisSnapshotSubscriber(
            ObjectMapper objectMapper,
            MarketDataDispatchService dispatchService
    ) {
        return new RedisSnapshotSubscriber(objectMapper, dispatchService);
    }

    @Bean
    public RedisMessageListenerContainer redisMessageListenerContainer(
            RedisConnectionFactory connectionFactory,
            RedisSnapshotSubscriber redisSnapshotSubscriber,
            MarketDataProperties properties
    ) {
        RedisMessageListenerContainer container = new RedisMessageListenerContainer();
        container.setConnectionFactory(connectionFactory);
        container.addMessageListener(redisSnapshotSubscriber, new PatternTopic(properties.redisChannel()));
        return container;
    }
}
