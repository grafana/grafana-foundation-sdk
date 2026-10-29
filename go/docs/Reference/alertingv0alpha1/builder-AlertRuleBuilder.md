---
title: <span class="badge builder"></span> AlertRuleBuilder
---
# <span class="badge builder"></span> AlertRuleBuilder

## Constructor

```go
func NewAlertRuleBuilder() *AlertRuleBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AlertRuleBuilder) Build() (AlertRuleSpec, error)
```

### <span class="badge object-method"></span> Annotations

```go
func (builder *AlertRuleBuilder) Annotations(annotations map[string]alertingv0alpha1.TemplateString) *AlertRuleBuilder
```

### <span class="badge object-method"></span> ExecErrState

```go
func (builder *AlertRuleBuilder) ExecErrState(execErrState alertingv0alpha1.ExecErrState) *AlertRuleBuilder
```

### <span class="badge object-method"></span> Expressions

```go
func (builder *AlertRuleBuilder) Expressions(expressions alertingv0alpha1.ExpressionMap) *AlertRuleBuilder
```

### <span class="badge object-method"></span> For

```go
func (builder *AlertRuleBuilder) For(forArg string) *AlertRuleBuilder
```

### <span class="badge object-method"></span> KeepFiringFor

```go
func (builder *AlertRuleBuilder) KeepFiringFor(keepFiringFor string) *AlertRuleBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *AlertRuleBuilder) Labels(labels map[string]alertingv0alpha1.TemplateString) *AlertRuleBuilder
```

### <span class="badge object-method"></span> MissingSeriesEvalsToResolve

```go
func (builder *AlertRuleBuilder) MissingSeriesEvalsToResolve(missingSeriesEvalsToResolve int64) *AlertRuleBuilder
```

### <span class="badge object-method"></span> NoDataState

```go
func (builder *AlertRuleBuilder) NoDataState(noDataState alertingv0alpha1.NoDataState) *AlertRuleBuilder
```

### <span class="badge object-method"></span> NotificationSettings

```go
func (builder *AlertRuleBuilder) NotificationSettings(notificationSettings alertingv0alpha1.NotificationSettings) *AlertRuleBuilder
```

### <span class="badge object-method"></span> PanelRef

```go
func (builder *AlertRuleBuilder) PanelRef(panelRef cog.Builder[alertingv0alpha1.PanelRef]) *AlertRuleBuilder
```

### <span class="badge object-method"></span> Paused

```go
func (builder *AlertRuleBuilder) Paused(paused bool) *AlertRuleBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *AlertRuleBuilder) Title(title string) *AlertRuleBuilder
```

### <span class="badge object-method"></span> Trigger

```go
func (builder *AlertRuleBuilder) Trigger(trigger cog.Builder[alertingv0alpha1.IntervalTrigger]) *AlertRuleBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AlertRuleSpec](./object-AlertRuleSpec.md)
