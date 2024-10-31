// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Controls exemplar options
public class ExemplarConfig {
    // Sets the color of the exemplar markers
    @JsonProperty("color")
    public String color;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExemplarConfig> {
        protected final ExemplarConfig internal;
        
        public Builder() {
            this.internal = new ExemplarConfig();
        }
    public Builder color(String color) {
    this.internal.color = color;
        return this;
    }
    public ExemplarConfig build() {
            return this.internal;
        }
    }
}
