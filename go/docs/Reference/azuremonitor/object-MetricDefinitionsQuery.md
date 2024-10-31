---
title: <span class="badge object-type-struct"></span> MetricDefinitionsQuery
---
# <span class="badge object-type-struct"></span> MetricDefinitionsQuery

@deprecated Use MetricNamespaceQuery instead

## Definition

```go
type MetricDefinitionsQuery struct {
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

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricDefinitionsQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metricDefinitionsQuery *MetricDefinitionsQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MetricDefinitionsQuery` objects.

```go
func (metricDefinitionsQuery *MetricDefinitionsQuery) Equals(other MetricDefinitionsQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MetricDefinitionsQuery` fields for violations and returns them.

```go
func (metricDefinitionsQuery *MetricDefinitionsQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetricDefinitionsQueryBuilder](./builder-MetricDefinitionsQueryBuilder.md)
