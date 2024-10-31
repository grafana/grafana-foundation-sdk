---
title: <span class="badge object-type-class"></span> BackgroundConfig
---
# <span class="badge object-type-class"></span> BackgroundConfig

## Definition

```python
class BackgroundConfig:
    color: typing.Optional[common.ColorDimensionConfig]
    image: typing.Optional[common.ResourceDimensionConfig]
    size: typing.Optional[canvas.BackgroundImageSize]
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

 * <span class="badge builder"></span> [BackgroundConfig](./builder-BackgroundConfig.md)
