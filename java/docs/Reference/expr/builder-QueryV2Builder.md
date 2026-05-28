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

### <span class="badge object-method"></span> typeClassicConditions

```java
public QueryV2Builder typeClassicConditions(com.grafana.foundation.cog.Builder<TypeClassicConditions> typeClassicConditions)
```

### <span class="badge object-method"></span> typeMath

```java
public QueryV2Builder typeMath(com.grafana.foundation.cog.Builder<TypeMath> typeMath)
```

### <span class="badge object-method"></span> typeReduce

```java
public QueryV2Builder typeReduce(com.grafana.foundation.cog.Builder<TypeReduce> typeReduce)
```

### <span class="badge object-method"></span> typeResample

```java
public QueryV2Builder typeResample(com.grafana.foundation.cog.Builder<TypeResample> typeResample)
```

### <span class="badge object-method"></span> typeSql

```java
public QueryV2Builder typeSql(com.grafana.foundation.cog.Builder<TypeSql> typeSql)
```

### <span class="badge object-method"></span> typeThreshold

```java
public QueryV2Builder typeThreshold(com.grafana.foundation.cog.Builder<TypeThreshold> typeThreshold)
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
