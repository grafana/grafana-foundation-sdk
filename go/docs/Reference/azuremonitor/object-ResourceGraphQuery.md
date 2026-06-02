---
title: <span class="badge object-type-struct"></span> ResourceGraphQuery
---
# <span class="badge object-type-struct"></span> ResourceGraphQuery

## Definition

```go
type ResourceGraphQuery struct {
    // Azure Resource Graph KQL query to be executed.
    Query *string `json:"query,omitempty"`
    // Specifies the format results should be returned as. Defaults to table.
    ResultFormat *string `json:"resultFormat,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ResourceGraphQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (resourceGraphQuery *ResourceGraphQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ResourceGraphQuery` objects.

```go
func (resourceGraphQuery *ResourceGraphQuery) Equals(other ResourceGraphQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ResourceGraphQuery` fields for violations and returns them.

```go
func (resourceGraphQuery *ResourceGraphQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [ResourceGraphQueryBuilder](./builder-ResourceGraphQueryBuilder.md)
