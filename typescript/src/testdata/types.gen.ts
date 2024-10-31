// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as dashboard from '../dashboard';


export interface CSVWave {
	labels?: string;
	name?: string;
	timeStep?: number;
	valuesCSV?: string;
}

export const defaultCSVWave = (): CSVWave => ({
});

export interface NodesQuery {
	count?: number;
	seed?: number;
	// Possible enum values:
	//  - `"random"` 
	//  - `"random edges"` 
	//  - `"response_medium"` 
	//  - `"response_small"` 
	//  - `"feature_showcase"` 
	type?: "random" | "random edges" | "response_medium" | "response_small" | "feature_showcase";
}

export const defaultNodesQuery = (): NodesQuery => ({
});

export interface PulseWaveQuery {
	offCount?: number;
	offValue?: number;
	onCount?: number;
	onValue?: number;
	timeStep?: number;
}

export const defaultPulseWaveQuery = (): PulseWaveQuery => ({
});

export interface ResultAssertions {
	// Maximum frame count
	maxFrames?: number;
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""` 
	//  - `"timeseries-wide"` 
	//  - `"timeseries-long"` 
	//  - `"timeseries-many"` 
	//  - `"timeseries-multi"` 
	//  - `"directory-listing"` 
	//  - `"table"` 
	//  - `"numeric-wide"` 
	//  - `"numeric-multi"` 
	//  - `"numeric-long"` 
	//  - `"log-lines"` 
	type?: "" | "timeseries-wide" | "timeseries-long" | "timeseries-many" | "timeseries-multi" | "directory-listing" | "table" | "numeric-wide" | "numeric-multi" | "numeric-long" | "log-lines";
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	typeVersion: number[];
}

export const defaultResultAssertions = (): ResultAssertions => ({
	typeVersion: [],
});

export interface Key {
	tick: number;
	type: string;
	uid?: string;
}

export const defaultKey = (): Key => ({
	tick: 0,
	type: "",
});

export interface SimulationQuery {
	config?: any;
	key: Key;
	last?: boolean;
	stream?: boolean;
}

export const defaultSimulationQuery = (): SimulationQuery => ({
	key: defaultKey(),
});

export interface StreamingQuery {
	bands?: number;
	noise: number;
	speed: number;
	spread: number;
	// Possible enum values:
	//  - `"fetch"` 
	//  - `"logs"` 
	//  - `"signal"` 
	//  - `"traces"` 
	type: "fetch" | "logs" | "signal" | "traces";
	url?: string;
}

export const defaultStreamingQuery = (): StreamingQuery => ({
	noise: 0,
	speed: 0,
	spread: 0,
	type: "fetch",
});

export interface TimeRange {
	// From is the start time of the query.
	from: string;
	// To is the end time of the query.
	to: string;
}

export const defaultTimeRange = (): TimeRange => ({
	from: "now-6h",
	to: "now",
});

export interface USAQuery {
	fields?: string[];
	mode?: string;
	period?: string;
	states?: string[];
}

export const defaultUSAQuery = (): USAQuery => ({
});

export interface dataquery {
	alias?: string;
	// Used for live query
	channel?: string;
	csvContent?: string;
	csvFileName?: string;
	csvWave?: CSVWave[];
	// The datasource
	datasource?: dashboard.DataSourceRef;
	// Drop percentage (the chance we will lose a point 0-100)
	dropPercent?: number;
	// Possible enum values:
	//  - `"plugin"` 
	//  - `"downstream"` 
	errorSource?: "plugin" | "downstream";
	// Possible enum values:
	//  - `"frontend_exception"` 
	//  - `"frontend_observable"` 
	//  - `"server_panic"` 
	errorType?: "frontend_exception" | "frontend_observable" | "server_panic";
	flamegraphDiff?: boolean;
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	hide?: boolean;
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	intervalMs?: number;
	labels?: string;
	levelColumn?: boolean;
	lines?: number;
	max?: number;
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	maxDataPoints?: number;
	min?: number;
	nodes?: NodesQuery;
	noise?: number;
	points?: any[][];
	pulseWave?: PulseWaveQuery;
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	queryType?: string;
	rawFrameContent?: string;
	// RefID is the unique identifier of the query, set by the frontend call.
	refId?: string;
	// Optionally define expected query result behavior
	resultAssertions?: ResultAssertions;
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
	scenarioId?: "annotations" | "arrow" | "csv_content" | "csv_file" | "csv_metric_values" | "datapoints_outside_range" | "error_with_source" | "exponential_heatmap_bucket_data" | "flame_graph" | "grafana_api" | "linear_heatmap_bucket_data" | "live" | "logs" | "manual_entry" | "no_data_points" | "node_graph" | "predictable_csv_wave" | "predictable_pulse" | "random_walk" | "random_walk_table" | "random_walk_with_error" | "raw_frame" | "server_error_500" | "simulation" | "slow_query" | "streaming_client" | "table_static" | "trace" | "usa" | "variables-query";
	seriesCount?: number;
	sim?: SimulationQuery;
	spanCount?: number;
	spread?: number;
	startValue?: number;
	stream?: StreamingQuery;
	// common parameter used by many query types
	stringInput?: string;
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	timeRange?: TimeRange;
	usa?: USAQuery;
	withNil?: boolean;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): dataquery => ({
	_implementsDataqueryVariant: () => {},
});

