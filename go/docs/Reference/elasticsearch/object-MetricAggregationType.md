---
title: <span class="badge object-type-enum"></span> MetricAggregationType
---
# <span class="badge object-type-enum"></span> MetricAggregationType

## Definition

```go
type MetricAggregationType string
const (
	MetricAggregationTypeCount MetricAggregationType = "count"
	MetricAggregationTypeAvg MetricAggregationType = "avg"
	MetricAggregationTypeSum MetricAggregationType = "sum"
	MetricAggregationTypeMin MetricAggregationType = "min"
	MetricAggregationTypeMax MetricAggregationType = "max"
	MetricAggregationTypeExtendedStats MetricAggregationType = "extended_stats"
	MetricAggregationTypePercentiles MetricAggregationType = "percentiles"
	MetricAggregationTypeCardinality MetricAggregationType = "cardinality"
	MetricAggregationTypeRawDocument MetricAggregationType = "raw_document"
	MetricAggregationTypeRawData MetricAggregationType = "raw_data"
	MetricAggregationTypeLogs MetricAggregationType = "logs"
	MetricAggregationTypeRate MetricAggregationType = "rate"
	MetricAggregationTypeTopMetrics MetricAggregationType = "top_metrics"
	MetricAggregationTypeMovingAvg MetricAggregationType = "moving_avg"
	MetricAggregationTypeMovingFn MetricAggregationType = "moving_fn"
	MetricAggregationTypeDerivative MetricAggregationType = "derivative"
	MetricAggregationTypeSerialDiff MetricAggregationType = "serial_diff"
	MetricAggregationTypeCumulativeSum MetricAggregationType = "cumulative_sum"
	MetricAggregationTypeBucketScript MetricAggregationType = "bucket_script"
)

```
