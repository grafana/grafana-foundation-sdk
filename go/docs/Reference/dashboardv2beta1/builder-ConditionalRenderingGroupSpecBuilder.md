---
title: <span class="badge builder"></span> ConditionalRenderingGroupSpecBuilder
---
# <span class="badge builder"></span> ConditionalRenderingGroupSpecBuilder

## Constructor

```go
func NewConditionalRenderingGroupSpecBuilder() *ConditionalRenderingGroupSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ConditionalRenderingGroupSpecBuilder) Build() (ConditionalRenderingGroupSpec, error)
```

### <span class="badge object-method"></span> Condition

```go
func (builder *ConditionalRenderingGroupSpecBuilder) Condition(condition dashboardv2beta1.ConditionalRenderingGroupSpecCondition) *ConditionalRenderingGroupSpecBuilder
```

### <span class="badge object-method"></span> Items

```go
func (builder *ConditionalRenderingGroupSpecBuilder) Items(items []dashboardv2beta1.ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) *ConditionalRenderingGroupSpecBuilder
```

### <span class="badge object-method"></span> Visibility

```go
func (builder *ConditionalRenderingGroupSpecBuilder) Visibility(visibility dashboardv2beta1.ConditionalRenderingGroupSpecVisibility) *ConditionalRenderingGroupSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ConditionalRenderingGroupSpec](./object-ConditionalRenderingGroupSpec.md)
