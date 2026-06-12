---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```python
class Dataquery(cogvariants.Dataquery):
    # Additional Ad-hoc filters that take precedence over Scope on conflict.
    adhoc_filters: typing.Optional[list[prometheus.AdhocFilters]]
    # The datasource
    datasource: typing.Optional[common.DataSourceRef]
    # what we should show in the editor
    # Possible enum values:
    #  - `"builder"` 
    #  - `"code"` 
    editor_mode: typing.Optional[prometheus.QueryEditorMode]
    # Execute an additional query to identify interesting raw samples relevant for the given expr
    exemplar: typing.Optional[bool]
    # The actual expression/query that will be evaluated by Prometheus
    expr: str
    # The response format
    # Possible enum values:
    #  - `"time_series"` 
    #  - `"table"` 
    #  - `"heatmap"` 
    format_val: typing.Optional[prometheus.PromQueryFormat]
    # Group By parameters to apply to aggregate expressions in the query
    group_by_keys: typing.Optional[list[str]]
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Returns only the latest value that Prometheus has scraped for the requested time series
    instant: typing.Optional[bool]
    # Used to specify how many times to divide max data points by. We use max data points under query options
    # See https://github.com/grafana/grafana/issues/48081
    # Deprecated: use interval
    interval_factor: typing.Optional[int]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    # Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    legend_format: typing.Optional[str]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    range_val: typing.Optional[bool]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: typing.Optional[str]
    # Optionally define expected query result behavior
    result_assertions: typing.Optional[prometheus.ResultAssertions]
    # A set of filters applied to apply to the query
    scopes: typing.Optional[list[prometheus.Scopes]]
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional[prometheus.TimeRange]
    # An additional lower limit for the step parameter of the Prometheus query and for the
    # `$__interval` and `$__rate_interval` variables.
    interval: typing.Optional[str]
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

 * <span class="badge builder"></span> [Dataquery](./builder-Dataquery.md)
