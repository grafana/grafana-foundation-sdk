---
title: <span class="badge builder"></span> ConditionalRenderingVariableSpecBuilder
---
# <span class="badge builder"></span> ConditionalRenderingVariableSpecBuilder

## Constructor

```go
func NewConditionalRenderingVariableSpecBuilder() *ConditionalRenderingVariableSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ConditionalRenderingVariableSpecBuilder) Build() (ConditionalRenderingVariableSpec, error)
```

### <span class="badge object-method"></span> Operator

```go
func (builder *ConditionalRenderingVariableSpecBuilder) Operator(operator dashboardv2beta1.ConditionalRenderingVariableSpecOperator) *ConditionalRenderingVariableSpecBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *ConditionalRenderingVariableSpecBuilder) Value(value string) *ConditionalRenderingVariableSpecBuilder
```

### <span class="badge object-method"></span> Variable

```go
func (builder *ConditionalRenderingVariableSpecBuilder) Variable(variable string) *ConditionalRenderingVariableSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ConditionalRenderingVariableSpec](./object-ConditionalRenderingVariableSpec.md)
