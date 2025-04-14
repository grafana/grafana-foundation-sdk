// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Controls tooltip options
public class HeatmapTooltip {
    // Controls if the tooltip is shown
    @JsonProperty("show")
    public Boolean show;
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
    public HeatmapTooltip(Boolean show,Boolean yHistogram,Boolean showColorScale) {
        this.show = show;
        this.yHistogram = yHistogram;
        this.showColorScale = showColorScale;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
