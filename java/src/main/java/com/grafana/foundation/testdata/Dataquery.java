// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.dashboard.DataSourceRef;

public class Dataquery implements com.grafana.foundation.cog.variants.Dataquery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alias")
    public String alias;
    // Used for live query
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("channel")
    public String channel;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("csvContent")
    public String csvContent;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("csvFileName")
    public String csvFileName;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("csvWave")
    public List<CSVWave> csvWave;
    // The datasource
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // Drop percentage (the chance we will lose a point 0-100)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dropPercent")
    public Double dropPercent;
    // Possible enum values:
    //  - `"frontend_exception"` 
    //  - `"frontend_observable"` 
    //  - `"server_panic"` 
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("errorType")
    public DataqueryErrorType errorType;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("flamegraphDiff")
    public Boolean flamegraphDiff;
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
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labels")
    public String labels;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("levelColumn")
    public Boolean levelColumn;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lines")
    public Long lines;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("max")
    public Double max;
    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxDataPoints")
    public Long maxDataPoints;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min")
    public Double min;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("nodes")
    public NodesQuery nodes;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("noise")
    public Double noise;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("points")
    public List<List<Object>> points;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pulseWave")
    public PulseWaveQuery pulseWave;
    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("rawFrameContent")
    public String rawFrameContent;
    // RefID is the unique identifier of the query, set by the frontend call.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("refId")
    public String refId;
    // Optionally define expected query result behavior
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resultAssertions")
    public ResultAssertions resultAssertions;
    // Possible enum values:
    //  - `"annotations"` 
    //  - `"arrow"` 
    //  - `"csv_content"` 
    //  - `"csv_file"` 
    //  - `"csv_metric_values"` 
    //  - `"datapoints_outside_range"` 
    //  - `"exponential_heatmap_bucket_data"` 
    //  - `"flame_graph"` 
    //  - `"grafana_api"` 
    //  - `"linear_heatmap_bucket_data"` 
    //  - `"live"` 
    //  - `"logs"` 
    //  - `"manual_entry"` 
    //  - `"no_data_points"` 
    //  - `"node_graph"` 
    //  - `"predictable_csv_wave"` 
    //  - `"predictable_pulse"` 
    //  - `"random_walk"` 
    //  - `"random_walk_table"` 
    //  - `"random_walk_with_error"` 
    //  - `"raw_frame"` 
    //  - `"server_error_500"` 
    //  - `"simulation"` 
    //  - `"slow_query"` 
    //  - `"streaming_client"` 
    //  - `"table_static"` 
    //  - `"trace"` 
    //  - `"usa"` 
    //  - `"variables-query"` 
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("scenarioId")
    public DataqueryScenarioId scenarioId;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("seriesCount")
    public Long seriesCount;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sim")
    public SimulationQuery sim;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("spanCount")
    public Long spanCount;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("spread")
    public Double spread;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("startValue")
    public Double startValue;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("stream")
    public StreamingQuery stream;
    // common parameter used by many query types
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("stringInput")
    public String stringInput;
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeRange")
    public TimeRange timeRange;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("usa")
    public USAQuery usa;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("withNil")
    public Boolean withNil;
    public String dataqueryName() {
        return "";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
        protected final Dataquery internal;
        
        public Builder() {
            this.internal = new Dataquery();
        }
    public Builder alias(String alias) {
    this.internal.alias = alias;
        return this;
    }
    
    public Builder channel(String channel) {
    this.internal.channel = channel;
        return this;
    }
    
    public Builder csvContent(String csvContent) {
    this.internal.csvContent = csvContent;
        return this;
    }
    
    public Builder csvFileName(String csvFileName) {
    this.internal.csvFileName = csvFileName;
        return this;
    }
    
    public Builder csvWave(com.grafana.foundation.cog.Builder<List<CSVWave>> csvWave) {
    this.internal.csvWave = csvWave.build();
        return this;
    }
    
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public Builder dropPercent(Double dropPercent) {
    this.internal.dropPercent = dropPercent;
        return this;
    }
    
    public Builder errorType(DataqueryErrorType errorType) {
    this.internal.errorType = errorType;
        return this;
    }
    
    public Builder flamegraphDiff(Boolean flamegraphDiff) {
    this.internal.flamegraphDiff = flamegraphDiff;
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
    
    public Builder labels(String labels) {
    this.internal.labels = labels;
        return this;
    }
    
    public Builder levelColumn(Boolean levelColumn) {
    this.internal.levelColumn = levelColumn;
        return this;
    }
    
    public Builder lines(Long lines) {
    this.internal.lines = lines;
        return this;
    }
    
    public Builder max(Double max) {
    this.internal.max = max;
        return this;
    }
    
    public Builder maxDataPoints(Long maxDataPoints) {
    this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public Builder min(Double min) {
    this.internal.min = min;
        return this;
    }
    
    public Builder nodes(com.grafana.foundation.cog.Builder<NodesQuery> nodes) {
    this.internal.nodes = nodes.build();
        return this;
    }
    
    public Builder noise(Double noise) {
    this.internal.noise = noise;
        return this;
    }
    
    public Builder points(List<List<Object>> points) {
    this.internal.points = points;
        return this;
    }
    
    public Builder pulseWave(com.grafana.foundation.cog.Builder<PulseWaveQuery> pulseWave) {
    this.internal.pulseWave = pulseWave.build();
        return this;
    }
    
    public Builder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public Builder rawFrameContent(String rawFrameContent) {
    this.internal.rawFrameContent = rawFrameContent;
        return this;
    }
    
    public Builder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public Builder resultAssertions(com.grafana.foundation.cog.Builder<ResultAssertions> resultAssertions) {
    this.internal.resultAssertions = resultAssertions.build();
        return this;
    }
    
    public Builder scenarioId(DataqueryScenarioId scenarioId) {
    this.internal.scenarioId = scenarioId;
        return this;
    }
    
    public Builder seriesCount(Long seriesCount) {
    this.internal.seriesCount = seriesCount;
        return this;
    }
    
    public Builder sim(com.grafana.foundation.cog.Builder<SimulationQuery> sim) {
    this.internal.sim = sim.build();
        return this;
    }
    
    public Builder spanCount(Long spanCount) {
    this.internal.spanCount = spanCount;
        return this;
    }
    
    public Builder spread(Double spread) {
    this.internal.spread = spread;
        return this;
    }
    
    public Builder startValue(Double startValue) {
    this.internal.startValue = startValue;
        return this;
    }
    
    public Builder stream(com.grafana.foundation.cog.Builder<StreamingQuery> stream) {
    this.internal.stream = stream.build();
        return this;
    }
    
    public Builder stringInput(String stringInput) {
    this.internal.stringInput = stringInput;
        return this;
    }
    
    public Builder timeRange(com.grafana.foundation.cog.Builder<TimeRange> timeRange) {
    this.internal.timeRange = timeRange.build();
        return this;
    }
    
    public Builder usa(com.grafana.foundation.cog.Builder<USAQuery> usa) {
    this.internal.usa = usa.build();
        return this;
    }
    
    public Builder withNil(Boolean withNil) {
    this.internal.withNil = withNil;
        return this;
    }
    public Dataquery build() {
            return this.internal;
        }
    }
}
