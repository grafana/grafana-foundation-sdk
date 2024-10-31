---
title: <span class="badge object-type-struct"></span> Nested
---
# <span class="badge object-type-struct"></span> Nested

## Definition

```go
type Nested struct {
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Type string `json:"type"`
    Settings any `json:"settings,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Nested` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (nested *Nested) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Nested` objects.

```go
func (nested *Nested) Equals(other Nested) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Nested` fields for violations and returns them.

```go
func (nested *Nested) Validate() error
```

## See also

 * <span class="badge builder"></span> [NestedBuilder](./builder-NestedBuilder.md)
