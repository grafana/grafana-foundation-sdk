---
title: <span class="badge object-type-class"></span> HideableFieldConfig
---
# <span class="badge object-type-class"></span> HideableFieldConfig

TODO docs

## Definition

```python
class HideableFieldConfig:
    """
    TODO docs
    """

    hide_from: typing.Optional[common.HideSeriesConfig]
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

 * <span class="badge builder"></span> [HideableFieldConfig](./builder-HideableFieldConfig.md)
