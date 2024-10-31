---
title: <span class="badge object-type-struct"></span> ExprTypeThresholdConditions
---
# <span class="badge object-type-struct"></span> ExprTypeThresholdConditions

## Definition

```go
type ExprTypeThresholdConditions struct {
    Evaluator expr.ExprTypeThresholdConditionsEvaluator `json:"evaluator"`
    LoadedDimensions any `json:"loadedDimensions,omitempty"`
    UnloadEvaluator *expr.ExprTypeThresholdConditionsUnloadEvaluator `json:"unloadEvaluator,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeThresholdConditions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (exprTypeThresholdConditions *ExprTypeThresholdConditions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ExprTypeThresholdConditions` objects.

```go
func (exprTypeThresholdConditions *ExprTypeThresholdConditions) Equals(other ExprTypeThresholdConditions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ExprTypeThresholdConditions` fields for violations and returns them.

```go
func (exprTypeThresholdConditions *ExprTypeThresholdConditions) Validate() error
```

## See also

 * <span class="badge builder"></span> [ExprTypeThresholdConditionsBuilder](./builder-ExprTypeThresholdConditionsBuilder.md)
