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
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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
    // PromQL sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("promQLQuery")
    public PromQLQuery promQLQuery;
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
    public String dataqueryName() {
        return "cloud-monitoring";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CloudMonitoringQuery> {
        protected final CloudMonitoringQuery internal;
        
        public Builder() {
            this.internal = new CloudMonitoringQuery();
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
    
    public Builder aliasBy(String aliasBy) {
    this.internal.aliasBy = aliasBy;
        return this;
    }
    
    public Builder timeSeriesList(com.grafana.foundation.cog.Builder<TimeSeriesList> timeSeriesList) {
    this.internal.timeSeriesList = timeSeriesList.build();
        return this;
    }
    
    public Builder timeSeriesQuery(com.grafana.foundation.cog.Builder<TimeSeriesQuery> timeSeriesQuery) {
    this.internal.timeSeriesQuery = timeSeriesQuery.build();
        return this;
    }
    
    public Builder sloQuery(com.grafana.foundation.cog.Builder<SLOQuery> sloQuery) {
    this.internal.sloQuery = sloQuery.build();
        return this;
    }
    
    public Builder promQLQuery(com.grafana.foundation.cog.Builder<PromQLQuery> promQLQuery) {
    this.internal.promQLQuery = promQLQuery.build();
        return this;
    }
    
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public Builder intervalMs(Double intervalMs) {
    this.internal.intervalMs = intervalMs;
        return this;
    }
    public CloudMonitoringQuery build() {
            return this.internal;
        }
    }
}
