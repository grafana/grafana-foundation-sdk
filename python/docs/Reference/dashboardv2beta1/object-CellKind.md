---
title: <span class="badge object-type-class"></span> CellKind
---
# <span class="badge object-type-class"></span> CellKind

A cell holds non-panel narrative content (markdown text, code) in a notebook layout.

Panel cells are not represented here — they reuse V2PanelKind.

## Definition

```python
class CellKind:
    """
    A cell holds non-panel narrative content (markdown text, code) in a notebook layout.
    Panel cells are not represented here — they reuse V2PanelKind.
    """

    kind: typing.Literal["Cell"]
    spec: dashboardv2beta1.CellSpec
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

 * <span class="badge builder"></span> [CellKind](./builder-CellKind.md)
