---
title: <span class="badge object-type-class"></span> CellValues
---
# <span class="badge object-type-class"></span> CellValues

Controls cell value options

## Definition

```python
class CellValues:
    """
    Controls cell value options
    """

    # Controls the cell value unit
    unit: typing.Optional[str]
    # Controls the number of decimals for cell values
    decimals: typing.Optional[float]
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

 * <span class="badge builder"></span> [CellValues](./builder-CellValues.md)
