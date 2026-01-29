---
title: <span class="badge builder"></span> ConditionalRenderingVariableBuilder
---
# <span class="badge builder"></span> ConditionalRenderingVariableBuilder

## Constructor

```go
func NewConditionalRenderingVariableBuilder() *ConditionalRenderingVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ConditionalRenderingVariableBuilder) Build() (ConditionalRenderingVariableKind, error)
```

### <span class="badge object-method"></span> Operator

```go
func (builder *ConditionalRenderingVariableBuilder) Operator(operator dashboardv2beta1.ConditionalRenderingVariableSpecOperator) *ConditionalRenderingVariableBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *ConditionalRenderingVariableBuilder) Spec(spec cog.Builder[dashboardv2beta1.ConditionalRenderingVariableSpec]) *ConditionalRenderingVariableBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *ConditionalRenderingVariableBuilder) Value(value string) *ConditionalRenderingVariableBuilder
```

### <span class="badge object-method"></span> Variable

```go
func (builder *ConditionalRenderingVariableBuilder) Variable(variable string) *ConditionalRenderingVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ConditionalRenderingVariableKind](./object-ConditionalRenderingVariableKind.md)
