// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class DashboardGraphPanelLegend {
    @JsonProperty("show")
    public Boolean show;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sort")
    public String sort;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sortDesc")
    public Boolean sortDesc;
    public DashboardGraphPanelLegend() {
        this.show = true;
    }
    public DashboardGraphPanelLegend(Boolean show,String sort,Boolean sortDesc) {
        this.show = show;
        this.sort = sort;
        this.sortDesc = sortDesc;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
