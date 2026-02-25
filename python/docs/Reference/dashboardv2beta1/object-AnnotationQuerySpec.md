---
title: <span class="badge object-type-class"></span> AnnotationQuerySpec
---
# <span class="badge object-type-class"></span> AnnotationQuerySpec

## Definition

```python
class AnnotationQuerySpec:
    query: dashboardv2beta1.DataQueryKind
    enable: bool
    hide: bool
    icon_color: str
    name: str
    built_in: typing.Optional[bool]
    filter_val: typing.Optional[dashboardv2beta1.AnnotationPanelFilter]
    # Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
    placement: str
    # Catch-all field for datasource-specific properties. Should not be available in as code tooling.
    legacy_options: typing.Optional[dict[str, object]]
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

