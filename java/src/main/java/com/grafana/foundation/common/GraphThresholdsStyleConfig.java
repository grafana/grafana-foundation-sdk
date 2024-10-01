// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class GraphThresholdsStyleConfig {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public GraphTresholdsStyleMode mode;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<GraphThresholdsStyleConfig> {
        protected final GraphThresholdsStyleConfig internal;
        
        public Builder() {
            this.internal = new GraphThresholdsStyleConfig();
        }
    public Builder mode(GraphTresholdsStyleMode mode) {
    this.internal.mode = mode;
        return this;
    }
    public GraphThresholdsStyleConfig build() {
            return this.internal;
        }
    }
}
