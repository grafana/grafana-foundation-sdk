---
title: <span class="badge builder"></span> ExprTypeThresholdConditionsEvaluatorBuilder
---
# <span class="badge builder"></span> ExprTypeThresholdConditionsEvaluatorBuilder

## Constructor

```go
func NewExprTypeThresholdConditionsEvaluatorBuilder() *ExprTypeThresholdConditionsEvaluatorBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) Build() (ExprTypeThresholdConditionsEvaluator, error)
```

### <span class="badge object-method"></span> Params

```go
func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) Params(params []float64) *ExprTypeThresholdConditionsEvaluatorBuilder
```

### <span class="badge object-method"></span> Type

e.g. "gt"

```go
func (builder *ExprTypeThresholdConditionsEvaluatorBuilder) Type(typeArg expr.TypeThresholdType) *ExprTypeThresholdConditionsEvaluatorBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeThresholdConditionsEvaluator](./object-ExprTypeThresholdConditionsEvaluator.md)
