// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Dashboardv2RangeMapOptions {
    // Min value of the range. It can be null which means -Infinity
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("from")
    public Double from;
    // Max value of the range. It can be null which means +Infinity
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("to")
    public Double to;
    // Config to apply when the value is within the range
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("result")
    public ValueMappingResult result;
    public Dashboardv2RangeMapOptions() {
        this.result = new com.grafana.foundation.dashboardv2.ValueMappingResult();
    }
    public Dashboardv2RangeMapOptions(Double from,Double to,ValueMappingResult result) {
        this.from = from;
        this.to = to;
        this.result = result;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
