// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Links to a resource (image/svg path)
public class ResourceDimensionConfig {
    @JsonProperty("mode")
    public ResourceDimensionMode mode;
    // fixed: T -- will be added by each element
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("field")
    public String field;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("fixed")
    public String fixed;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ResourceDimensionConfig> {
        private final ResourceDimensionConfig internal;
        
        public Builder() {
            this.internal = new ResourceDimensionConfig();
        }
    public Builder mode(ResourceDimensionMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public Builder fixed(String fixed) {
    this.internal.fixed = fixed;
        return this;
    }
    public ResourceDimensionConfig build() {
            return this.internal;
        }
    }
}
