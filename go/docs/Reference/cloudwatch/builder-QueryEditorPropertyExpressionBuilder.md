---
title: <span class="badge builder"></span> QueryEditorPropertyExpressionBuilder
---
# <span class="badge builder"></span> QueryEditorPropertyExpressionBuilder

## Constructor

```go
func NewQueryEditorPropertyExpressionBuilder() *QueryEditorPropertyExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryEditorPropertyExpressionBuilder) Build() (QueryEditorPropertyExpression, error)
```

### <span class="badge object-method"></span> Property

```go
func (builder *QueryEditorPropertyExpressionBuilder) Property(property cog.Builder[cloudwatch.QueryEditorProperty]) *QueryEditorPropertyExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [QueryEditorPropertyExpression](./object-QueryEditorPropertyExpression.md)
