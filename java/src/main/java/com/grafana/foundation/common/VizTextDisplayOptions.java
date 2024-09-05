// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// TODO docs
public class VizTextDisplayOptions {
    // Explicit title text size
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("titleSize")
    public Double titleSize;
    // Explicit value text size
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("valueSize")
    public Double valueSize;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<VizTextDisplayOptions> {
        private final VizTextDisplayOptions internal;
        
        public Builder() {
            this.internal = new VizTextDisplayOptions();
        }
    public Builder titleSize(Double titleSize) {
    this.internal.titleSize = titleSize;
        return this;
    }
    
    public Builder valueSize(Double valueSize) {
    this.internal.valueSize = valueSize;
        return this;
    }
    public VizTextDisplayOptions build() {
            return this.internal;
        }
    }
}
