---
title: <span class="badge object-type-struct"></span> QueryEditorArrayExpression
---
# <span class="badge object-type-struct"></span> QueryEditorArrayExpression

## Definition

```go
type QueryEditorArrayExpression struct {
    Type cloudwatch.QueryEditorArrayExpressionType `json:"type"`
    Expressions []cloudwatch.QueryEditorExpression `json:"expressions"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorArrayExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryEditorArrayExpression *QueryEditorArrayExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryEditorArrayExpression` objects.

```go
func (queryEditorArrayExpression *QueryEditorArrayExpression) Equals(other QueryEditorArrayExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryEditorArrayExpression` fields for violations and returns them.

```go
func (queryEditorArrayExpression *QueryEditorArrayExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryEditorArrayExpressionBuilder](./builder-QueryEditorArrayExpressionBuilder.md)
