// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.TooltipDisplayMode;

// Controls tooltip options
public class HeatmapTooltip {
    // Controls how the tooltip is shown
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public TooltipDisplayMode mode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxHeight")
    public Double maxHeight;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxWidth")
    public Double maxWidth;
    // Controls if the tooltip shows a histogram of the y-axis values
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("yHistogram")
    public Boolean yHistogram;
    // Controls if the tooltip shows a color scale in header
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showColorScale")
    public Boolean showColorScale;
    public HeatmapTooltip() {
    }
    
    public HeatmapTooltip(TooltipDisplayMode mode,Double maxHeight,Double maxWidth,Boolean yHistogram,Boolean showColorScale) {
        this.mode = mode;
        this.maxHeight = maxHeight;
        this.maxWidth = maxWidth;
        this.yHistogram = yHistogram;
        this.showColorScale = showColorScale;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
