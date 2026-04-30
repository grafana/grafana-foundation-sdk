---
title: <span class="badge object-type-struct"></span> ElementReference
---
# <span class="badge object-type-struct"></span> ElementReference

## Definition

```go
type ElementReference struct {
    Kind string `json:"kind"`
    Name string `json:"name"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElementReference` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elementReference *ElementReference) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElementReference` objects.

```go
func (elementReference *ElementReference) Equals(other ElementReference) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElementReference` fields for violations and returns them.

```go
func (elementReference *ElementReference) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElementReferenceBuilder](./builder-ElementReferenceBuilder.md)
