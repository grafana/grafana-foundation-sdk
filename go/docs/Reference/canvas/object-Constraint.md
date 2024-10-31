---
title: <span class="badge object-type-struct"></span> Constraint
---
# <span class="badge object-type-struct"></span> Constraint

## Definition

```go
type Constraint struct {
    Horizontal *canvas.HorizontalConstraint `json:"horizontal,omitempty"`
    Vertical *canvas.VerticalConstraint `json:"vertical,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Constraint` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (constraint *Constraint) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Constraint` objects.

```go
func (constraint *Constraint) Equals(other Constraint) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Constraint` fields for violations and returns them.

```go
func (constraint *Constraint) Validate() error
```

## See also

 * <span class="badge builder"></span> [ConstraintBuilder](./builder-ConstraintBuilder.md)
