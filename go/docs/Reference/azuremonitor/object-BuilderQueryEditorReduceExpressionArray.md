---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorReduceExpressionArray
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorReduceExpressionArray

## Definition

```go
type BuilderQueryEditorReduceExpressionArray struct {
    Expressions []azuremonitor.BuilderQueryEditorReduceExpression `json:"expressions"`
    Type azuremonitor.BuilderQueryEditorExpressionType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorReduceExpressionArray` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorReduceExpressionArray *BuilderQueryEditorReduceExpressionArray) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorReduceExpressionArray` objects.

```go
func (builderQueryEditorReduceExpressionArray *BuilderQueryEditorReduceExpressionArray) Equals(other BuilderQueryEditorReduceExpressionArray) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorReduceExpressionArray` fields for violations and returns them.

```go
func (builderQueryEditorReduceExpressionArray *BuilderQueryEditorReduceExpressionArray) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorReduceExpressionArrayBuilder](./builder-BuilderQueryEditorReduceExpressionArrayBuilder.md)
