---
title: <span class="badge object-type-struct"></span> BasePipelineMetricAggregation
---
# <span class="badge object-type-struct"></span> BasePipelineMetricAggregation

## Definition

```go
type BasePipelineMetricAggregation struct {
    PipelineAgg *string `json:"pipelineAgg,omitempty"`
    Field *string `json:"field,omitempty"`
    Type string `json:"type"`
    Id string `json:"id"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BasePipelineMetricAggregation` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (basePipelineMetricAggregation *BasePipelineMetricAggregation) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BasePipelineMetricAggregation` objects.

```go
func (basePipelineMetricAggregation *BasePipelineMetricAggregation) Equals(other BasePipelineMetricAggregation) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BasePipelineMetricAggregation` fields for violations and returns them.

```go
func (basePipelineMetricAggregation *BasePipelineMetricAggregation) Validate() error
```

## See also

 * <span class="badge builder"></span> [BasePipelineMetricAggregationBuilder](./builder-BasePipelineMetricAggregationBuilder.md)
