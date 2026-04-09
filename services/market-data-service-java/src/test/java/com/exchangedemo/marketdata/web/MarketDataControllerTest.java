package com.exchangedemo.marketdata.web;

import com.exchangedemo.marketdata.model.MarketDataSnapshot;
import com.exchangedemo.marketdata.service.MarketDataDispatchService;
import com.exchangedemo.marketdata.service.SseMarketDataPublisher;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.math.BigDecimal;
import java.time.Instant;
import java.util.List;
import java.util.Optional;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.web.servlet.mvc.method.annotation.SseEmitter;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.request;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(MarketDataController.class)
class MarketDataControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private ObjectMapper objectMapper;

    @MockBean
    private MarketDataDispatchService dispatchService;

    @MockBean
    private SseMarketDataPublisher publisher;

    @Test
    void listSnapshotsReturnsStoredSnapshots() throws Exception {
        when(dispatchService.list()).thenReturn(List.of(snapshot()));

        mockMvc.perform(get("/api/v1/market-data"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$[0].symbol").value("BTC/USDT"))
                .andExpect(jsonPath("$[0].sequence").value(42));
    }

    @Test
    void getSnapshotReturnsSingleSymbolView() throws Exception {
        MarketDataSnapshot snapshot = new MarketDataSnapshot(
                "BTC-USDT",
                new BigDecimal("60000"),
                new BigDecimal("60010"),
                new BigDecimal("1.5"),
                new BigDecimal("0.8"),
                new BigDecimal("60005"),
                42L,
                Instant.parse("2026-04-10T00:00:00Z"),
                "demo"
        );
        when(dispatchService.get("BTC-USDT")).thenReturn(Optional.of(snapshot));

        mockMvc.perform(get("/api/v1/market-data/BTC-USDT"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.bestBid").value(60000))
                .andExpect(jsonPath("$.bestAsk").value(60010));
    }

    @Test
    void ingestReturnsAcceptedSnapshotMetadata() throws Exception {
        when(dispatchService.ingest(any(MarketDataSnapshot.class))).thenReturn(snapshot());
        when(publisher.streamPath(eq("BTC/USDT"))).thenReturn("/api/v1/market-data/stream?symbol=BTC/USDT");

        mockMvc.perform(post("/internal/market-data/snapshots")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsBytes(snapshot())))
                .andExpect(status().isAccepted())
                .andExpect(jsonPath("$.symbol").value("BTC/USDT"))
                .andExpect(jsonPath("$.stream").value("/api/v1/market-data/stream?symbol=BTC/USDT"))
                .andExpect(jsonPath("$.sequence").value(42));
    }

    @Test
    void streamOpensSseEndpoint() throws Exception {
        when(publisher.subscribe("BTC/USDT")).thenReturn(new SseEmitter());

        mockMvc.perform(get("/api/v1/market-data/stream").param("symbol", "BTC/USDT"))
                .andExpect(status().isOk())
                .andExpect(request().asyncStarted());
    }

    private static MarketDataSnapshot snapshot() {
        return new MarketDataSnapshot(
                "BTC/USDT",
                new BigDecimal("60000"),
                new BigDecimal("60010"),
                new BigDecimal("1.5"),
                new BigDecimal("0.8"),
                new BigDecimal("60005"),
                42L,
                Instant.parse("2026-04-10T00:00:00Z"),
                "demo"
        );
    }
}
