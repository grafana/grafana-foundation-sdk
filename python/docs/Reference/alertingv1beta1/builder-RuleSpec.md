---
title: <span class="badge builder"></span> RuleSpec
---
# <span class="badge builder"></span> RuleSpec

## Constructor

```python
RuleSpec()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv1beta1.RuleSpec
```

### <span class="badge object-method"></span> expressions

```python
def expressions(expressions: alertingv1beta1.ExpressionMap) -> typing.Self
```

### <span class="badge object-method"></span> labels

```python
def labels(labels: dict[str, alertingv1beta1.TemplateString]) -> typing.Self
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
def trigger(trigger: cogbuilder.Builder[alertingv1beta1.IntervalTrigger]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [RuleSpec](./object-RuleSpec.md)
