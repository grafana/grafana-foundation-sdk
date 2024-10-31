---
title: <span class="badge object-type-class"></span> MetricQuery
---
# <span class="badge object-type-class"></span> MetricQuery

@deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.

## Definition

```python
class MetricQuery:
    """
    @deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.
    """

    # GCP project to execute the query against.
    project_name: str
    # Alignment function to be used. Defaults to ALIGN_MEAN.
    per_series_aligner: typing.Optional[str]
    # Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignment_period: typing.Optional[str]
    # Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
    alias_by: typing.Optional[str]
    editor_mode: str
    metric_type: str
    # Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    cross_series_reducer: str
    # Array of labels to group data by.
    group_bys: typing.Optional[list[str]]
    # Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    filters: typing.Optional[list[str]]
    metric_kind: typing.Optional[googlecloudmonitoring.MetricKind]
    value_type: typing.Optional[str]
    view: typing.Optional[str]
    # MQL query to be executed.
    query: str
    # Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    preprocessor: typing.Optional[googlecloudmonitoring.PreprocessorType]
    # To disable the graphPeriod, it should explictly be set to 'disabled'.
    graph_period: typing.Optional[str]
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

 * <span class="badge builder"></span> [MetricQuery](./builder-MetricQuery.md)
