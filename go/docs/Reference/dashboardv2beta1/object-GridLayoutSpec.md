---
title: <span class="badge object-type-struct"></span> GridLayoutSpec
---
# <span class="badge object-type-struct"></span> GridLayoutSpec

## Definition

```go
type GridLayoutSpec struct {
    Items []dashboardv2beta1.GridLayoutItemKind `json:"items"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (gridLayoutSpec *GridLayoutSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GridLayoutSpec` objects.

```go
func (gridLayoutSpec *GridLayoutSpec) Equals(other GridLayoutSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GridLayoutSpec` fields for violations and returns them.

```go
func (gridLayoutSpec *GridLayoutSpec) Validate() error
```

