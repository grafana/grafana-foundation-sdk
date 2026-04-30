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

### <span class="badge object-method"></span> column

```java
public QueryV2Builder column(String column)
```

### <span class="badge object-method"></span> connectionArgs

```java
public QueryV2Builder connectionArgs(com.grafana.foundation.cog.Builder<ConnectionArgs> connectionArgs)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> format

```java
public QueryV2Builder format(FormatOptions format)
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

### <span class="badge object-method"></span> queryID

```java
public QueryV2Builder queryID(String queryID)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public QueryV2Builder queryType(String queryType)
```

### <span class="badge object-method"></span> rawSQL

```java
public QueryV2Builder rawSQL(String rawSQL)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public QueryV2Builder refId(String refId)
```

### <span class="badge object-method"></span> table

```java
public QueryV2Builder table(String table)
```

### <span class="badge object-method"></span> version

```java
public QueryV2Builder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
