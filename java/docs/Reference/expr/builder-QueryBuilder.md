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

### <span class="badge object-method"></span> classicConditions

```java
public QueryBuilder classicConditions(com.grafana.foundation.cog.Builder<TypeClassicConditions> typeClassicConditions)
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

### <span class="badge object-method"></span> math

```java
public QueryBuilder math(com.grafana.foundation.cog.Builder<TypeMath> typeMath)
```

### <span class="badge object-method"></span> reduce

```java
public QueryBuilder reduce(com.grafana.foundation.cog.Builder<TypeReduce> typeReduce)
```

### <span class="badge object-method"></span> resample

```java
public QueryBuilder resample(com.grafana.foundation.cog.Builder<TypeResample> typeResample)
```

### <span class="badge object-method"></span> sql

```java
public QueryBuilder sql(com.grafana.foundation.cog.Builder<TypeSql> typeSql)
```

### <span class="badge object-method"></span> threshold

```java
public QueryBuilder threshold(com.grafana.foundation.cog.Builder<TypeThreshold> typeThreshold)
```

### <span class="badge object-method"></span> version

```java
public QueryBuilder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
