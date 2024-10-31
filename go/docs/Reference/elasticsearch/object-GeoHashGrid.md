---
title: <span class="badge object-type-struct"></span> GeoHashGrid
---
# <span class="badge object-type-struct"></span> GeoHashGrid

## Definition

```go
type GeoHashGrid struct {
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Type string `json:"type"`
    Settings *elasticsearch.ElasticsearchGeoHashGridSettings `json:"settings,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GeoHashGrid` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (geoHashGrid *GeoHashGrid) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GeoHashGrid` objects.

```go
func (geoHashGrid *GeoHashGrid) Equals(other GeoHashGrid) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GeoHashGrid` fields for violations and returns them.

```go
func (geoHashGrid *GeoHashGrid) Validate() error
```

## See also

 * <span class="badge builder"></span> [GeoHashGridBuilder](./builder-GeoHashGridBuilder.md)
