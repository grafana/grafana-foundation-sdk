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

### <span class="badge object-method"></span> direction

```java
public QueryBuilder direction(LokiQueryDirection direction)
```

### <span class="badge object-method"></span> editorMode

```java
public QueryBuilder editorMode(QueryEditorMode editorMode)
```

### <span class="badge object-method"></span> expr

The LogQL query.

```java
public QueryBuilder expr(String expr)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public QueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> instant

@deprecated, now use queryType.

```java
public QueryBuilder instant(Boolean instant)
```

### <span class="badge object-method"></span> legendFormat

Used to override the name of the series.

```java
public QueryBuilder legendFormat(String legendFormat)
```

### <span class="badge object-method"></span> maxLines

Used to limit the number of log rows returned.

```java
public QueryBuilder maxLines(Long maxLines)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public QueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> range

@deprecated, now use queryType.

```java
public QueryBuilder range(Boolean range)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public QueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> resolution

@deprecated, now use step.

```java
public QueryBuilder resolution(Long resolution)
```

### <span class="badge object-method"></span> step

Used to set step value for range queries.

```java
public QueryBuilder step(String step)
```

### <span class="badge object-method"></span> version

```java
public QueryBuilder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
