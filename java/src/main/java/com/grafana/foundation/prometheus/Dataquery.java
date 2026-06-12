// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.common.DataSourceRef;

public class Dataquery implements com.grafana.foundation.cog.variants.Dataquery {
    // Additional Ad-hoc filters that take precedence over Scope on conflict.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("adhocFilters")
    public List<AdhocFilters> adhocFilters;
    // The datasource
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // what we should show in the editor
    // Possible enum values:
    //  - `"builder"` 
    //  - `"code"` 
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("editorMode")
    public QueryEditorMode editorMode;
    // Execute an additional query to identify interesting raw samples relevant for the given expr
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("exemplar")
    public Boolean exemplar;
    // The actual expression/query that will be evaluated by Prometheus
    @JsonProperty("expr")
    public String expr;
    // The response format
    // Possible enum values:
    //  - `"time_series"` 
    //  - `"table"` 
    //  - `"heatmap"` 
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("format")
    public PromQueryFormat format;
    // Group By parameters to apply to aggregate expressions in the query
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("groupByKeys")
    public List<String> groupByKeys;
    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Returns only the latest value that Prometheus has scraped for the requested time series
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("instant")
    public Boolean instant;
    // Used to specify how many times to divide max data points by. We use max data points under query options
    // See https://github.com/grafana/grafana/issues/48081
    // Deprecated: use interval
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("intervalFactor")
    public Long intervalFactor;
    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("intervalMs")
    public Double intervalMs;
    // Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("legendFormat")
    public String legendFormat;
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
    // Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("range")
    public Boolean range;
    // RefID is the unique identifier of the query, set by the frontend call.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("refId")
    public String refId;
    // Optionally define expected query result behavior
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resultAssertions")
    public ResultAssertions resultAssertions;
    // A set of filters applied to apply to the query
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("scopes")
    public List<Scopes> scopes;
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeRange")
    public TimeRange timeRange;
    // An additional lower limit for the step parameter of the Prometheus query and for the
    // `$__interval` and `$__rate_interval` variables.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("interval")
    public String interval;
    public Dataquery() {
        this.expr = "";
    }
    public Dataquery(List<AdhocFilters> adhocFilters,DataSourceRef datasource,QueryEditorMode editorMode,Boolean exemplar,String expr,PromQueryFormat format,List<String> groupByKeys,Boolean hide,Boolean instant,Long intervalFactor,Double intervalMs,String legendFormat,Long maxDataPoints,String queryType,Boolean range,String refId,ResultAssertions resultAssertions,List<Scopes> scopes,TimeRange timeRange,String interval) {
        this.adhocFilters = adhocFilters;
        this.datasource = datasource;
        this.editorMode = editorMode;
        this.exemplar = exemplar;
        this.expr = expr;
        this.format = format;
        this.groupByKeys = groupByKeys;
        this.hide = hide;
        this.instant = instant;
        this.intervalFactor = intervalFactor;
        this.intervalMs = intervalMs;
        this.legendFormat = legendFormat;
        this.maxDataPoints = maxDataPoints;
        this.queryType = queryType;
        this.range = range;
        this.refId = refId;
        this.resultAssertions = resultAssertions;
        this.scopes = scopes;
        this.timeRange = timeRange;
        this.interval = interval;
    }
    public String dataqueryName() {
        return "prometheus";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
