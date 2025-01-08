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
    // If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dashboardTime")
    public Boolean dashboardTime;
    // If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeColumn")
    public String timeColumn;
    // Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("workspace")
    public String workspace;
    // @deprecated Use resources instead
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resource")
    public String resource;
    // @deprecated Use dashboardTime instead
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("intersectTime")
    public Boolean intersectTime;
    public AzureLogsQuery() {
    }
    
    public AzureLogsQuery(String query,ResultFormat resultFormat,List<String> resources,Boolean dashboardTime,String timeColumn,String workspace,String resource,Boolean intersectTime) {
        this.query = query;
        this.resultFormat = resultFormat;
        this.resources = resources;
        this.dashboardTime = dashboardTime;
        this.timeColumn = timeColumn;
        this.workspace = workspace;
        this.resource = resource;
        this.intersectTime = intersectTime;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
