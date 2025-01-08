// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Redefining this to avoid an import cycle
public class TimeRange {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("from")
    public String from;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("to")
    public String to;
    public TimeRange() {
    }
    
    public TimeRange(String from,String to) {
        this.from = from;
        this.to = to;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
