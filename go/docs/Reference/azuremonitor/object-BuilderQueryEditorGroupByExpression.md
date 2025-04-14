---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorGroupByExpression
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorGroupByExpression

## Definition

```go
type BuilderQueryEditorGroupByExpression struct {
    Property *azuremonitor.BuilderQueryEditorProperty `json:"property,omitempty"`
    Interval *azuremonitor.BuilderQueryEditorProperty `json:"interval,omitempty"`
    Focus *bool `json:"focus,omitempty"`
    Type *azuremonitor.BuilderQueryEditorExpressionType `json:"type,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorGroupByExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorGroupByExpression *BuilderQueryEditorGroupByExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorGroupByExpression` objects.

```go
func (builderQueryEditorGroupByExpression *BuilderQueryEditorGroupByExpression) Equals(other BuilderQueryEditorGroupByExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorGroupByExpression` fields for violations and returns them.

```go
func (builderQueryEditorGroupByExpression *BuilderQueryEditorGroupByExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorGroupByExpressionBuilder](./builder-BuilderQueryEditorGroupByExpressionBuilder.md)
