// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import java.util.List;

// Azure Monitor Logs sub-query properties
public class AzureLogsQuery {
    // KQL query to be executed.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("query")
    public String query;
    // Specifies the format results should be returned as.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("resultFormat")
    public ResultFormat resultFormat;
    // Array of resource URIs to be queried.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("resources")
    public List<String> resources;
    // If set to true the intersection of time ranges specified in the query and Grafana will be used. Otherwise the query time ranges will be used. Defaults to false
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("intersectTime")
    public Boolean intersectTime;
    // Workspace ID. This was removed in Grafana 8, but remains for backwards compat
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("workspace")
    public String workspace;
    // @deprecated Use resources instead
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("resource")
    public String resource;
    
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
    
    public Builder intersectTime(Boolean intersectTime) {
    this.internal.intersectTime = intersectTime;
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
    public AzureLogsQuery build() {
            return this.internal;
        }
    }
}
