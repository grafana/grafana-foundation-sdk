---
title: <span class="badge object-type-struct"></span> BaseMetricAggregation
---
# <span class="badge object-type-struct"></span> BaseMetricAggregation

## Definition

```go
type BaseMetricAggregation struct {
    Type elasticsearch.MetricAggregationType `json:"type"`
    Id string `json:"id"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BaseMetricAggregation` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (baseMetricAggregation *BaseMetricAggregation) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BaseMetricAggregation` objects.

```go
func (baseMetricAggregation *BaseMetricAggregation) Equals(other BaseMetricAggregation) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BaseMetricAggregation` fields for violations and returns them.

```go
func (baseMetricAggregation *BaseMetricAggregation) Validate() error
```

## See also

 * <span class="badge builder"></span> [BaseMetricAggregationBuilder](./builder-BaseMetricAggregationBuilder.md)
