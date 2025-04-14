---
title: <span class="badge object-type-struct"></span> BuilderQueryEditorFunctionParameterExpression
---
# <span class="badge object-type-struct"></span> BuilderQueryEditorFunctionParameterExpression

## Definition

```go
type BuilderQueryEditorFunctionParameterExpression struct {
    Value string `json:"value"`
    FieldType azuremonitor.BuilderQueryEditorPropertyType `json:"fieldType"`
    Type azuremonitor.BuilderQueryEditorExpressionType `json:"type"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BuilderQueryEditorFunctionParameterExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (builderQueryEditorFunctionParameterExpression *BuilderQueryEditorFunctionParameterExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BuilderQueryEditorFunctionParameterExpression` objects.

```go
func (builderQueryEditorFunctionParameterExpression *BuilderQueryEditorFunctionParameterExpression) Equals(other BuilderQueryEditorFunctionParameterExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BuilderQueryEditorFunctionParameterExpression` fields for violations and returns them.

```go
func (builderQueryEditorFunctionParameterExpression *BuilderQueryEditorFunctionParameterExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [BuilderQueryEditorFunctionParameterExpressionBuilder](./builder-BuilderQueryEditorFunctionParameterExpressionBuilder.md)
