---
title: <span class="badge object-type-struct"></span> Resource
---
# <span class="badge object-type-struct"></span> Resource

## Definition

```go
type Resource struct {
    Group string `json:"group"`
    Kind string `json:"kind"`
    // The set of resources
    // +listType=set
    Names []string `json:"names"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Resource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (resource *Resource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Resource` objects.

```go
func (resource *Resource) Equals(other Resource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Resource` fields for violations and returns them.

```go
func (resource *Resource) Validate() error
```

## See also

 * <span class="badge builder"></span> [ResourceBuilder](./builder-ResourceBuilder.md)
