---
title: <span class="badge builder"></span> TypeClassicConditionsBuilder
---
# <span class="badge builder"></span> TypeClassicConditionsBuilder

## Constructor

```java
new TypeClassicConditionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TypeClassicConditions build()
```

### <span class="badge object-method"></span> conditions

```java
public TypeClassicConditionsBuilder conditions(List<com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditions>> conditions)
```

### <span class="badge object-method"></span> datasource

The datasource

```java
public TypeClassicConditionsBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```java
public TypeClassicConditionsBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> intervalMs

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```java
public TypeClassicConditionsBuilder intervalMs(Double intervalMs)
```

### <span class="badge object-method"></span> maxDataPoints

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```java
public TypeClassicConditionsBuilder maxDataPoints(Long maxDataPoints)
```

### <span class="badge object-method"></span> queryType

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```java
public TypeClassicConditionsBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

RefID is the unique identifier of the query, set by the frontend call.

```java
public TypeClassicConditionsBuilder refId(String refId)
```

### <span class="badge object-method"></span> resultAssertions

Optionally define expected query result behavior

```java
public TypeClassicConditionsBuilder resultAssertions(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsResultAssertions> resultAssertions)
```

### <span class="badge object-method"></span> timeRange

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```java
public TypeClassicConditionsBuilder timeRange(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsTimeRange> timeRange)
```

## See also

 * <span class="badge object-type-class"></span> [TypeClassicConditions](./object-TypeClassicConditions.md)
