// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class DashboardDashboardTemplating {
    // List of configured template variables with their saved values along with some other metadata
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("list")
    public List<VariableModel> list;
    public DashboardDashboardTemplating() {
    }
    
    public DashboardDashboardTemplating(List<VariableModel> list) {
        this.list = list;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
