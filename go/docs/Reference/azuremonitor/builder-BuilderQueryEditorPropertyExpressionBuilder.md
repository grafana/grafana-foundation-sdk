---
title: <span class="badge builder"></span> BuilderQueryEditorPropertyExpressionBuilder
---
# <span class="badge builder"></span> BuilderQueryEditorPropertyExpressionBuilder

## Constructor

```go
func NewBuilderQueryEditorPropertyExpressionBuilder() *BuilderQueryEditorPropertyExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BuilderQueryEditorPropertyExpressionBuilder) Build() (BuilderQueryEditorPropertyExpression, error)
```

### <span class="badge object-method"></span> Property

```go
func (builder *BuilderQueryEditorPropertyExpressionBuilder) Property(property cog.Builder[azuremonitor.BuilderQueryEditorProperty]) *BuilderQueryEditorPropertyExpressionBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *BuilderQueryEditorPropertyExpressionBuilder) Type(typeArg azuremonitor.BuilderQueryEditorExpressionType) *BuilderQueryEditorPropertyExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BuilderQueryEditorPropertyExpression](./object-BuilderQueryEditorPropertyExpression.md)
