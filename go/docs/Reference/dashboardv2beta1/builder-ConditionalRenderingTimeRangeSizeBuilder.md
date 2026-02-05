---
title: <span class="badge builder"></span> ConditionalRenderingTimeRangeSizeBuilder
---
# <span class="badge builder"></span> ConditionalRenderingTimeRangeSizeBuilder

## Constructor

```go
func NewConditionalRenderingTimeRangeSizeBuilder() *ConditionalRenderingTimeRangeSizeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ConditionalRenderingTimeRangeSizeBuilder) Build() (ConditionalRenderingTimeRangeSizeKind, error)
```

### <span class="badge object-method"></span> Spec

```go
func (builder *ConditionalRenderingTimeRangeSizeBuilder) Spec(spec cog.Builder[dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec]) *ConditionalRenderingTimeRangeSizeBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *ConditionalRenderingTimeRangeSizeBuilder) Value(value string) *ConditionalRenderingTimeRangeSizeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ConditionalRenderingTimeRangeSizeKind](./object-ConditionalRenderingTimeRangeSizeKind.md)
