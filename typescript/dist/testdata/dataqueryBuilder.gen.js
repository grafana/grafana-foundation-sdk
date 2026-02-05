"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataqueryBuilder = void 0;
const tslib_1 = require("tslib");
const testdata = tslib_1.__importStar(require("../testdata"));
class DataqueryBuilder {
    constructor() {
        this.internal = testdata.defaultDataquery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    alias(alias) {
        this.internal.alias = alias;
        return this;
    }
    // Used for live query
    channel(channel) {
        this.internal.channel = channel;
        return this;
    }
    csvContent(csvContent) {
        this.internal.csvContent = csvContent;
        return this;
    }
    csvFileName(csvFileName) {
        this.internal.csvFileName = csvFileName;
        return this;
    }
    csvWave(csvWave) {
        const csvWaveResources = csvWave.map(builder1 => builder1.build());
        this.internal.csvWave = csvWaveResources;
        return this;
    }
    // The datasource
    datasource(datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    // Drop percentage (the chance we will lose a point 0-100)
    dropPercent(dropPercent) {
        this.internal.dropPercent = dropPercent;
        return this;
    }
    // Possible enum values:
    //  - `"plugin"` 
    //  - `"downstream"` 
    errorSource(errorSource) {
        this.internal.errorSource = errorSource;
        return this;
    }
    // Possible enum values:
    //  - `"frontend_exception"` 
    //  - `"frontend_observable"` 
    //  - `"server_panic"` 
    errorType(errorType) {
        this.internal.errorType = errorType;
        return this;
    }
    flamegraphDiff(flamegraphDiff) {
        this.internal.flamegraphDiff = flamegraphDiff;
        return this;
    }
    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    intervalMs(intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    labels(labels) {
        this.internal.labels = labels;
        return this;
    }
    levelColumn(levelColumn) {
        this.internal.levelColumn = levelColumn;
        return this;
    }
    lines(lines) {
        this.internal.lines = lines;
        return this;
    }
    max(max) {
        this.internal.max = max;
        return this;
    }
    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    maxDataPoints(maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    min(min) {
        this.internal.min = min;
        return this;
    }
    nodes(nodes) {
        const nodesResource = nodes.build();
        this.internal.nodes = nodesResource;
        return this;
    }
    noise(noise) {
        this.internal.noise = noise;
        return this;
    }
    points(points) {
        this.internal.points = points;
        return this;
    }
    pulseWave(pulseWave) {
        const pulseWaveResource = pulseWave.build();
        this.internal.pulseWave = pulseWaveResource;
        return this;
    }
    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    queryType(queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    rawFrameContent(rawFrameContent) {
        this.internal.rawFrameContent = rawFrameContent;
        return this;
    }
    // RefID is the unique identifier of the query, set by the frontend call.
    refId(refId) {
        this.internal.refId = refId;
        return this;
    }
    // Optionally define expected query result behavior
    resultAssertions(resultAssertions) {
        const resultAssertionsResource = resultAssertions.build();
        this.internal.resultAssertions = resultAssertionsResource;
        return this;
    }
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
    scenarioId(scenarioId) {
        this.internal.scenarioId = scenarioId;
        return this;
    }
    seriesCount(seriesCount) {
        this.internal.seriesCount = seriesCount;
        return this;
    }
    sim(sim) {
        const simResource = sim.build();
        this.internal.sim = simResource;
        return this;
    }
    spanCount(spanCount) {
        this.internal.spanCount = spanCount;
        return this;
    }
    spread(spread) {
        this.internal.spread = spread;
        return this;
    }
    startValue(startValue) {
        this.internal.startValue = startValue;
        return this;
    }
    stream(stream) {
        const streamResource = stream.build();
        this.internal.stream = streamResource;
        return this;
    }
    // common parameter used by many query types
    stringInput(stringInput) {
        this.internal.stringInput = stringInput;
        return this;
    }
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    timeRange(timeRange) {
        const timeRangeResource = timeRange.build();
        this.internal.timeRange = timeRangeResource;
        return this;
    }
    usa(usa) {
        const usaResource = usa.build();
        this.internal.usa = usaResource;
        return this;
    }
    withNil(withNil) {
        this.internal.withNil = withNil;
        return this;
    }
}
exports.DataqueryBuilder = DataqueryBuilder;
//# sourceMappingURL=dataqueryBuilder.gen.js.map