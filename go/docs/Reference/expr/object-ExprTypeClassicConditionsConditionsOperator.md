---
title: <span class="badge object-type-struct"></span> ExprTypeClassicConditionsConditionsOperator
---
# <span class="badge object-type-struct"></span> ExprTypeClassicConditionsConditionsOperator

## Definition

```go
type ExprTypeClassicConditionsConditionsOperator struct {
    Type expr.TypeClassicConditionsType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsConditionsOperator` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeClassicConditionsConditionsOperator *ExprTypeClassicConditionsConditionsOperator) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeClassicConditionsConditionsOperator` objects.

```go
func (exprTypeClassicConditionsConditionsOperator *ExprTypeClassicConditionsConditionsOperator) Equals(other ExprTypeClassicConditionsConditionsOperator) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsConditionsOperator` fields for violations and returns them.

```go
func (exprTypeClassicConditionsConditionsOperator *ExprTypeClassicConditionsConditionsOperator) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeClassicConditionsConditionsOperatorBuilder](./builder-ExprTypeClassicConditionsConditionsOperatorBuilder.md)
