---
title: <span class="badge object-type-class"></span> TableBarGaugeCellOptions
---
# <span class="badge object-type-class"></span> TableBarGaugeCellOptions

Gauge cell options

## Definition

```python
class TableBarGaugeCellOptions:
    """
    Gauge cell options
    """

    type_val: typing.Literal["gauge"]
    mode: typing.Optional[common.BarGaugeDisplayMode]
    value_display_mode: typing.Optional[common.BarGaugeValueMode]
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

 * <span class="badge builder"></span> [TableBarGaugeCellOptions](./builder-TableBarGaugeCellOptions.md)
