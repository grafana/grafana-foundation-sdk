---
title: <span class="badge builder"></span> TempoQueryBuilder
---
# <span class="badge builder"></span> TempoQueryBuilder

## Constructor

```go
func NewTempoQueryBuilder() *TempoQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TempoQueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *TempoQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *TempoQueryBuilder
```

### <span class="badge object-method"></span> Filters

```go
func (builder *TempoQueryBuilder) Filters(filters []cog.Builder[tempo.TraceqlFilter]) *TempoQueryBuilder
```

### <span class="badge object-method"></span> GroupBy

Filters that are used to query the metrics summary

```go
func (builder *TempoQueryBuilder) GroupBy(groupBy []cog.Builder[tempo.TraceqlFilter]) *TempoQueryBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *TempoQueryBuilder) Hide(hide bool) *TempoQueryBuilder
```

### <span class="badge object-method"></span> Limit

Defines the maximum number of traces that are returned from Tempo

```go
func (builder *TempoQueryBuilder) Limit(limit int64) *TempoQueryBuilder
```

### <span class="badge object-method"></span> MaxDuration

@deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms

```go
func (builder *TempoQueryBuilder) MaxDuration(maxDuration string) *TempoQueryBuilder
```

### <span class="badge object-method"></span> MinDuration

@deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms

```go
func (builder *TempoQueryBuilder) MinDuration(minDuration string) *TempoQueryBuilder
```

### <span class="badge object-method"></span> Query

TraceQL query or trace ID

```go
func (builder *TempoQueryBuilder) Query(query string) *TempoQueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *TempoQueryBuilder) QueryType(queryType string) *TempoQueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *TempoQueryBuilder) RefId(refId string) *TempoQueryBuilder
```

### <span class="badge object-method"></span> Search

@deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true

```go
func (builder *TempoQueryBuilder) Search(search string) *TempoQueryBuilder
```

### <span class="badge object-method"></span> ServiceMapIncludeNamespace

Use service.namespace in addition to service.name to uniquely identify a service.

```go
func (builder *TempoQueryBuilder) ServiceMapIncludeNamespace(serviceMapIncludeNamespace bool) *TempoQueryBuilder
```

### <span class="badge object-method"></span> ServiceMapQuery

Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.

```go
func (builder *TempoQueryBuilder) ServiceMapQuery(serviceMapQuery tempo.StringOrArrayOfString) *TempoQueryBuilder
```

### <span class="badge object-method"></span> ServiceName

@deprecated Query traces by service name

```go
func (builder *TempoQueryBuilder) ServiceName(serviceName string) *TempoQueryBuilder
```

### <span class="badge object-method"></span> SpanName

@deprecated Query traces by span name

```go
func (builder *TempoQueryBuilder) SpanName(spanName string) *TempoQueryBuilder
```

### <span class="badge object-method"></span> Spss

Defines the maximum number of spans per spanset that are returned from Tempo

```go
func (builder *TempoQueryBuilder) Spss(spss int64) *TempoQueryBuilder
```

### <span class="badge object-method"></span> TableType

The type of the table that is used to display the search results

```go
func (builder *TempoQueryBuilder) TableType(tableType tempo.SearchTableType) *TempoQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TempoQuery](./object-TempoQuery.md)
