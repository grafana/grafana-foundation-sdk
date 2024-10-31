---
title: <span class="badge object-type-struct"></span> SQLExpression
---
# <span class="badge object-type-struct"></span> SQLExpression

## Definition

```go
type SQLExpression struct {
    // SELECT part of the SQL expression
    Select *cloudwatch.QueryEditorFunctionExpression `json:"select,omitempty"`
    // FROM part of the SQL expression
    From *cloudwatch.QueryEditorPropertyExpressionOrQueryEditorFunctionExpression `json:"from,omitempty"`
    // WHERE part of the SQL expression
    Where *cloudwatch.QueryEditorArrayExpression `json:"where,omitempty"`
    // GROUP BY part of the SQL expression
    GroupBy *cloudwatch.QueryEditorArrayExpression `json:"groupBy,omitempty"`
    // ORDER BY part of the SQL expression
    OrderBy *cloudwatch.QueryEditorFunctionExpression `json:"orderBy,omitempty"`
    // The sort order of the SQL expression, `ASC` or `DESC`
    OrderByDirection *string `json:"orderByDirection,omitempty"`
    // LIMIT part of the SQL expression
    Limit *int64 `json:"limit,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SQLExpression` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (sQLExpression *SQLExpression) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SQLExpression` objects.

```go
func (sQLExpression *SQLExpression) Equals(other SQLExpression) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SQLExpression` fields for violations and returns them.

```go
func (sQLExpression *SQLExpression) Validate() error
```

## See also

 * <span class="badge builder"></span> [SQLExpressionBuilder](./builder-SQLExpressionBuilder.md)
