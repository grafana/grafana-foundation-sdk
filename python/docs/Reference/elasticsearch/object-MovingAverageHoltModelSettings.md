---
title: <span class="badge object-type-class"></span> MovingAverageHoltModelSettings
---
# <span class="badge object-type-class"></span> MovingAverageHoltModelSettings

## Definition

```python
class MovingAverageHoltModelSettings:
    model: typing.Literal["holt"]
    settings: elasticsearch.ElasticsearchMovingAverageHoltModelSettingsSettings
    window: str
    minimize: bool
    predict: str
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

 * <span class="badge builder"></span> [MovingAverageHoltModelSettings](./builder-MovingAverageHoltModelSettings.md)
