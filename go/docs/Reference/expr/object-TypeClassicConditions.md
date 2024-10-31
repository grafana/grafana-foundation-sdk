---
title: <span class="badge object-type-struct"></span> TypeClassicConditions
---
# <span class="badge object-type-struct"></span> TypeClassicConditions

## Definition

```go
type TypeClassicConditions struct {
    Conditions []expr.ExprTypeClassicConditionsConditions `json:"conditions"`
    // The datasource
    Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    Hide *bool `json:"hide,omitempty"`
    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    IntervalMs *float64 `json:"intervalMs,omitempty"`
    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    QueryType *string `json:"queryType,omitempty"`
    // RefID is the unique identifier of the query, set by the frontend call.
    RefId string `json:"refId"`
    // Optionally define expected query result behavior
    ResultAssertions *expr.ExprTypeClassicConditionsResultAssertions `json:"resultAssertions,omitempty"`
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    TimeRange *expr.ExprTypeClassicConditionsTimeRange `json:"timeRange,omitempty"`
    Type string `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TypeClassicConditions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (typeClassicConditions *TypeClassicConditions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TypeClassicConditions` objects.

```go
func (typeClassicConditions *TypeClassicConditions) Equals(other TypeClassicConditions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TypeClassicConditions` fields for violations and returns them.

```go
func (typeClassicConditions *TypeClassicConditions) Validate() error
```

## See also

 * <span class="badge builder"></span> [TypeClassicConditionsBuilder](./builder-TypeClassicConditionsBuilder.md)
