---
title: <span class="badge object-type-struct"></span> ResourceRef
---
# <span class="badge object-type-struct"></span> ResourceRef

## Definition

```go
type ResourceRef struct {
    Kind string `json:"kind"`
    Name string `json:"name"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ResourceRef` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (resourceRef *ResourceRef) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ResourceRef` objects.

```go
func (resourceRef *ResourceRef) Equals(other ResourceRef) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ResourceRef` fields for violations and returns them.

```go
func (resourceRef *ResourceRef) Validate() error
```

## See also

 * <span class="badge builder"></span> [ResourceRefBuilder](./builder-ResourceRefBuilder.md)
