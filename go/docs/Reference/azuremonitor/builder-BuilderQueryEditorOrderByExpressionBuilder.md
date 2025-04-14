---
title: <span class="badge builder"></span> BuilderQueryEditorOrderByExpressionBuilder
---
# <span class="badge builder"></span> BuilderQueryEditorOrderByExpressionBuilder

## Constructor

```go
func NewBuilderQueryEditorOrderByExpressionBuilder() *BuilderQueryEditorOrderByExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BuilderQueryEditorOrderByExpressionBuilder) Build() (BuilderQueryEditorOrderByExpression, error)
```

### <span class="badge object-method"></span> Order

```go
func (builder *BuilderQueryEditorOrderByExpressionBuilder) Order(order azuremonitor.BuilderQueryEditorOrderByOptions) *BuilderQueryEditorOrderByExpressionBuilder
```

### <span class="badge object-method"></span> Property

```go
func (builder *BuilderQueryEditorOrderByExpressionBuilder) Property(property cog.Builder[azuremonitor.BuilderQueryEditorProperty]) *BuilderQueryEditorOrderByExpressionBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *BuilderQueryEditorOrderByExpressionBuilder) Type(typeArg azuremonitor.BuilderQueryEditorExpressionType) *BuilderQueryEditorOrderByExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BuilderQueryEditorOrderByExpression](./object-BuilderQueryEditorOrderByExpression.md)
