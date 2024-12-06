---
title: <span class="badge object-type-struct"></span> MetricNamesQuery
---
# <span class="badge object-type-struct"></span> MetricNamesQuery

## Definition

```go
type MetricNamesQuery struct {
    RawQuery *string `json:"rawQuery,omitempty"`
    Kind string `json:"kind"`
    Subscription string `json:"subscription"`
    ResourceGroup string `json:"resourceGroup"`
    ResourceName string `json:"resourceName"`
    MetricNamespace string `json:"metricNamespace"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricNamesQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metricNamesQuery *MetricNamesQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MetricNamesQuery` objects.

```go
func (metricNamesQuery *MetricNamesQuery) Equals(other MetricNamesQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MetricNamesQuery` fields for violations and returns them.

```go
func (metricNamesQuery *MetricNamesQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetricNamesQueryBuilder](./builder-MetricNamesQueryBuilder.md)
