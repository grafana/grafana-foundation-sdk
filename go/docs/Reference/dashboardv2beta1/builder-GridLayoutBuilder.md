---
title: <span class="badge builder"></span> GridLayoutBuilder
---
# <span class="badge builder"></span> GridLayoutBuilder

## Constructor

```go
func NewGridLayoutBuilder() *GridLayoutBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *GridLayoutBuilder) Build() (GridLayoutKind, error)
```

### <span class="badge object-method"></span> Item

```go
func (builder *GridLayoutBuilder) Item(item cog.Builder[dashboardv2beta1.GridLayoutItemKind]) *GridLayoutBuilder
```

### <span class="badge object-method"></span> Items

```go
func (builder *GridLayoutBuilder) Items(items []cog.Builder[dashboardv2beta1.GridLayoutItemKind]) *GridLayoutBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [GridLayoutKind](./object-GridLayoutKind.md)
