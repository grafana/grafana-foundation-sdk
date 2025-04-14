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
    public CandlestickColors() {
        this.up = "green";
        this.down = "red";
        this.flat = "gray";
    }
    public CandlestickColors(String up,String down,String flat) {
        this.up = up;
        this.down = down;
        this.flat = flat;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
