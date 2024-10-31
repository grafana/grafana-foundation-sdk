---
title: <span class="badge object-type-struct"></span> ExprTypeThresholdConditionsUnloadEvaluator
---
# <span class="badge object-type-struct"></span> ExprTypeThresholdConditionsUnloadEvaluator

## Definition

```go
type ExprTypeThresholdConditionsUnloadEvaluator struct {
    Params []float64 `json:"params"`
    // e.g. "gt"
    Type expr.TypeThresholdType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeThresholdConditionsUnloadEvaluator` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeThresholdConditionsUnloadEvaluator *ExprTypeThresholdConditionsUnloadEvaluator) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeThresholdConditionsUnloadEvaluator` objects.

```go
func (exprTypeThresholdConditionsUnloadEvaluator *ExprTypeThresholdConditionsUnloadEvaluator) Equals(other ExprTypeThresholdConditionsUnloadEvaluator) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeThresholdConditionsUnloadEvaluator` fields for violations and returns them.

```go
func (exprTypeThresholdConditionsUnloadEvaluator *ExprTypeThresholdConditionsUnloadEvaluator) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeThresholdConditionsUnloadEvaluatorBuilder](./builder-ExprTypeThresholdConditionsUnloadEvaluatorBuilder.md)
