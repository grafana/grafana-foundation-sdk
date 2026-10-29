---
title: <span class="badge object-type-class"></span> SimplifiedRouting
---
# <span class="badge object-type-class"></span> SimplifiedRouting

## Definition

```python
class SimplifiedRouting:
    type_val: str
    receiver: str
    group_by: typing.Optional[list[str]]
    group_wait: typing.Optional[alertingv0alpha1.PromDuration]
    group_interval: typing.Optional[alertingv0alpha1.PromDuration]
    repeat_interval: typing.Optional[alertingv0alpha1.PromDuration]
    mute_time_intervals: typing.Optional[list[alertingv0alpha1.TimeIntervalRef]]
    active_time_intervals: typing.Optional[list[alertingv0alpha1.TimeIntervalRef]]
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

 * <span class="badge builder"></span> [SimplifiedRouting](./builder-SimplifiedRouting.md)
