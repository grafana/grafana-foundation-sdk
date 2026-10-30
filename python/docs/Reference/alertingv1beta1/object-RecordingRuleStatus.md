---
title: <span class="badge object-type-class"></span> RecordingRuleStatus
---
# <span class="badge object-type-class"></span> RecordingRuleStatus

## Definition

```python
class RecordingRuleStatus:
    health: typing.Optional[alertingv1beta1.RecordingRuleHealth]
    last_evaluation_time: typing.Optional[str]
    # duration of the last evaluation in seconds
    evaluation_duration: typing.Optional[float]
    last_error: typing.Optional[str]
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

 * <span class="badge builder"></span> [RecordingRuleStatus](./builder-RecordingRuleStatus.md)
