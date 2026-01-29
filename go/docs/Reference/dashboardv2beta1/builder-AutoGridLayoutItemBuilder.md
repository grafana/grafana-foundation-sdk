---
title: <span class="badge builder"></span> AutoGridLayoutItemBuilder
---
# <span class="badge builder"></span> AutoGridLayoutItemBuilder

## Constructor

```go
func NewAutoGridLayoutItemBuilder(name string) *AutoGridLayoutItemBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AutoGridLayoutItemBuilder) Build() (AutoGridLayoutItemKind, error)
```

### <span class="badge object-method"></span> ConditionalRendering

```go
func (builder *AutoGridLayoutItemBuilder) ConditionalRendering(conditionalRendering cog.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) *AutoGridLayoutItemBuilder
```

### <span class="badge object-method"></span> Element

```go
func (builder *AutoGridLayoutItemBuilder) Element(element cog.Builder[dashboardv2beta1.ElementReference]) *AutoGridLayoutItemBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *AutoGridLayoutItemBuilder) Name(name string) *AutoGridLayoutItemBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *AutoGridLayoutItemBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.AutoGridRepeatOptions]) *AutoGridLayoutItemBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AutoGridLayoutItemKind](./object-AutoGridLayoutItemKind.md)
