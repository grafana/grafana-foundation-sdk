---
title: <span class="badge object-type-struct"></span> NodesQuery
---
# <span class="badge object-type-struct"></span> NodesQuery

## Definition

```go
type NodesQuery struct {
    Type *testdata.NodesQueryType `json:"type,omitempty"`
    Count *int64 `json:"count,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NodesQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (nodesQuery *NodesQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NodesQuery` objects.

```go
func (nodesQuery *NodesQuery) Equals(other NodesQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NodesQuery` fields for violations and returns them.

```go
func (nodesQuery *NodesQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [NodesQueryBuilder](./builder-NodesQueryBuilder.md)
