// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class TimeRangeOption {
    @JsonProperty("display")
    public String display;
    @JsonProperty("from")
    public String from;
    @JsonProperty("to")
    public String to;
    public TimeRangeOption() {
        this.display = "Last 6 hours";
        this.from = "now-6h";
        this.to = "now";
    }
    public TimeRangeOption(String display,String from,String to) {
        this.display = display;
        this.from = from;
        this.to = to;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
