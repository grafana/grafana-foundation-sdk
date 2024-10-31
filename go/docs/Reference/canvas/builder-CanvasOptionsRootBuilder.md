---
title: <span class="badge builder"></span> CanvasOptionsRootBuilder
---
# <span class="badge builder"></span> CanvasOptionsRootBuilder

## Constructor

```go
func NewCanvasOptionsRootBuilder() *CanvasOptionsRootBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CanvasOptionsRootBuilder) Build() (CanvasOptionsRoot, error)
```

### <span class="badge object-method"></span> Elements

The list of canvas elements attached to the root element

```go
func (builder *CanvasOptionsRootBuilder) Elements(elements []cog.Builder[canvas.CanvasElementOptions]) *CanvasOptionsRootBuilder
```

### <span class="badge object-method"></span> Name

Name of the root element

```go
func (builder *CanvasOptionsRootBuilder) Name(name string) *CanvasOptionsRootBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CanvasOptionsRoot](./object-CanvasOptionsRoot.md)
