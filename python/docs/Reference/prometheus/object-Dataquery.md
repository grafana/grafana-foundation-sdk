---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```python
class Dataquery(cogvariants.Dataquery):
    # The actual expression/query that will be evaluated by Prometheus
    expr: str
    # Returns only the latest value that Prometheus has scraped for the requested time series
    instant: typing.Optional[bool]
    # Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    range_val: typing.Optional[bool]
    # Execute an additional query to identify interesting raw samples relevant for the given expr
    exemplar: typing.Optional[bool]
    # Specifies which editor is being used to prepare the query. It can be "code" or "builder"
    editor_mode: typing.Optional[prometheus.QueryEditorMode]
    # Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
    format_val: typing.Optional[prometheus.PromQueryFormat]
    # Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    legend_format: typing.Optional[str]
    # @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
    # See https://github.com/grafana/grafana/issues/48081
    interval_factor: typing.Optional[float]
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # Note this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[dashboard.DataSourceRef]
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
