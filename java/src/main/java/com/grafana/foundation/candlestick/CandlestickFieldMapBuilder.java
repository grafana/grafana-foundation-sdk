// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;


public class CandlestickFieldMapBuilder implements com.grafana.foundation.cog.Builder<CandlestickFieldMap> {
    protected final CandlestickFieldMap internal;
    
    public CandlestickFieldMapBuilder() {
        this.internal = new CandlestickFieldMap();
    }
    public CandlestickFieldMapBuilder open(String open) {
    this.internal.open = open;
        return this;
    }
    
    public CandlestickFieldMapBuilder high(String high) {
    this.internal.high = high;
        return this;
    }
    
    public CandlestickFieldMapBuilder low(String low) {
    this.internal.low = low;
        return this;
    }
    
    public CandlestickFieldMapBuilder close(String close) {
    this.internal.close = close;
        return this;
    }
    
    public CandlestickFieldMapBuilder volume(String volume) {
    this.internal.volume = volume;
        return this;
    }
    public CandlestickFieldMap build() {
        return this.internal;
    }
}
