---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorOrderByExpressionArray
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorOrderByExpressionArray

## Definition

```go
type BuilderQueryEditorOrderByExpressionArray struct {
    Expressions []azuremonitor.BuilderQueryEditorOrderByExpression `json:"expressions"`
    Type azuremonitor.BuilderQueryEditorExpressionType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorOrderByExpressionArray` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorOrderByExpressionArray *BuilderQueryEditorOrderByExpressionArray) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorOrderByExpressionArray` objects.

```go
func (builderQueryEditorOrderByExpressionArray *BuilderQueryEditorOrderByExpressionArray) Equals(other BuilderQueryEditorOrderByExpressionArray) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorOrderByExpressionArray` fields for violations and returns them.

```go
func (builderQueryEditorOrderByExpressionArray *BuilderQueryEditorOrderByExpressionArray) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorOrderByExpressionArrayBuilder](./builder-BuilderQueryEditorOrderByExpressionArrayBuilder.md)
