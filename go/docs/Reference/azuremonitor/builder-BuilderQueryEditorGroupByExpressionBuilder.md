---
title: <span class="badge builder"></span> BuilderQueryEditorGroupByExpressionBuilder
---
# <span class="badge builder"></span> BuilderQueryEditorGroupByExpressionBuilder

## Constructor

```go
func NewBuilderQueryEditorGroupByExpressionBuilder() *BuilderQueryEditorGroupByExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BuilderQueryEditorGroupByExpressionBuilder) Build() (BuilderQueryEditorGroupByExpression, error)
```

### <span class="badge object-method"></span> Focus

```go
func (builder *BuilderQueryEditorGroupByExpressionBuilder) Focus(focus bool) *BuilderQueryEditorGroupByExpressionBuilder
```

### <span class="badge object-method"></span> Interval

```go
func (builder *BuilderQueryEditorGroupByExpressionBuilder) Interval(interval cog.Builder[azuremonitor.BuilderQueryEditorProperty]) *BuilderQueryEditorGroupByExpressionBuilder
```

### <span class="badge object-method"></span> Property

```go
func (builder *BuilderQueryEditorGroupByExpressionBuilder) Property(property cog.Builder[azuremonitor.BuilderQueryEditorProperty]) *BuilderQueryEditorGroupByExpressionBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *BuilderQueryEditorGroupByExpressionBuilder) Type(typeArg azuremonitor.BuilderQueryEditorExpressionType) *BuilderQueryEditorGroupByExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BuilderQueryEditorGroupByExpression](./object-BuilderQueryEditorGroupByExpression.md)
