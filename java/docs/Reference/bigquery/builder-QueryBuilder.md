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

### <span class="badge object-method"></span> convertToUTC

```java
public QueryBuilder convertToUTC(Boolean convertToUTC)
```

### <span class="badge object-method"></span> dataset

```java
public QueryBuilder dataset(String dataset)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> editorMode

```java
public QueryBuilder editorMode(EditorMode editorMode)
```

### <span class="badge object-method"></span> format

```java
public QueryBuilder format(QueryFormat format)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public QueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> location

```java
public QueryBuilder location(String location)
```

### <span class="badge object-method"></span> partitioned

```java
public QueryBuilder partitioned(Boolean partitioned)
```

### <span class="badge object-method"></span> partitionedField

```java
public QueryBuilder partitionedField(String partitionedField)
```

### <span class="badge object-method"></span> project

```java
public QueryBuilder project(String project)
```

### <span class="badge object-method"></span> queryPriority

```java
public QueryBuilder queryPriority(QueryPriority queryPriority)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public QueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> rawQuery

```java
public QueryBuilder rawQuery(Boolean rawQuery)
```

### <span class="badge object-method"></span> rawSql

```java
public QueryBuilder rawSql(String rawSql)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public QueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> sharded

```java
public QueryBuilder sharded(Boolean sharded)
```

### <span class="badge object-method"></span> sql

```java
public QueryBuilder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql)
```

### <span class="badge object-method"></span> table

```java
public QueryBuilder table(String table)
```

### <span class="badge object-method"></span> timeShift

```java
public QueryBuilder timeShift(String timeShift)
```

### <span class="badge object-method"></span> version

```java
public QueryBuilder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
