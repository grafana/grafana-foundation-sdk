// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.dashboard.DataSourceRef;

public class TypeReduce implements com.grafana.foundation.cog.variants.Dataquery {
    // The datasource
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // Reference to single query result
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
    // The reducer
    // Possible enum values:
    //  - `"sum"` 
    //  - `"mean"` 
    //  - `"min"` 
    //  - `"max"` 
    //  - `"count"` 
    //  - `"last"` 
    //  - `"median"` 
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("reducer")
    public TypeReduceReducer reducer;
    // RefID is the unique identifier of the query, set by the frontend call.
    @JsonProperty("refId")
    public String refId;
    // Optionally define expected query result behavior
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resultAssertions")
    public ExprTypeReduceResultAssertions resultAssertions;
    // Reducer Options
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public ExprTypeReduceSettings settings;
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeRange")
    public ExprTypeReduceTimeRange timeRange;
    @JsonProperty("type")
    public String type;
    public TypeReduce() {
    }
    public TypeReduce(DataSourceRef datasource,String expression,Boolean hide,Double intervalMs,Long maxDataPoints,String queryType,TypeReduceReducer reducer,String refId,ExprTypeReduceResultAssertions resultAssertions,ExprTypeReduceSettings settings,ExprTypeReduceTimeRange timeRange,String type) {
        this.datasource = datasource;
        this.expression = expression;
        this.hide = hide;
        this.intervalMs = intervalMs;
        this.maxDataPoints = maxDataPoints;
        this.queryType = queryType;
        this.reducer = reducer;
        this.refId = refId;
        this.resultAssertions = resultAssertions;
        this.settings = settings;
        this.timeRange = timeRange;
        this.type = type;
    }
    public String dataqueryName() {
        return "__expr__";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
