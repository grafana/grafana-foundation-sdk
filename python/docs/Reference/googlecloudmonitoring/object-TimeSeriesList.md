---
title: <span class="badge object-type-class"></span> TimeSeriesList
---
# <span class="badge object-type-class"></span> TimeSeriesList

Time Series List sub-query properties.

## Definition

```python
class TimeSeriesList:
    """
    Time Series List sub-query properties.
    """

    # GCP project to execute the query against.
    project_name: str
    # Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    cross_series_reducer: str
    # Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignment_period: typing.Optional[str]
    # Alignment function to be used. Defaults to ALIGN_MEAN.
    per_series_aligner: typing.Optional[str]
    # Array of labels to group data by.
    group_bys: typing.Optional[list[str]]
    # Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    filters: typing.Optional[list[str]]
    # Data view, defaults to FULL.
    view: typing.Optional[str]
    # Annotation title.
    title: typing.Optional[str]
    # Annotation text.
    text: typing.Optional[str]
    # Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    secondary_cross_series_reducer: typing.Optional[str]
    # Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    secondary_alignment_period: typing.Optional[str]
    # Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
    secondary_per_series_aligner: typing.Optional[str]
    # Only present if a preprocessor is selected. Array of labels to group data by.
    secondary_group_bys: typing.Optional[list[str]]
    # Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    preprocessor: typing.Optional[googlecloudmonitoring.PreprocessorType]
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

 * <span class="badge builder"></span> [TimeSeriesList](./builder-TimeSeriesList.md)
