// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.dashboard.DataSourceRef;

public class CloudMonitoringQuery implements com.grafana.foundation.cog.variants.Dataquery {
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    @JsonProperty("refId")
    public String refId;
    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Specify the query flavor
    // TODO make this required and give it a default
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    // Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("aliasBy")
    public String aliasBy;
    // GCM query type.
    // queryType: #QueryType
    // Time Series List sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeSeriesList")
    public TimeSeriesList timeSeriesList;
    // Time Series sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeSeriesQuery")
    public TimeSeriesQuery timeSeriesQuery;
    // SLO sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sloQuery")
    public SLOQuery sloQuery;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // Time interval in milliseconds.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("intervalMs")
    public Double intervalMs;
    public CloudMonitoringQuery() {
        this.refId = "";
    }
    public CloudMonitoringQuery(String refId,Boolean hide,String queryType,String aliasBy,TimeSeriesList timeSeriesList,TimeSeriesQuery timeSeriesQuery,SLOQuery sloQuery,DataSourceRef datasource,Double intervalMs) {
        this.refId = refId;
        this.hide = hide;
        this.queryType = queryType;
        this.aliasBy = aliasBy;
        this.timeSeriesList = timeSeriesList;
        this.timeSeriesQuery = timeSeriesQuery;
        this.sloQuery = sloQuery;
        this.datasource = datasource;
        this.intervalMs = intervalMs;
    }
    public String dataqueryName() {
        return "cloud-monitoring";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
