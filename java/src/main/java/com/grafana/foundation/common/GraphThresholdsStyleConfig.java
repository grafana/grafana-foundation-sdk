// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// TODO docs
public class GraphThresholdsStyleConfig { 
    @JsonProperty("mode")
    public GraphThresholdsStyleMode mode;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> {
        private final GraphThresholdsStyleConfig internal;
        
        public Builder() {
            this.internal = new GraphThresholdsStyleConfig();
        }
    public Builder mode(GraphThresholdsStyleMode mode) {
    this.internal.mode = mode;
        return this;
    }
    public GraphThresholdsStyleConfig build() {
            return this.internal;
        }
    }
}
