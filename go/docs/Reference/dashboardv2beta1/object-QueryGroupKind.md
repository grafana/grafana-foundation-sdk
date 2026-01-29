---
title: <span class="badge object-type-struct"></span> QueryGroupKind
---
# <span class="badge object-type-struct"></span> QueryGroupKind

## Definition

```go
type QueryGroupKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.QueryGroupSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryGroupKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryGroupKind *QueryGroupKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryGroupKind` objects.

```go
func (queryGroupKind *QueryGroupKind) Equals(other QueryGroupKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryGroupKind` fields for violations and returns them.

```go
func (queryGroupKind *QueryGroupKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryGroupBuilder](./builder-QueryGroupBuilder.md)
