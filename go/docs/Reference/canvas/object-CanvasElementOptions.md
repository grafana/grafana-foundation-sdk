---
title: <span class="badge object-type-struct"></span> CanvasElementOptions
---
# <span class="badge object-type-struct"></span> CanvasElementOptions

## Definition

```go
type CanvasElementOptions struct {
    Name string `json:"name"`
    Type string `json:"type"`
    // TODO: figure out how to define this (element config(s))
    Config any `json:"config,omitempty"`
    Constraint *canvas.Constraint `json:"constraint,omitempty"`
    Placement *canvas.Placement `json:"placement,omitempty"`
    Background *canvas.BackgroundConfig `json:"background,omitempty"`
    Border *canvas.LineConfig `json:"border,omitempty"`
    Connections []canvas.CanvasConnection `json:"connections,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CanvasElementOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (canvasElementOptions *CanvasElementOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CanvasElementOptions` objects.

```go
func (canvasElementOptions *CanvasElementOptions) Equals(other CanvasElementOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CanvasElementOptions` fields for violations and returns them.

```go
func (canvasElementOptions *CanvasElementOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [CanvasElementOptionsBuilder](./builder-CanvasElementOptionsBuilder.md)
