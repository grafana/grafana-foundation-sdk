---
title: <span class="badge builder"></span> CloudMonitoringQueryBuilder
---
# <span class="badge builder"></span> CloudMonitoringQueryBuilder

## Constructor

```typescript
new CloudMonitoringQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> aliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```typescript
aliasBy(aliasBy: string)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```typescript
datasource(datasource: dashboard.DataSourceRef)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> intervalMs

Time interval in milliseconds.

```typescript
intervalMs(intervalMs: number)
```

### <span class="badge object-method"></span> promQLQuery

PromQL sub-query properties.

```typescript
promQLQuery(promQLQuery: cog.Builder<googlecloudmonitoring.PromQLQuery>)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```typescript
refId(refId: string)
```

### <span class="badge object-method"></span> sloQuery

SLO sub-query properties.

```typescript
sloQuery(sloQuery: cog.Builder<googlecloudmonitoring.SLOQuery>)
```

### <span class="badge object-method"></span> timeSeriesList

GCM query type.

queryType: #QueryType

Time Series List sub-query properties.

```typescript
timeSeriesList(timeSeriesList: cog.Builder<googlecloudmonitoring.TimeSeriesList>)
```

### <span class="badge object-method"></span> timeSeriesQuery

Time Series sub-query properties.

```typescript
timeSeriesQuery(timeSeriesQuery: cog.Builder<googlecloudmonitoring.TimeSeriesQuery>)
```

## See also

 * <span class="badge object-type-interface"></span> [CloudMonitoringQuery](./object-CloudMonitoringQuery.md)
