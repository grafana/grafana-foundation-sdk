---
title: <span class="badge builder"></span> QueryEditorFunctionExpressionBuilder
---
# <span class="badge builder"></span> QueryEditorFunctionExpressionBuilder

## Constructor

```go
func NewQueryEditorFunctionExpressionBuilder() *QueryEditorFunctionExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryEditorFunctionExpressionBuilder) Build() (QueryEditorFunctionExpression, error)
```

### <span class="badge object-method"></span> Name

```go
func (builder *QueryEditorFunctionExpressionBuilder) Name(name string) *QueryEditorFunctionExpressionBuilder
```

### <span class="badge object-method"></span> Parameters

```go
func (builder *QueryEditorFunctionExpressionBuilder) Parameters(parameters []cog.Builder[cloudwatch.QueryEditorFunctionParameterExpression]) *QueryEditorFunctionExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [QueryEditorFunctionExpression](./object-QueryEditorFunctionExpression.md)
