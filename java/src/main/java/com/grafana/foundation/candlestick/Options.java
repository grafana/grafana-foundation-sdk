// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;

import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options {
    // Sets which dimensions are used for the visualization 
    @JsonProperty("mode")
    public VizDisplayMode mode;
    // Sets the style of the candlesticks 
    @JsonProperty("candleStyle")
    public CandleStyle candleStyle;
    // Sets the color strategy for the candlesticks 
    @JsonProperty("colorStrategy")
    public ColorStrategy colorStrategy;
    // Map fields to appropriate dimension 
    @JsonProperty("fields")
    public CandlestickFieldMap fields;
    // Set which colors are used when the price movement is up or down 
    @JsonProperty("colors")
    public CandlestickColors colors; 
    @JsonProperty("legend")
    public VizLegendOptions legend; 
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    // When enabled, all fields will be sent to the graph 
    @JsonProperty("includeAllFields")
    public Boolean includeAllFields;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}