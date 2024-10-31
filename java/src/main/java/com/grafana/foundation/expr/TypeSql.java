// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.dashboard.DataSourceRef;

public class TypeSql implements com.grafana.foundation.cog.variants.Dataquery {
    // The datasource
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    @JsonProperty("expression")
    public String expression;
    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("intervalMs")
    public Double intervalMs;
    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxDataPoints")
    public Long maxDataPoints;
    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    // RefID is the unique identifier of the query, set by the frontend call.
    @JsonProperty("refId")
    public String refId;
    // Optionally define expected query result behavior
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resultAssertions")
    public ExprTypeSqlResultAssertions resultAssertions;
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeRange")
    public ExprTypeSqlTimeRange timeRange;
    @JsonProperty("type")
    public String type;
    public String dataqueryName() {
        return "__expr__";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TypeSql> {
        protected final TypeSql internal;
        
        public Builder() {
            this.internal = new TypeSql();
    this.internal.type = "sql";
        }
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
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
    
    public Builder resultAssertions(com.grafana.foundation.cog.Builder<ExprTypeSqlResultAssertions> resultAssertions) {
    this.internal.resultAssertions = resultAssertions.build();
        return this;
    }
    
    public Builder timeRange(com.grafana.foundation.cog.Builder<ExprTypeSqlTimeRange> timeRange) {
    this.internal.timeRange = timeRange.build();
        return this;
    }
    public TypeSql build() {
            return this.internal;
        }
    }
}
