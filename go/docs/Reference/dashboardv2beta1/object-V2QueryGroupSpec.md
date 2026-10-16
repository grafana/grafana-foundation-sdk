---
title: <span class="badge object-type-struct"></span> V2QueryGroupSpec
---
# <span class="badge object-type-struct"></span> V2QueryGroupSpec

## Definition

```go
type V2QueryGroupSpec struct {
    Queries []dashboardv2beta1.PanelQueryKind `json:"queries"`
    Transformations []dashboardv2beta1.V2TransformationKind `json:"transformations"`
    QueryOptions dashboardv2beta1.QueryOptionsSpec `json:"queryOptions"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `V2QueryGroupSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (v2QueryGroupSpec *V2QueryGroupSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `V2QueryGroupSpec` objects.

```go
func (v2QueryGroupSpec *V2QueryGroupSpec) Equals(other V2QueryGroupSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `V2QueryGroupSpec` fields for violations and returns them.

```go
func (v2QueryGroupSpec *V2QueryGroupSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [V2QueryGroupSpecBuilder](./builder-V2QueryGroupSpecBuilder.md)
