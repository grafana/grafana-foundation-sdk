---
title: <span class="badge object-type-struct"></span> QueryVariableKind
---
# <span class="badge object-type-struct"></span> QueryVariableKind

Query variable kind

## Definition

```go
type QueryVariableKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.QueryVariableSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryVariableKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryVariableKind *QueryVariableKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryVariableKind` objects.

```go
func (queryVariableKind *QueryVariableKind) Equals(other QueryVariableKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryVariableKind` fields for violations and returns them.

```go
func (queryVariableKind *QueryVariableKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryVariableBuilder](./builder-QueryVariableBuilder.md)
