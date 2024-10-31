---
title: <span class="badge object-type-class"></span> RowsHeatmapOptions
---
# <span class="badge object-type-class"></span> RowsHeatmapOptions

Controls frame rows options

## Definition

```python
class RowsHeatmapOptions:
    """
    Controls frame rows options
    """

    # Sets the name of the cell when not calculating from data
    value: typing.Optional[str]
    # Controls tick alignment when not calculating from data
    layout: typing.Optional[common.HeatmapCellLayout]
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

 * <span class="badge builder"></span> [RowsHeatmapOptions](./builder-RowsHeatmapOptions.md)
