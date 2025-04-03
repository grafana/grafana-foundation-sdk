---
title: <span class="badge object-type-class"></span> AnnotationQuery
---
# <span class="badge object-type-class"></span> AnnotationQuery

TODO docs

FROM: AnnotationQuery in grafana-data/src/types/annotations.ts

## Definition

```python
class AnnotationQuery:
    """
    TODO docs
    FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
    """

    # Name of annotation.
    name: str
    # Datasource where the annotations data is
    datasource: dashboard.DataSourceRef
    # When enabled the annotation query is issued with every dashboard refresh
    enable: bool
    # Annotation queries can be toggled on or off at the top of the dashboard.
    # When hide is true, the toggle is not shown in the dashboard.
    hide: typing.Optional[bool]
    # Color to use for the annotation event markers
    icon_color: str
    # Filters to apply when fetching annotations
    filter_val: typing.Optional[dashboard.AnnotationPanelFilter]
    # TODO.. this should just be a normal query target
    target: typing.Optional[dashboard.AnnotationTarget]
    # TODO -- this should not exist here, it is based on the --grafana-- datasource
    type_val: typing.Optional[str]
    # Set to 1 for the standard annotation query all dashboards have by default.
    built_in: typing.Optional[float]
    expr: typing.Optional[str]
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

 * <span class="badge builder"></span> [AnnotationQuery](./builder-AnnotationQuery.md)
