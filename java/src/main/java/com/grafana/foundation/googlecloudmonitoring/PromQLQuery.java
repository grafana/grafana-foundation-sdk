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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<PromQLQuery> {
        private final PromQLQuery internal;
        
        public Builder() {
            this.internal = new PromQLQuery();
        }
    public Builder projectName(String projectName) {
    this.internal.projectName = projectName;
        return this;
    }
    
    public Builder expr(String expr) {
    this.internal.expr = expr;
        return this;
    }
    
    public Builder step(String step) {
    this.internal.step = step;
        return this;
    }
    public PromQLQuery build() {
            return this.internal;
        }
    }
}
