---
title: <span class="badge object-type-class"></span> FieldColor
---
# <span class="badge object-type-class"></span> FieldColor

Map a field to a color.

## Definition

```python
class FieldColor:
    """
    Map a field to a color.
    """

    # The main color scheme mode.
    mode: dashboard.FieldColorModeId
    # The fixed color value for fixed or shades color modes.
    fixed_color: typing.Optional[str]
    # Some visualizations need to know how to assign a series color from by value color schemes.
    series_by: typing.Optional[dashboard.FieldColorSeriesByMode]
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

 * <span class="badge builder"></span> [FieldColor](./builder-FieldColor.md)
