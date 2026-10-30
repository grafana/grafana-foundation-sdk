---
title: <span class="badge object-type-class"></span> AlertRuleStatus
---
# <span class="badge object-type-class"></span> AlertRuleStatus

## Definition

```python
class AlertRuleStatus:
    health: typing.Optional[alertingv1beta1.AlertRuleHealth]
    state: typing.Optional[alertingv1beta1.AlertRuleState]
    state_reason: typing.Optional[alertingv1beta1.AlertRuleStateReason]
    last_evaluation_time: typing.Optional[str]
    # duration of the last evaluation in seconds
    evaluation_duration: typing.Optional[float]
    last_error: typing.Optional[str]
    totals: typing.Optional[alertingv1beta1.AlertRuleInstanceTotals]
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

 * <span class="badge builder"></span> [AlertRuleStatus](./builder-AlertRuleStatus.md)
