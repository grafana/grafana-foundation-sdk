---
title: <span class="badge object-type-struct"></span> CanvasConnection
---
# <span class="badge object-type-struct"></span> CanvasConnection

## Definition

```go
type CanvasConnection struct {
    Source canvas.ConnectionCoordinates `json:"source"`
    Target canvas.ConnectionCoordinates `json:"target"`
    TargetName *string `json:"targetName,omitempty"`
    Path canvas.ConnectionPath `json:"path"`
    Color *common.ColorDimensionConfig `json:"color,omitempty"`
    Size *common.ScaleDimensionConfig `json:"size,omitempty"`
    Vertices []canvas.ConnectionCoordinates `json:"vertices,omitempty"`
    SourceOriginal *canvas.ConnectionCoordinates `json:"sourceOriginal,omitempty"`
    TargetOriginal *canvas.ConnectionCoordinates `json:"targetOriginal,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CanvasConnection` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (canvasConnection *CanvasConnection) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CanvasConnection` objects.

```go
func (canvasConnection *CanvasConnection) Equals(other CanvasConnection) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CanvasConnection` fields for violations and returns them.

```go
func (canvasConnection *CanvasConnection) Validate() error
```

## See also

 * <span class="badge builder"></span> [CanvasConnectionBuilder](./builder-CanvasConnectionBuilder.md)
