---
title: <span class="badge object-type-struct"></span> MetricAggregationWithField
---
# <span class="badge object-type-struct"></span> MetricAggregationWithField

## Definition

```go
type MetricAggregationWithField struct {
    Field *string `json:"field,omitempty"`
    Type elasticsearch.MetricAggregationType `json:"type"`
    Id string `json:"id"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricAggregationWithField` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metricAggregationWithField *MetricAggregationWithField) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MetricAggregationWithField` objects.

```go
func (metricAggregationWithField *MetricAggregationWithField) Equals(other MetricAggregationWithField) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MetricAggregationWithField` fields for violations and returns them.

```go
func (metricAggregationWithField *MetricAggregationWithField) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetricAggregationWithFieldBuilder](./builder-MetricAggregationWithFieldBuilder.md)
