---
title: <span class="badge builder"></span> CloudMonitoringQueryBuilder
---
# <span class="badge builder"></span> CloudMonitoringQueryBuilder

## Constructor

```go
func NewCloudMonitoringQueryBuilder() *CloudMonitoringQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CloudMonitoringQueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> AliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```go
func (builder *CloudMonitoringQueryBuilder) AliasBy(aliasBy string) *CloudMonitoringQueryBuilder
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *CloudMonitoringQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *CloudMonitoringQueryBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *CloudMonitoringQueryBuilder) Hide(hide bool) *CloudMonitoringQueryBuilder
```

### <span class="badge object-method"></span> IntervalMs

Time interval in milliseconds.

```go
func (builder *CloudMonitoringQueryBuilder) IntervalMs(intervalMs float64) *CloudMonitoringQueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *CloudMonitoringQueryBuilder) QueryType(queryType string) *CloudMonitoringQueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *CloudMonitoringQueryBuilder) RefId(refId string) *CloudMonitoringQueryBuilder
```

### <span class="badge object-method"></span> SloQuery

SLO sub-query properties.

```go
func (builder *CloudMonitoringQueryBuilder) SloQuery(sloQuery cog.Builder[googlecloudmonitoring.SLOQuery]) *CloudMonitoringQueryBuilder
```

### <span class="badge object-method"></span> TimeSeriesList

GCM query type.

queryType: #QueryType

Time Series List sub-query properties.

```go
func (builder *CloudMonitoringQueryBuilder) TimeSeriesList(timeSeriesList cog.Builder[googlecloudmonitoring.TimeSeriesList]) *CloudMonitoringQueryBuilder
```

### <span class="badge object-method"></span> TimeSeriesQuery

Time Series sub-query properties.

```go
func (builder *CloudMonitoringQueryBuilder) TimeSeriesQuery(timeSeriesQuery cog.Builder[googlecloudmonitoring.TimeSeriesQuery]) *CloudMonitoringQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CloudMonitoringQuery](./object-CloudMonitoringQuery.md)