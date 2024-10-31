---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```python
class Dataquery(cogvariants.Dataquery):
    alias: typing.Optional[str]
    # Used for live query
    channel: typing.Optional[str]
    csv_content: typing.Optional[str]
    csv_file_name: typing.Optional[str]
    csv_wave: typing.Optional[list[testdata.CSVWave]]
    # The datasource
    datasource: typing.Optional[dashboard.DataSourceRef]
    # Drop percentage (the chance we will lose a point 0-100)
    drop_percent: typing.Optional[float]
    # Possible enum values:
    #  - `"frontend_exception"` 
    #  - `"frontend_observable"` 
    #  - `"server_panic"` 
    error_type: typing.Optional[typing.Literal["frontend_exception", "frontend_observable", "server_panic"]]
    flamegraph_diff: typing.Optional[bool]
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    labels: typing.Optional[str]
    level_column: typing.Optional[bool]
    lines: typing.Optional[int]
    max_val: typing.Optional[float]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    min_val: typing.Optional[float]
    nodes: typing.Optional[testdata.NodesQuery]
    noise: typing.Optional[float]
    points: typing.Optional[list[list[object]]]
    pulse_wave: typing.Optional[testdata.PulseWaveQuery]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    raw_frame_content: typing.Optional[str]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: typing.Optional[str]
    # Optionally define expected query result behavior
    result_assertions: typing.Optional[testdata.ResultAssertions]
    # Possible enum values:
    #  - `"annotations"` 
    #  - `"arrow"` 
    #  - `"csv_content"` 
    #  - `"csv_file"` 
    #  - `"csv_metric_values"` 
    #  - `"datapoints_outside_range"` 
    #  - `"exponential_heatmap_bucket_data"` 
    #  - `"flame_graph"` 
    #  - `"grafana_api"` 
    #  - `"linear_heatmap_bucket_data"` 
    #  - `"live"` 
    #  - `"logs"` 
    #  - `"manual_entry"` 
    #  - `"no_data_points"` 
    #  - `"node_graph"` 
    #  - `"predictable_csv_wave"` 
    #  - `"predictable_pulse"` 
    #  - `"random_walk"` 
    #  - `"random_walk_table"` 
    #  - `"random_walk_with_error"` 
    #  - `"raw_frame"` 
    #  - `"server_error_500"` 
    #  - `"simulation"` 
    #  - `"slow_query"` 
    #  - `"streaming_client"` 
    #  - `"table_static"` 
    #  - `"trace"` 
    #  - `"usa"` 
    #  - `"variables-query"` 
    scenario_id: typing.Optional[typing.Literal["annotations", "arrow", "csv_content", "csv_file", "csv_metric_values", "datapoints_outside_range", "exponential_heatmap_bucket_data", "flame_graph", "grafana_api", "linear_heatmap_bucket_data", "live", "logs", "manual_entry", "no_data_points", "node_graph", "predictable_csv_wave", "predictable_pulse", "random_walk", "random_walk_table", "random_walk_with_error", "raw_frame", "server_error_500", "simulation", "slow_query", "streaming_client", "table_static", "trace", "usa", "variables-query"]]
    series_count: typing.Optional[int]
    sim: typing.Optional[testdata.SimulationQuery]
    span_count: typing.Optional[int]
    spread: typing.Optional[float]
    start_value: typing.Optional[float]
    stream: typing.Optional[testdata.StreamingQuery]
    # common parameter used by many query types
    string_input: typing.Optional[str]
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional[testdata.TimeRange]
    usa: typing.Optional[testdata.USAQuery]
    with_nil: typing.Optional[bool]
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
