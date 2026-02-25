---
title: <span class="badge builder"></span> GridBuilder
---
# <span class="badge builder"></span> GridBuilder

## Constructor

```go
func NewGridBuilder() *GridBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *GridBuilder) Build() (GridLayoutKind, error)
```

### <span class="badge object-method"></span> Item

```go
func (builder *GridBuilder) Item(item cog.Builder[dashboardv2beta1.GridLayoutItemKind]) *GridBuilder
```

### <span class="badge object-method"></span> Items

```go
func (builder *GridBuilder) Items(items []cog.Builder[dashboardv2beta1.GridLayoutItemKind]) *GridBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [GridLayoutKind](./object-GridLayoutKind.md)
