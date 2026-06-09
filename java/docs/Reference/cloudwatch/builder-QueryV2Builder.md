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

### <span class="badge object-method"></span> annotationQuery

```java
public QueryV2Builder annotationQuery(com.grafana.foundation.cog.Builder<AnnotationQuery> annotationQuery)
```

### <span class="badge object-method"></span> logsQuery

```java
public QueryV2Builder logsQuery(com.grafana.foundation.cog.Builder<LogsQuery> logsQuery)
```

### <span class="badge object-method"></span> metricsQuery

```java
public QueryV2Builder metricsQuery(com.grafana.foundation.cog.Builder<MetricsQuery> metricsQuery)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> labels

```java
public QueryV2Builder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> version

```java
public QueryV2Builder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
