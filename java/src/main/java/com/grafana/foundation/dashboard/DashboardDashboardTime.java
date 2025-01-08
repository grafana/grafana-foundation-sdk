// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class DashboardDashboardTime {
    @JsonProperty("from")
    public String from;
    @JsonProperty("to")
    public String to;
    public DashboardDashboardTime() {
        this.from = "now-6h";
        this.to = "now";
    }
    
    public DashboardDashboardTime(String from,String to) {
        this.from = from;
        this.to = to;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
