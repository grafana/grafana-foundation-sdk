---
title: <span class="badge object-type-struct"></span> MetricNamespaceQuery
---
# <span class="badge object-type-struct"></span> MetricNamespaceQuery

## Definition

```go
type MetricNamespaceQuery struct {
    RawQuery *string `json:"rawQuery,omitempty"`
    Kind string `json:"kind"`
    Subscription string `json:"subscription"`
    ResourceGroup string `json:"resourceGroup"`
    MetricNamespace *string `json:"metricNamespace,omitempty"`
    ResourceName *string `json:"resourceName,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricNamespaceQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metricNamespaceQuery *MetricNamespaceQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MetricNamespaceQuery` objects.

```go
func (metricNamespaceQuery *MetricNamespaceQuery) Equals(other MetricNamespaceQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MetricNamespaceQuery` fields for violations and returns them.

```go
func (metricNamespaceQuery *MetricNamespaceQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetricNamespaceQueryBuilder](./builder-MetricNamespaceQueryBuilder.md)
