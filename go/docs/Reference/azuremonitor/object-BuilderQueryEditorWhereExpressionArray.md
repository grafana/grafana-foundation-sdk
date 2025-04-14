---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorWhereExpressionArray
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorWhereExpressionArray

## Definition

```go
type BuilderQueryEditorWhereExpressionArray struct {
    Expressions []azuremonitor.BuilderQueryEditorWhereExpression `json:"expressions"`
    Type azuremonitor.BuilderQueryEditorExpressionType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorWhereExpressionArray` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorWhereExpressionArray *BuilderQueryEditorWhereExpressionArray) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorWhereExpressionArray` objects.

```go
func (builderQueryEditorWhereExpressionArray *BuilderQueryEditorWhereExpressionArray) Equals(other BuilderQueryEditorWhereExpressionArray) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorWhereExpressionArray` fields for violations and returns them.

```go
func (builderQueryEditorWhereExpressionArray *BuilderQueryEditorWhereExpressionArray) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorWhereExpressionArrayBuilder](./builder-BuilderQueryEditorWhereExpressionArrayBuilder.md)
