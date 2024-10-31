---
title: <span class="badge builder"></span> CloudWatchMetricsQueryBuilder
---
# <span class="badge builder"></span> CloudWatchMetricsQueryBuilder

## Constructor

```go
func NewCloudWatchMetricsQueryBuilder() *CloudWatchMetricsQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CloudWatchMetricsQueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> AccountId

The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.

```go
func (builder *CloudWatchMetricsQueryBuilder) AccountId(accountId string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Alias

Deprecated: use label

@deprecated use label

```go
func (builder *CloudWatchMetricsQueryBuilder) Alias(alias string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *CloudWatchMetricsQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Dimensions

The dimensions of the metric

```go
func (builder *CloudWatchMetricsQueryBuilder) Dimensions(dimensions cloudwatch.Dimensions) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Expression

Math expression query

```go
func (builder *CloudWatchMetricsQueryBuilder) Expression(expression string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *CloudWatchMetricsQueryBuilder) Hide(hide bool) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Id

ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.

```go
func (builder *CloudWatchMetricsQueryBuilder) Id(id string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Label

Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.

```go
func (builder *CloudWatchMetricsQueryBuilder) Label(label string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> MatchExact

Only show metrics that exactly match all defined dimension names.

```go
func (builder *CloudWatchMetricsQueryBuilder) MatchExact(matchExact bool) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> MetricEditorMode

Whether to use the query builder or code editor to create the query

```go
func (builder *CloudWatchMetricsQueryBuilder) MetricEditorMode(metricEditorMode cloudwatch.MetricEditorMode) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> MetricName

Name of the metric

```go
func (builder *CloudWatchMetricsQueryBuilder) MetricName(metricName string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> MetricQueryType

Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.

```go
func (builder *CloudWatchMetricsQueryBuilder) MetricQueryType(metricQueryType cloudwatch.MetricQueryType) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Namespace

A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.

```go
func (builder *CloudWatchMetricsQueryBuilder) Namespace(namespace string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Period

The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes

```go
func (builder *CloudWatchMetricsQueryBuilder) Period(period string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> QueryMode

Whether a query is a Metrics, Logs, or Annotations query

```go
func (builder *CloudWatchMetricsQueryBuilder) QueryMode(queryMode cloudwatch.CloudWatchQueryMode) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *CloudWatchMetricsQueryBuilder) QueryType(queryType string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *CloudWatchMetricsQueryBuilder) RefId(refId string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Region

AWS region to query for the metric

```go
func (builder *CloudWatchMetricsQueryBuilder) Region(region string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Sql

When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.

```go
func (builder *CloudWatchMetricsQueryBuilder) Sql(sql cog.Builder[cloudwatch.SQLExpression]) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> SqlExpression

When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.

```go
func (builder *CloudWatchMetricsQueryBuilder) SqlExpression(sqlExpression string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Statistic

Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.

```go
func (builder *CloudWatchMetricsQueryBuilder) Statistic(statistic string) *CloudWatchMetricsQueryBuilder
```

### <span class="badge object-method"></span> Statistics

@deprecated use statistic

```go
func (builder *CloudWatchMetricsQueryBuilder) Statistics(statistics []string) *CloudWatchMetricsQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CloudWatchMetricsQuery](./object-CloudWatchMetricsQuery.md)
