---
title: <span class="badge object-type-class"></span> MetricFindValue
---
# <span class="badge object-type-class"></span> MetricFindValue

Define the MetricFindValue type

## Definition

```python
class MetricFindValue:
    """
    Define the MetricFindValue type
    """

    text: str
    value: typing.Optional[typing.Union[str, float]]
    group: typing.Optional[str]
    expandable: typing.Optional[bool]
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

 * <span class="badge builder"></span> [MetricFindValue](./builder-MetricFindValue.md)
