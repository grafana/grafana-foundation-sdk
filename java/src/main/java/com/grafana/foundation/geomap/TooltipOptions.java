// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class TooltipOptions {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public TooltipMode mode;
    public TooltipOptions() {
        this.mode = TooltipMode.NONE;
    }
    public TooltipOptions(TooltipMode mode) {
        this.mode = mode;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
