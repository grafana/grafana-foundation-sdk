---
title: <span class="badge object-type-class"></span> NotebookLayoutItemSpec
---
# <span class="badge object-type-class"></span> NotebookLayoutItemSpec

One ordered item in a notebook layout. `element` references either a CellKind

(markdown/code content) or a V2PanelKind in the notebook's elements map. `source`

records who authored the cell; `collapsed` hides the body in the UI.

## Definition

```python
class NotebookLayoutItemSpec:
    """
    One ordered item in a notebook layout. `element` references either a CellKind
    (markdown/code content) or a V2PanelKind in the notebook's elements map. `source`
    records who authored the cell; `collapsed` hides the body in the UI.
    """

    element: dashboardv2beta1.ElementReference
    source: typing.Literal["assistant", "user"]
    collapsed: typing.Optional[bool]
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

 * <span class="badge builder"></span> [NotebookLayoutItemSpec](./builder-NotebookLayoutItemSpec.md)
