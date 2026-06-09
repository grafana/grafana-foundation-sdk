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

### <span class="badge object-method"></span> annotationQuery

```java
public QueryBuilder annotationQuery(com.grafana.foundation.cog.Builder<AnnotationQuery> annotationQuery)
```

### <span class="badge object-method"></span> logsQuery

```java
public QueryBuilder logsQuery(com.grafana.foundation.cog.Builder<LogsQuery> logsQuery)
```

### <span class="badge object-method"></span> metricsQuery

```java
public QueryBuilder metricsQuery(com.grafana.foundation.cog.Builder<MetricsQuery> metricsQuery)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> labels

```java
public QueryBuilder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> version

```java
public QueryBuilder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
