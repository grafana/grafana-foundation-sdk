---
title: <span class="badge object-type-struct"></span> TimeSeriesQuery
---
# <span class="badge object-type-struct"></span> TimeSeriesQuery

Time Series sub-query properties.

## Definition

```go
type TimeSeriesQuery struct {
    // GCP project to execute the query against.
    ProjectName string `json:"projectName"`
    // MQL query to be executed.
    Query string `json:"query"`
    // To disable the graphPeriod, it should explictly be set to 'disabled'.
    GraphPeriod *string `json:"graphPeriod,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeSeriesQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (timeSeriesQuery *TimeSeriesQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TimeSeriesQuery` objects.

```go
func (timeSeriesQuery *TimeSeriesQuery) Equals(other TimeSeriesQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TimeSeriesQuery` fields for violations and returns them.

```go
func (timeSeriesQuery *TimeSeriesQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [TimeSeriesQueryBuilder](./builder-TimeSeriesQueryBuilder.md)
