// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.debug;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Options {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public DebugMode mode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("counters")
    public UpdateConfig counters;
    public Options() {
    }
    
    public Options(DebugMode mode,UpdateConfig counters) {
        this.mode = mode;
        this.counters = counters;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
