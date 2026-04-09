package com.exchangedemo.marketdata.service;

import com.exchangedemo.marketdata.model.MarketDataSnapshot;

public interface MarketDataPublisher {

    void publish(MarketDataSnapshot snapshot);
}
