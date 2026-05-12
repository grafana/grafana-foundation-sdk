---
title: <span class="badge object-type-struct"></span> TransformationKind
---
# <span class="badge object-type-struct"></span> TransformationKind

## Definition

```go
type TransformationKind struct {
    Kind string `json:"kind"`
    // The group is the transformation ID
    Group string `json:"group"`
    Spec dashboardv2.TransformationSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TransformationKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (transformationKind *TransformationKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TransformationKind` objects.

```go
func (transformationKind *TransformationKind) Equals(other TransformationKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TransformationKind` fields for violations and returns them.

```go
func (transformationKind *TransformationKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [TransformationBuilder](./builder-TransformationBuilder.md)
