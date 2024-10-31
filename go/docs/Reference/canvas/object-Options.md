---
title: <span class="badge object-type-struct"></span> Options
---
# <span class="badge object-type-struct"></span> Options

## Definition

```go
type Options struct {
    // Enable inline editing
    InlineEditing bool `json:"inlineEditing"`
    // Show all available element types
    ShowAdvancedTypes bool `json:"showAdvancedTypes"`
    // Enable pan and zoom
    PanZoom bool `json:"panZoom"`
    // Enable infinite pan
    InfinitePan bool `json:"infinitePan"`
    // The root element of canvas (frame), where all canvas elements are nested
    // TODO: Figure out how to define a default value for this
    Root canvas.CanvasOptionsRoot `json:"root"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (options *Options) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Options` objects.

```go
func (options *Options) Equals(other Options) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.

```go
func (options *Options) Validate() error
```

