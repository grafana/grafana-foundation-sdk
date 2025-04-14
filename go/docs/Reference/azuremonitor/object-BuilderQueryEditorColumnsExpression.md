---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorColumnsExpression
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorColumnsExpression

## Definition

```go
type BuilderQueryEditorColumnsExpression struct {
    Columns []string `json:"columns,omitempty"`
    Type azuremonitor.BuilderQueryEditorExpressionType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorColumnsExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorColumnsExpression *BuilderQueryEditorColumnsExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorColumnsExpression` objects.

```go
func (builderQueryEditorColumnsExpression *BuilderQueryEditorColumnsExpression) Equals(other BuilderQueryEditorColumnsExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorColumnsExpression` fields for violations and returns them.

```go
func (builderQueryEditorColumnsExpression *BuilderQueryEditorColumnsExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorColumnsExpressionBuilder](./builder-BuilderQueryEditorColumnsExpressionBuilder.md)
