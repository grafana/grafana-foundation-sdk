---
title: <span class="badge object-type-struct"></span> PromQLQuery
---
# <span class="badge object-type-struct"></span> PromQLQuery

PromQL sub-query properties.

## Definition

```go
type PromQLQuery struct {
    // GCP project to execute the query against.
    ProjectName string `json:"projectName"`
    // PromQL expression/query to be executed.
    Expr string `json:"expr"`
    // PromQL min step
    Step string `json:"step"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PromQLQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (promQLQuery *PromQLQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PromQLQuery` objects.

```go
func (promQLQuery *PromQLQuery) Equals(other PromQLQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PromQLQuery` fields for violations and returns them.

```go
func (promQLQuery *PromQLQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [PromQLQueryBuilder](./builder-PromQLQueryBuilder.md)
