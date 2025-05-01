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

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```java
public CloudMonitoringQueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> intervalMs

Time interval in milliseconds.

```java
public CloudMonitoringQueryBuilder intervalMs(Double intervalMs)
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
public CloudMonitoringQueryBuilder sloQuery(SLOQuery sloQuery)
```

### <span class="badge object-method"></span> timeSeriesList

GCM query type.

queryType: #QueryType

Time Series List sub-query properties.

```java
public CloudMonitoringQueryBuilder timeSeriesList(TimeSeriesList timeSeriesList)
```

### <span class="badge object-method"></span> timeSeriesQuery

Time Series sub-query properties.

```java
public CloudMonitoringQueryBuilder timeSeriesQuery(TimeSeriesQuery timeSeriesQuery)
```

## See also

 * <span class="badge object-type-class"></span> [CloudMonitoringQuery](./object-CloudMonitoringQuery.md)
