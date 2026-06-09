---
title: <span class="badge builder"></span> LogsQueryBuilder
---
# <span class="badge builder"></span> LogsQueryBuilder

## Constructor

```java
new LogsQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public LogsQuery build()
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public LogsQueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> expression

The CloudWatch Logs Insights query to execute

```java
public LogsQueryBuilder expression(String expression)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public LogsQueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> id

```java
public LogsQueryBuilder id(String id)
```

### <span class="badge object-method"></span> logGroupNames

@deprecated use logGroups

```java
public LogsQueryBuilder logGroupNames(List<String> logGroupNames)
```

### <span class="badge object-method"></span> logGroups

Log groups to query

```java
public LogsQueryBuilder logGroups(List<com.grafana.foundation.cog.Builder<LogGroup>> logGroups)
```

### <span class="badge object-method"></span> queryLanguage

Language used for querying logs, can be CWLI, SQL, or PPL. If empty, the default language is CWLI.

```java
public LogsQueryBuilder queryLanguage(LogsQueryLanguage queryLanguage)
```

### <span class="badge object-method"></span> queryMode

Whether a query is a Metrics, Logs, or Annotations query

```java
public LogsQueryBuilder queryMode(QueryMode queryMode)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public LogsQueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public LogsQueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> region

AWS region to query for the logs

```java
public LogsQueryBuilder region(String region)
```

### <span class="badge object-method"></span> statsGroups

Fields to group the results by, this field is automatically populated whenever the query is updated

```java
public LogsQueryBuilder statsGroups(List<String> statsGroups)
```

## See also

 * <span class="badge object-type-class"></span> [LogsQuery](./object-LogsQuery.md)
