// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class DashboardGraphPanelLegend {
    @JsonProperty("show")
    public Boolean show;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("sort")
    public String sort;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("sortDesc")
    public Boolean sortDesc;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DashboardGraphPanelLegend> {
        private final DashboardGraphPanelLegend internal;
        
        public Builder() {
            this.internal = new DashboardGraphPanelLegend();
        this.show(true);
        }
    public Builder show(Boolean show) {
    this.internal.show = show;
        return this;
    }
    
    public Builder sort(String sort) {
    this.internal.sort = sort;
        return this;
    }
    
    public Builder sortDesc(Boolean sortDesc) {
    this.internal.sortDesc = sortDesc;
        return this;
    }
    public DashboardGraphPanelLegend build() {
            return this.internal;
        }
    }
}
