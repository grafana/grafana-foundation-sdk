---
title: <span class="badge builder"></span> DataQueryKindBuilder
---
# <span class="badge builder"></span> DataQueryKindBuilder

## Constructor

```java
new DataQueryKindBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public DataQueryKind build()
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public DataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> group

```java
public DataQueryKindBuilder group(String group)
```

### <span class="badge object-method"></span> spec

```java
public DataQueryKindBuilder spec(Object spec)
```

### <span class="badge object-method"></span> version

```java
public DataQueryKindBuilder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [DataQueryKind](./object-DataQueryKind.md)
