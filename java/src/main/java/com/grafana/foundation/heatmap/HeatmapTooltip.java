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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<HeatmapTooltip> {
        protected final HeatmapTooltip internal;
        
        public Builder() {
            this.internal = new HeatmapTooltip();
        }
    public Builder show(Boolean show) {
    this.internal.show = show;
        return this;
    }
    
    public Builder yHistogram(Boolean yHistogram) {
    this.internal.yHistogram = yHistogram;
        return this;
    }
    public HeatmapTooltip build() {
            return this.internal;
        }
    }
}
