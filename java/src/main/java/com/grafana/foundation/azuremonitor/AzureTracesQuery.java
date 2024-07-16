// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Application Insights Traces sub-query properties
public class AzureTracesQuery {
    // Specifies the format results should be returned as. 
    @JsonProperty("resultFormat")
    public ResultFormat resultFormat;
    // Array of resource URIs to be queried. 
    @JsonProperty("resources")
    public List<String> resources;
    // Operation ID. Used only for Traces queries. 
    @JsonProperty("operationId")
    public String operationId;
    // Types of events to filter by. 
    @JsonProperty("traceTypes")
    public List<String> traceTypes;
    // Filters for property values. 
    @JsonProperty("filters")
    public List<AzureTracesFilter> filters;
    // KQL query to be executed. 
    @JsonProperty("query")
    public String query;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AzureTracesQuery> {
        private final AzureTracesQuery internal;
        
        public Builder() {
            this.internal = new AzureTracesQuery();
        }
    public Builder resultFormat(ResultFormat resultFormat) {
    this.internal.resultFormat = resultFormat;
        return this;
    }
    
    public Builder resources(List<String> resources) {
    this.internal.resources = resources;
        return this;
    }
    
    public Builder operationId(String operationId) {
    this.internal.operationId = operationId;
        return this;
    }
    
    public Builder traceTypes(List<String> traceTypes) {
    this.internal.traceTypes = traceTypes;
        return this;
    }
    
    public Builder filters(com.grafana.foundation.cog.Builder<List<AzureTracesFilter>> filters) {
    this.internal.filters = filters.build();
        return this;
    }
    
    public Builder query(String query) {
    this.internal.query = query;
        return this;
    }
    public AzureTracesQuery build() {
            return this.internal;
        }
    }
}
