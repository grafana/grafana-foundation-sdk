// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class Options {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("seriesMapping")
    public SeriesMapping seriesMapping;
    // Table Mode (auto)
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("dims")
    public XYDimensionConfig dims;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    // Manual Mode
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("series")
    public List<ScatterSeriesConfig> series;
    public Options() {
        this.dims = new com.grafana.foundation.xychart.XYDimensionConfigBuilder().build();
        this.legend = new com.grafana.foundation.common.VizLegendOptionsBuilder().build();
        this.tooltip = new com.grafana.foundation.common.VizTooltipOptionsBuilder().build();
        this.series = new LinkedList<>();
    }
    public Options(SeriesMapping seriesMapping,XYDimensionConfig dims,VizLegendOptions legend,VizTooltipOptions tooltip,List<ScatterSeriesConfig> series) {
        this.seriesMapping = seriesMapping;
        this.dims = dims;
        this.legend = legend;
        this.tooltip = tooltip;
        this.series = series;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
