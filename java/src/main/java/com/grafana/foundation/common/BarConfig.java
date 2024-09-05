// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// TODO docs
public class BarConfig {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("barAlignment")
    public BarAlignment barAlignment;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("barWidthFactor")
    public Double barWidthFactor;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("barMaxWidth")
    public Double barMaxWidth;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<BarConfig> {
        private final BarConfig internal;
        
        public Builder() {
            this.internal = new BarConfig();
        }
    public Builder barAlignment(BarAlignment barAlignment) {
    this.internal.barAlignment = barAlignment;
        return this;
    }
    
    public Builder barWidthFactor(Double barWidthFactor) {
    this.internal.barWidthFactor = barWidthFactor;
        return this;
    }
    
    public Builder barMaxWidth(Double barMaxWidth) {
    this.internal.barMaxWidth = barMaxWidth;
        return this;
    }
    public BarConfig build() {
            return this.internal;
        }
    }
}
