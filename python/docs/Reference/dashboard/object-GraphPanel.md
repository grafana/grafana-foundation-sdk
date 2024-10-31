---
title: <span class="badge object-type-class"></span> GraphPanel
---
# <span class="badge object-type-class"></span> GraphPanel

Support for legacy graph panel.

@deprecated this a deprecated panel type

## Definition

```python
class GraphPanel:
    """
    Support for legacy graph panel.
    @deprecated this a deprecated panel type
    """

    type_val: typing.Literal["graph"]
    # @deprecated this is part of deprecated graph panel
    legend: typing.Optional[dashboard.DashboardGraphPanelLegend]
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

 * <span class="badge builder"></span> [GraphPanel](./builder-GraphPanel.md)
