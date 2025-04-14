---
title: <span class="badge builder"></span> BuilderQueryEditorWhereExpressionItemsBuilder
---
# <span class="badge builder"></span> BuilderQueryEditorWhereExpressionItemsBuilder

## Constructor

```go
func NewBuilderQueryEditorWhereExpressionItemsBuilder() *BuilderQueryEditorWhereExpressionItemsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BuilderQueryEditorWhereExpressionItemsBuilder) Build() (BuilderQueryEditorWhereExpressionItems, error)
```

### <span class="badge object-method"></span> Operator

```go
func (builder *BuilderQueryEditorWhereExpressionItemsBuilder) Operator(operator cog.Builder[azuremonitor.BuilderQueryEditorOperator]) *BuilderQueryEditorWhereExpressionItemsBuilder
```

### <span class="badge object-method"></span> Property

```go
func (builder *BuilderQueryEditorWhereExpressionItemsBuilder) Property(property cog.Builder[azuremonitor.BuilderQueryEditorProperty]) *BuilderQueryEditorWhereExpressionItemsBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *BuilderQueryEditorWhereExpressionItemsBuilder) Type(typeArg azuremonitor.BuilderQueryEditorExpressionType) *BuilderQueryEditorWhereExpressionItemsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BuilderQueryEditorWhereExpressionItems](./object-BuilderQueryEditorWhereExpressionItems.md)
