---
title: <span class="badge object-type-struct"></span> QueryEditorPropertyExpression
---
# <span class="badge object-type-struct"></span> QueryEditorPropertyExpression

## Definition

```go
type QueryEditorPropertyExpression struct {
    Type string `json:"type"`
    Property cloudwatch.QueryEditorProperty `json:"property"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorPropertyExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryEditorPropertyExpression *QueryEditorPropertyExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryEditorPropertyExpression` objects.

```go
func (queryEditorPropertyExpression *QueryEditorPropertyExpression) Equals(other QueryEditorPropertyExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryEditorPropertyExpression` fields for violations and returns them.

```go
func (queryEditorPropertyExpression *QueryEditorPropertyExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryEditorPropertyExpressionBuilder](./builder-QueryEditorPropertyExpressionBuilder.md)
