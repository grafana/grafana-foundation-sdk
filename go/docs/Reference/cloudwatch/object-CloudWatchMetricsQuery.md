---
title: <span class="badge object-type-struct"></span> CloudWatchMetricsQuery
---
# <span class="badge object-type-struct"></span> CloudWatchMetricsQuery

Shape of a CloudWatch Metrics query

## Definition

```go
type CloudWatchMetricsQuery struct {
    // Whether a query is a Metrics, Logs, or Annotations query
    QueryMode cloudwatch.CloudWatchQueryMode `json:"queryMode"`
    // Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.
    MetricQueryType *cloudwatch.MetricQueryType `json:"metricQueryType,omitempty"`
    // Whether to use the query builder or code editor to create the query
    MetricEditorMode *cloudwatch.MetricEditorMode `json:"metricEditorMode,omitempty"`
    // ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
    Id string `json:"id"`
    // Deprecated: use label
    // @deprecated use label
    Alias *string `json:"alias,omitempty"`
    // Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
    Label *string `json:"label,omitempty"`
    // Math expression query
    Expression *string `json:"expression,omitempty"`
    // When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
    SqlExpression *string `json:"sqlExpression,omitempty"`
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    RefId string `json:"refId"`
    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    Hide *bool `json:"hide,omitempty"`
    // Specify the query flavor
    // TODO make this required and give it a default
    QueryType *string `json:"queryType,omitempty"`
    // AWS region to query for the metric
    Region string `json:"region"`
    // A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
    Namespace string `json:"namespace"`
    // Name of the metric
    MetricName *string `json:"metricName,omitempty"`
    // The dimensions of the metric
    Dimensions *cloudwatch.Dimensions `json:"dimensions,omitempty"`
    // Only show metrics that exactly match all defined dimension names.
    MatchExact *bool `json:"matchExact,omitempty"`
    // The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
    Period *string `json:"period,omitempty"`
    // The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
    AccountId *string `json:"accountId,omitempty"`
    // Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
    Statistic *string `json:"statistic,omitempty"`
    // When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
    Sql *cloudwatch.SQLExpression `json:"sql,omitempty"`
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
    // @deprecated use statistic
    Statistics []string `json:"statistics,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CloudWatchMetricsQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (cloudWatchMetricsQuery *CloudWatchMetricsQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CloudWatchMetricsQuery` objects.

```go
func (cloudWatchMetricsQuery *CloudWatchMetricsQuery) Equals(other CloudWatchMetricsQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CloudWatchMetricsQuery` fields for violations and returns them.

```go
func (cloudWatchMetricsQuery *CloudWatchMetricsQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [CloudWatchMetricsQueryBuilder](./builder-CloudWatchMetricsQueryBuilder.md)
