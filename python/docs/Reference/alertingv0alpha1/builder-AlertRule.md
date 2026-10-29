---
title: <span class="badge builder"></span> AlertRule
---
# <span class="badge builder"></span> AlertRule

## Constructor

```python
AlertRule()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv0alpha1.AlertRuleSpec
```

### <span class="badge object-method"></span> annotations

```python
def annotations(annotations: dict[str, alertingv0alpha1.TemplateString]) -> typing.Self
```

### <span class="badge object-method"></span> exec_err_state

```python
def exec_err_state(exec_err_state: alertingv0alpha1.ExecErrState) -> typing.Self
```

### <span class="badge object-method"></span> expressions

```python
def expressions(expressions: alertingv0alpha1.ExpressionMap) -> typing.Self
```

### <span class="badge object-method"></span> for_val

```python
def for_val(for_val: str) -> typing.Self
```

### <span class="badge object-method"></span> keep_firing_for

```python
def keep_firing_for(keep_firing_for: str) -> typing.Self
```

### <span class="badge object-method"></span> labels

```python
def labels(labels: dict[str, alertingv0alpha1.TemplateString]) -> typing.Self
```

### <span class="badge object-method"></span> missing_series_evals_to_resolve

```python
def missing_series_evals_to_resolve(missing_series_evals_to_resolve: int) -> typing.Self
```

### <span class="badge object-method"></span> no_data_state

```python
def no_data_state(no_data_state: alertingv0alpha1.NoDataState) -> typing.Self
```

### <span class="badge object-method"></span> notification_settings

```python
def notification_settings(notification_settings: cogbuilder.Builder[alertingv0alpha1.NotificationSettings]) -> typing.Self
```

### <span class="badge object-method"></span> panel_ref

```python
def panel_ref(panel_ref: cogbuilder.Builder[alertingv0alpha1.PanelRef]) -> typing.Self
```

### <span class="badge object-method"></span> paused

```python
def paused(paused: bool) -> typing.Self
```

### <span class="badge object-method"></span> title

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> trigger

```python
def trigger(trigger: cogbuilder.Builder[alertingv0alpha1.IntervalTrigger]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AlertRuleSpec](./object-AlertRuleSpec.md)
