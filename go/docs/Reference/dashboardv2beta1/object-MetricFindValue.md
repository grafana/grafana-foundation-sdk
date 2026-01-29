---
title: <span class="badge object-type-struct"></span> MetricFindValue
---
# <span class="badge object-type-struct"></span> MetricFindValue

Define the MetricFindValue type

## Definition

```go
type MetricFindValue struct {
    Text string `json:"text"`
    Value *dashboardv2beta1.StringOrFloat64 `json:"value,omitempty"`
    Group *string `json:"group,omitempty"`
    Expandable *bool `json:"expandable,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricFindValue` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metricFindValue *MetricFindValue) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MetricFindValue` objects.

```go
func (metricFindValue *MetricFindValue) Equals(other MetricFindValue) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MetricFindValue` fields for violations and returns them.

```go
func (metricFindValue *MetricFindValue) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetricFindValueBuilder](./builder-MetricFindValueBuilder.md)
