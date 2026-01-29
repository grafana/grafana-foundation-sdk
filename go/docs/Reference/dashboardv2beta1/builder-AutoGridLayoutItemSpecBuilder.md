---
title: <span class="badge builder"></span> AutoGridLayoutItemSpecBuilder
---
# <span class="badge builder"></span> AutoGridLayoutItemSpecBuilder

## Constructor

```go
func NewAutoGridLayoutItemSpecBuilder() *AutoGridLayoutItemSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AutoGridLayoutItemSpecBuilder) Build() (AutoGridLayoutItemSpec, error)
```

### <span class="badge object-method"></span> ConditionalRendering

```go
func (builder *AutoGridLayoutItemSpecBuilder) ConditionalRendering(conditionalRendering cog.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) *AutoGridLayoutItemSpecBuilder
```

### <span class="badge object-method"></span> Element

```go
func (builder *AutoGridLayoutItemSpecBuilder) Element(element cog.Builder[dashboardv2beta1.ElementReference]) *AutoGridLayoutItemSpecBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *AutoGridLayoutItemSpecBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.AutoGridRepeatOptions]) *AutoGridLayoutItemSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AutoGridLayoutItemSpec](./object-AutoGridLayoutItemSpec.md)
