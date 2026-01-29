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

### <span class="badge object-method"></span> AliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```go
func (builder *QueryBuilder) AliasBy(aliasBy string) *QueryBuilder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder
```

### <span class="badge object-method"></span> IntervalMs

Time interval in milliseconds.

```go
func (builder *QueryBuilder) IntervalMs(intervalMs float64) *QueryBuilder
```

### <span class="badge object-method"></span> OldDatasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *QueryBuilder) OldDatasource(datasource common.DataSourceRef) *QueryBuilder
```

### <span class="badge object-method"></span> PromQLQuery

PromQL sub-query properties.

```go
func (builder *QueryBuilder) PromQLQuery(promQLQuery cog.Builder[googlecloudmonitoring.PromQLQuery]) *QueryBuilder
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

### <span class="badge object-method"></span> SloQuery

SLO sub-query properties.

```go
func (builder *QueryBuilder) SloQuery(sloQuery cog.Builder[googlecloudmonitoring.SLOQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> TimeSeriesList

GCM query type.

queryType: #QueryType

Time Series List sub-query properties.

```go
func (builder *QueryBuilder) TimeSeriesList(timeSeriesList cog.Builder[googlecloudmonitoring.TimeSeriesList]) *QueryBuilder
```

### <span class="badge object-method"></span> TimeSeriesQuery

Time Series sub-query properties.

```go
func (builder *QueryBuilder) TimeSeriesQuery(timeSeriesQuery cog.Builder[googlecloudmonitoring.TimeSeriesQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
