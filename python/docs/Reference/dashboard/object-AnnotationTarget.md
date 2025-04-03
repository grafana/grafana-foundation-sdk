---
title: <span class="badge object-type-class"></span> AnnotationTarget
---
# <span class="badge object-type-class"></span> AnnotationTarget

TODO: this should be a regular DataQuery that depends on the selected dashboard

these match the properties of the "grafana" datasouce that is default in most dashboards

## Definition

```python
class AnnotationTarget:
    """
    TODO: this should be a regular DataQuery that depends on the selected dashboard
    these match the properties of the "grafana" datasouce that is default in most dashboards
    """

    # Only required/valid for the grafana datasource...
    # but code+tests is already depending on it so hard to change
    limit: int
    # Only required/valid for the grafana datasource...
    # but code+tests is already depending on it so hard to change
    match_any: bool
    # Only required/valid for the grafana datasource...
    # but code+tests is already depending on it so hard to change
    tags: list[str]
    # Only required/valid for the grafana datasource...
    # but code+tests is already depending on it so hard to change
    type_val: str
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

 * <span class="badge builder"></span> [AnnotationTarget](./builder-AnnotationTarget.md)
