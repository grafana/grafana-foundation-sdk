---
title: <span class="badge object-type-struct"></span> FrameGeometrySource
---
# <span class="badge object-type-struct"></span> FrameGeometrySource

## Definition

```go
type FrameGeometrySource struct {
    Mode common.FrameGeometrySourceMode `json:"mode"`
    // Field mappings
    Geohash *string `json:"geohash,omitempty"`
    Latitude *string `json:"latitude,omitempty"`
    Longitude *string `json:"longitude,omitempty"`
    Wkt *string `json:"wkt,omitempty"`
    Lookup *string `json:"lookup,omitempty"`
    // Path to Gazetteer
    Gazetteer *string `json:"gazetteer,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FrameGeometrySource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (frameGeometrySource *FrameGeometrySource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `FrameGeometrySource` objects.

```go
func (frameGeometrySource *FrameGeometrySource) Equals(other FrameGeometrySource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FrameGeometrySource` fields for violations and returns them.

```go
func (frameGeometrySource *FrameGeometrySource) Validate() error
```

## See also

 * <span class="badge builder"></span> [FrameGeometrySourceBuilder](./builder-FrameGeometrySourceBuilder.md)
