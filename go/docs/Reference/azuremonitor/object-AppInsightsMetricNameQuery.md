---
title: <span class="badge object-type-struct"></span> AppInsightsMetricNameQuery
---
# <span class="badge object-type-struct"></span> AppInsightsMetricNameQuery

## Definition

```go
type AppInsightsMetricNameQuery struct {
    RawQuery *string `json:"rawQuery,omitempty"`
    Kind string `json:"kind"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AppInsightsMetricNameQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (appInsightsMetricNameQuery *AppInsightsMetricNameQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AppInsightsMetricNameQuery` objects.

```go
func (appInsightsMetricNameQuery *AppInsightsMetricNameQuery) Equals(other AppInsightsMetricNameQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AppInsightsMetricNameQuery` fields for violations and returns them.

```go
func (appInsightsMetricNameQuery *AppInsightsMetricNameQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [AppInsightsMetricNameQueryBuilder](./builder-AppInsightsMetricNameQueryBuilder.md)
