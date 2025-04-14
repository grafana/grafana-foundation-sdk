---
title: <span class="badge builder"></span> BuilderQueryEditorReduceExpressionArrayBuilder
---
# <span class="badge builder"></span> BuilderQueryEditorReduceExpressionArrayBuilder

## Constructor

```go
func NewBuilderQueryEditorReduceExpressionArrayBuilder() *BuilderQueryEditorReduceExpressionArrayBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BuilderQueryEditorReduceExpressionArrayBuilder) Build() (BuilderQueryEditorReduceExpressionArray, error)
```

### <span class="badge object-method"></span> Expressions

```go
func (builder *BuilderQueryEditorReduceExpressionArrayBuilder) Expressions(expressions []cog.Builder[azuremonitor.BuilderQueryEditorReduceExpression]) *BuilderQueryEditorReduceExpressionArrayBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *BuilderQueryEditorReduceExpressionArrayBuilder) Type(typeArg azuremonitor.BuilderQueryEditorExpressionType) *BuilderQueryEditorReduceExpressionArrayBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BuilderQueryEditorReduceExpressionArray](./object-BuilderQueryEditorReduceExpressionArray.md)
