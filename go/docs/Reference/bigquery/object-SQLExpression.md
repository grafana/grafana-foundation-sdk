---
title: <span class="badge object-type-struct"></span> SQLExpression
---
# <span class="badge object-type-struct"></span> SQLExpression

## Definition

```go
type SQLExpression struct {
    Columns []bigquery.QueryEditorFunctionExpression `json:"columns,omitempty"`
    From *string `json:"from,omitempty"`
    // whereJsonTree?: _
    WhereString *string `json:"whereString,omitempty"`
    GroupBy []bigquery.QueryEditorGroupByExpression `json:"groupBy,omitempty"`
    OrderBy *bigquery.QueryEditorPropertyExpression `json:"orderBy,omitempty"`
    OrderByDirection *bigquery.OrderByDirection `json:"orderByDirection,omitempty"`
    Limit *int64 `json:"limit,omitempty"`
    Offset *int64 `json:"offset,omitempty"`
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
