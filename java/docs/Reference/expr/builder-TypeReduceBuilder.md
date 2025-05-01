---
title: <span class="badge builder"></span> TypeReduceBuilder
---
# <span class="badge builder"></span> TypeReduceBuilder

## Constructor

```java
new TypeReduceBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TypeReduce build()
```

### <span class="badge object-method"></span> datasource

The datasource

```java
public TypeReduceBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> expression

Reference to single query result

```java
public TypeReduceBuilder expression(String expression)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```java
public TypeReduceBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```java
public TypeReduceBuilder intervalMs(Double intervalMs)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```java
public TypeReduceBuilder maxDataPoints(Long maxDataPoints)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```java
public TypeReduceBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> reducer

The reducer

Possible enum values:

 - `"sum"` 

 - `"mean"` 

 - `"min"` 

 - `"max"` 

 - `"count"` 

 - `"last"` 

 - `"median"` 

```java
public TypeReduceBuilder reducer(TypeReduceReducer reducer)
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```java
public TypeReduceBuilder refId(String refId)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

```java
public TypeReduceBuilder resultAssertions(ExprTypeReduceResultAssertions resultAssertions)
```

### <span class="badge object-method"></span> settings

Reducer Options

```java
public TypeReduceBuilder settings(ExprTypeReduceSettings settings)
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```java
public TypeReduceBuilder timeRange(ExprTypeReduceTimeRange timeRange)
```

## See also

 * <span class="badge object-type-class"></span> [TypeReduce](./object-TypeReduce.md)
