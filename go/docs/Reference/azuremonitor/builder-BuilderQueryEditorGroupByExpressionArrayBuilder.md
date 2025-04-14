---
title: <span class="badge builder"></span> BuilderQueryEditorGroupByExpressionArrayBuilder
---
# <span class="badge builder"></span> BuilderQueryEditorGroupByExpressionArrayBuilder

## Constructor

```go
func NewBuilderQueryEditorGroupByExpressionArrayBuilder() *BuilderQueryEditorGroupByExpressionArrayBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BuilderQueryEditorGroupByExpressionArrayBuilder) Build() (BuilderQueryEditorGroupByExpressionArray, error)
```

### <span class="badge object-method"></span> Expressions

```go
func (builder *BuilderQueryEditorGroupByExpressionArrayBuilder) Expressions(expressions []cog.Builder[azuremonitor.BuilderQueryEditorGroupByExpression]) *BuilderQueryEditorGroupByExpressionArrayBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *BuilderQueryEditorGroupByExpressionArrayBuilder) Type(typeArg azuremonitor.BuilderQueryEditorExpressionType) *BuilderQueryEditorGroupByExpressionArrayBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BuilderQueryEditorGroupByExpressionArray](./object-BuilderQueryEditorGroupByExpressionArray.md)
