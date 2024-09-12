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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DashboardDashboardTemplating> {
        private final DashboardDashboardTemplating internal;
        
        public Builder() {
            this.internal = new DashboardDashboardTemplating();
        }
    public Builder list(com.grafana.foundation.cog.Builder<List<VariableModel>> list) {
    this.internal.list = list.build();
        return this;
    }
    public DashboardDashboardTemplating build() {
            return this.internal;
        }
    }
}
