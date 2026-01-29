---
title: <span class="badge builder"></span> ConditionalRenderingGroupBuilder
---
# <span class="badge builder"></span> ConditionalRenderingGroupBuilder

## Constructor

```go
func NewConditionalRenderingGroupBuilder() *ConditionalRenderingGroupBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ConditionalRenderingGroupBuilder) Build() (ConditionalRenderingGroupKind, error)
```

### <span class="badge object-method"></span> Condition

```go
func (builder *ConditionalRenderingGroupBuilder) Condition(condition dashboardv2beta1.ConditionalRenderingGroupSpecCondition) *ConditionalRenderingGroupBuilder
```

### <span class="badge object-method"></span> Items

```go
func (builder *ConditionalRenderingGroupBuilder) Items(items []dashboardv2beta1.ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) *ConditionalRenderingGroupBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *ConditionalRenderingGroupBuilder) Spec(spec cog.Builder[dashboardv2beta1.ConditionalRenderingGroupSpec]) *ConditionalRenderingGroupBuilder
```

### <span class="badge object-method"></span> Visibility

```go
func (builder *ConditionalRenderingGroupBuilder) Visibility(visibility dashboardv2beta1.ConditionalRenderingGroupSpecVisibility) *ConditionalRenderingGroupBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ConditionalRenderingGroupKind](./object-ConditionalRenderingGroupKind.md)
