---
title: <span class="badge object-type-class"></span> AnnotationContainer
---
# <span class="badge object-type-class"></span> AnnotationContainer

Contains the list of annotations that are associated with the dashboard.

Annotations are used to overlay event markers and overlay event tags on graphs.

Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.

See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/

## Definition

```python
class AnnotationContainer:
    """
    Contains the list of annotations that are associated with the dashboard.
    Annotations are used to overlay event markers and overlay event tags on graphs.
    Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
    See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
    """

    # List of annotations
    list_val: typing.Optional[list[dashboard.AnnotationQuery]]
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

