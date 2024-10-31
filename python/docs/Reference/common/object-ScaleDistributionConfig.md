---
title: <span class="badge object-type-class"></span> ScaleDistributionConfig
---
# <span class="badge object-type-class"></span> ScaleDistributionConfig

TODO docs

## Definition

```python
class ScaleDistributionConfig:
    """
    TODO docs
    """

    type_val: common.ScaleDistribution
    log: typing.Optional[float]
    linear_threshold: typing.Optional[float]
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

 * <span class="badge builder"></span> [ScaleDistributionConfig](./builder-ScaleDistributionConfig.md)
