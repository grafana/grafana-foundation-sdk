---
title: <span class="badge object-type-interface"></span> dataquery
---
# <span class="badge object-type-interface"></span> dataquery

## Definition

```typescript
export interface dataquery {
	alias?: string;
	// Used for live query
	channel?: string;
	csvContent?: string;
	csvFileName?: string;
	csvWave?: testdata.CSVWave[];
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
	nodes?: testdata.NodesQuery;
	noise?: number;
	points?: any[][];
	pulseWave?: testdata.PulseWaveQuery;
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	queryType?: string;
	rawFrameContent?: string;
	// RefID is the unique identifier of the query, set by the frontend call.
	refId?: string;
	// Optionally define expected query result behavior
	resultAssertions?: testdata.ResultAssertions;
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
	sim?: testdata.SimulationQuery;
	spanCount?: number;
	spread?: number;
	startValue?: number;
	stream?: testdata.StreamingQuery;
	// common parameter used by many query types
	stringInput?: string;
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	timeRange?: testdata.TimeRange;
	usa?: testdata.USAQuery;
	withNil?: boolean;
	_implementsDataqueryVariant(): void;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
