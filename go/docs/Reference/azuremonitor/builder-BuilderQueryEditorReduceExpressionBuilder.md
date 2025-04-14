---
title: <span class="badge builder"></span> BuilderQueryEditorReduceExpressionBuilder
---
# <span class="badge builder"></span> BuilderQueryEditorReduceExpressionBuilder

## Constructor

```go
func NewBuilderQueryEditorReduceExpressionBuilder() *BuilderQueryEditorReduceExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BuilderQueryEditorReduceExpressionBuilder) Build() (BuilderQueryEditorReduceExpression, error)
```

### <span class="badge object-method"></span> Focus

```go
func (builder *BuilderQueryEditorReduceExpressionBuilder) Focus(focus bool) *BuilderQueryEditorReduceExpressionBuilder
```

### <span class="badge object-method"></span> Parameters

```go
func (builder *BuilderQueryEditorReduceExpressionBuilder) Parameters(parameters []cog.Builder[azuremonitor.BuilderQueryEditorFunctionParameterExpression]) *BuilderQueryEditorReduceExpressionBuilder
```

### <span class="badge object-method"></span> Property

```go
func (builder *BuilderQueryEditorReduceExpressionBuilder) Property(property cog.Builder[azuremonitor.BuilderQueryEditorProperty]) *BuilderQueryEditorReduceExpressionBuilder
```

### <span class="badge object-method"></span> Reduce

```go
func (builder *BuilderQueryEditorReduceExpressionBuilder) Reduce(reduce cog.Builder[azuremonitor.BuilderQueryEditorProperty]) *BuilderQueryEditorReduceExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BuilderQueryEditorReduceExpression](./object-BuilderQueryEditorReduceExpression.md)
