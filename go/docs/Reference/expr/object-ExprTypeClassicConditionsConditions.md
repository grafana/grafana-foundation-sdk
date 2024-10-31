---
title: <span class="badge object-type-struct"></span> ExprTypeClassicConditionsConditions
---
# <span class="badge object-type-struct"></span> ExprTypeClassicConditionsConditions

## Definition

```go
type ExprTypeClassicConditionsConditions struct {
    Evaluator expr.ExprTypeClassicConditionsConditionsEvaluator `json:"evaluator"`
    Operator expr.ExprTypeClassicConditionsConditionsOperator `json:"operator"`
    Query expr.ExprTypeClassicConditionsConditionsQuery `json:"query"`
    Reducer expr.ExprTypeClassicConditionsConditionsReducer `json:"reducer"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsConditions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeClassicConditionsConditions *ExprTypeClassicConditionsConditions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeClassicConditionsConditions` objects.

```go
func (exprTypeClassicConditionsConditions *ExprTypeClassicConditionsConditions) Equals(other ExprTypeClassicConditionsConditions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsConditions` fields for violations and returns them.

```go
func (exprTypeClassicConditionsConditions *ExprTypeClassicConditionsConditions) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeClassicConditionsConditionsBuilder](./builder-ExprTypeClassicConditionsConditionsBuilder.md)
