---
title: <span class="badge object-type-struct"></span> QueryEditorPropertyExpressionOrQueryEditorFunctionExpression
---
# <span class="badge object-type-struct"></span> QueryEditorPropertyExpressionOrQueryEditorFunctionExpression

## Definition

```go
type QueryEditorPropertyExpressionOrQueryEditorFunctionExpression struct {
    QueryEditorPropertyExpression *cloudwatch.QueryEditorPropertyExpression `json:"QueryEditorPropertyExpression,omitempty"`
    QueryEditorFunctionExpression *cloudwatch.QueryEditorFunctionExpression `json:"QueryEditorFunctionExpression,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `QueryEditorPropertyExpressionOrQueryEditorFunctionExpression` as JSON.

```go
func (queryEditorPropertyExpressionOrQueryEditorFunctionExpression *QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `QueryEditorPropertyExpressionOrQueryEditorFunctionExpression` from JSON.

```go
func (queryEditorPropertyExpressionOrQueryEditorFunctionExpression *QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorPropertyExpressionOrQueryEditorFunctionExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryEditorPropertyExpressionOrQueryEditorFunctionExpression *QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryEditorPropertyExpressionOrQueryEditorFunctionExpression` objects.

```go
func (queryEditorPropertyExpressionOrQueryEditorFunctionExpression *QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) Equals(other QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryEditorPropertyExpressionOrQueryEditorFunctionExpression` fields for violations and returns them.

```go
func (queryEditorPropertyExpressionOrQueryEditorFunctionExpression *QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) Validate() error
```

