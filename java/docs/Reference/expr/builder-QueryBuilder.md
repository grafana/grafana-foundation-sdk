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

### <span class="badge object-method"></span> typeClassicConditions

```java
public QueryBuilder typeClassicConditions(com.grafana.foundation.cog.Builder<TypeClassicConditions> typeClassicConditions)
```

### <span class="badge object-method"></span> typeMath

```java
public QueryBuilder typeMath(com.grafana.foundation.cog.Builder<TypeMath> typeMath)
```

### <span class="badge object-method"></span> typeReduce

```java
public QueryBuilder typeReduce(com.grafana.foundation.cog.Builder<TypeReduce> typeReduce)
```

### <span class="badge object-method"></span> typeResample

```java
public QueryBuilder typeResample(com.grafana.foundation.cog.Builder<TypeResample> typeResample)
```

### <span class="badge object-method"></span> typeSql

```java
public QueryBuilder typeSql(com.grafana.foundation.cog.Builder<TypeSql> typeSql)
```

### <span class="badge object-method"></span> typeThreshold

```java
public QueryBuilder typeThreshold(com.grafana.foundation.cog.Builder<TypeThreshold> typeThreshold)
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
