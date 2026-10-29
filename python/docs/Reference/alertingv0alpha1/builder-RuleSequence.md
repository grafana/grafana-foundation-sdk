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
def build() -> alertingv0alpha1.RuleSequenceSpec
```

### <span class="badge object-method"></span> alerting_rules

```python
def alerting_rules(alerting_rules: list[cogbuilder.Builder[alertingv0alpha1.RuleRef]]) -> typing.Self
```

### <span class="badge object-method"></span> recording_rules

```python
def recording_rules(recording_rules: list[cogbuilder.Builder[alertingv0alpha1.RuleRef]]) -> typing.Self
```

### <span class="badge object-method"></span> trigger

```python
def trigger(trigger: cogbuilder.Builder[alertingv0alpha1.IntervalTrigger]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [RuleSequenceSpec](./object-RuleSequenceSpec.md)
