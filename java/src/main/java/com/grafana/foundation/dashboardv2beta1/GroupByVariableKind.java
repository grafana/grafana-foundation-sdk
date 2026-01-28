// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Group variable kind
public class GroupByVariableKind {
    @JsonProperty("kind")
    public String kind;
    @JsonProperty("group")
    public String group;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public Dashboardv2beta1GroupByVariableKindDatasource datasource;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("spec")
    public GroupByVariableSpec spec;
    public GroupByVariableKind() {
        this.kind = "";
        this.group = "";
        this.spec = new com.grafana.foundation.dashboardv2beta1.GroupByVariableSpec();
    }
    public GroupByVariableKind(String kind,String group,Dashboardv2beta1GroupByVariableKindDatasource datasource,GroupByVariableSpec spec) {
        this.kind = kind;
        this.group = group;
        this.datasource = datasource;
        this.spec = spec;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
