// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// PromQL sub-query properties.
public class PromQLQuery {
    // GCP project to execute the query against.
    @JsonProperty("projectName")
    public String projectName;
    // PromQL expression/query to be executed.
    @JsonProperty("expr")
    public String expr;
    // PromQL min step
    @JsonProperty("step")
    public String step;
    public PromQLQuery() {
        this.projectName = "";
        this.expr = "";
        this.step = "";
    }
    public PromQLQuery(String projectName,String expr,String step) {
        this.projectName = projectName;
        this.expr = expr;
        this.step = step;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
