# <span class="badge package-variant-dataquery"></span> bigquery

## Objects

 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
 * <span class="badge object-type-enum"></span> [EditorMode](./object-EditorMode.md)
 * <span class="badge object-type-enum"></span> [OrderByDirection](./object-OrderByDirection.md)
 * <span class="badge object-type-enum"></span> [QueryEditorExpressionType](./object-QueryEditorExpressionType.md)
 * <span class="badge object-type-struct"></span> [QueryEditorFunctionExpression](./object-QueryEditorFunctionExpression.md)
 * <span class="badge object-type-struct"></span> [QueryEditorFunctionParameterExpression](./object-QueryEditorFunctionParameterExpression.md)
 * <span class="badge object-type-struct"></span> [QueryEditorGroupByExpression](./object-QueryEditorGroupByExpression.md)
 * <span class="badge object-type-struct"></span> [QueryEditorProperty](./object-QueryEditorProperty.md)
 * <span class="badge object-type-struct"></span> [QueryEditorPropertyExpression](./object-QueryEditorPropertyExpression.md)
 * <span class="badge object-type-enum"></span> [QueryEditorPropertyType](./object-QueryEditorPropertyType.md)
 * <span class="badge object-type-enum"></span> [QueryFormat](./object-QueryFormat.md)
 * <span class="badge object-type-enum"></span> [QueryPriority](./object-QueryPriority.md)
 * <span class="badge object-type-struct"></span> [SQLExpression](./object-SQLExpression.md)
## Builders

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
 * <span class="badge builder"></span> [QueryEditorFunctionExpressionBuilder](./builder-QueryEditorFunctionExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorFunctionParameterExpressionBuilder](./builder-QueryEditorFunctionParameterExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorGroupByExpressionBuilder](./builder-QueryEditorGroupByExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorPropertyBuilder](./builder-QueryEditorPropertyBuilder.md)
 * <span class="badge builder"></span> [QueryEditorPropertyExpressionBuilder](./builder-QueryEditorPropertyExpressionBuilder.md)
 * <span class="badge builder"></span> [SQLExpressionBuilder](./builder-SQLExpressionBuilder.md)
## Functions

### <span class="badge function"></span> NewSQLExpression

NewSQLExpression creates a new SQLExpression object.

```go
func NewSQLExpression() *SQLExpression
```

### <span class="badge function"></span> NewQueryEditorFunctionExpression

NewQueryEditorFunctionExpression creates a new QueryEditorFunctionExpression object.

```go
func NewQueryEditorFunctionExpression() *QueryEditorFunctionExpression
```

### <span class="badge function"></span> NewQueryEditorFunctionParameterExpression

NewQueryEditorFunctionParameterExpression creates a new QueryEditorFunctionParameterExpression object.

```go
func NewQueryEditorFunctionParameterExpression() *QueryEditorFunctionParameterExpression
```

### <span class="badge function"></span> NewQueryEditorGroupByExpression

NewQueryEditorGroupByExpression creates a new QueryEditorGroupByExpression object.

```go
func NewQueryEditorGroupByExpression() *QueryEditorGroupByExpression
```

### <span class="badge function"></span> NewQueryEditorProperty

NewQueryEditorProperty creates a new QueryEditorProperty object.

```go
func NewQueryEditorProperty() *QueryEditorProperty
```

### <span class="badge function"></span> NewQueryEditorPropertyExpression

NewQueryEditorPropertyExpression creates a new QueryEditorPropertyExpression object.

```go
func NewQueryEditorPropertyExpression() *QueryEditorPropertyExpression
```

### <span class="badge function"></span> NewDataquery

NewDataquery creates a new Dataquery object.

```go
func NewDataquery() *Dataquery
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to grafana-bigquery-datasource dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> SQLExpressionConverter

SQLExpressionConverter accepts a `SQLExpression` object and generates the Go code to build this object using builders.

```go
func SQLExpressionConverter(input SQLExpression) string
```

### <span class="badge function"></span> QueryEditorFunctionExpressionConverter

QueryEditorFunctionExpressionConverter accepts a `QueryEditorFunctionExpression` object and generates the Go code to build this object using builders.

```go
func QueryEditorFunctionExpressionConverter(input QueryEditorFunctionExpression) string
```

### <span class="badge function"></span> QueryEditorFunctionParameterExpressionConverter

QueryEditorFunctionParameterExpressionConverter accepts a `QueryEditorFunctionParameterExpression` object and generates the Go code to build this object using builders.

```go
func QueryEditorFunctionParameterExpressionConverter(input QueryEditorFunctionParameterExpression) string
```

### <span class="badge function"></span> QueryEditorGroupByExpressionConverter

QueryEditorGroupByExpressionConverter accepts a `QueryEditorGroupByExpression` object and generates the Go code to build this object using builders.

```go
func QueryEditorGroupByExpressionConverter(input QueryEditorGroupByExpression) string
```

### <span class="badge function"></span> QueryEditorPropertyConverter

QueryEditorPropertyConverter accepts a `QueryEditorProperty` object and generates the Go code to build this object using builders.

```go
func QueryEditorPropertyConverter(input QueryEditorProperty) string
```

### <span class="badge function"></span> QueryEditorPropertyExpressionConverter

QueryEditorPropertyExpressionConverter accepts a `QueryEditorPropertyExpression` object and generates the Go code to build this object using builders.

```go
func QueryEditorPropertyExpressionConverter(input QueryEditorPropertyExpression) string
```

### <span class="badge function"></span> DataqueryConverter

DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.

```go
func DataqueryConverter(input Dataquery) string
```

