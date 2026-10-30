---
title: <span class="badge object-type-class"></span> RecordingRuleSpec
---
# <span class="badge object-type-class"></span> RecordingRuleSpec

## Definition

```python
class RecordingRuleSpec:
    title: str
    paused: typing.Optional[bool]
    trigger: alertingv1beta1.IntervalTrigger
    labels: typing.Optional[dict[str, alertingv1beta1.TemplateString]]
    metric: alertingv1beta1.MetricName
    expressions: alertingv1beta1.ExpressionMap
    target_datasource_uid: alertingv1beta1.DatasourceUID
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

 * <span class="badge builder"></span> [RecordingRule](./builder-RecordingRule.md)
