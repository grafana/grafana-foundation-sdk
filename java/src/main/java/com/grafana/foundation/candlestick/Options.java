// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.VizLegendOptions;

public class Options {
    // Sets which dimensions are used for the visualization
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public VizDisplayMode mode;
    // Sets the style of the candlesticks
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("candleStyle")
    public CandleStyle candleStyle;
    // Sets the color strategy for the candlesticks
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("colorStrategy")
    public ColorStrategy colorStrategy;
    // Map fields to appropriate dimension
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("fields")
    public CandlestickFieldMap fields;
    // Set which colors are used when the price movement is up or down
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("colors")
    public CandlestickColors colors;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    // When enabled, all fields will be sent to the graph
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("includeAllFields")
    public Boolean includeAllFields;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
