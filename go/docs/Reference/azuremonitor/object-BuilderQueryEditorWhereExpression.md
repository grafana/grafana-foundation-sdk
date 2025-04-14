---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorWhereExpression
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorWhereExpression

## Definition

```go
type BuilderQueryEditorWhereExpression struct {
    Type azuremonitor.BuilderQueryEditorExpressionType `json:"type"`
    Expressions []azuremonitor.BuilderQueryEditorWhereExpressionItems `json:"expressions"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorWhereExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorWhereExpression *BuilderQueryEditorWhereExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorWhereExpression` objects.

```go
func (builderQueryEditorWhereExpression *BuilderQueryEditorWhereExpression) Equals(other BuilderQueryEditorWhereExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorWhereExpression` fields for violations and returns them.

```go
func (builderQueryEditorWhereExpression *BuilderQueryEditorWhereExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorWhereExpressionBuilder](./builder-BuilderQueryEditorWhereExpressionBuilder.md)
