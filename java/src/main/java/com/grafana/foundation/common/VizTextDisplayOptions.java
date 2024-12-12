// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class VizTextDisplayOptions {
    // Explicit title text size
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("titleSize")
    public Double titleSize;
    // Explicit value text size
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("valueSize")
    public Double valueSize;
    // Explicit percent text size
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("percentSize")
    public Double percentSize;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<VizTextDisplayOptions> {
        protected final VizTextDisplayOptions internal;
        
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
    
    public Builder percentSize(Double percentSize) {
    this.internal.percentSize = percentSize;
        return this;
    }
    public VizTextDisplayOptions build() {
            return this.internal;
        }
    }
}
