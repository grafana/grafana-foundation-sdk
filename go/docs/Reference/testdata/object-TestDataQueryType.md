---
title: <span class="badge object-type-enum"></span> TestDataQueryType
---
# <span class="badge object-type-enum"></span> TestDataQueryType

## Definition

```go
type TestDataQueryType string
const (
	TestDataQueryTypeRandomWalk TestDataQueryType = "random_walk"
	TestDataQueryTypeSlowQuery TestDataQueryType = "slow_query"
	TestDataQueryTypeRandomWalkWithError TestDataQueryType = "random_walk_with_error"
	TestDataQueryTypeRandomWalkTable TestDataQueryType = "random_walk_table"
	TestDataQueryTypeExponentialHeatmapBucketData TestDataQueryType = "exponential_heatmap_bucket_data"
	TestDataQueryTypeLinearHeatmapBucketData TestDataQueryType = "linear_heatmap_bucket_data"
	TestDataQueryTypeNoDataPoints TestDataQueryType = "no_data_points"
	TestDataQueryTypeDataPointsOutsideRange TestDataQueryType = "datapoints_outside_range"
	TestDataQueryTypeCSVMetricValues TestDataQueryType = "csv_metric_values"
	TestDataQueryTypePredictablePulse TestDataQueryType = "predictable_pulse"
	TestDataQueryTypePredictableCSVWave TestDataQueryType = "predictable_csv_wave"
	TestDataQueryTypeStreamingClient TestDataQueryType = "streaming_client"
	TestDataQueryTypeSimulation TestDataQueryType = "simulation"
	TestDataQueryTypeUSA TestDataQueryType = "usa"
	TestDataQueryTypeLive TestDataQueryType = "live"
	TestDataQueryTypeGrafanaAPI TestDataQueryType = "grafana_api"
	TestDataQueryTypeArrow TestDataQueryType = "arrow"
	TestDataQueryTypeAnnotations TestDataQueryType = "annotations"
	TestDataQueryTypeTableStatic TestDataQueryType = "table_static"
	TestDataQueryTypeServerError500 TestDataQueryType = "server_error_500"
	TestDataQueryTypeLogs TestDataQueryType = "logs"
	TestDataQueryTypeNodeGraph TestDataQueryType = "node_graph"
	TestDataQueryTypeFlameGraph TestDataQueryType = "flame_graph"
	TestDataQueryTypeRawFrame TestDataQueryType = "raw_frame"
	TestDataQueryTypeCSVFile TestDataQueryType = "csv_file"
	TestDataQueryTypeCSVContent TestDataQueryType = "csv_content"
	TestDataQueryTypeTrace TestDataQueryType = "trace"
	TestDataQueryTypeManualEntry TestDataQueryType = "manual_entry"
	TestDataQueryTypeVariablesQuery TestDataQueryType = "variables-query"
)

```
