---
title: <span class="badge object-type-struct"></span> MapViewConfig
---
# <span class="badge object-type-struct"></span> MapViewConfig

## Definition

```go
type MapViewConfig struct {
    Id string `json:"id"`
    Lat *int64 `json:"lat,omitempty"`
    Lon *int64 `json:"lon,omitempty"`
    Zoom *int64 `json:"zoom,omitempty"`
    MinZoom *int64 `json:"minZoom,omitempty"`
    MaxZoom *int64 `json:"maxZoom,omitempty"`
    Padding *int64 `json:"padding,omitempty"`
    AllLayers *bool `json:"allLayers,omitempty"`
    LastOnly *bool `json:"lastOnly,omitempty"`
    Layer *string `json:"layer,omitempty"`
    Shared *bool `json:"shared,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MapViewConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (mapViewConfig *MapViewConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MapViewConfig` objects.

```go
func (mapViewConfig *MapViewConfig) Equals(other MapViewConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MapViewConfig` fields for violations and returns them.

```go
func (mapViewConfig *MapViewConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [MapViewConfigBuilder](./builder-MapViewConfigBuilder.md)
