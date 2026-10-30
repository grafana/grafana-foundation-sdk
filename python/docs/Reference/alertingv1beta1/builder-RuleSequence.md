---
title: <span class="badge builder"></span> RuleSequence
---
# <span class="badge builder"></span> RuleSequence

## Constructor

```python
RuleSequence()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv1beta1.RuleSequenceSpec
```

### <span class="badge object-method"></span> alert_rules

```python
def alert_rules(alert_rules: list[cogbuilder.Builder[alertingv1beta1.RuleRef]]) -> typing.Self
```

### <span class="badge object-method"></span> recording_rules

```python
def recording_rules(recording_rules: list[cogbuilder.Builder[alertingv1beta1.RuleRef]]) -> typing.Self
```

### <span class="badge object-method"></span> trigger

```python
def trigger(trigger: cogbuilder.Builder[alertingv1beta1.IntervalTrigger]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [RuleSequenceSpec](./object-RuleSequenceSpec.md)
