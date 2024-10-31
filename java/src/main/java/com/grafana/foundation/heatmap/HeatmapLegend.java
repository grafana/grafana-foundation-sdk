// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Controls legend options
public class HeatmapLegend {
    // Controls if the legend is shown
    @JsonProperty("show")
    public Boolean show;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<HeatmapLegend> {
        protected final HeatmapLegend internal;
        
        public Builder() {
            this.internal = new HeatmapLegend();
        }
    public Builder show(Boolean show) {
    this.internal.show = show;
        return this;
    }
    public HeatmapLegend build() {
            return this.internal;
        }
    }
}
