// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class VizTooltipOptions {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public TooltipDisplayMode mode;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("sort")
    public SortOrder sort;
    public VizTooltipOptions() {
    }
    public VizTooltipOptions(TooltipDisplayMode mode,SortOrder sort) {
        this.mode = mode;
        this.sort = sort;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
