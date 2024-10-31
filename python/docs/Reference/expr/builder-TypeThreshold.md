---
title: <span class="badge builder"></span> TypeThreshold
---
# <span class="badge builder"></span> TypeThreshold

## Constructor

```python
TypeThreshold()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> expr.TypeThreshold
```

### <span class="badge object-method"></span> conditions

Threshold Conditions

```python
def conditions(conditions: list[cogbuilder.Builder[expr.ExprTypeThresholdConditions]]) -> typing.Self
```

### <span class="badge object-method"></span> datasource

The datasource

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> expression

Reference to single query result

```python
def expression(expression: str) -> typing.Self
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> interval_ms

Interval is the suggested duration between time points in a time series query.

NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated

from the interval required to fill a pixels in the visualization

```python
def interval_ms(interval_ms: float) -> typing.Self
```

### <span class="badge object-method"></span> max_data_points

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```python
def max_data_points(max_data_points: int) -> typing.Self
```

### <span class="badge object-method"></span> query_type

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

RefID is the unique identifier of the query, set by the frontend call.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> result_assertions

Optionally define expected query result behavior

```python
def result_assertions(result_assertions: cogbuilder.Builder[expr.ExprTypeThresholdResultAssertions]) -> typing.Self
```

### <span class="badge object-method"></span> time_range

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```python
def time_range(time_range: cogbuilder.Builder[expr.ExprTypeThresholdTimeRange]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TypeThreshold](./object-TypeThreshold.md)
