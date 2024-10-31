---
title: <span class="badge builder"></span> MapViewConfigBuilder
---
# <span class="badge builder"></span> MapViewConfigBuilder

## Constructor

```go
func NewMapViewConfigBuilder() *MapViewConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MapViewConfigBuilder) Build() (MapViewConfig, error)
```

### <span class="badge object-method"></span> AllLayers

```go
func (builder *MapViewConfigBuilder) AllLayers(allLayers bool) *MapViewConfigBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *MapViewConfigBuilder) Id(id string) *MapViewConfigBuilder
```

### <span class="badge object-method"></span> LastOnly

```go
func (builder *MapViewConfigBuilder) LastOnly(lastOnly bool) *MapViewConfigBuilder
```

### <span class="badge object-method"></span> Lat

```go
func (builder *MapViewConfigBuilder) Lat(lat int64) *MapViewConfigBuilder
```

### <span class="badge object-method"></span> Layer

```go
func (builder *MapViewConfigBuilder) Layer(layer string) *MapViewConfigBuilder
```

### <span class="badge object-method"></span> Lon

```go
func (builder *MapViewConfigBuilder) Lon(lon int64) *MapViewConfigBuilder
```

### <span class="badge object-method"></span> MaxZoom

```go
func (builder *MapViewConfigBuilder) MaxZoom(maxZoom int64) *MapViewConfigBuilder
```

### <span class="badge object-method"></span> MinZoom

```go
func (builder *MapViewConfigBuilder) MinZoom(minZoom int64) *MapViewConfigBuilder
```

### <span class="badge object-method"></span> Padding

```go
func (builder *MapViewConfigBuilder) Padding(padding int64) *MapViewConfigBuilder
```

### <span class="badge object-method"></span> Shared

```go
func (builder *MapViewConfigBuilder) Shared(shared bool) *MapViewConfigBuilder
```

### <span class="badge object-method"></span> Zoom

```go
func (builder *MapViewConfigBuilder) Zoom(zoom int64) *MapViewConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MapViewConfig](./object-MapViewConfig.md)
