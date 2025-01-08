// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ExprTypeThresholdTimeRange {
    // From is the start time of the query.
    @JsonProperty("from")
    public String from;
    // To is the end time of the query.
    @JsonProperty("to")
    public String to;
    public ExprTypeThresholdTimeRange() {
        this.from = "now-6h";
        this.to = "now";
    }
    
    public ExprTypeThresholdTimeRange(String from,String to) {
        this.from = from;
        this.to = to;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
