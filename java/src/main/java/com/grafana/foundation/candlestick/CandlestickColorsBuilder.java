// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;


public class CandlestickColorsBuilder implements com.grafana.foundation.cog.Builder<CandlestickColors> {
    protected final CandlestickColors internal;
    
    public CandlestickColorsBuilder() {
        this.internal = new CandlestickColors();
    }
    public CandlestickColorsBuilder up(String up) {
    this.internal.up = up;
        return this;
    }
    
    public CandlestickColorsBuilder down(String down) {
    this.internal.down = down;
        return this;
    }
    
    public CandlestickColorsBuilder flat(String flat) {
    this.internal.flat = flat;
        return this;
    }
    public CandlestickColors build() {
        return this.internal;
    }
}
