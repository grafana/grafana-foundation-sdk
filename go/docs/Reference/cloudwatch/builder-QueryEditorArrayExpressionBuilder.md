---
title: <span class="badge builder"></span> QueryEditorArrayExpressionBuilder
---
# <span class="badge builder"></span> QueryEditorArrayExpressionBuilder

## Constructor

```go
func NewQueryEditorArrayExpressionBuilder() *QueryEditorArrayExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryEditorArrayExpressionBuilder) Build() (QueryEditorArrayExpression, error)
```

### <span class="badge object-method"></span> Expressions

```go
func (builder *QueryEditorArrayExpressionBuilder) Expressions(expressions []cloudwatch.QueryEditorExpression) *QueryEditorArrayExpressionBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *QueryEditorArrayExpressionBuilder) Type(typeArg cloudwatch.QueryEditorArrayExpressionType) *QueryEditorArrayExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [QueryEditorArrayExpression](./object-QueryEditorArrayExpression.md)
