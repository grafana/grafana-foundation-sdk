---
title: <span class="badge builder"></span> RecordingRuleStatus
---
# <span class="badge builder"></span> RecordingRuleStatus

## Constructor

```python
RecordingRuleStatus()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv1beta1.RecordingRuleStatus
```

### <span class="badge object-method"></span> evaluation_duration

duration of the last evaluation in seconds

```python
def evaluation_duration(evaluation_duration: float) -> typing.Self
```

### <span class="badge object-method"></span> health

```python
def health(health: alertingv1beta1.RecordingRuleHealth) -> typing.Self
```

### <span class="badge object-method"></span> last_error

```python
def last_error(last_error: str) -> typing.Self
```

### <span class="badge object-method"></span> last_evaluation_time

```python
def last_evaluation_time(last_evaluation_time: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [RecordingRuleStatus](./object-RecordingRuleStatus.md)
