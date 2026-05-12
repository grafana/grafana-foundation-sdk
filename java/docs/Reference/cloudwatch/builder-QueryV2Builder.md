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

### <span class="badge object-method"></span> cloudWatchAnnotationQuery

```java
public QueryV2Builder cloudWatchAnnotationQuery(com.grafana.foundation.cog.Builder<CloudWatchAnnotationQuery> cloudWatchAnnotationQuery)
```

### <span class="badge object-method"></span> cloudWatchLogsQuery

```java
public QueryV2Builder cloudWatchLogsQuery(com.grafana.foundation.cog.Builder<CloudWatchLogsQuery> cloudWatchLogsQuery)
```

### <span class="badge object-method"></span> cloudWatchMetricsQuery

```java
public QueryV2Builder cloudWatchMetricsQuery(com.grafana.foundation.cog.Builder<CloudWatchMetricsQuery> cloudWatchMetricsQuery)
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
