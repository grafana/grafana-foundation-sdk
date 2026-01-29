---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```go
func NewQueryBuilder() *QueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error)
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder
```

### <span class="badge object-method"></span> Exemplars

For metric queries, how many exemplars to request, 0 means no exemplars

```go
func (builder *QueryBuilder) Exemplars(exemplars int64) *QueryBuilder
```

### <span class="badge object-method"></span> Filters

```go
func (builder *QueryBuilder) Filters(filters []cog.Builder[tempo.TraceqlFilter]) *QueryBuilder
```

### <span class="badge object-method"></span> GroupBy

deprecated Filters that are used to query the metrics summary

```go
func (builder *QueryBuilder) GroupBy(groupBy []cog.Builder[tempo.TraceqlFilter]) *QueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder
```

### <span class="badge object-method"></span> Limit

Defines the maximum number of traces that are returned from Tempo

```go
func (builder *QueryBuilder) Limit(limit int64) *QueryBuilder
```

### <span class="badge object-method"></span> MaxDuration

@deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms

```go
func (builder *QueryBuilder) MaxDuration(maxDuration string) *QueryBuilder
```

### <span class="badge object-method"></span> MetricsQueryType

For metric queries, whether to run instant or range queries

```go
func (builder *QueryBuilder) MetricsQueryType(metricsQueryType tempo.MetricsQueryType) *QueryBuilder
```

### <span class="badge object-method"></span> MinDuration

@deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms

```go
func (builder *QueryBuilder) MinDuration(minDuration string) *QueryBuilder
```

### <span class="badge object-method"></span> OldDatasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *QueryBuilder) OldDatasource(datasource common.DataSourceRef) *QueryBuilder
```

### <span class="badge object-method"></span> Query

TraceQL query or trace ID

```go
func (builder *QueryBuilder) Query(query string) *QueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder
```

### <span class="badge object-method"></span> Search

@deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true

```go
func (builder *QueryBuilder) Search(search string) *QueryBuilder
```

### <span class="badge object-method"></span> ServiceMapIncludeNamespace

Use service.namespace in addition to service.name to uniquely identify a service.

```go
func (builder *QueryBuilder) ServiceMapIncludeNamespace(serviceMapIncludeNamespace bool) *QueryBuilder
```

### <span class="badge object-method"></span> ServiceMapQuery

Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.

```go
func (builder *QueryBuilder) ServiceMapQuery(serviceMapQuery tempo.StringOrArrayOfString) *QueryBuilder
```

### <span class="badge object-method"></span> ServiceName

@deprecated Query traces by service name

```go
func (builder *QueryBuilder) ServiceName(serviceName string) *QueryBuilder
```

### <span class="badge object-method"></span> SpanName

@deprecated Query traces by span name

```go
func (builder *QueryBuilder) SpanName(spanName string) *QueryBuilder
```

### <span class="badge object-method"></span> Spss

Defines the maximum number of spans per spanset that are returned from Tempo

```go
func (builder *QueryBuilder) Spss(spss int64) *QueryBuilder
```

### <span class="badge object-method"></span> Step

For metric queries, the step size to use

```go
func (builder *QueryBuilder) Step(step string) *QueryBuilder
```

### <span class="badge object-method"></span> TableType

The type of the table that is used to display the search results

```go
func (builder *QueryBuilder) TableType(tableType tempo.SearchTableType) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
