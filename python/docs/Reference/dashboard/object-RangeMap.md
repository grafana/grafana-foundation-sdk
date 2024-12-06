---
title: <span class="badge object-type-class"></span> RangeMap
---
# <span class="badge object-type-class"></span> RangeMap

Maps numerical ranges to a display text and color.

For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.

## Definition

```python
class RangeMap:
    """
    Maps numerical ranges to a display text and color.
    For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
    """

    type_val: typing.Literal["range"]
    # Range to match against and the result to apply when the value is within the range
    options: dashboard.DashboardRangeMapOptions
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

 * <span class="badge builder"></span> [RangeMap](./builder-RangeMap.md)
