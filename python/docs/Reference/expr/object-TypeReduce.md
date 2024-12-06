---
title: <span class="badge object-type-class"></span> TypeReduce
---
# <span class="badge object-type-class"></span> TypeReduce

## Definition

```python
class TypeReduce(cogvariants.Dataquery):
    # The datasource
    datasource: typing.Optional[dashboard.DataSourceRef]
    # Reference to single query result
    expression: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # The reducer
    # Possible enum values:
    #  - `"sum"` 
    #  - `"mean"` 
    #  - `"min"` 
    #  - `"max"` 
    #  - `"count"` 
    #  - `"last"` 
    #  - `"median"` 
    reducer: typing.Literal["sum", "mean", "min", "max", "count", "last", "median"]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: str
    # Optionally define expected query result behavior
    result_assertions: typing.Optional[expr.ExprTypeReduceResultAssertions]
    # Reducer Options
    settings: typing.Optional[expr.ExprTypeReduceSettings]
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional[expr.ExprTypeReduceTimeRange]
    type_val: typing.Literal["reduce"]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [TypeReduce](./builder-TypeReduce.md)
