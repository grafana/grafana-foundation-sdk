// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class ScaleDistributionConfig {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public ScaleDistribution type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("log")
    public Double log;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("linearThreshold")
    public Double linearThreshold;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ScaleDistributionConfig> {
        private final ScaleDistributionConfig internal;
        
        public Builder() {
            this.internal = new ScaleDistributionConfig();
        }
    public Builder type(ScaleDistribution type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder log(Double log) {
    this.internal.log = log;
        return this;
    }
    
    public Builder linearThreshold(Double linearThreshold) {
    this.internal.linearThreshold = linearThreshold;
        return this;
    }
    public ScaleDistributionConfig build() {
            return this.internal;
        }
    }
}
