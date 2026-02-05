import * as common from '../common';
export interface CSVWave {
    labels?: string;
    name?: string;
    timeStep?: number;
    valuesCSV?: string;
}
export declare const defaultCSVWave: () => CSVWave;
export interface NodesQuery {
    count?: number;
    seed?: number;
    type?: "random" | "random edges" | "response_medium" | "response_small" | "feature_showcase";
}
export declare const defaultNodesQuery: () => NodesQuery;
export interface PulseWaveQuery {
    offCount?: number;
    offValue?: number;
    onCount?: number;
    onValue?: number;
    timeStep?: number;
}
export declare const defaultPulseWaveQuery: () => PulseWaveQuery;
export interface ResultAssertions {
    maxFrames?: number;
    type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
    typeVersion: number[];
}
export declare const defaultResultAssertions: () => ResultAssertions;
export interface Key {
    tick: number;
    type: string;
    uid?: string;
}
export declare const defaultKey: () => Key;
export interface SimulationQuery {
    config?: any;
    key: Key;
    last?: boolean;
    stream?: boolean;
}
export declare const defaultSimulationQuery: () => SimulationQuery;
export interface StreamingQuery {
    bands?: number;
    noise: number;
    speed: number;
    spread: number;
    type: "fetch" | "logs" | "signal" | "traces";
    url?: string;
}
export declare const defaultStreamingQuery: () => StreamingQuery;
export interface TimeRange {
    from: string;
    to: string;
}
export declare const defaultTimeRange: () => TimeRange;
export interface USAQuery {
    fields?: string[];
    mode?: string;
    period?: string;
    states?: string[];
}
export declare const defaultUSAQuery: () => USAQuery;
export interface dataquery {
    alias?: string;
    channel?: string;
    csvContent?: string;
    csvFileName?: string;
    csvWave?: CSVWave[];
    datasource?: common.DataSourceRef;
    dropPercent?: number;
    errorSource?: "plugin" | "downstream";
    errorType?: "frontend_exception" | "frontend_observable" | "server_panic";
    flamegraphDiff?: boolean;
    hide?: boolean;
    intervalMs?: number;
    labels?: string;
    levelColumn?: boolean;
    lines?: number;
    max?: number;
    maxDataPoints?: number;
    min?: number;
    nodes?: NodesQuery;
    noise?: number;
    points?: any[][];
    pulseWave?: PulseWaveQuery;
    queryType?: string;
    rawFrameContent?: string;
    refId?: string;
    resultAssertions?: ResultAssertions;
    scenarioId?: "annotations" | "arrow" | "csv_content" | "csv_file" | "csv_metric_values" | "datapoints_outside_range" | "error_with_source" | "exponential_heatmap_bucket_data" | "flame_graph" | "grafana_api" | "linear_heatmap_bucket_data" | "live" | "logs" | "manual_entry" | "no_data_points" | "node_graph" | "predictable_csv_wave" | "predictable_pulse" | "random_walk" | "random_walk_table" | "random_walk_with_error" | "raw_frame" | "server_error_500" | "simulation" | "slow_query" | "streaming_client" | "table_static" | "trace" | "usa" | "variables-query";
    seriesCount?: number;
    sim?: SimulationQuery;
    spanCount?: number;
    spread?: number;
    startValue?: number;
    stream?: StreamingQuery;
    stringInput?: string;
    timeRange?: TimeRange;
    usa?: USAQuery;
    withNil?: boolean;
    _implementsDataqueryVariant(): void;
}
export declare const defaultDataquery: () => dataquery;
