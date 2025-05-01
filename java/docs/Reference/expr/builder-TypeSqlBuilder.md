---
title: <span class="badge builder"></span> TypeSqlBuilder
---
# <span class="badge builder"></span> TypeSqlBuilder

## Constructor

```java
new TypeSqlBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TypeSql build()
```

### <span class="badge object-method"></span> datasource

The datasource

```java
public TypeSqlBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> expression

```java
public TypeSqlBuilder expression(String expression)
```

### <span class="badge object-method"></span> format

```java
public TypeSqlBuilder format(String format)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```java
public TypeSqlBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```java
public TypeSqlBuilder intervalMs(Double intervalMs)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```java
public TypeSqlBuilder maxDataPoints(Long maxDataPoints)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```java
public TypeSqlBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```java
public TypeSqlBuilder refId(String refId)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

```java
public TypeSqlBuilder resultAssertions(ExprTypeSqlResultAssertions resultAssertions)
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```java
public TypeSqlBuilder timeRange(ExprTypeSqlTimeRange timeRange)
```

## See also

 * <span class="badge object-type-class"></span> [TypeSql](./object-TypeSql.md)
