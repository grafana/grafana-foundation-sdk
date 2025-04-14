---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorWhereExpressionItems
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorWhereExpressionItems

## Definition

```go
type BuilderQueryEditorWhereExpressionItems struct {
    Property azuremonitor.BuilderQueryEditorProperty `json:"property"`
    Operator azuremonitor.BuilderQueryEditorOperator `json:"operator"`
    Type azuremonitor.BuilderQueryEditorExpressionType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorWhereExpressionItems` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorWhereExpressionItems *BuilderQueryEditorWhereExpressionItems) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorWhereExpressionItems` objects.

```go
func (builderQueryEditorWhereExpressionItems *BuilderQueryEditorWhereExpressionItems) Equals(other BuilderQueryEditorWhereExpressionItems) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorWhereExpressionItems` fields for violations and returns them.

```go
func (builderQueryEditorWhereExpressionItems *BuilderQueryEditorWhereExpressionItems) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorWhereExpressionItemsBuilder](./builder-BuilderQueryEditorWhereExpressionItemsBuilder.md)
