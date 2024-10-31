// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class CandlestickFieldMap {
    // Corresponds to the starting value of the given period
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("open")
    public String open;
    // Corresponds to the highest value of the given period
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("high")
    public String high;
    // Corresponds to the lowest value of the given period
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("low")
    public String low;
    // Corresponds to the final (end) value of the given period
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("close")
    public String close;
    // Corresponds to the sample count in the given period. (e.g. number of trades)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("volume")
    public String volume;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CandlestickFieldMap> {
        protected final CandlestickFieldMap internal;
        
        public Builder() {
            this.internal = new CandlestickFieldMap();
        }
    public Builder open(String open) {
    this.internal.open = open;
        return this;
    }
    
    public Builder high(String high) {
    this.internal.high = high;
        return this;
    }
    
    public Builder low(String low) {
    this.internal.low = low;
        return this;
    }
    
    public Builder close(String close) {
    this.internal.close = close;
        return this;
    }
    
    public Builder volume(String volume) {
    this.internal.volume = volume;
        return this;
    }
    public CandlestickFieldMap build() {
            return this.internal;
        }
    }
}
