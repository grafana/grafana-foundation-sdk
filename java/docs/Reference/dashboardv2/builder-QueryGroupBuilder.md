---
title: <span class="badge builder"></span> QueryGroupBuilder
---
# <span class="badge builder"></span> QueryGroupBuilder

## Constructor

```java
new QueryGroupBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public QueryGroup build()
```

### <span class="badge object-method"></span> queryOptions

```java
public QueryGroupBuilder queryOptions(com.grafana.foundation.cog.Builder<QueryOptionsSpec> queryOptions)
```

### <span class="badge object-method"></span> target

```java
public QueryGroupBuilder target(com.grafana.foundation.cog.Builder<PanelQueryKind> querie)
```

### <span class="badge object-method"></span> targets

```java
public QueryGroupBuilder targets(List<com.grafana.foundation.cog.Builder<PanelQueryKind>> queries)
```

### <span class="badge object-method"></span> transformation

```java
public QueryGroupBuilder transformation(com.grafana.foundation.cog.Builder<TransformationKind> transformation)
```

### <span class="badge object-method"></span> transformations

```java
public QueryGroupBuilder transformations(List<com.grafana.foundation.cog.Builder<TransformationKind>> transformations)
```

## See also

 * <span class="badge object-type-class"></span> [QueryGroupKind](./object-QueryGroupKind.md)
