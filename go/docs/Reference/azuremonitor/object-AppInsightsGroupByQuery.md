---
title: <span class="badge object-type-struct"></span> AppInsightsGroupByQuery
---
# <span class="badge object-type-struct"></span> AppInsightsGroupByQuery

## Definition

```go
type AppInsightsGroupByQuery struct {
    RawQuery *string `json:"rawQuery,omitempty"`
    Kind string `json:"kind"`
    MetricName string `json:"metricName"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AppInsightsGroupByQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (appInsightsGroupByQuery *AppInsightsGroupByQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AppInsightsGroupByQuery` objects.

```go
func (appInsightsGroupByQuery *AppInsightsGroupByQuery) Equals(other AppInsightsGroupByQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AppInsightsGroupByQuery` fields for violations and returns them.

```go
func (appInsightsGroupByQuery *AppInsightsGroupByQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [AppInsightsGroupByQueryBuilder](./builder-AppInsightsGroupByQueryBuilder.md)
