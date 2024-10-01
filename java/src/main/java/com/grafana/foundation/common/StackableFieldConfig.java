// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class StackableFieldConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("stacking")
    public StackingConfig stacking;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<StackableFieldConfig> {
        protected final StackableFieldConfig internal;
        
        public Builder() {
            this.internal = new StackableFieldConfig();
        }
    public Builder stacking(com.grafana.foundation.cog.Builder<StackingConfig> stacking) {
    this.internal.stacking = stacking.build();
        return this;
    }
    public StackableFieldConfig build() {
            return this.internal;
        }
    }
}
