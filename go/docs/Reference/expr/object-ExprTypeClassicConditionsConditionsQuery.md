---
title: <span class="badge object-type-struct"></span> ExprTypeClassicConditionsConditionsQuery
---
# <span class="badge object-type-struct"></span> ExprTypeClassicConditionsConditionsQuery

## Definition

```go
type ExprTypeClassicConditionsConditionsQuery struct {
    Params []string `json:"params"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsConditionsQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeClassicConditionsConditionsQuery *ExprTypeClassicConditionsConditionsQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeClassicConditionsConditionsQuery` objects.

```go
func (exprTypeClassicConditionsConditionsQuery *ExprTypeClassicConditionsConditionsQuery) Equals(other ExprTypeClassicConditionsConditionsQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsConditionsQuery` fields for violations and returns them.

```go
func (exprTypeClassicConditionsConditionsQuery *ExprTypeClassicConditionsConditionsQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeClassicConditionsConditionsQueryBuilder](./builder-ExprTypeClassicConditionsConditionsQueryBuilder.md)
