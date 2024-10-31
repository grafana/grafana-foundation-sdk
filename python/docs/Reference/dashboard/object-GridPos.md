---
title: <span class="badge object-type-class"></span> GridPos
---
# <span class="badge object-type-class"></span> GridPos

Position and dimensions of a panel in the grid

## Definition

```python
class GridPos:
    """
    Position and dimensions of a panel in the grid
    """

    # Panel height. The height is the number of rows from the top edge of the panel.
    h: int
    # Panel width. The width is the number of columns from the left edge of the panel.
    w: int
    # Panel x. The x coordinate is the number of columns from the left edge of the grid
    x: int
    # Panel y. The y coordinate is the number of rows from the top edge of the grid
    y: int
    # Whether the panel is fixed within the grid. If true, the panel will not be affected by other panels' interactions
    static: typing.Optional[bool]
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

