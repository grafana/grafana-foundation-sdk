// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.common.DataSourceRef;

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
    //  - `"plugin"` 
    //  - `"downstream"` 
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("errorSource")
    public DataqueryErrorSource errorSource;
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
    //  - `"error_with_source"` 
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
    public Dataquery() {
    }
    public Dataquery(String alias,String channel,String csvContent,String csvFileName,List<CSVWave> csvWave,DataSourceRef datasource,Double dropPercent,DataqueryErrorSource errorSource,DataqueryErrorType errorType,Boolean flamegraphDiff,Boolean hide,Double intervalMs,String labels,Boolean levelColumn,Long lines,Double max,Long maxDataPoints,Double min,NodesQuery nodes,Double noise,List<List<Object>> points,PulseWaveQuery pulseWave,String queryType,String rawFrameContent,String refId,ResultAssertions resultAssertions,DataqueryScenarioId scenarioId,Long seriesCount,SimulationQuery sim,Long spanCount,Double spread,Double startValue,StreamingQuery stream,String stringInput,TimeRange timeRange,USAQuery usa,Boolean withNil) {
        this.alias = alias;
        this.channel = channel;
        this.csvContent = csvContent;
        this.csvFileName = csvFileName;
        this.csvWave = csvWave;
        this.datasource = datasource;
        this.dropPercent = dropPercent;
        this.errorSource = errorSource;
        this.errorType = errorType;
        this.flamegraphDiff = flamegraphDiff;
        this.hide = hide;
        this.intervalMs = intervalMs;
        this.labels = labels;
        this.levelColumn = levelColumn;
        this.lines = lines;
        this.max = max;
        this.maxDataPoints = maxDataPoints;
        this.min = min;
        this.nodes = nodes;
        this.noise = noise;
        this.points = points;
        this.pulseWave = pulseWave;
        this.queryType = queryType;
        this.rawFrameContent = rawFrameContent;
        this.refId = refId;
        this.resultAssertions = resultAssertions;
        this.scenarioId = scenarioId;
        this.seriesCount = seriesCount;
        this.sim = sim;
        this.spanCount = spanCount;
        this.spread = spread;
        this.startValue = startValue;
        this.stream = stream;
        this.stringInput = stringInput;
        this.timeRange = timeRange;
        this.usa = usa;
        this.withNil = withNil;
    }
    public String dataqueryName() {
        return "testdata";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
