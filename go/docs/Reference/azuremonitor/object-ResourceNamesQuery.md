---
title: <span class="badge object-type-struct"></span> ResourceNamesQuery
---
# <span class="badge object-type-struct"></span> ResourceNamesQuery

## Definition

```go
type ResourceNamesQuery struct {
    RawQuery *string `json:"rawQuery,omitempty"`
    Kind string `json:"kind"`
    Subscription string `json:"subscription"`
    ResourceGroup string `json:"resourceGroup"`
    MetricNamespace string `json:"metricNamespace"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ResourceNamesQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (resourceNamesQuery *ResourceNamesQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ResourceNamesQuery` objects.

```go
func (resourceNamesQuery *ResourceNamesQuery) Equals(other ResourceNamesQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ResourceNamesQuery` fields for violations and returns them.

```go
func (resourceNamesQuery *ResourceNamesQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [ResourceNamesQueryBuilder](./builder-ResourceNamesQueryBuilder.md)
