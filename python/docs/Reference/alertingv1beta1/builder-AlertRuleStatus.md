---
title: <span class="badge builder"></span> AlertRuleStatus
---
# <span class="badge builder"></span> AlertRuleStatus

## Constructor

```python
AlertRuleStatus()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv1beta1.AlertRuleStatus
```

### <span class="badge object-method"></span> evaluation_duration

duration of the last evaluation in seconds

```python
def evaluation_duration(evaluation_duration: float) -> typing.Self
```

### <span class="badge object-method"></span> health

```python
def health(health: alertingv1beta1.AlertRuleHealth) -> typing.Self
```

### <span class="badge object-method"></span> last_error

```python
def last_error(last_error: str) -> typing.Self
```

### <span class="badge object-method"></span> last_evaluation_time

```python
def last_evaluation_time(last_evaluation_time: str) -> typing.Self
```

### <span class="badge object-method"></span> state

```python
def state(state: alertingv1beta1.AlertRuleState) -> typing.Self
```

### <span class="badge object-method"></span> state_reason

```python
def state_reason(state_reason: alertingv1beta1.AlertRuleStateReason) -> typing.Self
```

### <span class="badge object-method"></span> totals

```python
def totals(totals: cogbuilder.Builder[alertingv1beta1.AlertRuleInstanceTotals]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AlertRuleStatus](./object-AlertRuleStatus.md)
