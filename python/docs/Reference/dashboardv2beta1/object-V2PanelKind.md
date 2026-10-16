---
title: <span class="badge object-type-class"></span> V2PanelKind
---
# <span class="badge object-type-class"></span> V2PanelKind

The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,

which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has

to be forked rather than shared because PanelKind reaches TransformationKind through

QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not

change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.

## Definition

```python
class V2PanelKind:
    """
    The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,
    which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has
    to be forked rather than shared because PanelKind reaches TransformationKind through
    QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not
    change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.
    """

    kind: typing.Literal["Panel"]
    spec: dashboardv2beta1.V2PanelSpec
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

 * <span class="badge builder"></span> [V2PanelKind](./builder-V2PanelKind.md)
