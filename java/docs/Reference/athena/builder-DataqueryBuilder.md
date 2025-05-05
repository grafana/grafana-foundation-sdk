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

### <span class="badge object-method"></span> column

```java
public DataqueryBuilder column(String column)
```

### <span class="badge object-method"></span> connectionArgs

```java
public DataqueryBuilder connectionArgs(com.grafana.foundation.cog.Builder<ConnectionArgs> connectionArgs)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public DataqueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> format

```java
public DataqueryBuilder format(FormatOptions format)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public DataqueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> queryID

```java
public DataqueryBuilder queryID(String queryID)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public DataqueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> rawSQL

```java
public DataqueryBuilder rawSQL(String rawSQL)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public DataqueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> table

```java
public DataqueryBuilder table(String table)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
