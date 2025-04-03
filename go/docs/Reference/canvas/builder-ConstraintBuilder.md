---
title: <span class="badge builder"></span> ConstraintBuilder
---
# <span class="badge builder"></span> ConstraintBuilder

## Constructor

```go
func NewConstraintBuilder() *ConstraintBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ConstraintBuilder) Build() (Constraint, error)
```

### <span class="badge object-method"></span> Horizontal

```go
func (builder *ConstraintBuilder) Horizontal(horizontal canvas.HorizontalConstraint) *ConstraintBuilder
```

### <span class="badge object-method"></span> Vertical

```go
func (builder *ConstraintBuilder) Vertical(vertical canvas.VerticalConstraint) *ConstraintBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Constraint](./object-Constraint.md)
