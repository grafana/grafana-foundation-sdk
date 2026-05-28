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

### <span class="badge object-method"></span> convertToUTC

```java
public QueryV2Builder convertToUTC(Boolean convertToUTC)
```

### <span class="badge object-method"></span> dataset

```java
public QueryV2Builder dataset(String dataset)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> editorMode

```java
public QueryV2Builder editorMode(EditorMode editorMode)
```

### <span class="badge object-method"></span> format

```java
public QueryV2Builder format(QueryFormat format)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public QueryV2Builder hide(Boolean hide)
```

### <span class="badge object-method"></span> labels

```java
public QueryV2Builder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> location

```java
public QueryV2Builder location(String location)
```

### <span class="badge object-method"></span> partitioned

```java
public QueryV2Builder partitioned(Boolean partitioned)
```

### <span class="badge object-method"></span> partitionedField

```java
public QueryV2Builder partitionedField(String partitionedField)
```

### <span class="badge object-method"></span> project

```java
public QueryV2Builder project(String project)
```

### <span class="badge object-method"></span> queryPriority

```java
public QueryV2Builder queryPriority(QueryPriority queryPriority)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public QueryV2Builder queryType(String queryType)
```

### <span class="badge object-method"></span> rawQuery

```java
public QueryV2Builder rawQuery(Boolean rawQuery)
```

### <span class="badge object-method"></span> rawSql

```java
public QueryV2Builder rawSql(String rawSql)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public QueryV2Builder refId(String refId)
```

### <span class="badge object-method"></span> sharded

```java
public QueryV2Builder sharded(Boolean sharded)
```

### <span class="badge object-method"></span> sql

```java
public QueryV2Builder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql)
```

### <span class="badge object-method"></span> table

```java
public QueryV2Builder table(String table)
```

### <span class="badge object-method"></span> timeShift

```java
public QueryV2Builder timeShift(String timeShift)
```

### <span class="badge object-method"></span> version

```java
public QueryV2Builder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
