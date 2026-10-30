---
title: <span class="badge builder"></span> RuleRefBuilder
---
# <span class="badge builder"></span> RuleRefBuilder

## Constructor

```go
func NewRuleRefBuilder() *RuleRefBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RuleRefBuilder) Build() (RuleRef, error)
```

### <span class="badge object-method"></span> Name

name is the metadata.name of an AlertRule or RecordingRule resource.

```go
func (builder *RuleRefBuilder) Name(name alertingv1beta1.RuleUID) *RuleRefBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RuleRef](./object-RuleRef.md)
