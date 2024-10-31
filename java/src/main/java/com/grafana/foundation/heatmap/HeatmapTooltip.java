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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<HeatmapTooltip> {
        protected final HeatmapTooltip internal;
        
        public Builder() {
            this.internal = new HeatmapTooltip();
        }
    public Builder mode(TooltipDisplayMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder maxHeight(Double maxHeight) {
    this.internal.maxHeight = maxHeight;
        return this;
    }
    
    public Builder maxWidth(Double maxWidth) {
    this.internal.maxWidth = maxWidth;
        return this;
    }
    
    public Builder yHistogram(Boolean yHistogram) {
    this.internal.yHistogram = yHistogram;
        return this;
    }
    
    public Builder showColorScale(Boolean showColorScale) {
    this.internal.showColorScale = showColorScale;
        return this;
    }
    public HeatmapTooltip build() {
            return this.internal;
        }
    }
}
