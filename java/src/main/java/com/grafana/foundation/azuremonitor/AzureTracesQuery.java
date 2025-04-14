// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Application Insights Traces sub-query properties
public class AzureTracesQuery {
    // Specifies the format results should be returned as.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resultFormat")
    public ResultFormat resultFormat;
    // Array of resource URIs to be queried.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resources")
    public List<String> resources;
    // Operation ID. Used only for Traces queries.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("operationId")
    public String operationId;
    // Types of events to filter by.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("traceTypes")
    public List<String> traceTypes;
    // Filters for property values.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filters")
    public List<AzureTracesFilter> filters;
    // KQL query to be executed.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("query")
    public String query;
    public AzureTracesQuery() {
    }
    public AzureTracesQuery(ResultFormat resultFormat,List<String> resources,String operationId,List<String> traceTypes,List<AzureTracesFilter> filters,String query) {
        this.resultFormat = resultFormat;
        this.resources = resources;
        this.operationId = operationId;
        this.traceTypes = traceTypes;
        this.filters = filters;
        this.query = query;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
