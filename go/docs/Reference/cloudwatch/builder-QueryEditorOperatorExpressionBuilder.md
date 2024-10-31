---
title: <span class="badge builder"></span> QueryEditorOperatorExpressionBuilder
---
# <span class="badge builder"></span> QueryEditorOperatorExpressionBuilder

## Constructor

```go
func NewQueryEditorOperatorExpressionBuilder() *QueryEditorOperatorExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryEditorOperatorExpressionBuilder) Build() (QueryEditorOperatorExpression, error)
```

### <span class="badge object-method"></span> Operator

TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer

```go
func (builder *QueryEditorOperatorExpressionBuilder) Operator(operator cog.Builder[cloudwatch.QueryEditorOperator]) *QueryEditorOperatorExpressionBuilder
```

### <span class="badge object-method"></span> Property

```go
func (builder *QueryEditorOperatorExpressionBuilder) Property(property cog.Builder[cloudwatch.QueryEditorProperty]) *QueryEditorOperatorExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [QueryEditorOperatorExpression](./object-QueryEditorOperatorExpression.md)
