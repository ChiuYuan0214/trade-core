package com.exchangedemo.marketdata.web;

import com.exchangedemo.marketdata.model.MarketDataSnapshot;
import com.exchangedemo.marketdata.service.MarketDataDispatchService;
import com.exchangedemo.marketdata.service.SseMarketDataPublisher;
import jakarta.validation.Valid;
import java.util.Collection;
import java.util.Map;
import java.util.Optional;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;
import org.springframework.web.servlet.mvc.method.annotation.SseEmitter;

@RestController
@RequestMapping
public class MarketDataController {

    private final MarketDataDispatchService dispatchService;
    private final SseMarketDataPublisher publisher;

    public MarketDataController(MarketDataDispatchService dispatchService, SseMarketDataPublisher publisher) {
        this.dispatchService = dispatchService;
        this.publisher = publisher;
    }

    @GetMapping("/api/v1/market-data")
    public Collection<MarketDataSnapshot> listSnapshots() {
        return dispatchService.list();
    }

    @GetMapping("/api/v1/market-data/{symbol}")
    public MarketDataSnapshot getSnapshot(@PathVariable String symbol) {
        Optional<MarketDataSnapshot> snapshot = dispatchService.get(symbol);
        return snapshot.orElseThrow(() -> new ResponseStatusException(HttpStatus.NOT_FOUND, "snapshot not found"));
    }

    @GetMapping("/api/v1/market-data/stream")
    public SseEmitter stream(@RequestParam String symbol) {
        return publisher.subscribe(symbol);
    }

    @PostMapping("/internal/market-data/snapshots")
    public ResponseEntity<?> ingest(@Valid @RequestBody MarketDataSnapshot snapshot) {
        MarketDataSnapshot stored = dispatchService.ingest(snapshot);
        return ResponseEntity.accepted().body(Map.of(
                "symbol", stored.symbol(),
                "stream", publisher.streamPath(stored.symbol()),
                "sequence", stored.sequence()
        ));
    }
}
