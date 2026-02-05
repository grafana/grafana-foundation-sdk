import * as cog from '../cog';
import * as testdata from '../testdata';
import * as common from '../common';
export declare class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: testdata.dataquery;
    constructor();
    /**
     * Builds the object.
     */
    build(): testdata.dataquery;
    alias(alias: string): this;
    channel(channel: string): this;
    csvContent(csvContent: string): this;
    csvFileName(csvFileName: string): this;
    csvWave(csvWave: cog.Builder<testdata.CSVWave>[]): this;
    datasource(datasource: common.DataSourceRef): this;
    dropPercent(dropPercent: number): this;
    errorSource(errorSource: "plugin" | "downstream"): this;
    errorType(errorType: "frontend_exception" | "frontend_observable" | "server_panic"): this;
    flamegraphDiff(flamegraphDiff: boolean): this;
    hide(hide: boolean): this;
    intervalMs(intervalMs: number): this;
    labels(labels: string): this;
    levelColumn(levelColumn: boolean): this;
    lines(lines: number): this;
    max(max: number): this;
    maxDataPoints(maxDataPoints: number): this;
    min(min: number): this;
    nodes(nodes: cog.Builder<testdata.NodesQuery>): this;
    noise(noise: number): this;
    points(points: any[][]): this;
    pulseWave(pulseWave: cog.Builder<testdata.PulseWaveQuery>): this;
    queryType(queryType: string): this;
    rawFrameContent(rawFrameContent: string): this;
    refId(refId: string): this;
    resultAssertions(resultAssertions: cog.Builder<testdata.ResultAssertions>): this;
    scenarioId(scenarioId: "annotations" | "arrow" | "csv_content" | "csv_file" | "csv_metric_values" | "datapoints_outside_range" | "error_with_source" | "exponential_heatmap_bucket_data" | "flame_graph" | "grafana_api" | "linear_heatmap_bucket_data" | "live" | "logs" | "manual_entry" | "no_data_points" | "node_graph" | "predictable_csv_wave" | "predictable_pulse" | "random_walk" | "random_walk_table" | "random_walk_with_error" | "raw_frame" | "server_error_500" | "simulation" | "slow_query" | "streaming_client" | "table_static" | "trace" | "usa" | "variables-query"): this;
    seriesCount(seriesCount: number): this;
    sim(sim: cog.Builder<testdata.SimulationQuery>): this;
    spanCount(spanCount: number): this;
    spread(spread: number): this;
    startValue(startValue: number): this;
    stream(stream: cog.Builder<testdata.StreamingQuery>): this;
    stringInput(stringInput: string): this;
    timeRange(timeRange: cog.Builder<testdata.TimeRange>): this;
    usa(usa: cog.Builder<testdata.USAQuery>): this;
    withNil(withNil: boolean): this;
}
