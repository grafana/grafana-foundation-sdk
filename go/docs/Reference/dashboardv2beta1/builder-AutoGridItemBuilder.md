---
title: <span class="badge builder"></span> AutoGridItemBuilder
---
# <span class="badge builder"></span> AutoGridItemBuilder

## Constructor

```go
func NewAutoGridItemBuilder() *AutoGridItemBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AutoGridItemBuilder) Build() (AutoGridLayoutItemKind, error)
```

### <span class="badge object-method"></span> ConditionalRendering

```go
func (builder *AutoGridItemBuilder) ConditionalRendering(conditionalRendering cog.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) *AutoGridItemBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *AutoGridItemBuilder) Name(name string) *AutoGridItemBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *AutoGridItemBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.AutoGridRepeatOptions]) *AutoGridItemBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AutoGridLayoutItemKind](./object-AutoGridLayoutItemKind.md)
