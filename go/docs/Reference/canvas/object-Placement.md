---
title: <span class="badge object-type-struct"></span> Placement
---
# <span class="badge object-type-struct"></span> Placement

## Definition

```go
type Placement struct {
    Top *float64 `json:"top,omitempty"`
    Left *float64 `json:"left,omitempty"`
    Right *float64 `json:"right,omitempty"`
    Bottom *float64 `json:"bottom,omitempty"`
    Width *float64 `json:"width,omitempty"`
    Height *float64 `json:"height,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Placement` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (placement *Placement) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Placement` objects.

```go
func (placement *Placement) Equals(other Placement) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Placement` fields for violations and returns them.

```go
func (placement *Placement) Validate() error
```

## See also

 * <span class="badge builder"></span> [PlacementBuilder](./builder-PlacementBuilder.md)
