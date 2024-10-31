---
title: <span class="badge object-type-enum"></span> TestDataQueryType
---
# <span class="badge object-type-enum"></span> TestDataQueryType

## Definition

```typescript
export enum TestDataQueryType {
	RandomWalk = "random_walk",
	SlowQuery = "slow_query",
	RandomWalkWithError = "random_walk_with_error",
	RandomWalkTable = "random_walk_table",
	ExponentialHeatmapBucketData = "exponential_heatmap_bucket_data",
	LinearHeatmapBucketData = "linear_heatmap_bucket_data",
	NoDataPoints = "no_data_points",
	DataPointsOutsideRange = "datapoints_outside_range",
	CSVMetricValues = "csv_metric_values",
	PredictablePulse = "predictable_pulse",
	PredictableCSVWave = "predictable_csv_wave",
	StreamingClient = "streaming_client",
	Simulation = "simulation",
	USA = "usa",
	Live = "live",
	GrafanaAPI = "grafana_api",
	Arrow = "arrow",
	Annotations = "annotations",
	TableStatic = "table_static",
	ServerError500 = "server_error_500",
	Logs = "logs",
	NodeGraph = "node_graph",
	FlameGraph = "flame_graph",
	RawFrame = "raw_frame",
	CSVFile = "csv_file",
	CSVContent = "csv_content",
	Trace = "trace",
	ManualEntry = "manual_entry",
	VariablesQuery = "variables-query",
}

```
