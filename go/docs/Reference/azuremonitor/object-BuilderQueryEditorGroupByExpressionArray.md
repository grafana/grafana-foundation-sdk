---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorGroupByExpressionArray
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorGroupByExpressionArray

## Definition

```go
type BuilderQueryEditorGroupByExpressionArray struct {
    Expressions []azuremonitor.BuilderQueryEditorGroupByExpression `json:"expressions"`
    Type azuremonitor.BuilderQueryEditorExpressionType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorGroupByExpressionArray` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorGroupByExpressionArray *BuilderQueryEditorGroupByExpressionArray) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorGroupByExpressionArray` objects.

```go
func (builderQueryEditorGroupByExpressionArray *BuilderQueryEditorGroupByExpressionArray) Equals(other BuilderQueryEditorGroupByExpressionArray) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorGroupByExpressionArray` fields for violations and returns them.

```go
func (builderQueryEditorGroupByExpressionArray *BuilderQueryEditorGroupByExpressionArray) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorGroupByExpressionArrayBuilder](./builder-BuilderQueryEditorGroupByExpressionArrayBuilder.md)
