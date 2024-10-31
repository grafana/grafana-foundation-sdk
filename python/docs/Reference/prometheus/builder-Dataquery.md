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
def build() -> prometheus.Dataquery
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> editor_mode

Specifies which editor is being used to prepare the query. It can be "code" or "builder"

```python
def editor_mode(editor_mode: prometheus.QueryEditorMode) -> typing.Self
```

### <span class="badge object-method"></span> exemplar

Execute an additional query to identify interesting raw samples relevant for the given expr

```python
def exemplar(exemplar: bool) -> typing.Self
```

### <span class="badge object-method"></span> expr

The actual expression/query that will be evaluated by Prometheus

```python
def expr(expr: str) -> typing.Self
```

### <span class="badge object-method"></span> format_val

Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"

```python
def format_val(format_val: prometheus.PromQueryFormat) -> typing.Self
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> instant

Returns only the latest value that Prometheus has scraped for the requested time series

```python
def instant() -> typing.Self
```

### <span class="badge object-method"></span> interval

An additional lower limit for the step parameter of the Prometheus query and for the

`$__interval` and `$__rate_interval` variables.

```python
def interval(interval: str) -> typing.Self
```

### <span class="badge object-method"></span> interval_factor

@deprecated Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

```python
def interval_factor(interval_factor: float) -> typing.Self
```

### <span class="badge object-method"></span> legend_format

Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname

```python
def legend_format(legend_format: str) -> typing.Self
```

### <span class="badge object-method"></span> query_type

Specify the query flavor

TODO make this required and give it a default

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> range_val

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```python
def range_val() -> typing.Self
```

### <span class="badge object-method"></span> range_and_instant

```python
def range_and_instant() -> typing.Self
```

### <span class="badge object-method"></span> ref_id

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```python
def ref_id(ref_id: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
