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

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> groupBy

Allows to group the results.

```java
public QueryBuilder groupBy(List<String> groupBy)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public QueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> labelSelector

Specifies the query label selectors.

```java
public QueryBuilder labelSelector(String labelSelector)
```

### <span class="badge object-method"></span> limit

Sets the maximum number of time series.

```java
public QueryBuilder limit(Long limit)
```

### <span class="badge object-method"></span> maxNodes

Sets the maximum number of nodes in the flamegraph.

```java
public QueryBuilder maxNodes(Long maxNodes)
```

### <span class="badge object-method"></span> profileTypeId

Specifies the type of profile to query.

```java
public QueryBuilder profileTypeId(String profileTypeId)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public QueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public QueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> spanSelector

Specifies the query span selectors.

```java
public QueryBuilder spanSelector(List<String> spanSelector)
```

### <span class="badge object-method"></span> version

```java
public QueryBuilder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
