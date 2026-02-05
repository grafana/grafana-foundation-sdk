// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as testdata from '../testdata';

export class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "testdata";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind {
        return this.internal;
    }

    version(version: string): this {
        this.internal.version = version;
        return this;
    }

    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource(ref: {
	name?: string;
}): this {
        this.internal.datasource = ref;
        return this;
    }

    alias(alias: string): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.alias = alias;
        return this;
    }

    // Used for live query
    channel(channel: string): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.channel = channel;
        return this;
    }

    csvContent(csvContent: string): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.csvContent = csvContent;
        return this;
    }

    csvFileName(csvFileName: string): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.csvFileName = csvFileName;
        return this;
    }

    csvWave(csvWave: cog.Builder<testdata.CSVWave>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        const csvWaveResources = csvWave.map(builder1 => builder1.build());
        this.internal.spec.csvWave = csvWaveResources;
        return this;
    }

    // Drop percentage (the chance we will lose a point 0-100)
    dropPercent(dropPercent: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.dropPercent = dropPercent;
        return this;
    }

    // Possible enum values:
    //  - `"plugin"` 
    //  - `"downstream"` 
    errorSource(errorSource: "plugin" | "downstream"): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.errorSource = errorSource;
        return this;
    }

    // Possible enum values:
    //  - `"frontend_exception"` 
    //  - `"frontend_observable"` 
    //  - `"server_panic"` 
    errorType(errorType: "frontend_exception" | "frontend_observable" | "server_panic"): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.errorType = errorType;
        return this;
    }

    flamegraphDiff(flamegraphDiff: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.flamegraphDiff = flamegraphDiff;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    intervalMs(intervalMs: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.intervalMs = intervalMs;
        return this;
    }

    labels(labels: string): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.labels = labels;
        return this;
    }

    levelColumn(levelColumn: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.levelColumn = levelColumn;
        return this;
    }

    lines(lines: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.lines = lines;
        return this;
    }

    max(max: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.max = max;
        return this;
    }

    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    maxDataPoints(maxDataPoints: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.maxDataPoints = maxDataPoints;
        return this;
    }

    min(min: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.min = min;
        return this;
    }

    nodes(nodes: cog.Builder<testdata.NodesQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        const nodesResource = nodes.build();
        this.internal.spec.nodes = nodesResource;
        return this;
    }

    noise(noise: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.noise = noise;
        return this;
    }

    points(points: any[][]): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.points = points;
        return this;
    }

    pulseWave(pulseWave: cog.Builder<testdata.PulseWaveQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        const pulseWaveResource = pulseWave.build();
        this.internal.spec.pulseWave = pulseWaveResource;
        return this;
    }

    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    queryType(queryType: string): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }

    rawFrameContent(rawFrameContent: string): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.rawFrameContent = rawFrameContent;
        return this;
    }

    // RefID is the unique identifier of the query, set by the frontend call.
    refId(refId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.refId = refId;
        return this;
    }

    // Optionally define expected query result behavior
    resultAssertions(resultAssertions: cog.Builder<testdata.ResultAssertions>): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        const resultAssertionsResource = resultAssertions.build();
        this.internal.spec.resultAssertions = resultAssertionsResource;
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
    scenarioId(scenarioId: "annotations" | "arrow" | "csv_content" | "csv_file" | "csv_metric_values" | "datapoints_outside_range" | "error_with_source" | "exponential_heatmap_bucket_data" | "flame_graph" | "grafana_api" | "linear_heatmap_bucket_data" | "live" | "logs" | "manual_entry" | "no_data_points" | "node_graph" | "predictable_csv_wave" | "predictable_pulse" | "random_walk" | "random_walk_table" | "random_walk_with_error" | "raw_frame" | "server_error_500" | "simulation" | "slow_query" | "streaming_client" | "table_static" | "trace" | "usa" | "variables-query"): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.scenarioId = scenarioId;
        return this;
    }

    seriesCount(seriesCount: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.seriesCount = seriesCount;
        return this;
    }

    sim(sim: cog.Builder<testdata.SimulationQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        const simResource = sim.build();
        this.internal.spec.sim = simResource;
        return this;
    }

    spanCount(spanCount: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.spanCount = spanCount;
        return this;
    }

    spread(spread: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.spread = spread;
        return this;
    }

    startValue(startValue: number): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.startValue = startValue;
        return this;
    }

    stream(stream: cog.Builder<testdata.StreamingQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        const streamResource = stream.build();
        this.internal.spec.stream = streamResource;
        return this;
    }

    // common parameter used by many query types
    stringInput(stringInput: string): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.stringInput = stringInput;
        return this;
    }

    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    timeRange(timeRange: cog.Builder<testdata.TimeRange>): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        const timeRangeResource = timeRange.build();
        this.internal.spec.timeRange = timeRangeResource;
        return this;
    }

    usa(usa: cog.Builder<testdata.USAQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        const usaResource = usa.build();
        this.internal.spec.usa = usaResource;
        return this;
    }

    withNil(withNil: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = testdata.defaultDataquery();
        }
        this.internal.spec.withNil = withNil;
        return this;
    }
}

