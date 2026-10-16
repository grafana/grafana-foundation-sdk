---
title: <span class="badge object-type-struct"></span> V2QueryGroupKind
---
# <span class="badge object-type-struct"></span> V2QueryGroupKind

## Definition

```go
type V2QueryGroupKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.V2QueryGroupSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `V2QueryGroupKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (v2QueryGroupKind *V2QueryGroupKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `V2QueryGroupKind` objects.

```go
func (v2QueryGroupKind *V2QueryGroupKind) Equals(other V2QueryGroupKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `V2QueryGroupKind` fields for violations and returns them.

```go
func (v2QueryGroupKind *V2QueryGroupKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [V2QueryGroupKindBuilder](./builder-V2QueryGroupKindBuilder.md)
