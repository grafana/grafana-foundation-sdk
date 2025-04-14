// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Azure Monitor Logs sub-query properties
public class AzureLogsQuery {
    // KQL query to be executed.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("query")
    public String query;
    // Specifies the format results should be returned as.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resultFormat")
    public ResultFormat resultFormat;
    // Array of resource URIs to be queried.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resources")
    public List<String> resources;
    // If set to true the intersection of time ranges specified in the query and Grafana will be used. Otherwise the query time ranges will be used. Defaults to false
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("intersectTime")
    public Boolean intersectTime;
    // Workspace ID. This was removed in Grafana 8, but remains for backwards compat
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("workspace")
    public String workspace;
    // @deprecated Use resources instead
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resource")
    public String resource;
    public AzureLogsQuery() {
    }
    public AzureLogsQuery(String query,ResultFormat resultFormat,List<String> resources,Boolean intersectTime,String workspace,String resource) {
        this.query = query;
        this.resultFormat = resultFormat;
        this.resources = resources;
        this.intersectTime = intersectTime;
        this.workspace = workspace;
        this.resource = resource;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
