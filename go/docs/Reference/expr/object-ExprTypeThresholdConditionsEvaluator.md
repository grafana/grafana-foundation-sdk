---
title: <span class="badge object-type-struct"></span> ExprTypeThresholdConditionsEvaluator
---
# <span class="badge object-type-struct"></span> ExprTypeThresholdConditionsEvaluator

## Definition

```go
type ExprTypeThresholdConditionsEvaluator struct {
    Params []float64 `json:"params"`
    // e.g. "gt"
    Type expr.TypeThresholdType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeThresholdConditionsEvaluator` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeThresholdConditionsEvaluator *ExprTypeThresholdConditionsEvaluator) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeThresholdConditionsEvaluator` objects.

```go
func (exprTypeThresholdConditionsEvaluator *ExprTypeThresholdConditionsEvaluator) Equals(other ExprTypeThresholdConditionsEvaluator) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeThresholdConditionsEvaluator` fields for violations and returns them.

```go
func (exprTypeThresholdConditionsEvaluator *ExprTypeThresholdConditionsEvaluator) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeThresholdConditionsEvaluatorBuilder](./builder-ExprTypeThresholdConditionsEvaluatorBuilder.md)
