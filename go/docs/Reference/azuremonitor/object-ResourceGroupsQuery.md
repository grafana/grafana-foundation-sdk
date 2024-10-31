---
title: <span class="badge object-type-struct"></span> ResourceGroupsQuery
---
# <span class="badge object-type-struct"></span> ResourceGroupsQuery

## Definition

```go
type ResourceGroupsQuery struct {
    RawQuery *string `json:"rawQuery,omitempty"`
    Kind string `json:"kind"`
    Subscription string `json:"subscription"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ResourceGroupsQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (resourceGroupsQuery *ResourceGroupsQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ResourceGroupsQuery` objects.

```go
func (resourceGroupsQuery *ResourceGroupsQuery) Equals(other ResourceGroupsQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ResourceGroupsQuery` fields for violations and returns them.

```go
func (resourceGroupsQuery *ResourceGroupsQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [ResourceGroupsQueryBuilder](./builder-ResourceGroupsQueryBuilder.md)
