---
title: <span class="badge builder"></span> GridItemBuilder
---
# <span class="badge builder"></span> GridItemBuilder

## Constructor

```go
func NewGridItemBuilder() *GridItemBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *GridItemBuilder) Build() (GridLayoutItemKind, error)
```

### <span class="badge object-method"></span> Height

```go
func (builder *GridItemBuilder) Height(height int64) *GridItemBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *GridItemBuilder) Name(name string) *GridItemBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *GridItemBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.RepeatOptions]) *GridItemBuilder
```

### <span class="badge object-method"></span> Width

```go
func (builder *GridItemBuilder) Width(width int64) *GridItemBuilder
```

### <span class="badge object-method"></span> X

```go
func (builder *GridItemBuilder) X(x int64) *GridItemBuilder
```

### <span class="badge object-method"></span> Y

```go
func (builder *GridItemBuilder) Y(y int64) *GridItemBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [GridLayoutItemKind](./object-GridLayoutItemKind.md)
