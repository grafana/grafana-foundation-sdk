---
title: <span class="badge object-type-struct"></span> TracesQuery
---
# <span class="badge object-type-struct"></span> TracesQuery

Application Insights Traces sub-query properties

## Definition

```go
type TracesQuery struct {
    // Specifies the format results should be returned as.
    ResultFormat *azuremonitor.ResultFormat `json:"resultFormat,omitempty"`
    // Array of resource URIs to be queried.
    Resources []string `json:"resources,omitempty"`
    // Operation ID. Used only for Traces queries.
    OperationId *string `json:"operationId,omitempty"`
    // Types of events to filter by.
    TraceTypes []string `json:"traceTypes,omitempty"`
    // Filters for property values.
    Filters []azuremonitor.TracesFilter `json:"filters,omitempty"`
    // KQL query to be executed.
    Query *string `json:"query,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TracesQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tracesQuery *TracesQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TracesQuery` objects.

```go
func (tracesQuery *TracesQuery) Equals(other TracesQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TracesQuery` fields for violations and returns them.

```go
func (tracesQuery *TracesQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [TracesQueryBuilder](./builder-TracesQueryBuilder.md)
