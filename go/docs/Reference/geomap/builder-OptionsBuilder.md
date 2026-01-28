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

### <span class="badge object-method"></span> Basemap

```go
func (builder *OptionsBuilder) Basemap(basemap cog.Builder[common.MapLayerOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Controls

```go
func (builder *OptionsBuilder) Controls(controls cog.Builder[geomap.ControlsOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Layers

```go
func (builder *OptionsBuilder) Layers(layers []cog.Builder[common.MapLayerOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Tooltip

```go
func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[geomap.TooltipOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> View

```go
func (builder *OptionsBuilder) View(view cog.Builder[geomap.MapViewConfig]) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
