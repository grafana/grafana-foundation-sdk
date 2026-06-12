---
title: <span class="badge builder"></span> QueryV2
---
# <span class="badge builder"></span> QueryV2

## Constructor

```python
QueryV2()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2.DataQueryKind
```

### <span class="badge object-method"></span> adhoc_filters

Additional Ad-hoc filters that take precedence over Scope on conflict.

```python
def adhoc_filters(adhoc_filters: list[cogbuilder.Builder[prometheus.AdhocFilters]]) -> typing.Self
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```python
def datasource(ref: cogbuilder.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) -> typing.Self
```

### <span class="badge object-method"></span> editor_mode

what we should show in the editor

Possible enum values:

 - `"builder"` 

 - `"code"` 

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

The response format

Possible enum values:

 - `"time_series"` 

 - `"table"` 

 - `"heatmap"` 

```python
def format_val(format_val: prometheus.PromQueryFormat) -> typing.Self
```

### <span class="badge object-method"></span> group_by_keys

Group By parameters to apply to aggregate expressions in the query

```python
def group_by_keys(group_by_keys: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

NOTE: this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> instant

Returns only the latest value that Prometheus has scraped for the requested time series

```python
def instant(instant: bool) -> typing.Self
```

### <span class="badge object-method"></span> interval

An additional lower limit for the step parameter of the Prometheus query and for the

`$__interval` and `$__rate_interval` variables.

```python
def interval(interval: str) -> typing.Self
```

### <span class="badge object-method"></span> interval_factor

Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

Deprecated: use interval

```python
def interval_factor(interval_factor: int) -> typing.Self
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
def labels(labels: dict[str, str]) -> typing.Self
```

### <span class="badge object-method"></span> legend_format

Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname

```python
def legend_format(legend_format: str) -> typing.Self
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

### <span class="badge object-method"></span> range_val

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```python
def range_val(range_val: bool) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

RefID is the unique identifier of the query, set by the frontend call.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> result_assertions

Optionally define expected query result behavior

```python
def result_assertions(result_assertions: cogbuilder.Builder[prometheus.ResultAssertions]) -> typing.Self
```

### <span class="badge object-method"></span> scopes

A set of filters applied to apply to the query

```python
def scopes(scopes: list[cogbuilder.Builder[prometheus.Scopes]]) -> typing.Self
```

### <span class="badge object-method"></span> time_range

TimeRange represents the query range

NOTE: unlike generic /ds/query, we can now send explicit time values in each query

NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly

```python
def time_range(time_range: cogbuilder.Builder[prometheus.TimeRange]) -> typing.Self
```

### <span class="badge object-method"></span> version

```python
def version(version: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
