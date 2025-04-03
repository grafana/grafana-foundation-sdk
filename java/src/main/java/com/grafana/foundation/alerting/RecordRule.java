// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class RecordRule {
    // Which expression node should be used as the input for the recorded metric.
    @JsonProperty("from")
    public String from;
    // Name of the recorded metric.
    @JsonProperty("metric")
    public String metric;
    public RecordRule() {
    }
    public RecordRule(String from,String metric) {
        this.from = from;
        this.metric = metric;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
