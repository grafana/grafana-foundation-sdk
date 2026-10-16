---
title: <span class="badge object-type-class"></span> NotebookSpec
---
# <span class="badge object-type-class"></span> NotebookSpec

A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time

settings, and a map of elements (panels and cells) referenced by the layout.

## Definition

```python
class NotebookSpec:
    """
    A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time
    settings, and a map of elements (panels and cells) referenced by the layout.
    """

    title: str
    description: typing.Optional[str]
    tags: list[str]
    time_settings: dashboardv2beta1.TimeSettingsSpec
    elements: dict[str, dashboardv2beta1.NotebookElement]
    layout: dashboardv2beta1.NotebookLayoutKind
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

 * <span class="badge builder"></span> [NotebookSpec](./builder-NotebookSpec.md)
