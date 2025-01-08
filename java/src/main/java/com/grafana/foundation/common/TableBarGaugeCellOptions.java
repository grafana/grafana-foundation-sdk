// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Gauge cell options
public class TableBarGaugeCellOptions {
    @JsonProperty("type")
    public String type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mode")
    public BarGaugeDisplayMode mode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("valueDisplayMode")
    public BarGaugeValueMode valueDisplayMode;
    public TableBarGaugeCellOptions() {
    }
    
    public TableBarGaugeCellOptions(String type,BarGaugeDisplayMode mode,BarGaugeValueMode valueDisplayMode) {
        this.type = type;
        this.mode = mode;
        this.valueDisplayMode = valueDisplayMode;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
