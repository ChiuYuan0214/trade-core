package com.exchangedemo.marketdata.model;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import java.math.BigDecimal;
import java.time.Instant;

public record MarketDataSnapshot(
        @NotBlank String symbol,
        @NotNull BigDecimal bestBid,
        @NotNull BigDecimal bestAsk,
        @NotNull BigDecimal bidSize,
        @NotNull BigDecimal askSize,
        @NotNull BigDecimal lastPrice,
        long sequence,
        @NotNull Instant occurredAt,
        String source
) {
}
