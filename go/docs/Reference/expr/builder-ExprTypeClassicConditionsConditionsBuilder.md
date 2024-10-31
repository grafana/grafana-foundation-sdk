---
title: <span class="badge builder"></span> ExprTypeClassicConditionsConditionsBuilder
---
# <span class="badge builder"></span> ExprTypeClassicConditionsConditionsBuilder

## Constructor

```go
func NewExprTypeClassicConditionsConditionsBuilder() *ExprTypeClassicConditionsConditionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprTypeClassicConditionsConditionsBuilder) Build() (ExprTypeClassicConditionsConditions, error)
```

### <span class="badge object-method"></span> Evaluator

```go
func (builder *ExprTypeClassicConditionsConditionsBuilder) Evaluator(evaluator cog.Builder[expr.ExprTypeClassicConditionsConditionsEvaluator]) *ExprTypeClassicConditionsConditionsBuilder
```

### <span class="badge object-method"></span> Operator

```go
func (builder *ExprTypeClassicConditionsConditionsBuilder) Operator(operator cog.Builder[expr.ExprTypeClassicConditionsConditionsOperator]) *ExprTypeClassicConditionsConditionsBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *ExprTypeClassicConditionsConditionsBuilder) Query(query cog.Builder[expr.ExprTypeClassicConditionsConditionsQuery]) *ExprTypeClassicConditionsConditionsBuilder
```

### <span class="badge object-method"></span> Reducer

```go
func (builder *ExprTypeClassicConditionsConditionsBuilder) Reducer(reducer cog.Builder[expr.ExprTypeClassicConditionsConditionsReducer]) *ExprTypeClassicConditionsConditionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExprTypeClassicConditionsConditions](./object-ExprTypeClassicConditionsConditions.md)
