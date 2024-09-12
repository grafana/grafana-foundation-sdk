// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class FillConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillColor")
    public String fillColor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillOpacity")
    public Double fillOpacity;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fillBelowTo")
    public String fillBelowTo;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<FillConfig> {
        private final FillConfig internal;
        
        public Builder() {
            this.internal = new FillConfig();
        }
    public Builder fillColor(String fillColor) {
    this.internal.fillColor = fillColor;
        return this;
    }
    
    public Builder fillOpacity(Double fillOpacity) {
    this.internal.fillOpacity = fillOpacity;
        return this;
    }
    
    public Builder fillBelowTo(String fillBelowTo) {
    this.internal.fillBelowTo = fillBelowTo;
        return this;
    }
    public FillConfig build() {
            return this.internal;
        }
    }
}
