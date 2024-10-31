---
title: <span class="badge object-type-class"></span> ThresholdsConfig
---
# <span class="badge object-type-class"></span> ThresholdsConfig

Thresholds configuration for the panel

## Definition

```python
class ThresholdsConfig:
    """
    Thresholds configuration for the panel
    """

    # Thresholds mode.
    mode: dashboard.ThresholdsMode
    # Must be sorted by 'value', first value is always -Infinity
    steps: list[dashboard.Threshold]
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

 * <span class="badge builder"></span> [ThresholdsConfig](./builder-ThresholdsConfig.md)
