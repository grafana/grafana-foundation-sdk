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

### <span class="badge object-method"></span> classicConditions

```java
public QueryV2Builder classicConditions(com.grafana.foundation.cog.Builder<TypeClassicConditions> typeClassicConditions)
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

### <span class="badge object-method"></span> math

```java
public QueryV2Builder math(com.grafana.foundation.cog.Builder<TypeMath> typeMath)
```

### <span class="badge object-method"></span> reduce

```java
public QueryV2Builder reduce(com.grafana.foundation.cog.Builder<TypeReduce> typeReduce)
```

### <span class="badge object-method"></span> resample

```java
public QueryV2Builder resample(com.grafana.foundation.cog.Builder<TypeResample> typeResample)
```

### <span class="badge object-method"></span> sql

```java
public QueryV2Builder sql(com.grafana.foundation.cog.Builder<TypeSql> typeSql)
```

### <span class="badge object-method"></span> threshold

```java
public QueryV2Builder threshold(com.grafana.foundation.cog.Builder<TypeThreshold> typeThreshold)
```

### <span class="badge object-method"></span> version

```java
public QueryV2Builder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
