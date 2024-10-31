---
title: <span class="badge object-type-struct"></span> ExprTypeClassicConditionsConditionsReducer
---
# <span class="badge object-type-struct"></span> ExprTypeClassicConditionsConditionsReducer

## Definition

```go
type ExprTypeClassicConditionsConditionsReducer struct {
    Type string `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsConditionsReducer` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeClassicConditionsConditionsReducer *ExprTypeClassicConditionsConditionsReducer) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeClassicConditionsConditionsReducer` objects.

```go
func (exprTypeClassicConditionsConditionsReducer *ExprTypeClassicConditionsConditionsReducer) Equals(other ExprTypeClassicConditionsConditionsReducer) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsConditionsReducer` fields for violations and returns them.

```go
func (exprTypeClassicConditionsConditionsReducer *ExprTypeClassicConditionsConditionsReducer) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeClassicConditionsConditionsReducerBuilder](./builder-ExprTypeClassicConditionsConditionsReducerBuilder.md)
