---
title: <span class="badge object-type-struct"></span> UnknownQuery
---
# <span class="badge object-type-struct"></span> UnknownQuery

## Definition

```go
type UnknownQuery struct {
    RawQuery *string `json:"rawQuery,omitempty"`
    Kind string `json:"kind"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `UnknownQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (unknownQuery *UnknownQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `UnknownQuery` objects.

```go
func (unknownQuery *UnknownQuery) Equals(other UnknownQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `UnknownQuery` fields for violations and returns them.

```go
func (unknownQuery *UnknownQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [UnknownQueryBuilder](./builder-UnknownQueryBuilder.md)
