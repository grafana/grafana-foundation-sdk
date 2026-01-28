---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```java
new QueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Query build()
```

### <span class="badge object-method"></span> aliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```java
public QueryBuilder aliasBy(String aliasBy)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public QueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> intervalMs

Time interval in milliseconds.

```java
public QueryBuilder intervalMs(Double intervalMs)
```

### <span class="badge object-method"></span> oldDatasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public QueryBuilder oldDatasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> promQLQuery

PromQL sub-query properties.

```java
public QueryBuilder promQLQuery(com.grafana.foundation.cog.Builder<PromQLQuery> promQLQuery)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public QueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public QueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> sloQuery

SLO sub-query properties.

```java
public QueryBuilder sloQuery(com.grafana.foundation.cog.Builder<SLOQuery> sloQuery)
```

### <span class="badge object-method"></span> timeSeriesList

GCM query type.

queryType: #QueryType

Time Series List sub-query properties.

```java
public QueryBuilder timeSeriesList(com.grafana.foundation.cog.Builder<TimeSeriesList> timeSeriesList)
```

### <span class="badge object-method"></span> timeSeriesQuery

Time Series sub-query properties.

```java
public QueryBuilder timeSeriesQuery(com.grafana.foundation.cog.Builder<TimeSeriesQuery> timeSeriesQuery)
```

### <span class="badge object-method"></span> version

```java
public QueryBuilder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
