// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Azure Monitor Logs sub-query properties
public class AzureLogsQuery {
    // KQL query to be executed. 
    @JsonProperty("query")
    public String query;
    // Specifies the format results should be returned as. 
    @JsonProperty("resultFormat")
    public ResultFormat resultFormat;
    // Array of resource URIs to be queried. 
    @JsonProperty("resources")
    public List<String> resources;
    // If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false. 
    @JsonProperty("dashboardTime")
    public Boolean dashboardTime;
    // If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated 
    @JsonProperty("timeColumn")
    public String timeColumn;
    // If set to true the query will be run as a basic logs query 
    @JsonProperty("basicLogsQuery")
    public Boolean basicLogsQuery;
    // Workspace ID. This was removed in Grafana 8, but remains for backwards compat. 
    @JsonProperty("workspace")
    public String workspace;
    // @deprecated Use resources instead 
    @JsonProperty("resource")
    public String resource;
    // @deprecated Use dashboardTime instead 
    @JsonProperty("intersectTime")
    public Boolean intersectTime;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AzureLogsQuery> {
        private final AzureLogsQuery internal;
        
        public Builder() {
            this.internal = new AzureLogsQuery();
        }
    public Builder query(String query) {
    this.internal.query = query;
        return this;
    }
    
    public Builder resultFormat(ResultFormat resultFormat) {
    this.internal.resultFormat = resultFormat;
        return this;
    }
    
    public Builder resources(List<String> resources) {
    this.internal.resources = resources;
        return this;
    }
    
    public Builder dashboardTime(Boolean dashboardTime) {
    this.internal.dashboardTime = dashboardTime;
        return this;
    }
    
    public Builder timeColumn(String timeColumn) {
    this.internal.timeColumn = timeColumn;
        return this;
    }
    
    public Builder basicLogsQuery(Boolean basicLogsQuery) {
    this.internal.basicLogsQuery = basicLogsQuery;
        return this;
    }
    
    public Builder workspace(String workspace) {
    this.internal.workspace = workspace;
        return this;
    }
    
    public Builder resource(String resource) {
    this.internal.resource = resource;
        return this;
    }
    
    public Builder intersectTime(Boolean intersectTime) {
    this.internal.intersectTime = intersectTime;
        return this;
    }
    public AzureLogsQuery build() {
            return this.internal;
        }
    }
}
