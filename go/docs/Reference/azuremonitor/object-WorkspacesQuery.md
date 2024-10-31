---
title: <span class="badge object-type-struct"></span> WorkspacesQuery
---
# <span class="badge object-type-struct"></span> WorkspacesQuery

## Definition

```go
type WorkspacesQuery struct {
    RawQuery *string `json:"rawQuery,omitempty"`
    Kind string `json:"kind"`
    Subscription string `json:"subscription"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `WorkspacesQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (workspacesQuery *WorkspacesQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `WorkspacesQuery` objects.

```go
func (workspacesQuery *WorkspacesQuery) Equals(other WorkspacesQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `WorkspacesQuery` fields for violations and returns them.

```go
func (workspacesQuery *WorkspacesQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [WorkspacesQueryBuilder](./builder-WorkspacesQueryBuilder.md)
