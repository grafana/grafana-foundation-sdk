---
title: <span class="badge builder"></span> CanvasConnectionBuilder
---
# <span class="badge builder"></span> CanvasConnectionBuilder

## Constructor

```go
func NewCanvasConnectionBuilder() *CanvasConnectionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CanvasConnectionBuilder) Build() (CanvasConnection, error)
```

### <span class="badge object-method"></span> Color

```go
func (builder *CanvasConnectionBuilder) Color(color cog.Builder[common.ColorDimensionConfig]) *CanvasConnectionBuilder
```

### <span class="badge object-method"></span> Path

```go
func (builder *CanvasConnectionBuilder) Path(path canvas.ConnectionPath) *CanvasConnectionBuilder
```

### <span class="badge object-method"></span> Size

```go
func (builder *CanvasConnectionBuilder) Size(size cog.Builder[common.ScaleDimensionConfig]) *CanvasConnectionBuilder
```

### <span class="badge object-method"></span> Source

```go
func (builder *CanvasConnectionBuilder) Source(source cog.Builder[canvas.ConnectionCoordinates]) *CanvasConnectionBuilder
```

### <span class="badge object-method"></span> Target

```go
func (builder *CanvasConnectionBuilder) Target(target cog.Builder[canvas.ConnectionCoordinates]) *CanvasConnectionBuilder
```

### <span class="badge object-method"></span> TargetName

```go
func (builder *CanvasConnectionBuilder) TargetName(targetName string) *CanvasConnectionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CanvasConnection](./object-CanvasConnection.md)
