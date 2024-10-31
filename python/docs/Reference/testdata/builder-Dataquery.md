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

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> drop_percent

Drop percentage (the chance we will lose a point 0-100)

```python
def drop_percent(drop_percent: float) -> typing.Self
```

### <span class="badge object-method"></span> error_type

```python
def error_type(error_type: typing.Literal["server_panic", "frontend_exception", "frontend_observable"]) -> typing.Self
```

### <span class="badge object-method"></span> flamegraph_diff

```python
def flamegraph_diff(flamegraph_diff: bool) -> typing.Self
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```python
def hide(hide: bool) -> typing.Self
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

### <span class="badge object-method"></span> nodes

```python
def nodes(nodes: cogbuilder.Builder[testdata.NodesQuery]) -> typing.Self
```

### <span class="badge object-method"></span> points

```python
def points(points: list[list[typing.Union[str, int]]]) -> typing.Self
```

### <span class="badge object-method"></span> pulse_wave

```python
def pulse_wave(pulse_wave: cogbuilder.Builder[testdata.PulseWaveQuery]) -> typing.Self
```

### <span class="badge object-method"></span> query_type

Specify the query flavor

TODO make this required and give it a default

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> raw_frame_content

```python
def raw_frame_content(raw_frame_content: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> scenario_id

```python
def scenario_id(scenario_id: testdata.TestDataQueryType) -> typing.Self
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

### <span class="badge object-method"></span> stream

```python
def stream(stream: cogbuilder.Builder[testdata.StreamingQuery]) -> typing.Self
```

### <span class="badge object-method"></span> string_input

```python
def string_input(string_input: str) -> typing.Self
```

### <span class="badge object-method"></span> usa

```python
def usa(usa: cogbuilder.Builder[testdata.USAQuery]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
