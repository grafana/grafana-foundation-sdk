// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.candlestick;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;

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
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    // When enabled, all fields will be sent to the graph
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("includeAllFields")
    public Boolean includeAllFields;
    public Options() {
        this.mode = VizDisplayMode.CANDLES_VOLUME;
        this.candleStyle = CandleStyle.CANDLES;
        this.colorStrategy = ColorStrategy.OPEN_CLOSE;
        this.fields = new com.grafana.foundation.candlestick.CandlestickFieldMap();
        this.colors = new CandlestickColors("green", "red", "gray");
        this.legend = new com.grafana.foundation.common.VizLegendOptions();
        this.tooltip = new com.grafana.foundation.common.VizTooltipOptions();
        this.includeAllFields = false;
    }
    public Options(VizDisplayMode mode,CandleStyle candleStyle,ColorStrategy colorStrategy,CandlestickFieldMap fields,CandlestickColors colors,VizLegendOptions legend,VizTooltipOptions tooltip,Boolean includeAllFields) {
        this.mode = mode;
        this.candleStyle = candleStyle;
        this.colorStrategy = colorStrategy;
        this.fields = fields;
        this.colors = colors;
        this.legend = legend;
        this.tooltip = tooltip;
        this.includeAllFields = includeAllFields;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
