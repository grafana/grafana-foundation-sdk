---
title: <span class="badge builder"></span> PipelineMetricAggregationWithMultipleBucketPathsBuilder
---
# <span class="badge builder"></span> PipelineMetricAggregationWithMultipleBucketPathsBuilder

## Constructor

```go
func NewPipelineMetricAggregationWithMultipleBucketPathsBuilder() *PipelineMetricAggregationWithMultipleBucketPathsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) Build() (PipelineMetricAggregationWithMultipleBucketPaths, error)
```

### <span class="badge object-method"></span> Hide

```go
func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) Hide(hide bool) *PipelineMetricAggregationWithMultipleBucketPathsBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) Id(id string) *PipelineMetricAggregationWithMultipleBucketPathsBuilder
```

### <span class="badge object-method"></span> PipelineVariables

```go
func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) PipelineVariables(pipelineVariables []cog.Builder[elasticsearch.PipelineVariable]) *PipelineMetricAggregationWithMultipleBucketPathsBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) Type(typeArg cog.Builder[elasticsearch.MetricAggregationType]) *PipelineMetricAggregationWithMultipleBucketPathsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PipelineMetricAggregationWithMultipleBucketPaths](./object-PipelineMetricAggregationWithMultipleBucketPaths.md)
