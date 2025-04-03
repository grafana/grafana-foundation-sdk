// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.dashboard.DataSourceRef;

public class TypeResample implements com.grafana.foundation.cog.variants.Dataquery {
    // The datasource
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // The downsample function
    // Possible enum values:
    //  - `"sum"` 
    //  - `"mean"` 
    //  - `"min"` 
    //  - `"max"` 
    //  - `"count"` 
    //  - `"last"` 
    //  - `"median"` 
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("downsampler")
    public TypeResampleDownsampler downsampler;
    // The math expression
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
    public ExprTypeResampleResultAssertions resultAssertions;
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeRange")
    public ExprTypeResampleTimeRange timeRange;
    @JsonProperty("type")
    public String type;
    // The upsample function
    // Possible enum values:
    //  - `"pad"` Use the last seen value
    //  - `"backfilling"` backfill
    //  - `"fillna"` Do not fill values (nill)
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("upsampler")
    public TypeResampleUpsampler upsampler;
    // The time duration
    @JsonProperty("window")
    public String window;
    public TypeResample() {
    }
    public TypeResample(DataSourceRef datasource,TypeResampleDownsampler downsampler,String expression,Boolean hide,Double intervalMs,Long maxDataPoints,String queryType,String refId,ExprTypeResampleResultAssertions resultAssertions,ExprTypeResampleTimeRange timeRange,String type,TypeResampleUpsampler upsampler,String window) {
        this.datasource = datasource;
        this.downsampler = downsampler;
        this.expression = expression;
        this.hide = hide;
        this.intervalMs = intervalMs;
        this.maxDataPoints = maxDataPoints;
        this.queryType = queryType;
        this.refId = refId;
        this.resultAssertions = resultAssertions;
        this.timeRange = timeRange;
        this.type = type;
        this.upsampler = upsampler;
        this.window = window;
    }
    public String dataqueryName() {
        return "__expr__";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
