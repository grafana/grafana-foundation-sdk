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

### <span class="badge object-method"></span> From

FROM part of the SQL expression

```go
func (builder *SQLExpressionBuilder) From(from cloudwatch.QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> GroupBy

GROUP BY part of the SQL expression

```go
func (builder *SQLExpressionBuilder) GroupBy(groupBy cog.Builder[cloudwatch.QueryEditorArrayExpression]) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> Limit

LIMIT part of the SQL expression

```go
func (builder *SQLExpressionBuilder) Limit(limit int64) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> OrderBy

ORDER BY part of the SQL expression

```go
func (builder *SQLExpressionBuilder) OrderBy(orderBy cog.Builder[cloudwatch.QueryEditorFunctionExpression]) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> OrderByDirection

The sort order of the SQL expression, `ASC` or `DESC`

```go
func (builder *SQLExpressionBuilder) OrderByDirection(orderByDirection string) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> Select

SELECT part of the SQL expression

```go
func (builder *SQLExpressionBuilder) Select(selectArg cog.Builder[cloudwatch.QueryEditorFunctionExpression]) *SQLExpressionBuilder
```

### <span class="badge object-method"></span> Where

WHERE part of the SQL expression

```go
func (builder *SQLExpressionBuilder) Where(where cog.Builder[cloudwatch.QueryEditorArrayExpression]) *SQLExpressionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SQLExpression](./object-SQLExpression.md)
