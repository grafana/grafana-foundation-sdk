---
title: <span class="badge object-type-struct"></span> USAQuery
---
# <span class="badge object-type-struct"></span> USAQuery

## Definition

```go
type USAQuery struct {
    Fields []string `json:"fields,omitempty"`
    Mode *string `json:"mode,omitempty"`
    Period *string `json:"period,omitempty"`
    States []string `json:"states,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `USAQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (uSAQuery *USAQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `USAQuery` objects.

```go
func (uSAQuery *USAQuery) Equals(other USAQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `USAQuery` fields for violations and returns them.

```go
func (uSAQuery *USAQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [USAQueryBuilder](./builder-USAQueryBuilder.md)
