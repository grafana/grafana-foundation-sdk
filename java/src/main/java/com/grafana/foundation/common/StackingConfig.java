// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class StackingConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mode")
    public StackingMode mode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group")
    public String group;
    public StackingConfig() {
    }
    
    public StackingConfig(StackingMode mode,String group) {
        this.mode = mode;
        this.group = group;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
