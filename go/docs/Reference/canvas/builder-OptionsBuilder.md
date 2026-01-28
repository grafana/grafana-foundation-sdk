---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```go
func NewOptionsBuilder() *OptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *OptionsBuilder) Build() (Options, error)
```

### <span class="badge object-method"></span> InfinitePan

Enable infinite pan

```go
func (builder *OptionsBuilder) InfinitePan(infinitePan bool) *OptionsBuilder
```

### <span class="badge object-method"></span> InlineEditing

Enable inline editing

```go
func (builder *OptionsBuilder) InlineEditing(inlineEditing bool) *OptionsBuilder
```

### <span class="badge object-method"></span> PanZoom

Enable pan and zoom

```go
func (builder *OptionsBuilder) PanZoom(panZoom bool) *OptionsBuilder
```

### <span class="badge object-method"></span> Root

The root element of canvas (frame), where all canvas elements are nested

TODO: Figure out how to define a default value for this

```go
func (builder *OptionsBuilder) Root(root cog.Builder[canvas.CanvasOptionsRoot]) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowAdvancedTypes

Show all available element types

```go
func (builder *OptionsBuilder) ShowAdvancedTypes(showAdvancedTypes bool) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
