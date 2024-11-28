---
title: <span class="badge builder"></span> CloudWatchAnnotationQueryBuilder
---
# <span class="badge builder"></span> CloudWatchAnnotationQueryBuilder

## Constructor

```go
func NewCloudWatchAnnotationQueryBuilder() *CloudWatchAnnotationQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CloudWatchAnnotationQueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> AccountId

The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.

```go
func (builder *CloudWatchAnnotationQueryBuilder) AccountId(accountId string) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> ActionPrefix

Use this parameter to filter the results of the operation to only those alarms

that use a certain alarm action. For example, you could specify the ARN of

an SNS topic to find all alarms that send notifications to that topic.

e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`

but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`

```go
func (builder *CloudWatchAnnotationQueryBuilder) ActionPrefix(actionPrefix string) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> AlarmNamePrefix

An alarm name prefix. If you specify this parameter, you receive information

about all alarms that have names that start with this prefix.

e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`

```go
func (builder *CloudWatchAnnotationQueryBuilder) AlarmNamePrefix(alarmNamePrefix string) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *CloudWatchAnnotationQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> Dimensions

The dimensions of the metric

```go
func (builder *CloudWatchAnnotationQueryBuilder) Dimensions(dimensions cloudwatch.Dimensions) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *CloudWatchAnnotationQueryBuilder) Hide(hide bool) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> MatchExact

Only show metrics that exactly match all defined dimension names.

```go
func (builder *CloudWatchAnnotationQueryBuilder) MatchExact(matchExact bool) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> MetricName

Name of the metric

```go
func (builder *CloudWatchAnnotationQueryBuilder) MetricName(metricName string) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> Namespace

A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.

```go
func (builder *CloudWatchAnnotationQueryBuilder) Namespace(namespace string) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> Period

The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes

```go
func (builder *CloudWatchAnnotationQueryBuilder) Period(period string) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> PrefixMatching

Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix

```go
func (builder *CloudWatchAnnotationQueryBuilder) PrefixMatching(prefixMatching bool) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> QueryMode

Whether a query is a Metrics, Logs, or Annotations query

```go
func (builder *CloudWatchAnnotationQueryBuilder) QueryMode(queryMode cloudwatch.CloudWatchQueryMode) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *CloudWatchAnnotationQueryBuilder) QueryType(queryType string) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *CloudWatchAnnotationQueryBuilder) RefId(refId string) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> Region

AWS region to query for the metric

```go
func (builder *CloudWatchAnnotationQueryBuilder) Region(region string) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> Statistic

Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.

```go
func (builder *CloudWatchAnnotationQueryBuilder) Statistic(statistic string) *CloudWatchAnnotationQueryBuilder
```

### <span class="badge object-method"></span> Statistics

@deprecated use statistic

```go
func (builder *CloudWatchAnnotationQueryBuilder) Statistics(statistics []string) *CloudWatchAnnotationQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CloudWatchAnnotationQuery](./object-CloudWatchAnnotationQuery.md)