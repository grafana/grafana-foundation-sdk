---
title: <span class="badge object-type-class"></span> MovingAverage
---
# <span class="badge object-type-class"></span> MovingAverage

#MovingAverage's settings are overridden in types.ts

## Definition

```python
class MovingAverage:
    """
    #MovingAverage's settings are overridden in types.ts
    """

    pipeline_agg: typing.Optional[str]
    field: typing.Optional[str]
    type_val: typing.Literal["moving_avg"]
    id_val: str
    settings: typing.Optional[dict[str, object]]
    hide: typing.Optional[bool]
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

 * <span class="badge builder"></span> [MovingAverage](./builder-MovingAverage.md)
