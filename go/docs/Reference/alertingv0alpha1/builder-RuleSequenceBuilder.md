---
title: <span class="badge builder"></span> RuleSequenceBuilder
---
# <span class="badge builder"></span> RuleSequenceBuilder

## Constructor

```go
func NewRuleSequenceBuilder() *RuleSequenceBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RuleSequenceBuilder) Build() (RuleSequenceSpec, error)
```

### <span class="badge object-method"></span> AlertingRules

```go
func (builder *RuleSequenceBuilder) AlertingRules(alertingRules []cog.Builder[alertingv0alpha1.RuleRef]) *RuleSequenceBuilder
```

### <span class="badge object-method"></span> RecordingRules

```go
func (builder *RuleSequenceBuilder) RecordingRules(recordingRules []cog.Builder[alertingv0alpha1.RuleRef]) *RuleSequenceBuilder
```

### <span class="badge object-method"></span> Trigger

```go
func (builder *RuleSequenceBuilder) Trigger(trigger cog.Builder[alertingv0alpha1.IntervalTrigger]) *RuleSequenceBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RuleSequenceSpec](./object-RuleSequenceSpec.md)
