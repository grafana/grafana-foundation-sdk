// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class TypeThreshold implements com.grafana.foundation.cog.variants.Dataquery {
    // Threshold Conditions 
    @JsonProperty("conditions")
    public List<ExprTypeThresholdConditions> conditions;
    // The datasource 
    @JsonProperty("datasource")
    public ExprTypeThresholdDatasource datasource;
    // Reference to single query result 
    @JsonProperty("expression")
    public String expression;
    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc) 
    @JsonProperty("hide")
    public Boolean hide;
    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization 
    @JsonProperty("intervalMs")
    public Double intervalMs;
    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization 
    @JsonProperty("maxDataPoints")
    public Long maxDataPoints;
    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries. 
    @JsonProperty("queryType")
    public String queryType;
    // RefID is the unique identifier of the query, set by the frontend call. 
    @JsonProperty("refId")
    public String refId;
    // Optionally define expected query result behavior 
    @JsonProperty("resultAssertions")
    public ExprTypeThresholdResultAssertions resultAssertions;
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly 
    @JsonProperty("timeRange")
    public ExprTypeThresholdTimeRange timeRange; 
    @JsonProperty("type")
    public String type;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TypeThreshold> {
        private final TypeThreshold internal;
        
        public Builder() {
            this.internal = new TypeThreshold();
    this.internal.type = "threshold";
        }
    public Builder conditions(com.grafana.foundation.cog.Builder<List<ExprTypeThresholdConditions>> conditions) {
    this.internal.conditions = conditions.build();
        return this;
    }
    
    public Builder datasource(com.grafana.foundation.cog.Builder<ExprTypeThresholdDatasource> datasource) {
    this.internal.datasource = datasource.build();
        return this;
    }
    
    public Builder expression(String expression) {
        if (!(expression.length() >= 1)) {
            throw new IllegalArgumentException("expression.length() must be >= 1");
        }
    this.internal.expression = expression;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public Builder intervalMs(Double intervalMs) {
    this.internal.intervalMs = intervalMs;
        return this;
    }
    
    public Builder maxDataPoints(Long maxDataPoints) {
    this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public Builder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public Builder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public Builder resultAssertions(com.grafana.foundation.cog.Builder<ExprTypeThresholdResultAssertions> resultAssertions) {
    this.internal.resultAssertions = resultAssertions.build();
        return this;
    }
    
    public Builder timeRange(com.grafana.foundation.cog.Builder<ExprTypeThresholdTimeRange> timeRange) {
    this.internal.timeRange = timeRange.build();
        return this;
    }
    public TypeThreshold build() {
            return this.internal;
        }
    }
}
