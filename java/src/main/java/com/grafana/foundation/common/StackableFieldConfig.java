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
    public StackableFieldConfig() {
    }
    public StackableFieldConfig(StackingConfig stacking) {
        this.stacking = stacking;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
