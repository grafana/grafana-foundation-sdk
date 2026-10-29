---
title: <span class="badge builder"></span> RecordingRuleBuilder
---
# <span class="badge builder"></span> RecordingRuleBuilder

## Constructor

```go
func NewRecordingRuleBuilder() *RecordingRuleBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RecordingRuleBuilder) Build() (RecordingRuleSpec, error)
```

### <span class="badge object-method"></span> Expressions

```go
func (builder *RecordingRuleBuilder) Expressions(expressions alertingv0alpha1.ExpressionMap) *RecordingRuleBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *RecordingRuleBuilder) Labels(labels map[string]alertingv0alpha1.TemplateString) *RecordingRuleBuilder
```

### <span class="badge object-method"></span> Metric

```go
func (builder *RecordingRuleBuilder) Metric(metric alertingv0alpha1.MetricName) *RecordingRuleBuilder
```

### <span class="badge object-method"></span> Paused

```go
func (builder *RecordingRuleBuilder) Paused(paused bool) *RecordingRuleBuilder
```

### <span class="badge object-method"></span> TargetDatasourceUID

```go
func (builder *RecordingRuleBuilder) TargetDatasourceUID(targetDatasourceUID alertingv0alpha1.DatasourceUID) *RecordingRuleBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *RecordingRuleBuilder) Title(title string) *RecordingRuleBuilder
```

### <span class="badge object-method"></span> Trigger

```go
func (builder *RecordingRuleBuilder) Trigger(trigger cog.Builder[alertingv0alpha1.IntervalTrigger]) *RecordingRuleBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RecordingRuleSpec](./object-RecordingRuleSpec.md)
