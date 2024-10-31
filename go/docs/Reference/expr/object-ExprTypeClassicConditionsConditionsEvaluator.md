---
title: <span class="badge object-type-struct"></span> ExprTypeClassicConditionsConditionsEvaluator
---
# <span class="badge object-type-struct"></span> ExprTypeClassicConditionsConditionsEvaluator

## Definition

```go
type ExprTypeClassicConditionsConditionsEvaluator struct {
    Params []float64 `json:"params"`
    // e.g. "gt"
    Type string `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsConditionsEvaluator` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeClassicConditionsConditionsEvaluator *ExprTypeClassicConditionsConditionsEvaluator) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeClassicConditionsConditionsEvaluator` objects.

```go
func (exprTypeClassicConditionsConditionsEvaluator *ExprTypeClassicConditionsConditionsEvaluator) Equals(other ExprTypeClassicConditionsConditionsEvaluator) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsConditionsEvaluator` fields for violations and returns them.

```go
func (exprTypeClassicConditionsConditionsEvaluator *ExprTypeClassicConditionsConditionsEvaluator) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeClassicConditionsConditionsEvaluatorBuilder](./builder-ExprTypeClassicConditionsConditionsEvaluatorBuilder.md)
