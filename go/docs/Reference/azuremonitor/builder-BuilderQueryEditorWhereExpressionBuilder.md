---
title: <span class="badge builder"></span> BuilderQueryEditorWhereExpressionBuilder
---
# <span class="badge builder"></span> BuilderQueryEditorWhereExpressionBuilder

## Constructor

```go
func NewBuilderQueryEditorWhereExpressionBuilder() *BuilderQueryEditorWhereExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BuilderQueryEditorWhereExpressionBuilder) Build() (BuilderQueryEditorWhereExpression, error)
```

### <span class="badge object-method"></span> Expressions

```go
func (builder *BuilderQueryEditorWhereExpressionBuilder) Expressions(expressions []cog.Builder[azuremonitor.BuilderQueryEditorWhereExpressionItems]) *BuilderQueryEditorWhereExpressionBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *BuilderQueryEditorWhereExpressionBuilder) Type(typeArg azuremonitor.BuilderQueryEditorExpressionType) *BuilderQueryEditorWhereExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BuilderQueryEditorWhereExpression](./object-BuilderQueryEditorWhereExpression.md)
