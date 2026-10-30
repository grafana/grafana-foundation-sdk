---
title: <span class="badge object-type-class"></span> RuleSequenceSpec
---
# <span class="badge object-type-class"></span> RuleSequenceSpec

## Definition

```python
class RuleSequenceSpec:
    trigger: alertingv1beta1.IntervalTrigger
    recording_rules: list[alertingv1beta1.RuleRef]
    alert_rules: typing.Optional[list[alertingv1beta1.RuleRef]]
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

 * <span class="badge builder"></span> [RuleSequence](./builder-RuleSequence.md)
