// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// TODO docs
public class HideSeriesConfig { 
    @JsonProperty("tooltip")
    public Boolean tooltip; 
    @JsonProperty("legend")
    public Boolean legend; 
    @JsonProperty("viz")
    public Boolean viz;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<HideSeriesConfig> {
        private final HideSeriesConfig internal;
        
        public Builder() {
            this.internal = new HideSeriesConfig();
        }
    public Builder tooltip(Boolean tooltip) {
    this.internal.tooltip = tooltip;
        return this;
    }
    
    public Builder legend(Boolean legend) {
    this.internal.legend = legend;
        return this;
    }
    
    public Builder viz(Boolean viz) {
    this.internal.viz = viz;
        return this;
    }
    public HideSeriesConfig build() {
            return this.internal;
        }
    }
}
