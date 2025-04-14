---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorPropertyExpression
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorPropertyExpression

## Definition

```go
type BuilderQueryEditorPropertyExpression struct {
    Property azuremonitor.BuilderQueryEditorProperty `json:"property"`
    Type azuremonitor.BuilderQueryEditorExpressionType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorPropertyExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorPropertyExpression *BuilderQueryEditorPropertyExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorPropertyExpression` objects.

```go
func (builderQueryEditorPropertyExpression *BuilderQueryEditorPropertyExpression) Equals(other BuilderQueryEditorPropertyExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorPropertyExpression` fields for violations and returns them.

```go
func (builderQueryEditorPropertyExpression *BuilderQueryEditorPropertyExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorPropertyExpressionBuilder](./builder-BuilderQueryEditorPropertyExpressionBuilder.md)
