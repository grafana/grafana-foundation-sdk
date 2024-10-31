---
title: <span class="badge object-type-struct"></span> QueryEditorFunctionExpression
---
# <span class="badge object-type-struct"></span> QueryEditorFunctionExpression

## Definition

```go
type QueryEditorFunctionExpression struct {
    Type string `json:"type"`
    Name *string `json:"name,omitempty"`
    Parameters []cloudwatch.QueryEditorFunctionParameterExpression `json:"parameters,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorFunctionExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryEditorFunctionExpression *QueryEditorFunctionExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryEditorFunctionExpression` objects.

```go
func (queryEditorFunctionExpression *QueryEditorFunctionExpression) Equals(other QueryEditorFunctionExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryEditorFunctionExpression` fields for violations and returns them.

```go
func (queryEditorFunctionExpression *QueryEditorFunctionExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryEditorFunctionExpressionBuilder](./builder-QueryEditorFunctionExpressionBuilder.md)
