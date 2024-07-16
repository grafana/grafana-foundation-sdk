// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ExprTypeReduceSettings {
    // Non-number reduce behavior
    // Possible enum values:
    //  - `"dropNN"` Drop non-numbers
    //  - `"replaceNN"` Replace non-numbers 
    @JsonProperty("mode")
    public TypeReduceMode mode;
    // Only valid when mode is replace 
    @JsonProperty("replaceWithValue")
    public Double replaceWithValue;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExprTypeReduceSettings> {
        private final ExprTypeReduceSettings internal;
        
        public Builder() {
            this.internal = new ExprTypeReduceSettings();
        }
    public Builder mode(TypeReduceMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder replaceWithValue(Double replaceWithValue) {
    this.internal.replaceWithValue = replaceWithValue;
        return this;
    }
    public ExprTypeReduceSettings build() {
            return this.internal;
        }
    }
}