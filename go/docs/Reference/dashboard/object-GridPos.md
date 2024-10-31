---
title: <span class="badge object-type-struct"></span> GridPos
---
# <span class="badge object-type-struct"></span> GridPos

Position and dimensions of a panel in the grid

## Definition

```go
type GridPos struct {
    // Panel height. The height is the number of rows from the top edge of the panel.
    H uint32 `json:"h"`
    // Panel width. The width is the number of columns from the left edge of the panel.
    W uint32 `json:"w"`
    // Panel x. The x coordinate is the number of columns from the left edge of the grid
    X uint32 `json:"x"`
    // Panel y. The y coordinate is the number of rows from the top edge of the grid
    Y uint32 `json:"y"`
    // Whether the panel is fixed within the grid. If true, the panel will not be affected by other panels' interactions
    Static *bool `json:"static,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridPos` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (gridPos *GridPos) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GridPos` objects.

```go
func (gridPos *GridPos) Equals(other GridPos) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GridPos` fields for violations and returns them.

```go
func (gridPos *GridPos) Validate() error
```

