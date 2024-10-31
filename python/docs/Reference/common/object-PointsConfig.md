---
title: <span class="badge object-type-class"></span> PointsConfig
---
# <span class="badge object-type-class"></span> PointsConfig

TODO docs

## Definition

```python
class PointsConfig:
    """
    TODO docs
    """

    show_points: typing.Optional[common.VisibilityMode]
    point_size: typing.Optional[float]
    point_color: typing.Optional[str]
    point_symbol: typing.Optional[str]
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

 * <span class="badge builder"></span> [PointsConfig](./builder-PointsConfig.md)
