---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```go
func NewQueryV2Builder() *QueryV2Builder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error)
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder
```

### <span class="badge object-method"></span> Exemplars

For metric queries, how many exemplars to request, 0 means no exemplars

```go
func (builder *QueryV2Builder) Exemplars(exemplars int64) *QueryV2Builder
```

### <span class="badge object-method"></span> Filters

```go
func (builder *QueryV2Builder) Filters(filters []cog.Builder[tempo.TraceqlFilter]) *QueryV2Builder
```

### <span class="badge object-method"></span> GroupBy

Filters that are used to query the metrics summary

```go
func (builder *QueryV2Builder) GroupBy(groupBy []cog.Builder[tempo.TraceqlFilter]) *QueryV2Builder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder
```

### <span class="badge object-method"></span> Limit

Defines the maximum number of traces that are returned from Tempo

```go
func (builder *QueryV2Builder) Limit(limit int64) *QueryV2Builder
```

### <span class="badge object-method"></span> MaxDuration

@deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms

```go
func (builder *QueryV2Builder) MaxDuration(maxDuration string) *QueryV2Builder
```

### <span class="badge object-method"></span> MetricsQueryType

For metric queries, whether to run instant or range queries

```go
func (builder *QueryV2Builder) MetricsQueryType(metricsQueryType tempo.MetricsQueryType) *QueryV2Builder
```

### <span class="badge object-method"></span> MinDuration

@deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms

```go
func (builder *QueryV2Builder) MinDuration(minDuration string) *QueryV2Builder
```

### <span class="badge object-method"></span> Query

TraceQL query or trace ID

```go
func (builder *QueryV2Builder) Query(query string) *QueryV2Builder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder
```

### <span class="badge object-method"></span> Search

@deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true

```go
func (builder *QueryV2Builder) Search(search string) *QueryV2Builder
```

### <span class="badge object-method"></span> ServiceMapIncludeNamespace

Use service.namespace in addition to service.name to uniquely identify a service.

```go
func (builder *QueryV2Builder) ServiceMapIncludeNamespace(serviceMapIncludeNamespace bool) *QueryV2Builder
```

### <span class="badge object-method"></span> ServiceMapQuery

Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.

```go
func (builder *QueryV2Builder) ServiceMapQuery(serviceMapQuery tempo.StringOrArrayOfString) *QueryV2Builder
```

### <span class="badge object-method"></span> ServiceName

@deprecated Query traces by service name

```go
func (builder *QueryV2Builder) ServiceName(serviceName string) *QueryV2Builder
```

### <span class="badge object-method"></span> SpanName

@deprecated Query traces by span name

```go
func (builder *QueryV2Builder) SpanName(spanName string) *QueryV2Builder
```

### <span class="badge object-method"></span> Spss

Defines the maximum number of spans per spanset that are returned from Tempo

```go
func (builder *QueryV2Builder) Spss(spss int64) *QueryV2Builder
```

### <span class="badge object-method"></span> Step

For metric queries, the step size to use

```go
func (builder *QueryV2Builder) Step(step string) *QueryV2Builder
```

### <span class="badge object-method"></span> TableType

The type of the table that is used to display the search results

```go
func (builder *QueryV2Builder) TableType(tableType tempo.SearchTableType) *QueryV2Builder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryV2Builder) Version(version string) *QueryV2Builder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
