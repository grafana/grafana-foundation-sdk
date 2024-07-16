// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// TODO docs
public class StackingConfig { 
    @JsonProperty("mode")
    public StackingMode mode; 
    @JsonProperty("group")
    public String group;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<StackingConfig> {
        private final StackingConfig internal;
        
        public Builder() {
            this.internal = new StackingConfig();
        }
    public Builder mode(StackingMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder group(String group) {
    this.internal.group = group;
        return this;
    }
    public StackingConfig build() {
            return this.internal;
        }
    }
}
