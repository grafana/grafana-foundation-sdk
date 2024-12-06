---
title: <span class="badge object-type-struct"></span> MapLayerOptions
---
# <span class="badge object-type-struct"></span> MapLayerOptions

## Definition

```go
type MapLayerOptions struct {
    Type string `json:"type"`
    // configured unique display name
    Name string `json:"name"`
    // Custom options depending on the type
    Config any `json:"config,omitempty"`
    // Common method to define geometry fields
    Location *common.FrameGeometrySource `json:"location,omitempty"`
    // Defines a frame MatcherConfig that may filter data for the given layer
    FilterData any `json:"filterData,omitempty"`
    // Common properties:
    // https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
    // Layer opacity (0-1)
    Opacity *int64 `json:"opacity,omitempty"`
    // Check tooltip (defaults to true)
    Tooltip *bool `json:"tooltip,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MapLayerOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (mapLayerOptions *MapLayerOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MapLayerOptions` objects.

```go
func (mapLayerOptions *MapLayerOptions) Equals(other MapLayerOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MapLayerOptions` fields for violations and returns them.

```go
func (mapLayerOptions *MapLayerOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [MapLayerOptionsBuilder](./builder-MapLayerOptionsBuilder.md)
