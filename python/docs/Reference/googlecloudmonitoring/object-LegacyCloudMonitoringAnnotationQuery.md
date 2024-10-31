---
title: <span class="badge object-type-class"></span> LegacyCloudMonitoringAnnotationQuery
---
# <span class="badge object-type-class"></span> LegacyCloudMonitoringAnnotationQuery

@deprecated Use AnnotationQuery instead. Legacy annotation query properties for migration purposes.

## Definition

```python
class LegacyCloudMonitoringAnnotationQuery:
    """
    @deprecated Use AnnotationQuery instead. Legacy annotation query properties for migration purposes.
    """

    # GCP project to execute the query against.
    project_name: str
    metric_type: str
    # Query refId.
    ref_id: str
    # Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    filters: list[str]
    metric_kind: googlecloudmonitoring.MetricKind
    value_type: str
    # Annotation title.
    title: str
    # Annotation text.
    text: str
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

 * <span class="badge builder"></span> [LegacyCloudMonitoringAnnotationQuery](./builder-LegacyCloudMonitoringAnnotationQuery.md)
