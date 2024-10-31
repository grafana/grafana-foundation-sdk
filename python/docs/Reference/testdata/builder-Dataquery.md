---
title: <span class="badge builder"></span> Dataquery
---
# <span class="badge builder"></span> Dataquery

## Constructor

```python
Dataquery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> testdata.Dataquery
```

### <span class="badge object-method"></span> alias

```python
def alias(alias: str) -> typing.Self
```

### <span class="badge object-method"></span> channel

Used for live query

```python
def channel(channel: str) -> typing.Self
```

### <span class="badge object-method"></span> csv_content

```python
def csv_content(csv_content: str) -> typing.Self
```

### <span class="badge object-method"></span> csv_file_name

```python
def csv_file_name(csv_file_name: str) -> typing.Self
```

### <span class="badge object-method"></span> csv_wave

```python
def csv_wave(csv_wave: list[cogbuilder.Builder[testdata.CSVWave]]) -> typing.Self
```

### <span class="badge object-method"></span> datasource

The datasource

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> drop_percent

Drop percentage (the chance we will lose a point 0-100)

```python
def drop_percent(drop_percent: float) -> typing.Self
```

### <span class="badge object-method"></span> error_type

Possible enum values:

 - `"frontend_exception"` 

 - `"frontend_observable"` 

 - `"server_panic"` 

```python
def error_type(error_type: typing.Literal["frontend_exception", "frontend_observable", "server_panic"]) -> typing.Self
```

### <span class="badge object-method"></span> flamegraph_diff

```python
def flamegraph_diff(flamegraph_diff: bool) -> typing.Self
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

### <span class="badge object-method"></span> labels

```python
def labels(labels: str) -> typing.Self
```

### <span class="badge object-method"></span> level_column

```python
def level_column(level_column: bool) -> typing.Self
```

### <span class="badge object-method"></span> lines

```python
def lines(lines: int) -> typing.Self
```

### <span class="badge object-method"></span> max_val

```python
def max_val(max_val: float) -> typing.Self
```

### <span class="badge object-method"></span> max_data_points

MaxDataPoints is the maximum number of data points that should be returned from a time series query.

NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated

from the number of pixels visible in a visualization

```python
def max_data_points(max_data_points: int) -> typing.Self
```

### <span class="badge object-method"></span> min_val

```python
def min_val(min_val: float) -> typing.Self
```

### <span class="badge object-method"></span> nodes

```python
def nodes(nodes: cogbuilder.Builder[testdata.NodesQuery]) -> typing.Self
```

### <span class="badge object-method"></span> noise

```python
def noise(noise: float) -> typing.Self
```

### <span class="badge object-method"></span> points

```python
def points(points: list[list[object]]) -> typing.Self
```

### <span class="badge object-method"></span> pulse_wave

```python
def pulse_wave(pulse_wave: cogbuilder.Builder[testdata.PulseWaveQuery]) -> typing.Self
```

### <span class="badge object-method"></span> query_type

QueryType is an optional identifier for the type of query.

It can be used to distinguish different types of queries.

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> raw_frame_content

```python
def raw_frame_content(raw_frame_content: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

RefID is the unique identifier of the query, set by the frontend call.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> result_assertions

Optionally define expected query result behavior

```python
def result_assertions(result_assertions: cogbuilder.Builder[testdata.ResultAssertions]) -> typing.Self
```

### <span class="badge object-method"></span> scenario_id

Possible enum values:

 - `"annotations"` 

 - `"arrow"` 

 - `"csv_content"` 

 - `"csv_file"` 

 - `"csv_metric_values"` 

 - `"datapoints_outside_range"` 

 - `"exponential_heatmap_bucket_data"` 

 - `"flame_graph"` 

 - `"grafana_api"` 

 - `"linear_heatmap_bucket_data"` 

 - `"live"` 

 - `"logs"` 

 - `"manual_entry"` 

 - `"no_data_points"` 

 - `"node_graph"` 

 - `"predictable_csv_wave"` 

 - `"predictable_pulse"` 

 - `"random_walk"` 

 - `"random_walk_table"` 

 - `"random_walk_with_error"` 

 - `"raw_frame"` 

 - `"server_error_500"` 

 - `"simulation"` 

 - `"slow_query"` 

 - `"streaming_client"` 

 - `"table_static"` 

 - `"trace"` 

 - `"usa"` 

 - `"variables-query"` 

```python
def scenario_id(scenario_id: typing.Literal["annotations", "arrow", "csv_content", "csv_file", "csv_metric_values", "datapoints_outside_range", "exponential_heatmap_bucket_data", "flame_graph", "grafana_api", "linear_heatmap_bucket_data", "live", "logs", "manual_entry", "no_data_points", "node_graph", "predictable_csv_wave", "predictable_pulse", "random_walk", "random_walk_table", "random_walk_with_error", "raw_frame", "server_error_500", "simulation", "slow_query", "streaming_client", "table_static", "trace", "usa", "variables-query"]) -> typing.Self
```

### <span class="badge object-method"></span> series_count

```python
def series_count(series_count: int) -> typing.Self
```

### <span class="badge object-method"></span> sim

```python
def sim(sim: cogbuilder.Builder[testdata.SimulationQuery]) -> typing.Self
```

### <span class="badge object-method"></span> span_count

```python
def span_count(span_count: int) -> typing.Self
```

### <span class="badge object-method"></span> spread

```python
def spread(spread: float) -> typing.Self
```

### <span class="badge object-method"></span> start_value

```python
def start_value(start_value: float) -> typing.Self
```

### <span class="badge object-method"></span> stream

```python
def stream(stream: cogbuilder.Builder[testdata.StreamingQuery]) -> typing.Self
```

### <span class="badge object-method"></span> string_input

common parameter used by many query types

```python
def string_input(string_input: str) -> typing.Self
```

### <span class="badge object-method"></span> time_range

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```python
def time_range(time_range: cogbuilder.Builder[testdata.TimeRange]) -> typing.Self
```

### <span class="badge object-method"></span> usa

```python
def usa(usa: cogbuilder.Builder[testdata.USAQuery]) -> typing.Self
```

### <span class="badge object-method"></span> with_nil

```python
def with_nil(with_nil: bool) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
