package com.exchangedemo.marketdata.service;

import com.exchangedemo.marketdata.config.MarketDataProperties;
import com.exchangedemo.marketdata.model.MarketDataSnapshot;
import java.io.IOException;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.CopyOnWriteArrayList;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.mvc.method.annotation.SseEmitter;

@Component
public class SseMarketDataPublisher implements MarketDataPublisher {

    private static final Logger logger = LoggerFactory.getLogger(SseMarketDataPublisher.class);

    private final MarketDataProperties properties;
    private final Map<String, CopyOnWriteArrayList<SseEmitter>> emitters = new ConcurrentHashMap<>();

    public SseMarketDataPublisher(MarketDataProperties properties) {
        this.properties = properties;
    }

    public SseEmitter subscribe(String symbol) {
        SseEmitter emitter = new SseEmitter(0L);
        emitters.computeIfAbsent(symbol, ignored -> new CopyOnWriteArrayList<>()).add(emitter);
        emitter.onCompletion(() -> remove(symbol, emitter));
        emitter.onTimeout(() -> remove(symbol, emitter));
        emitter.onError(ignored -> remove(symbol, emitter));
        return emitter;
    }

    @Override
    public void publish(MarketDataSnapshot snapshot) {
        List<SseEmitter> symbolEmitters = emitters.get(snapshot.symbol());
        if (symbolEmitters == null || symbolEmitters.isEmpty()) {
            return;
        }

        for (SseEmitter emitter : symbolEmitters) {
            try {
                emitter.send(SseEmitter.event()
                        .name("market-data")
                        .id(Long.toString(snapshot.sequence()))
                        .data(snapshot));
            } catch (IOException error) {
                logger.debug("removing broken market data emitter symbol={}", snapshot.symbol(), error);
                remove(snapshot.symbol(), emitter);
            }
        }
    }

    public String streamPath(String symbol) {
        return properties.streamPrefix() + "?symbol=" + symbol;
    }

    private void remove(String symbol, SseEmitter emitter) {
        List<SseEmitter> symbolEmitters = emitters.get(symbol);
        if (symbolEmitters == null) {
            return;
        }
        symbolEmitters.remove(emitter);
        if (symbolEmitters.isEmpty()) {
            emitters.remove(symbol);
        }
    }
}
