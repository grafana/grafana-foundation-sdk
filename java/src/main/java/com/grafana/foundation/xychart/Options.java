// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options { 
    @JsonProperty("seriesMapping")
    public SeriesMapping seriesMapping; 
    @JsonProperty("dims")
    public XYDimensionConfig dims; 
    @JsonProperty("legend")
    public VizLegendOptions legend; 
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip; 
    @JsonProperty("series")
    public List<ScatterSeriesConfig> series;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
