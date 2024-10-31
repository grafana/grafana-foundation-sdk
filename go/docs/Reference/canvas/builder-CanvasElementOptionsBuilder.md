---
title: <span class="badge builder"></span> CanvasElementOptionsBuilder
---
# <span class="badge builder"></span> CanvasElementOptionsBuilder

## Constructor

```go
func NewCanvasElementOptionsBuilder() *CanvasElementOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CanvasElementOptionsBuilder) Build() (CanvasElementOptions, error)
```

### <span class="badge object-method"></span> Background

```go
func (builder *CanvasElementOptionsBuilder) Background(background cog.Builder[canvas.BackgroundConfig]) *CanvasElementOptionsBuilder
```

### <span class="badge object-method"></span> Border

```go
func (builder *CanvasElementOptionsBuilder) Border(border cog.Builder[canvas.LineConfig]) *CanvasElementOptionsBuilder
```

### <span class="badge object-method"></span> Config

TODO: figure out how to define this (element config(s))

```go
func (builder *CanvasElementOptionsBuilder) Config(config any) *CanvasElementOptionsBuilder
```

### <span class="badge object-method"></span> Connections

```go
func (builder *CanvasElementOptionsBuilder) Connections(connections []cog.Builder[canvas.CanvasConnection]) *CanvasElementOptionsBuilder
```

### <span class="badge object-method"></span> Constraint

```go
func (builder *CanvasElementOptionsBuilder) Constraint(constraint cog.Builder[canvas.Constraint]) *CanvasElementOptionsBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *CanvasElementOptionsBuilder) Name(name string) *CanvasElementOptionsBuilder
```

### <span class="badge object-method"></span> Placement

```go
func (builder *CanvasElementOptionsBuilder) Placement(placement cog.Builder[canvas.Placement]) *CanvasElementOptionsBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *CanvasElementOptionsBuilder) Type(typeArg string) *CanvasElementOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CanvasElementOptions](./object-CanvasElementOptions.md)
