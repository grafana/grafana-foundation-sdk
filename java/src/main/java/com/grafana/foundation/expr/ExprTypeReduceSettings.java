// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ExprTypeReduceSettings {
    // Non-number reduce behavior
    // Possible enum values:
    //  - `"dropNN"` Drop non-numbers
    //  - `"replaceNN"` Replace non-numbers
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public TypeReduceMode mode;
    // Only valid when mode is replace
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("replaceWithValue")
    public Double replaceWithValue;
    public ExprTypeReduceSettings() {
    }
    
    public ExprTypeReduceSettings(TypeReduceMode mode,Double replaceWithValue) {
        this.mode = mode;
        this.replaceWithValue = replaceWithValue;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
