---
title: <span class="badge builder"></span> GeoHashGridBuilder
---
# <span class="badge builder"></span> GeoHashGridBuilder

## Constructor

```go
func NewGeoHashGridBuilder() *GeoHashGridBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *GeoHashGridBuilder) Build() (GeoHashGrid, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *GeoHashGridBuilder) Field(field string) *GeoHashGridBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *GeoHashGridBuilder) Id(id string) *GeoHashGridBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *GeoHashGridBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchGeoHashGridSettings]) *GeoHashGridBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [GeoHashGrid](./object-GeoHashGrid.md)
