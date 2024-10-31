---
title: <span class="badge builder"></span> MapLayerOptionsBuilder
---
# <span class="badge builder"></span> MapLayerOptionsBuilder

## Constructor

```go
func NewMapLayerOptionsBuilder() *MapLayerOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MapLayerOptionsBuilder) Build() (MapLayerOptions, error)
```

### <span class="badge object-method"></span> Config

Custom options depending on the type

```go
func (builder *MapLayerOptionsBuilder) Config(config any) *MapLayerOptionsBuilder
```

### <span class="badge object-method"></span> FilterData

Defines a frame MatcherConfig that may filter data for the given layer

```go
func (builder *MapLayerOptionsBuilder) FilterData(filterData any) *MapLayerOptionsBuilder
```

### <span class="badge object-method"></span> Location

Common method to define geometry fields

```go
func (builder *MapLayerOptionsBuilder) Location(location cog.Builder[common.FrameGeometrySource]) *MapLayerOptionsBuilder
```

### <span class="badge object-method"></span> Name

configured unique display name

```go
func (builder *MapLayerOptionsBuilder) Name(name string) *MapLayerOptionsBuilder
```

### <span class="badge object-method"></span> Opacity

Common properties:

https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html

Layer opacity (0-1)

```go
func (builder *MapLayerOptionsBuilder) Opacity(opacity int64) *MapLayerOptionsBuilder
```

### <span class="badge object-method"></span> Tooltip

Check tooltip (defaults to true)

```go
func (builder *MapLayerOptionsBuilder) Tooltip(tooltip bool) *MapLayerOptionsBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *MapLayerOptionsBuilder) Type(typeArg string) *MapLayerOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MapLayerOptions](./object-MapLayerOptions.md)
