---
title: <span class="badge object-type-class"></span> LineConfig
---
# <span class="badge object-type-class"></span> LineConfig

TODO docs

## Definition

```python
class LineConfig:
    """
    TODO docs
    """

    line_color: typing.Optional[str]
    line_width: typing.Optional[float]
    line_interpolation: typing.Optional[common.LineInterpolation]
    line_style: typing.Optional[common.LineStyle]
    # Indicate if null values should be treated as gaps or connected.
    # When the value is a number, it represents the maximum delta in the
    # X axis that should be considered connected.  For timeseries, this is milliseconds
    span_nulls: typing.Optional[typing.Union[bool, float]]
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

 * <span class="badge builder"></span> [LineConfig](./builder-LineConfig.md)
