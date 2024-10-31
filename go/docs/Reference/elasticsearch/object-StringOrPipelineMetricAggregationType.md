---
title: <span class="badge object-type-struct"></span> StringOrPipelineMetricAggregationType
---
# <span class="badge object-type-struct"></span> StringOrPipelineMetricAggregationType

## Definition

```go
type StringOrPipelineMetricAggregationType struct {
    String *string `json:"String,omitempty"`
    PipelineMetricAggregationType *elasticsearch.PipelineMetricAggregationType `json:"PipelineMetricAggregationType,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrPipelineMetricAggregationType` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrPipelineMetricAggregationType *StringOrPipelineMetricAggregationType) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrPipelineMetricAggregationType` objects.

```go
func (stringOrPipelineMetricAggregationType *StringOrPipelineMetricAggregationType) Equals(other StringOrPipelineMetricAggregationType) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrPipelineMetricAggregationType` fields for violations and returns them.

```go
func (stringOrPipelineMetricAggregationType *StringOrPipelineMetricAggregationType) Validate() error
```

## See also

 * <span class="badge builder"></span> [StringOrPipelineMetricAggregationTypeBuilder](./builder-StringOrPipelineMetricAggregationTypeBuilder.md)
