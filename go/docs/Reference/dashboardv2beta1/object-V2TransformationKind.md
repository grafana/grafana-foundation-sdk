---
title: <span class="badge object-type-struct"></span> V2TransformationKind
---
# <span class="badge object-type-struct"></span> V2TransformationKind

Dashboard v2 shape: the transformation ID moved from `kind` to `group`.

## Definition

```go
type V2TransformationKind struct {
    Kind string `json:"kind"`
    // The group is the transformation ID
    Group string `json:"group"`
    Spec dashboardv2beta1.V2TransformationSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `V2TransformationKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (v2TransformationKind *V2TransformationKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `V2TransformationKind` objects.

```go
func (v2TransformationKind *V2TransformationKind) Equals(other V2TransformationKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `V2TransformationKind` fields for violations and returns them.

```go
func (v2TransformationKind *V2TransformationKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [V2TransformationKindBuilder](./builder-V2TransformationKindBuilder.md)
