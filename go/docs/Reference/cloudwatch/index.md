# <span class="badge package-variant-dataquery"></span> cloudwatch

## Objects

 * <span class="badge object-type-struct"></span> [CloudWatchAnnotationQuery](./object-CloudWatchAnnotationQuery.md)
 * <span class="badge object-type-struct"></span> [CloudWatchLogsQuery](./object-CloudWatchLogsQuery.md)
 * <span class="badge object-type-struct"></span> [CloudWatchMetricsQuery](./object-CloudWatchMetricsQuery.md)
 * <span class="badge object-type-struct"></span> [CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery](./object-CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery.md)
 * <span class="badge object-type-ref"></span> [CloudWatchQuery](./object-CloudWatchQuery.md)
 * <span class="badge object-type-enum"></span> [CloudWatchQueryMode](./object-CloudWatchQueryMode.md)
 * <span class="badge object-type-map"></span> [Dimensions](./object-Dimensions.md)
 * <span class="badge object-type-struct"></span> [LogGroup](./object-LogGroup.md)
 * <span class="badge object-type-enum"></span> [MetricEditorMode](./object-MetricEditorMode.md)
 * <span class="badge object-type-enum"></span> [MetricQueryType](./object-MetricQueryType.md)
 * <span class="badge object-type-struct"></span> [MetricStat](./object-MetricStat.md)
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
 * <span class="badge object-type-struct"></span> [SQLExpression](./object-SQLExpression.md)
 * <span class="badge object-type-struct"></span> [StringOrArrayOfString](./object-StringOrArrayOfString.md)
 * <span class="badge object-type-struct"></span> [StringOrBoolOrInt64](./object-StringOrBoolOrInt64.md)
 * <span class="badge object-type-struct"></span> [StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType](./object-StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType.md)
## Builders

 * <span class="badge builder"></span> [CloudWatchAnnotationQueryBuilder](./builder-CloudWatchAnnotationQueryBuilder.md)
 * <span class="badge builder"></span> [CloudWatchLogsQueryBuilder](./builder-CloudWatchLogsQueryBuilder.md)
 * <span class="badge builder"></span> [CloudWatchMetricsQueryBuilder](./builder-CloudWatchMetricsQueryBuilder.md)
 * <span class="badge builder"></span> [LogGroupBuilder](./builder-LogGroupBuilder.md)
 * <span class="badge builder"></span> [MetricStatBuilder](./builder-MetricStatBuilder.md)
 * <span class="badge builder"></span> [QueryEditorArrayExpressionBuilder](./builder-QueryEditorArrayExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorFunctionExpressionBuilder](./builder-QueryEditorFunctionExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorFunctionParameterExpressionBuilder](./builder-QueryEditorFunctionParameterExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorGroupByExpressionBuilder](./builder-QueryEditorGroupByExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorOperatorBuilder](./builder-QueryEditorOperatorBuilder.md)
 * <span class="badge builder"></span> [QueryEditorOperatorExpressionBuilder](./builder-QueryEditorOperatorExpressionBuilder.md)
 * <span class="badge builder"></span> [QueryEditorPropertyBuilder](./builder-QueryEditorPropertyBuilder.md)
 * <span class="badge builder"></span> [QueryEditorPropertyExpressionBuilder](./builder-QueryEditorPropertyExpressionBuilder.md)
 * <span class="badge builder"></span> [SQLExpressionBuilder](./builder-SQLExpressionBuilder.md)
## Functions

### <span class="badge function"></span> NewMetricStat

NewMetricStat creates a new MetricStat object.

```go
func NewMetricStat() *MetricStat
```

### <span class="badge function"></span> NewCloudWatchMetricsQuery

NewCloudWatchMetricsQuery creates a new CloudWatchMetricsQuery object.

```go
func NewCloudWatchMetricsQuery() *CloudWatchMetricsQuery
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

### <span class="badge function"></span> NewQueryEditorOperatorValueType

NewQueryEditorOperatorValueType creates a new QueryEditorOperatorValueType object.

```go
func NewQueryEditorOperatorValueType() *QueryEditorOperatorValueType
```

### <span class="badge function"></span> NewQueryEditorOperatorType

NewQueryEditorOperatorType creates a new QueryEditorOperatorType object.

```go
func NewQueryEditorOperatorType() *QueryEditorOperatorType
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

### <span class="badge function"></span> NewCloudWatchLogsQuery

NewCloudWatchLogsQuery creates a new CloudWatchLogsQuery object.

```go
func NewCloudWatchLogsQuery() *CloudWatchLogsQuery
```

### <span class="badge function"></span> NewLogGroup

NewLogGroup creates a new LogGroup object.

```go
func NewLogGroup() *LogGroup
```

### <span class="badge function"></span> NewCloudWatchAnnotationQuery

NewCloudWatchAnnotationQuery creates a new CloudWatchAnnotationQuery object.

```go
func NewCloudWatchAnnotationQuery() *CloudWatchAnnotationQuery
```

### <span class="badge function"></span> NewCloudWatchQuery

NewCloudWatchQuery creates a new CloudWatchQuery object.

```go
func NewCloudWatchQuery() *CloudWatchQuery
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to cloudwatch dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
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

### <span class="badge function"></span> NewQueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression

NewQueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression creates a new QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression object.

```go
func NewQueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression() *QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression
```

### <span class="badge function"></span> NewCloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery

NewCloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery creates a new CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery object.

```go
func NewCloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery() *CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery
```

### <span class="badge function"></span> MetricStatConverter

MetricStatConverter accepts a `MetricStat` object and generates the Go code to build this object using builders.

```go
func MetricStatConverter(input MetricStat) string
```

### <span class="badge function"></span> CloudWatchMetricsQueryConverter

CloudWatchMetricsQueryConverter accepts a `CloudWatchMetricsQuery` object and generates the Go code to build this object using builders.

```go
func CloudWatchMetricsQueryConverter(input CloudWatchMetricsQuery) string
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

### <span class="badge function"></span> CloudWatchLogsQueryConverter

CloudWatchLogsQueryConverter accepts a `CloudWatchLogsQuery` object and generates the Go code to build this object using builders.

```go
func CloudWatchLogsQueryConverter(input CloudWatchLogsQuery) string
```

### <span class="badge function"></span> LogGroupConverter

LogGroupConverter accepts a `LogGroup` object and generates the Go code to build this object using builders.

```go
func LogGroupConverter(input LogGroup) string
```

### <span class="badge function"></span> CloudWatchAnnotationQueryConverter

CloudWatchAnnotationQueryConverter accepts a `CloudWatchAnnotationQuery` object and generates the Go code to build this object using builders.

```go
func CloudWatchAnnotationQueryConverter(input CloudWatchAnnotationQuery) string
```

