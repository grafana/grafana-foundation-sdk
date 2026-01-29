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

### <span class="badge object-method"></span> cloudWatchAnnotationQuery

```java
public QueryBuilder cloudWatchAnnotationQuery(com.grafana.foundation.cog.Builder<CloudWatchAnnotationQuery> cloudWatchAnnotationQuery)
```

### <span class="badge object-method"></span> cloudWatchLogsQuery

```java
public QueryBuilder cloudWatchLogsQuery(com.grafana.foundation.cog.Builder<CloudWatchLogsQuery> cloudWatchLogsQuery)
```

### <span class="badge object-method"></span> cloudWatchMetricsQuery

```java
public QueryBuilder cloudWatchMetricsQuery(com.grafana.foundation.cog.Builder<CloudWatchMetricsQuery> cloudWatchMetricsQuery)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> version

```java
public QueryBuilder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
