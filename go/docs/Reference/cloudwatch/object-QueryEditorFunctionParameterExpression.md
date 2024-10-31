---
title: <span class="badge object-type-struct"></span> QueryEditorFunctionParameterExpression
---
# <span class="badge object-type-struct"></span> QueryEditorFunctionParameterExpression

## Definition

```go
type QueryEditorFunctionParameterExpression struct {
    Type string `json:"type"`
    Name *string `json:"name,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorFunctionParameterExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryEditorFunctionParameterExpression *QueryEditorFunctionParameterExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryEditorFunctionParameterExpression` objects.

```go
func (queryEditorFunctionParameterExpression *QueryEditorFunctionParameterExpression) Equals(other QueryEditorFunctionParameterExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryEditorFunctionParameterExpression` fields for violations and returns them.

```go
func (queryEditorFunctionParameterExpression *QueryEditorFunctionParameterExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryEditorFunctionParameterExpressionBuilder](./builder-QueryEditorFunctionParameterExpressionBuilder.md)
