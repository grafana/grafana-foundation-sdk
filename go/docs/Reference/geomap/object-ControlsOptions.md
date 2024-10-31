---
title: <span class="badge object-type-struct"></span> ControlsOptions
---
# <span class="badge object-type-struct"></span> ControlsOptions

## Definition

```go
type ControlsOptions struct {
    // Zoom (upper left)
    ShowZoom *bool `json:"showZoom,omitempty"`
    // let the mouse wheel zoom
    MouseWheelZoom *bool `json:"mouseWheelZoom,omitempty"`
    // Lower right
    ShowAttribution *bool `json:"showAttribution,omitempty"`
    // Scale options
    ShowScale *bool `json:"showScale,omitempty"`
    // Show debug
    ShowDebug *bool `json:"showDebug,omitempty"`
    // Show measure
    ShowMeasure *bool `json:"showMeasure,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ControlsOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (controlsOptions *ControlsOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ControlsOptions` objects.

```go
func (controlsOptions *ControlsOptions) Equals(other ControlsOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ControlsOptions` fields for violations and returns them.

```go
func (controlsOptions *ControlsOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [ControlsOptionsBuilder](./builder-ControlsOptionsBuilder.md)
