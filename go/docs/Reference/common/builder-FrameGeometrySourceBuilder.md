---
title: <span class="badge builder"></span> FrameGeometrySourceBuilder
---
# <span class="badge builder"></span> FrameGeometrySourceBuilder

## Constructor

```go
func NewFrameGeometrySourceBuilder() *FrameGeometrySourceBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FrameGeometrySourceBuilder) Build() (FrameGeometrySource, error)
```

### <span class="badge object-method"></span> Gazetteer

Path to Gazetteer

```go
func (builder *FrameGeometrySourceBuilder) Gazetteer(gazetteer string) *FrameGeometrySourceBuilder
```

### <span class="badge object-method"></span> Geohash

Field mappings

```go
func (builder *FrameGeometrySourceBuilder) Geohash(geohash string) *FrameGeometrySourceBuilder
```

### <span class="badge object-method"></span> Latitude

```go
func (builder *FrameGeometrySourceBuilder) Latitude(latitude string) *FrameGeometrySourceBuilder
```

### <span class="badge object-method"></span> Longitude

```go
func (builder *FrameGeometrySourceBuilder) Longitude(longitude string) *FrameGeometrySourceBuilder
```

### <span class="badge object-method"></span> Lookup

```go
func (builder *FrameGeometrySourceBuilder) Lookup(lookup string) *FrameGeometrySourceBuilder
```

### <span class="badge object-method"></span> Mode

```go
func (builder *FrameGeometrySourceBuilder) Mode(mode common.FrameGeometrySourceMode) *FrameGeometrySourceBuilder
```

### <span class="badge object-method"></span> Wkt

```go
func (builder *FrameGeometrySourceBuilder) Wkt(wkt string) *FrameGeometrySourceBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FrameGeometrySource](./object-FrameGeometrySource.md)
