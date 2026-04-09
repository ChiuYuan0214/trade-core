package com.exchangedemo.marketdata.redis;

import com.exchangedemo.marketdata.model.MarketDataSnapshot;
import com.exchangedemo.marketdata.service.MarketDataDispatchService;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.nio.charset.StandardCharsets;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.data.redis.connection.Message;
import org.springframework.data.redis.connection.MessageListener;

public class RedisSnapshotSubscriber implements MessageListener {

    private static final Logger logger = LoggerFactory.getLogger(RedisSnapshotSubscriber.class);

    private final ObjectMapper objectMapper;
    private final MarketDataDispatchService dispatchService;

    public RedisSnapshotSubscriber(ObjectMapper objectMapper, MarketDataDispatchService dispatchService) {
        this.objectMapper = objectMapper;
        this.dispatchService = dispatchService;
    }

    @Override
    public void onMessage(Message message, byte[] pattern) {
        try {
            MarketDataSnapshot snapshot = objectMapper.readValue(message.getBody(), MarketDataSnapshot.class);
            dispatchService.ingest(snapshot);
        } catch (Exception error) {
            logger.warn("failed to consume redis market-data message payload={}",
                    new String(message.getBody(), StandardCharsets.UTF_8), error);
        }
    }
}
