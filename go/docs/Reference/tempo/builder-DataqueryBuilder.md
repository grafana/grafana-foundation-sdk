---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```go
func NewDataqueryBuilder() *DataqueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DataqueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *DataqueryBuilder) Datasource(datasource common.DataSourceRef) *DataqueryBuilder
```

### <span class="badge object-method"></span> Exemplars

For metric queries, how many exemplars to request, 0 means no exemplars

```go
func (builder *DataqueryBuilder) Exemplars(exemplars int64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Filters

```go
func (builder *DataqueryBuilder) Filters(filters []cog.Builder[tempo.TraceqlFilter]) *DataqueryBuilder
```

### <span class="badge object-method"></span> GroupBy

Filters that are used to query the metrics summary

```go
func (builder *DataqueryBuilder) GroupBy(groupBy []cog.Builder[tempo.TraceqlFilter]) *DataqueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> Limit

Defines the maximum number of traces that are returned from Tempo

```go
func (builder *DataqueryBuilder) Limit(limit int64) *DataqueryBuilder
```

### <span class="badge object-method"></span> MaxDuration

@deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms

```go
func (builder *DataqueryBuilder) MaxDuration(maxDuration string) *DataqueryBuilder
```

### <span class="badge object-method"></span> MetricsQueryType

For metric queries, whether to run instant or range queries

```go
func (builder *DataqueryBuilder) MetricsQueryType(metricsQueryType tempo.MetricsQueryType) *DataqueryBuilder
```

### <span class="badge object-method"></span> MinDuration

@deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms

```go
func (builder *DataqueryBuilder) MinDuration(minDuration string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Query

TraceQL query or trace ID

```go
func (builder *DataqueryBuilder) Query(query string) *DataqueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Search

@deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true

```go
func (builder *DataqueryBuilder) Search(search string) *DataqueryBuilder
```

### <span class="badge object-method"></span> ServiceMapIncludeNamespace

Use service.namespace in addition to service.name to uniquely identify a service.

```go
func (builder *DataqueryBuilder) ServiceMapIncludeNamespace(serviceMapIncludeNamespace bool) *DataqueryBuilder
```

### <span class="badge object-method"></span> ServiceMapQuery

Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.

```go
func (builder *DataqueryBuilder) ServiceMapQuery(serviceMapQuery tempo.StringOrArrayOfString) *DataqueryBuilder
```

### <span class="badge object-method"></span> ServiceName

@deprecated Query traces by service name

```go
func (builder *DataqueryBuilder) ServiceName(serviceName string) *DataqueryBuilder
```

### <span class="badge object-method"></span> SpanName

@deprecated Query traces by span name

```go
func (builder *DataqueryBuilder) SpanName(spanName string) *DataqueryBuilder
```

### <span class="badge object-method"></span> Spss

Defines the maximum number of spans per spanset that are returned from Tempo

```go
func (builder *DataqueryBuilder) Spss(spss int64) *DataqueryBuilder
```

### <span class="badge object-method"></span> Step

For metric queries, the step size to use

```go
func (builder *DataqueryBuilder) Step(step string) *DataqueryBuilder
```

### <span class="badge object-method"></span> TableType

The type of the table that is used to display the search results

```go
func (builder *DataqueryBuilder) TableType(tableType tempo.SearchTableType) *DataqueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
