# <span class="badge package-variant-dataquery"></span> cloudwatch

## Objects

 * <span class="badge object-type-struct"></span> [AnnotationQuery](./object-AnnotationQuery.md)
 * <span class="badge object-type-map"></span> [Dimensions](./object-Dimensions.md)
 * <span class="badge object-type-struct"></span> [LogGroup](./object-LogGroup.md)
 * <span class="badge object-type-struct"></span> [LogsQuery](./object-LogsQuery.md)
 * <span class="badge object-type-enum"></span> [LogsQueryLanguage](./object-LogsQueryLanguage.md)
 * <span class="badge object-type-enum"></span> [MetricEditorMode](./object-MetricEditorMode.md)
 * <span class="badge object-type-enum"></span> [MetricQueryType](./object-MetricQueryType.md)
 * <span class="badge object-type-struct"></span> [MetricStat](./object-MetricStat.md)
 * <span class="badge object-type-struct"></span> [MetricsQuery](./object-MetricsQuery.md)
 * <span class="badge object-type-struct"></span> [MetricsQueryOrLogsQueryOrAnnotationQuery](./object-MetricsQueryOrLogsQueryOrAnnotationQuery.md)
 * <span class="badge object-type-struct"></span> [QueryEditorArrayExpression](./object-QueryEditorArrayExpression.md)
 * <span class="badge object-type-struct"></span> [QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression](./object-QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression.md)
 * <span class="badge object-type-enum"></span> [QueryEditorArrayExpressionType](./object-QueryEditorArrayExpressionType.md)
 * <span class="badge object-type-ref"></span> [QueryEditorExpression](./object-QueryEditorExpression.md)
 * <span class="badge object-type-enum"></span> [QueryEditorExpressionType](./object-QueryEditorExpressionType.md)
 * <span class="badge object-type-struct"></span> [QueryEditorFunctionExpression](./object-QueryEditorFunctionExpression.md)
 * <span class="badge object-type-struct"></span> [QueryEditorFunctionParameterExpression](./object-QueryEditorFunctionParameterExpression.md)
 * <span class="badge object-type-struct"></span> [QueryEditorGroupByExpression](./object-QueryEditorGroupByExpression.md)
 * <span class="badge object-type-struct"></span> [QueryEditorOperator](./object-QueryEditorOperator.md)
 * <span class="badge object-type-struct"></span> [QueryEditorOperatorExpression](./object-QueryEditorOperatorExpression.md)
 * <span class="badge object-type-ref"></span> [QueryEditorOperatorType](./object-QueryEditorOperatorType.md)
 * <span class="badge object-type-ref"></span> [QueryEditorOperatorValueType](./object-QueryEditorOperatorValueType.md)
 * <span class="badge object-type-struct"></span> [QueryEditorProperty](./object-QueryEditorProperty.md)
 * <span class="badge object-type-struct"></span> [QueryEditorPropertyExpression](./object-QueryEditorPropertyExpression.md)
 * <span class="badge object-type-struct"></span> [QueryEditorPropertyExpressionOrQueryEditorFunctionExpression](./object-QueryEditorPropertyExpressionOrQueryEditorFunctionExpression.md)
 * <span class="badge object-type-enum"></span> [QueryEditorPropertyType](./object-QueryEditorPropertyType.md)
 * <span class="badge object-type-enum"></span> [QueryMode](./object-QueryMode.md)
 * <span class="badge object-type-ref"></span> [Request](./object-Request.md)
 * <span class="badge object-type-struct"></span> [SQLExpression](./object-SQLExpression.md)
 * <span class="badge object-type-struct"></span> [StringOrArrayOfString](./object-StringOrArrayOfString.md)
 * <span class="badge object-type-struct"></span> [StringOrBoolOrInt64](./object-StringOrBoolOrInt64.md)
 * <span class="badge object-type-struct"></span> [StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType](./object-StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType.md)
