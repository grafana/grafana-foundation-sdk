---
title: <span class="badge builder"></span> CloudMonitoringQueryBuilder
---
# <span class="badge builder"></span> CloudMonitoringQueryBuilder

## Constructor

```java
new CloudMonitoringQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public CloudMonitoringQuery build()
```

### <span class="badge object-method"></span> aliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```java
public CloudMonitoringQueryBuilder aliasBy(String aliasBy)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public CloudMonitoringQueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public CloudMonitoringQueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> intervalMs

Time interval in milliseconds.

```java
public CloudMonitoringQueryBuilder intervalMs(Double intervalMs)
```

### <span class="badge object-method"></span> promQLQuery

PromQL sub-query properties.

```java
public CloudMonitoringQueryBuilder promQLQuery(com.grafana.foundation.cog.Builder<PromQLQuery> promQLQuery)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public CloudMonitoringQueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public CloudMonitoringQueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> sloQuery

SLO sub-query properties.

```java
public CloudMonitoringQueryBuilder sloQuery(com.grafana.foundation.cog.Builder<SLOQuery> sloQuery)
```

### <span class="badge object-method"></span> timeSeriesList

GCM query type.

queryType: #QueryType

Time Series List sub-query properties.

```java
public CloudMonitoringQueryBuilder timeSeriesList(com.grafana.foundation.cog.Builder<TimeSeriesList> timeSeriesList)
```

### <span class="badge object-method"></span> timeSeriesQuery

Time Series sub-query properties.

```java
public CloudMonitoringQueryBuilder timeSeriesQuery(com.grafana.foundation.cog.Builder<TimeSeriesQuery> timeSeriesQuery)
```

## See also

 * <span class="badge object-type-class"></span> [CloudMonitoringQuery](./object-CloudMonitoringQuery.md)
