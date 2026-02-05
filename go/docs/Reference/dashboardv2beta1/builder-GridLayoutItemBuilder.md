---
title: <span class="badge builder"></span> GridLayoutItemBuilder
---
# <span class="badge builder"></span> GridLayoutItemBuilder

## Constructor

```go
func NewGridLayoutItemBuilder(name string) *GridLayoutItemBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *GridLayoutItemBuilder) Build() (GridLayoutItemKind, error)
```

### <span class="badge object-method"></span> Element

```go
func (builder *GridLayoutItemBuilder) Element(name string) *GridLayoutItemBuilder
```

### <span class="badge object-method"></span> Height

```go
func (builder *GridLayoutItemBuilder) Height(height int64) *GridLayoutItemBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *GridLayoutItemBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.RepeatOptions]) *GridLayoutItemBuilder
```

### <span class="badge object-method"></span> Width

```go
func (builder *GridLayoutItemBuilder) Width(width int64) *GridLayoutItemBuilder
```

### <span class="badge object-method"></span> X

```go
func (builder *GridLayoutItemBuilder) X(x int64) *GridLayoutItemBuilder
```

### <span class="badge object-method"></span> Y

```go
func (builder *GridLayoutItemBuilder) Y(y int64) *GridLayoutItemBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [GridLayoutItemKind](./object-GridLayoutItemKind.md)
