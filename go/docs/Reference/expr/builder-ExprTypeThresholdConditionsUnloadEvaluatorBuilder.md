---
title: <span class="badge builder"></span> ExprTypeThresholdConditionsUnloadEvaluatorBuilder
---
# <span class="badge builder"></span> ExprTypeThresholdConditionsUnloadEvaluatorBuilder

## Constructor

```go
func NewExprTypeThresholdConditionsUnloadEvaluatorBuilder() *ExprTypeThresholdConditionsUnloadEvaluatorBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeThresholdConditionsUnloadEvaluatorBuilder) Build() (ExprTypeThresholdConditionsUnloadEvaluator, error)
```

### <span class="badge object-method"></span> Params

```go
func (builder *ExprTypeThresholdConditionsUnloadEvaluatorBuilder) Params(params []float64) *ExprTypeThresholdConditionsUnloadEvaluatorBuilder
```

### <span class="badge object-method"></span> Type

e.g. "gt"

```go
func (builder *ExprTypeThresholdConditionsUnloadEvaluatorBuilder) Type(typeArg expr.TypeThresholdType) *ExprTypeThresholdConditionsUnloadEvaluatorBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeThresholdConditionsUnloadEvaluator](./object-ExprTypeThresholdConditionsUnloadEvaluator.md)
