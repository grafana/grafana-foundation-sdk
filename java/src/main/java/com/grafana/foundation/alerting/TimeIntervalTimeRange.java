// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class TimeIntervalTimeRange {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("end_time")
    public String endTime;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("start_time")
    public String startTime;
    public TimeIntervalTimeRange() {
    }
    
    public TimeIntervalTimeRange(String endTime,String startTime) {
        this.endTime = endTime;
        this.startTime = startTime;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
