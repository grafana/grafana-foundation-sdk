---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```java
new QueryV2Builder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public QueryV2 build()
```

### <span class="badge object-method"></span> aliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```java
public QueryV2Builder aliasBy(String aliasBy)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public QueryV2Builder hide(Boolean hide)
```

### <span class="badge object-method"></span> intervalMs

Time interval in milliseconds.

```java
public QueryV2Builder intervalMs(Double intervalMs)
```

### <span class="badge object-method"></span> labels

```java
public QueryV2Builder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> promQLQuery

PromQL sub-query properties.

```java
public QueryV2Builder promQLQuery(com.grafana.foundation.cog.Builder<PromQLQuery> promQLQuery)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public QueryV2Builder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public QueryV2Builder refId(String refId)
```

### <span class="badge object-method"></span> sloQuery

SLO sub-query properties.

```java
public QueryV2Builder sloQuery(com.grafana.foundation.cog.Builder<SLOQuery> sloQuery)
```

### <span class="badge object-method"></span> timeSeriesList

GCM query type.

queryType: #QueryType

Time Series List sub-query properties.

```java
public QueryV2Builder timeSeriesList(com.grafana.foundation.cog.Builder<TimeSeriesList> timeSeriesList)
```

### <span class="badge object-method"></span> timeSeriesQuery

Time Series sub-query properties.

```java
public QueryV2Builder timeSeriesQuery(com.grafana.foundation.cog.Builder<TimeSeriesQuery> timeSeriesQuery)
```

### <span class="badge object-method"></span> version

```java
public QueryV2Builder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
