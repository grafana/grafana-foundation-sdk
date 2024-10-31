// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as testdata from '../testdata';
import * as dashboard from '../dashboard';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: testdata.dataquery;

    constructor() {
        this.internal = testdata.defaultDataquery();
    }

    /**
     * Builds the object.
     */
    build(): testdata.dataquery {
        return this.internal;
    }

    alias(alias: string): this {
        this.internal.alias = alias;
        return this;
    }

    // Used for live query
    channel(channel: string): this {
        this.internal.channel = channel;
        return this;
    }

    csvContent(csvContent: string): this {
        this.internal.csvContent = csvContent;
        return this;
    }

    csvFileName(csvFileName: string): this {
        this.internal.csvFileName = csvFileName;
        return this;
    }

    csvWave(csvWave: cog.Builder<testdata.CSVWave>[]): this {
        const csvWaveResources = csvWave.map(builder1 => builder1.build());
        this.internal.csvWave = csvWaveResources;
        return this;
    }

    // The datasource
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // Drop percentage (the chance we will lose a point 0-100)
    dropPercent(dropPercent: number): this {
        this.internal.dropPercent = dropPercent;
        return this;
    }

    // Possible enum values:
    //  - `"frontend_exception"` 
    //  - `"frontend_observable"` 
    //  - `"server_panic"` 
    errorType(errorType: "frontend_exception" | "frontend_observable" | "server_panic"): this {
        this.internal.errorType = errorType;
        return this;
    }

    flamegraphDiff(flamegraphDiff: boolean): this {
        this.internal.flamegraphDiff = flamegraphDiff;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    intervalMs(intervalMs: number): this {
        this.internal.intervalMs = intervalMs;
        return this;
    }

    labels(labels: string): this {
        this.internal.labels = labels;
        return this;
    }

    levelColumn(levelColumn: boolean): this {
        this.internal.levelColumn = levelColumn;
        return this;
    }

    lines(lines: number): this {
        this.internal.lines = lines;
        return this;
    }

    max(max: number): this {
        this.internal.max = max;
        return this;
    }

    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    maxDataPoints(maxDataPoints: number): this {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }

    min(min: number): this {
        this.internal.min = min;
        return this;
    }

    nodes(nodes: cog.Builder<testdata.NodesQuery>): this {
        const nodesResource = nodes.build();
        this.internal.nodes = nodesResource;
        return this;
    }

    noise(noise: number): this {
        this.internal.noise = noise;
        return this;
    }

    points(points: any[][]): this {
        this.internal.points = points;
        return this;
    }

    pulseWave(pulseWave: cog.Builder<testdata.PulseWaveQuery>): this {
        const pulseWaveResource = pulseWave.build();
        this.internal.pulseWave = pulseWaveResource;
        return this;
    }

    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    rawFrameContent(rawFrameContent: string): this {
        this.internal.rawFrameContent = rawFrameContent;
        return this;
    }

    // RefID is the unique identifier of the query, set by the frontend call.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // Optionally define expected query result behavior
    resultAssertions(resultAssertions: cog.Builder<testdata.ResultAssertions>): this {
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
    scenarioId(scenarioId: "annotations" | "arrow" | "csv_content" | "csv_file" | "csv_metric_values" | "datapoints_outside_range" | "exponential_heatmap_bucket_data" | "flame_graph" | "grafana_api" | "linear_heatmap_bucket_data" | "live" | "logs" | "manual_entry" | "no_data_points" | "node_graph" | "predictable_csv_wave" | "predictable_pulse" | "random_walk" | "random_walk_table" | "random_walk_with_error" | "raw_frame" | "server_error_500" | "simulation" | "slow_query" | "streaming_client" | "table_static" | "trace" | "usa" | "variables-query"): this {
        this.internal.scenarioId = scenarioId;
        return this;
    }

    seriesCount(seriesCount: number): this {
        this.internal.seriesCount = seriesCount;
        return this;
    }

    sim(sim: cog.Builder<testdata.SimulationQuery>): this {
        const simResource = sim.build();
        this.internal.sim = simResource;
        return this;
    }

    spanCount(spanCount: number): this {
        this.internal.spanCount = spanCount;
        return this;
    }

    spread(spread: number): this {
        this.internal.spread = spread;
        return this;
    }

    startValue(startValue: number): this {
        this.internal.startValue = startValue;
        return this;
    }

    stream(stream: cog.Builder<testdata.StreamingQuery>): this {
        const streamResource = stream.build();
        this.internal.stream = streamResource;
        return this;
    }

    // common parameter used by many query types
    stringInput(stringInput: string): this {
        this.internal.stringInput = stringInput;
        return this;
    }

    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    timeRange(timeRange: cog.Builder<testdata.TimeRange>): this {
        const timeRangeResource = timeRange.build();
        this.internal.timeRange = timeRangeResource;
        return this;
    }

    usa(usa: cog.Builder<testdata.USAQuery>): this {
        const usaResource = usa.build();
        this.internal.usa = usaResource;
        return this;
    }

    withNil(withNil: boolean): this {
        this.internal.withNil = withNil;
        return this;
    }
}
