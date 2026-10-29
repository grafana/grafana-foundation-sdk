---
title: <span class="badge builder"></span> RuleSpecBuilder
---
# <span class="badge builder"></span> RuleSpecBuilder

## Constructor

```go
func NewRuleSpecBuilder() *RuleSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RuleSpecBuilder) Build() (RuleSpec, error)
```

### <span class="badge object-method"></span> Expressions

```go
func (builder *RuleSpecBuilder) Expressions(expressions alertingv0alpha1.ExpressionMap) *RuleSpecBuilder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *RuleSpecBuilder) Labels(labels map[string]alertingv0alpha1.TemplateString) *RuleSpecBuilder
```

### <span class="badge object-method"></span> Paused

```go
func (builder *RuleSpecBuilder) Paused(paused bool) *RuleSpecBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *RuleSpecBuilder) Title(title string) *RuleSpecBuilder
```

### <span class="badge object-method"></span> Trigger

```go
func (builder *RuleSpecBuilder) Trigger(trigger cog.Builder[alertingv0alpha1.IntervalTrigger]) *RuleSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RuleSpec](./object-RuleSpec.md)
