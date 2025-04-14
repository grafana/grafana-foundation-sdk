---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorReduceExpression
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorReduceExpression

## Definition

```go
type BuilderQueryEditorReduceExpression struct {
    Property *azuremonitor.BuilderQueryEditorProperty `json:"property,omitempty"`
    Reduce *azuremonitor.BuilderQueryEditorProperty `json:"reduce,omitempty"`
    Parameters []azuremonitor.BuilderQueryEditorFunctionParameterExpression `json:"parameters,omitempty"`
    Focus *bool `json:"focus,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorReduceExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorReduceExpression *BuilderQueryEditorReduceExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorReduceExpression` objects.

```go
func (builderQueryEditorReduceExpression *BuilderQueryEditorReduceExpression) Equals(other BuilderQueryEditorReduceExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorReduceExpression` fields for violations and returns them.

```go
func (builderQueryEditorReduceExpression *BuilderQueryEditorReduceExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorReduceExpressionBuilder](./builder-BuilderQueryEditorReduceExpressionBuilder.md)
