// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Thresholds configuration for the panel
public class ThresholdsConfig {
    // Thresholds mode. 
    @JsonProperty("mode")
    public ThresholdsMode mode;
    // Must be sorted by 'value', first value is always -Infinity 
    @JsonProperty("steps")
    public List<Threshold> steps;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ThresholdsConfig> {
        private final ThresholdsConfig internal;
        
        public Builder() {
            this.internal = new ThresholdsConfig();
        }
    public Builder mode(ThresholdsMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder steps(List<Threshold> steps) {
    this.internal.steps = steps;
        return this;
    }
    public ThresholdsConfig build() {
            return this.internal;
        }
    }
}
