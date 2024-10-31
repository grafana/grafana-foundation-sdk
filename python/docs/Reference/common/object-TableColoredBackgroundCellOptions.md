---
title: <span class="badge object-type-class"></span> TableColoredBackgroundCellOptions
---
# <span class="badge object-type-class"></span> TableColoredBackgroundCellOptions

Colored background cell options

## Definition

```python
class TableColoredBackgroundCellOptions:
    """
    Colored background cell options
    """

    type_val: typing.Literal["color-background"]
    mode: typing.Optional[common.TableCellBackgroundDisplayMode]
    apply_to_row: typing.Optional[bool]
    wrap_text: typing.Optional[bool]
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

 * <span class="badge builder"></span> [TableColoredBackgroundCellOptions](./builder-TableColoredBackgroundCellOptions.md)
