---
title: <span class="badge builder"></span> RecordingRule
---
# <span class="badge builder"></span> RecordingRule

## Constructor

```python
RecordingRule()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv1beta1.RecordingRuleSpec
```

### <span class="badge object-method"></span> expressions

```python
def expressions(expressions: alertingv1beta1.ExpressionMap) -> typing.Self
```

### <span class="badge object-method"></span> labels

```python
def labels(labels: dict[str, alertingv1beta1.TemplateString]) -> typing.Self
```

### <span class="badge object-method"></span> metric

```python
def metric(metric: alertingv1beta1.MetricName) -> typing.Self
```

### <span class="badge object-method"></span> paused

```python
def paused(paused: bool) -> typing.Self
```

### <span class="badge object-method"></span> target_datasource_uid

```python
def target_datasource_uid(target_datasource_uid: alertingv1beta1.DatasourceUID) -> typing.Self
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

 * <span class="badge object-type-class"></span> [RecordingRuleSpec](./object-RecordingRuleSpec.md)
