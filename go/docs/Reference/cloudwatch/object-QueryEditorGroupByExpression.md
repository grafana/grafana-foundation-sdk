---
title: <span class="badge object-type-struct"></span> QueryEditorGroupByExpression
---
# <span class="badge object-type-struct"></span> QueryEditorGroupByExpression

## Definition

```go
type QueryEditorGroupByExpression struct {
    Type string `json:"type"`
    Property cloudwatch.QueryEditorProperty `json:"property"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorGroupByExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (queryEditorGroupByExpression *QueryEditorGroupByExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `QueryEditorGroupByExpression` objects.

```go
func (queryEditorGroupByExpression *QueryEditorGroupByExpression) Equals(other QueryEditorGroupByExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `QueryEditorGroupByExpression` fields for violations and returns them.

```go
func (queryEditorGroupByExpression *QueryEditorGroupByExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [QueryEditorGroupByExpressionBuilder](./builder-QueryEditorGroupByExpressionBuilder.md)