## Builders

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
 * <span class="badge builder"></span> [LogGroupBuilder](./builder-LogGroupBuilder.md)
 * <span class="badge builder"></span> [LogsQueryBuilder](./builder-LogsQueryBuilder.md)
 * <span class="badge builder"></span> [MetricStatBuilder](./builder-MetricStatBuilder.md)
 * <span class="badge builder"></span> [MetricsQueryBuilder](./builder-MetricsQueryBuilder.md)
 * <span class="badge builder"></span> [MetricsQueryOrLogsQueryOrAnnotationQueryBuilder](./builder-MetricsQueryOrLogsQueryOrAnnotationQueryBuilder.md)
 * <span class="badge builder"></span> [QueryBuilder](./builder-QueryBuilder.md)
 * <span class="badge builder"></span> [QueryEditorArrayExpressionBuilder](./builder-QueryEditorArrayExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorFunctionExpressionBuilder](./builder-QueryEditorFunctionExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorFunctionParameterExpressionBuilder](./builder-QueryEditorFunctionParameterExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorGroupByExpressionBuilder](./builder-QueryEditorGroupByExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorOperatorBuilder](./builder-QueryEditorOperatorBuilder.md)
 * <span class="badge builder"></span> [QueryEditorOperatorExpressionBuilder](./builder-QueryEditorOperatorExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorPropertyBuilder](./builder-QueryEditorPropertyBuilder.md)
 * <span class="badge builder"></span> [QueryEditorPropertyExpressionBuilder](./builder-QueryEditorPropertyExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryV2Builder](./builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [RequestBuilder](./builder-RequestBuilder.md)
 * <span class="badge builder"></span> [SQLExpressionBuilder](./builder-SQLExpressionBuilder.md)
## Functions

### <span class="badge function"></span> NewMetricStat

NewMetricStat creates a new MetricStat object.

```go
func NewMetricStat() *MetricStat
```

### <span class="badge function"></span> NewMetricsQuery

NewMetricsQuery creates a new MetricsQuery object.

```go
func NewMetricsQuery() *MetricsQuery
```

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

### <span class="badge function"></span> NewQueryEditorPropertyExpression

NewQueryEditorPropertyExpression creates a new QueryEditorPropertyExpression object.

```go
func NewQueryEditorPropertyExpression() *QueryEditorPropertyExpression
```

### <span class="badge function"></span> NewQueryEditorProperty

NewQueryEditorProperty creates a new QueryEditorProperty object.

```go
func NewQueryEditorProperty() *QueryEditorProperty
```

### <span class="badge function"></span> NewQueryEditorArrayExpression

NewQueryEditorArrayExpression creates a new QueryEditorArrayExpression object.

```go
func NewQueryEditorArrayExpression() *QueryEditorArrayExpression
```

### <span class="badge function"></span> NewQueryEditorExpression

NewQueryEditorExpression creates a new QueryEditorExpression object.

```go
func NewQueryEditorExpression() *QueryEditorExpression
```

### <span class="badge function"></span> NewQueryEditorGroupByExpression

NewQueryEditorGroupByExpression creates a new QueryEditorGroupByExpression object.

```go
func NewQueryEditorGroupByExpression() *QueryEditorGroupByExpression
```

### <span class="badge function"></span> NewQueryEditorOperatorExpression

NewQueryEditorOperatorExpression creates a new QueryEditorOperatorExpression object.

```go
func NewQueryEditorOperatorExpression() *QueryEditorOperatorExpression
```

### <span class="badge function"></span> NewQueryEditorOperator

NewQueryEditorOperator creates a new QueryEditorOperator object.

```go
func NewQueryEditorOperator() *QueryEditorOperator
```

### <span class="badge function"></span> NewQueryEditorOperatorType

NewQueryEditorOperatorType creates a new QueryEditorOperatorType object.

```go
func NewQueryEditorOperatorType() *QueryEditorOperatorType
```

### <span class="badge function"></span> NewQueryEditorOperatorValueType

NewQueryEditorOperatorValueType creates a new QueryEditorOperatorValueType object.

```go
func NewQueryEditorOperatorValueType() *QueryEditorOperatorValueType
```

### <span class="badge function"></span> NewLogsQuery

NewLogsQuery creates a new LogsQuery object.

```go
func NewLogsQuery() *LogsQuery
```

### <span class="badge function"></span> NewLogGroup

NewLogGroup creates a new LogGroup object.

```go
func NewLogGroup() *LogGroup
```

### <span class="badge function"></span> NewAnnotationQuery

NewAnnotationQuery creates a new AnnotationQuery object.

```go
func NewAnnotationQuery() *AnnotationQuery
```

### <span class="badge function"></span> NewRequest

NewRequest creates a new Request object.

```go
func NewRequest() *Request
```

### <span class="badge function"></span> NewStringOrArrayOfString

NewStringOrArrayOfString creates a new StringOrArrayOfString object.

```go
func NewStringOrArrayOfString() *StringOrArrayOfString
```

### <span class="badge function"></span> NewQueryEditorPropertyExpressionOrQueryEditorFunctionExpression

NewQueryEditorPropertyExpressionOrQueryEditorFunctionExpression creates a new QueryEditorPropertyExpressionOrQueryEditorFunctionExpression object.

```go
func NewQueryEditorPropertyExpressionOrQueryEditorFunctionExpression() *QueryEditorPropertyExpressionOrQueryEditorFunctionExpression
```

### <span class="badge function"></span> NewQueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression

NewQueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression creates a new QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression object.

```go
func NewQueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression() *QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression
```

### <span class="badge function"></span> NewStringOrBoolOrInt64OrArrayOfQueryEditorOperatorType

NewStringOrBoolOrInt64OrArrayOfQueryEditorOperatorType creates a new StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType object.

```go
func NewStringOrBoolOrInt64OrArrayOfQueryEditorOperatorType() *StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType
```

### <span class="badge function"></span> NewStringOrBoolOrInt64

NewStringOrBoolOrInt64 creates a new StringOrBoolOrInt64 object.

```go
func NewStringOrBoolOrInt64() *StringOrBoolOrInt64
```

### <span class="badge function"></span> NewMetricsQueryOrLogsQueryOrAnnotationQuery

NewMetricsQueryOrLogsQueryOrAnnotationQuery creates a new MetricsQueryOrLogsQueryOrAnnotationQuery object.

```go
func NewMetricsQueryOrLogsQueryOrAnnotationQuery() *MetricsQueryOrLogsQueryOrAnnotationQuery
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to cloudwatch dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> QueryV2Converter

QueryV2Converter accepts a `QueryV2` object and generates the Go code to build this object using builders.

```go
func QueryV2Converter(input dashboardv2.DataQueryKind) string
```

### <span class="badge function"></span> MetricStatConverter

MetricStatConverter accepts a `MetricStat` object and generates the Go code to build this object using builders.

```go
func MetricStatConverter(input MetricStat) string
```

### <span class="badge function"></span> MetricsQueryConverter

MetricsQueryConverter accepts a `MetricsQuery` object and generates the Go code to build this object using builders.

```go
func MetricsQueryConverter(input MetricsQuery) string
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

### <span class="badge function"></span> QueryEditorPropertyExpressionConverter

QueryEditorPropertyExpressionConverter accepts a `QueryEditorPropertyExpression` object and generates the Go code to build this object using builders.

```go
func QueryEditorPropertyExpressionConverter(input QueryEditorPropertyExpression) string
```

### <span class="badge function"></span> QueryEditorPropertyConverter

QueryEditorPropertyConverter accepts a `QueryEditorProperty` object and generates the Go code to build this object using builders.

```go
func QueryEditorPropertyConverter(input QueryEditorProperty) string
```

### <span class="badge function"></span> QueryEditorArrayExpressionConverter

QueryEditorArrayExpressionConverter accepts a `QueryEditorArrayExpression` object and generates the Go code to build this object using builders.

```go
func QueryEditorArrayExpressionConverter(input QueryEditorArrayExpression) string
```

### <span class="badge function"></span> QueryEditorGroupByExpressionConverter

QueryEditorGroupByExpressionConverter accepts a `QueryEditorGroupByExpression` object and generates the Go code to build this object using builders.

```go
func QueryEditorGroupByExpressionConverter(input QueryEditorGroupByExpression) string
```

### <span class="badge function"></span> QueryEditorOperatorExpressionConverter

QueryEditorOperatorExpressionConverter accepts a `QueryEditorOperatorExpression` object and generates the Go code to build this object using builders.

```go
func QueryEditorOperatorExpressionConverter(input QueryEditorOperatorExpression) string
```

### <span class="badge function"></span> QueryEditorOperatorConverter

QueryEditorOperatorConverter accepts a `QueryEditorOperator` object and generates the Go code to build this object using builders.

```go
func QueryEditorOperatorConverter(input QueryEditorOperator) string
```

### <span class="badge function"></span> LogsQueryConverter

LogsQueryConverter accepts a `LogsQuery` object and generates the Go code to build this object using builders.

```go
func LogsQueryConverter(input LogsQuery) string
```

### <span class="badge function"></span> LogGroupConverter

LogGroupConverter accepts a `LogGroup` object and generates the Go code to build this object using builders.

```go
func LogGroupConverter(input LogGroup) string
```

### <span class="badge function"></span> AnnotationQueryConverter

AnnotationQueryConverter accepts a `AnnotationQuery` object and generates the Go code to build this object using builders.

```go
func AnnotationQueryConverter(input AnnotationQuery) string
```

### <span class="badge function"></span> RequestConverter

RequestConverter accepts a `Request` object and generates the Go code to build this object using builders.

```go
func RequestConverter(input Request) string
```

### <span class="badge function"></span> MetricsQueryOrLogsQueryOrAnnotationQueryConverter

MetricsQueryOrLogsQueryOrAnnotationQueryConverter accepts a `MetricsQueryOrLogsQueryOrAnnotationQuery` object and generates the Go code to build this object using builders.

```go
func MetricsQueryOrLogsQueryOrAnnotationQueryConverter(input MetricsQueryOrLogsQueryOrAnnotationQuery) string
```

### <span class="badge function"></span> QueryConverter

QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.

```go
func QueryConverter(input dashboardv2beta1.DataQueryKind) string
```

