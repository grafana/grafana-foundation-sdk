---
title: <span class="badge object-type-enum"></span> PipelineMetricAggregationType
---
# <span class="badge object-type-enum"></span> PipelineMetricAggregationType

## Definition

```go
type PipelineMetricAggregationType string
const (
	PipelineMetricAggregationTypeMovingAvg PipelineMetricAggregationType = "moving_avg"
	PipelineMetricAggregationTypeMovingFn PipelineMetricAggregationType = "moving_fn"
	PipelineMetricAggregationTypeDerivative PipelineMetricAggregationType = "derivative"
	PipelineMetricAggregationTypeSerialDiff PipelineMetricAggregationType = "serial_diff"
	PipelineMetricAggregationTypeCumulativeSum PipelineMetricAggregationType = "cumulative_sum"
	PipelineMetricAggregationTypeBucketScript PipelineMetricAggregationType = "bucket_script"
)

```
