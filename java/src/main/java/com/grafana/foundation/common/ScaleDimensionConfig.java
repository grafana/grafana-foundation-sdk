// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class ScaleDimensionConfig {
    @JsonProperty("min")
    public Double min;
    @JsonProperty("max")
    public Double max;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("fixed")
    public Double fixed;
    // fixed: T -- will be added by each element
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("field")
    public String field;
    // | *"linear"
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("mode")
    public ScaleDimensionMode mode;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ScaleDimensionConfig> {
        private final ScaleDimensionConfig internal;
        
        public Builder() {
            this.internal = new ScaleDimensionConfig();
        }
    public Builder min(Double min) {
    this.internal.min = min;
        return this;
    }
    
    public Builder max(Double max) {
    this.internal.max = max;
        return this;
    }
    
    public Builder fixed(Double fixed) {
    this.internal.fixed = fixed;
        return this;
    }
    
    public Builder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public Builder mode(ScaleDimensionMode mode) {
    this.internal.mode = mode;
        return this;
    }
    public ScaleDimensionConfig build() {
            return this.internal;
        }
    }
}
