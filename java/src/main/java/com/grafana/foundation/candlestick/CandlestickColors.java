// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class CandlestickColors {
    @JsonProperty("up")
    public String up;
    @JsonProperty("down")
    public String down;
    @JsonProperty("flat")
    public String flat;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CandlestickColors> {
        protected final CandlestickColors internal;
        
        public Builder() {
            this.internal = new CandlestickColors();
        this.up("green");
        this.down("red");
        this.flat("gray");
        }
    public Builder up(String up) {
    this.internal.up = up;
        return this;
    }
    
    public Builder down(String down) {
    this.internal.down = down;
        return this;
    }
    
    public Builder flat(String flat) {
    this.internal.flat = flat;
        return this;
    }
    public CandlestickColors build() {
            return this.internal;
        }
    }
}
