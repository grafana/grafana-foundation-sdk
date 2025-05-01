---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```java
new DataqueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Dataquery build()
```

### <span class="badge object-method"></span> convertToUTC

```java
public DataqueryBuilder convertToUTC(Boolean convertToUTC)
```

### <span class="badge object-method"></span> dataset

```java
public DataqueryBuilder dataset(String dataset)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public DataqueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> editorMode

```java
public DataqueryBuilder editorMode(EditorMode editorMode)
```

### <span class="badge object-method"></span> format

```java
public DataqueryBuilder format(QueryFormat format)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public DataqueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> location

```java
public DataqueryBuilder location(String location)
```

### <span class="badge object-method"></span> partitioned

```java
public DataqueryBuilder partitioned(Boolean partitioned)
```

### <span class="badge object-method"></span> partitionedField

```java
public DataqueryBuilder partitionedField(String partitionedField)
```

### <span class="badge object-method"></span> project

```java
public DataqueryBuilder project(String project)
```

### <span class="badge object-method"></span> queryPriority

```java
public DataqueryBuilder queryPriority(QueryPriority queryPriority)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public DataqueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> rawQuery

```java
public DataqueryBuilder rawQuery(Boolean rawQuery)
```

### <span class="badge object-method"></span> rawSql

```java
public DataqueryBuilder rawSql(String rawSql)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public DataqueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> sharded

```java
public DataqueryBuilder sharded(Boolean sharded)
```

### <span class="badge object-method"></span> sql

```java
public DataqueryBuilder sql(SQLExpression sql)
```

### <span class="badge object-method"></span> table

```java
public DataqueryBuilder table(String table)
```

### <span class="badge object-method"></span> timeShift

```java
public DataqueryBuilder timeShift(String timeShift)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
