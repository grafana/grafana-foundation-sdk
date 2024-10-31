---
title: <span class="badge object-type-class"></span> CloudMonitoringQuery
---
# <span class="badge object-type-class"></span> CloudMonitoringQuery

## Definition

```python
class CloudMonitoringQuery(cogvariants.Dataquery):
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
    # Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
    alias_by: typing.Optional[str]
    # GCM query type.
    # queryType: #QueryType
    # Time Series List sub-query properties.
    time_series_list: typing.Optional[googlecloudmonitoring.TimeSeriesList]
    # Time Series sub-query properties.
    time_series_query: typing.Optional[googlecloudmonitoring.TimeSeriesQuery]
    # SLO sub-query properties.
    slo_query: typing.Optional[googlecloudmonitoring.SLOQuery]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[dashboard.DataSourceRef]
    # Time interval in milliseconds.
    interval_ms: typing.Optional[float]
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

 * <span class="badge builder"></span> [CloudMonitoringQuery](./builder-CloudMonitoringQuery.md)
