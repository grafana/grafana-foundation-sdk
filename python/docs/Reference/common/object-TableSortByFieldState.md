---
title: <span class="badge object-type-class"></span> TableSortByFieldState
---
# <span class="badge object-type-class"></span> TableSortByFieldState

Sort by field state

## Definition

```python
class TableSortByFieldState:
    """
    Sort by field state
    """

    # Sets the display name of the field to sort by
    display_name: str
    # Flag used to indicate descending sort order
    desc: typing.Optional[bool]
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

 * <span class="badge builder"></span> [TableSortByFieldState](./builder-TableSortByFieldState.md)
