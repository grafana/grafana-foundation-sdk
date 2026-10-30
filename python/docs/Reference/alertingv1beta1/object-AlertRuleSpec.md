---
title: <span class="badge object-type-class"></span> AlertRuleSpec
---
# <span class="badge object-type-class"></span> AlertRuleSpec

## Definition

```python
class AlertRuleSpec:
    title: str
    paused: typing.Optional[bool]
    trigger: alertingv1beta1.IntervalTrigger
    labels: typing.Optional[dict[str, alertingv1beta1.TemplateString]]
    annotations: typing.Optional[dict[str, alertingv1beta1.TemplateString]]
    for_val: typing.Optional[str]
    keep_firing_for: typing.Optional[str]
    missing_series_evals_to_resolve: typing.Optional[int]
    no_data_state: alertingv1beta1.NoDataState
    exec_err_state: alertingv1beta1.ExecErrState
    notification_settings: typing.Optional[alertingv1beta1.NotificationSettings]
    expressions: alertingv1beta1.ExpressionMap
    panel_ref: typing.Optional[alertingv1beta1.PanelRef]
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

 * <span class="badge builder"></span> [AlertRule](./builder-AlertRule.md)
