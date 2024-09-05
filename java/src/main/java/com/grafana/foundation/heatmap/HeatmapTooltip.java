// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.TooltipDisplayMode;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Controls tooltip options
public class HeatmapTooltip {
    // Controls how the tooltip is shown
    @JsonProperty("mode")
    public TooltipDisplayMode mode;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("maxHeight")
    public Double maxHeight;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("maxWidth")
    public Double maxWidth;
    // Controls if the tooltip shows a histogram of the y-axis values
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("yHistogram")
    public Boolean yHistogram;
    // Controls if the tooltip shows a color scale in header
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("showColorScale")
    public Boolean showColorScale;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
