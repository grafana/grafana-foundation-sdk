// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class CandlestickFieldMap {
    // Corresponds to the starting value of the given period 
    @JsonProperty("open")
    public String open;
    // Corresponds to the highest value of the given period 
    @JsonProperty("high")
    public String high;
    // Corresponds to the lowest value of the given period 
    @JsonProperty("low")
    public String low;
    // Corresponds to the final (end) value of the given period 
    @JsonProperty("close")
    public String close;
    // Corresponds to the sample count in the given period. (e.g. number of trades) 
    @JsonProperty("volume")
    public String volume;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
