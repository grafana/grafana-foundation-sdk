// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Shape of a CloudWatch Logs query
public class CloudWatchLogsQuery implements com.grafana.foundation.cog.variants.Dataquery {
    // Whether a query is a Metrics, Logs, or Annotations query 
    @JsonProperty("queryMode")
    public CloudWatchQueryMode queryMode; 
    @JsonProperty("id")
    public String id;
    // AWS region to query for the logs 
    @JsonProperty("region")
    public String region;
    // The CloudWatch Logs Insights query to execute 
    @JsonProperty("expression")
    public String expression;
    // Fields to group the results by, this field is automatically populated whenever the query is updated 
    @JsonProperty("statsGroups")
    public List<String> statsGroups;
    // Log groups to query 
    @JsonProperty("logGroups")
    public List<LogGroup> logGroups;
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful. 
    @JsonProperty("refId")
    public String refId;
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel. 
    @JsonProperty("hide")
    public Boolean hide;
    // Specify the query flavor
    // TODO make this required and give it a default 
    @JsonProperty("queryType")
    public String queryType;
    // @deprecated use logGroups 
    @JsonProperty("logGroupNames")
    public List<String> logGroupNames;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null 
    @JsonProperty("datasource")
    public Object datasource;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CloudWatchLogsQuery> {
        private final CloudWatchLogsQuery internal;
        
        public Builder() {
            this.internal = new CloudWatchLogsQuery();
        this.queryMode(CloudWatchQueryMode.LOGS);
        }
    public Builder queryMode(CloudWatchQueryMode queryMode) {
    this.internal.queryMode = queryMode;
        return this;
    }
    
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder region(String region) {
    this.internal.region = region;
        return this;
    }
    
    public Builder expression(String expression) {
    this.internal.expression = expression;
        return this;
    }
    
    public Builder statsGroups(List<String> statsGroups) {
    this.internal.statsGroups = statsGroups;
        return this;
    }
    
    public Builder logGroups(com.grafana.foundation.cog.Builder<List<LogGroup>> logGroups) {
    this.internal.logGroups = logGroups.build();
        return this;
    }
    
    public Builder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public Builder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public Builder logGroupNames(List<String> logGroupNames) {
    this.internal.logGroupNames = logGroupNames;
        return this;
    }
    
    public Builder datasource(Object datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    public CloudWatchLogsQuery build() {
            return this.internal;
        }
    }
}
