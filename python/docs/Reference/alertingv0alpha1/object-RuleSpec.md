---
title: <span class="badge object-type-class"></span> RuleSpec
---
# <span class="badge object-type-class"></span> RuleSpec

## Definition

```python
class RuleSpec:
    title: str
    paused: typing.Optional[bool]
    trigger: alertingv0alpha1.IntervalTrigger
    labels: typing.Optional[dict[str, alertingv0alpha1.TemplateString]]
    expressions: alertingv0alpha1.ExpressionMap
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

 * <span class="badge builder"></span> [RuleSpec](./builder-RuleSpec.md)
