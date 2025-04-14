// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.dashboard.DataSourceRef;

// Shape of a CloudWatch Logs query
public class CloudWatchLogsQuery implements com.grafana.foundation.cog.variants.Dataquery {
    // Whether a query is a Metrics, Logs, or Annotations query
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("queryMode")
    public CloudWatchQueryMode queryMode;
    @JsonProperty("id")
    public String id;
    // AWS region to query for the logs
    @JsonProperty("region")
    public String region;
    // The CloudWatch Logs Insights query to execute
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("expression")
    public String expression;
    // Fields to group the results by, this field is automatically populated whenever the query is updated
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("statsGroups")
    public List<String> statsGroups;
    // Log groups to query
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("logGroups")
    public List<LogGroup> logGroups;
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    @JsonProperty("refId")
    public String refId;
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Specify the query flavor
    // TODO make this required and give it a default
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    // @deprecated use logGroups
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("logGroupNames")
    public List<String> logGroupNames;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    public CloudWatchLogsQuery() {
        this.queryMode = CloudWatchQueryMode.LOGS;
    }
    public CloudWatchLogsQuery(CloudWatchQueryMode queryMode,String id,String region,String expression,List<String> statsGroups,List<LogGroup> logGroups,String refId,Boolean hide,String queryType,List<String> logGroupNames,DataSourceRef datasource) {
        this.queryMode = queryMode;
        this.id = id;
        this.region = region;
        this.expression = expression;
        this.statsGroups = statsGroups;
        this.logGroups = logGroups;
        this.refId = refId;
        this.hide = hide;
        this.queryType = queryType;
        this.logGroupNames = logGroupNames;
        this.datasource = datasource;
    }
    public String dataqueryName() {
        return "cloudwatch";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
