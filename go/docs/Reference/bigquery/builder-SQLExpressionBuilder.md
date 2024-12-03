---
title: <span class="badge builder"></span> SQLExpressionBuilder
---
# <span class="badge builder"></span> SQLExpressionBuilder

## Constructor

```go
func NewSQLExpressionBuilder() *SQLExpressionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SQLExpressionBuilder) Build() (SQLExpression, error)
```

### <span class="badge object-method"></span> Columns

```go
func (builder *SQLExpressionBuilder) Columns(columns []cog.Builder[bigquery.QueryEditorFunctionExpression]) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> From

```go
func (builder *SQLExpressionBuilder) From(from string) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> GroupBy

```go
func (builder *SQLExpressionBuilder) GroupBy(groupBy []cog.Builder[bigquery.QueryEditorGroupByExpression]) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> Limit

```go
func (builder *SQLExpressionBuilder) Limit(limit int64) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> Offset

```go
func (builder *SQLExpressionBuilder) Offset(offset int64) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> OrderBy

```go
func (builder *SQLExpressionBuilder) OrderBy(orderBy cog.Builder[bigquery.QueryEditorPropertyExpression]) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> OrderByDirection

```go
func (builder *SQLExpressionBuilder) OrderByDirection(orderByDirection bigquery.OrderByDirection) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> WhereString

whereJsonTree?: _

```go
func (builder *SQLExpressionBuilder) WhereString(whereString string) *SQLExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SQLExpression](./object-SQLExpression.md)
