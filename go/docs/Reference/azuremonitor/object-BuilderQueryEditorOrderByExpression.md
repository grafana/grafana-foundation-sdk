---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorOrderByExpression
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorOrderByExpression

## Definition

```go
type BuilderQueryEditorOrderByExpression struct {
    Property azuremonitor.BuilderQueryEditorProperty `json:"property"`
    Order azuremonitor.BuilderQueryEditorOrderByOptions `json:"order"`
    Type azuremonitor.BuilderQueryEditorExpressionType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorOrderByExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorOrderByExpression *BuilderQueryEditorOrderByExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorOrderByExpression` objects.

```go
func (builderQueryEditorOrderByExpression *BuilderQueryEditorOrderByExpression) Equals(other BuilderQueryEditorOrderByExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorOrderByExpression` fields for violations and returns them.

```go
func (builderQueryEditorOrderByExpression *BuilderQueryEditorOrderByExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorOrderByExpressionBuilder](./builder-BuilderQueryEditorOrderByExpressionBuilder.md)
