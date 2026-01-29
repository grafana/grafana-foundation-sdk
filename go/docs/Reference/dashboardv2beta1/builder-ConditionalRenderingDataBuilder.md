---
title: <span class="badge builder"></span> ConditionalRenderingDataBuilder
---
# <span class="badge builder"></span> ConditionalRenderingDataBuilder

## Constructor

```go
func NewConditionalRenderingDataBuilder() *ConditionalRenderingDataBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ConditionalRenderingDataBuilder) Build() (ConditionalRenderingDataKind, error)
```

### <span class="badge object-method"></span> Spec

```go
func (builder *ConditionalRenderingDataBuilder) Spec(spec cog.Builder[dashboardv2beta1.ConditionalRenderingDataSpec]) *ConditionalRenderingDataBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *ConditionalRenderingDataBuilder) Value(value bool) *ConditionalRenderingDataBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ConditionalRenderingDataKind](./object-ConditionalRenderingDataKind.md)
