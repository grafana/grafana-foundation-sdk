---
title: <span class="badge object-type-struct"></span> PipelineMetricAggregationWithMultipleBucketPaths
---
# <span class="badge object-type-struct"></span> PipelineMetricAggregationWithMultipleBucketPaths

## Definition

```go
type PipelineMetricAggregationWithMultipleBucketPaths struct {
    PipelineVariables []elasticsearch.PipelineVariable `json:"pipelineVariables,omitempty"`
    Type elasticsearch.MetricAggregationType `json:"type"`
    Id string `json:"id"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PipelineMetricAggregationWithMultipleBucketPaths` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (pipelineMetricAggregationWithMultipleBucketPaths *PipelineMetricAggregationWithMultipleBucketPaths) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PipelineMetricAggregationWithMultipleBucketPaths` objects.

```go
func (pipelineMetricAggregationWithMultipleBucketPaths *PipelineMetricAggregationWithMultipleBucketPaths) Equals(other PipelineMetricAggregationWithMultipleBucketPaths) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PipelineMetricAggregationWithMultipleBucketPaths` fields for violations and returns them.

```go
func (pipelineMetricAggregationWithMultipleBucketPaths *PipelineMetricAggregationWithMultipleBucketPaths) Validate() error
```

## See also

 * <span class="badge builder"></span> [PipelineMetricAggregationWithMultipleBucketPathsBuilder](./builder-PipelineMetricAggregationWithMultipleBucketPathsBuilder.md)
