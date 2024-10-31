---
title: <span class="badge object-type-struct"></span> TempoQuery
---
# <span class="badge object-type-struct"></span> TempoQuery

## Definition

```go
type TempoQuery struct {
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    RefId string `json:"refId"`
    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    Hide *bool `json:"hide,omitempty"`
    // Specify the query flavor
    // TODO make this required and give it a default
    QueryType *string `json:"queryType,omitempty"`
    // TraceQL query or trace ID
    Query *string `json:"query,omitempty"`
    // @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
    Search *string `json:"search,omitempty"`
    // @deprecated Query traces by service name
    ServiceName *string `json:"serviceName,omitempty"`
    // @deprecated Query traces by span name
    SpanName *string `json:"spanName,omitempty"`
    // @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
    MinDuration *string `json:"minDuration,omitempty"`
    // @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
    MaxDuration *string `json:"maxDuration,omitempty"`
    // Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
    ServiceMapQuery *tempo.StringOrArrayOfString `json:"serviceMapQuery,omitempty"`
    // Use service.namespace in addition to service.name to uniquely identify a service.
    ServiceMapIncludeNamespace *bool `json:"serviceMapIncludeNamespace,omitempty"`
    // Defines the maximum number of traces that are returned from Tempo
    Limit *int64 `json:"limit,omitempty"`
    // Defines the maximum number of spans per spanset that are returned from Tempo
    Spss *int64 `json:"spss,omitempty"`
    Filters []tempo.TraceqlFilter `json:"filters"`
    // Filters that are used to query the metrics summary
    GroupBy []tempo.TraceqlFilter `json:"groupBy,omitempty"`
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
    // The type of the table that is used to display the search results
    TableType *tempo.SearchTableType `json:"tableType,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TempoQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tempoQuery *TempoQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TempoQuery` objects.

```go
func (tempoQuery *TempoQuery) Equals(other TempoQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TempoQuery` fields for violations and returns them.

```go
func (tempoQuery *TempoQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [TempoQueryBuilder](./builder-TempoQueryBuilder.md)
