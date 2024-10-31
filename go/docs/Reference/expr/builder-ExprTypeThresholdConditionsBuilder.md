---
title: <span class="badge builder"></span> ExprTypeThresholdConditionsBuilder
---
# <span class="badge builder"></span> ExprTypeThresholdConditionsBuilder

## Constructor

```go
func NewExprTypeThresholdConditionsBuilder() *ExprTypeThresholdConditionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeThresholdConditionsBuilder) Build() (ExprTypeThresholdConditions, error)
```

### <span class="badge object-method"></span> Evaluator

```go
func (builder *ExprTypeThresholdConditionsBuilder) Evaluator(evaluator cog.Builder[expr.ExprTypeThresholdConditionsEvaluator]) *ExprTypeThresholdConditionsBuilder
```

### <span class="badge object-method"></span> LoadedDimensions

```go
func (builder *ExprTypeThresholdConditionsBuilder) LoadedDimensions(loadedDimensions any) *ExprTypeThresholdConditionsBuilder
```

### <span class="badge object-method"></span> UnloadEvaluator

```go
func (builder *ExprTypeThresholdConditionsBuilder) UnloadEvaluator(unloadEvaluator cog.Builder[expr.ExprTypeThresholdConditionsUnloadEvaluator]) *ExprTypeThresholdConditionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeThresholdConditions](./object-ExprTypeThresholdConditions.md)
