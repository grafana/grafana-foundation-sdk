---
title: <span class="badge object-type-struct"></span> CanvasOptionsRoot
---
# <span class="badge object-type-struct"></span> CanvasOptionsRoot

## Definition

```go
type CanvasOptionsRoot struct {
    // Name of the root element
    Name string `json:"name"`
    // Type of root element (frame)
    Type string `json:"type"`
    // The list of canvas elements attached to the root element
    Elements []canvas.CanvasElementOptions `json:"elements"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CanvasOptionsRoot` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (canvasOptionsRoot *CanvasOptionsRoot) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CanvasOptionsRoot` objects.

```go
func (canvasOptionsRoot *CanvasOptionsRoot) Equals(other CanvasOptionsRoot) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CanvasOptionsRoot` fields for violations and returns them.

```go
func (canvasOptionsRoot *CanvasOptionsRoot) Validate() error
```

## See also

 * <span class="badge builder"></span> [CanvasOptionsRootBuilder](./builder-CanvasOptionsRootBuilder.md)
