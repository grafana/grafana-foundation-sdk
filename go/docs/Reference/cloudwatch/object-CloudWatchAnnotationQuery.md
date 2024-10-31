---
title: <span class="badge object-type-struct"></span> CloudWatchAnnotationQuery
---
# <span class="badge object-type-struct"></span> CloudWatchAnnotationQuery

Shape of a CloudWatch Annotation query

TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer

#CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")

## Definition

```go
type CloudWatchAnnotationQuery struct {
    // Whether a query is a Metrics, Logs, or Annotations query
    QueryMode cloudwatch.CloudWatchQueryMode `json:"queryMode"`
    // Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
    PrefixMatching *bool `json:"prefixMatching,omitempty"`
    // Use this parameter to filter the results of the operation to only those alarms
    // that use a certain alarm action. For example, you could specify the ARN of
    // an SNS topic to find all alarms that send notifications to that topic.
    // e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
    // but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
    ActionPrefix *string `json:"actionPrefix,omitempty"`
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
    // An alarm name prefix. If you specify this parameter, you receive information
    // about all alarms that have names that start with this prefix.
    // e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
    AlarmNamePrefix *string `json:"alarmNamePrefix,omitempty"`
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

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CloudWatchAnnotationQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (cloudWatchAnnotationQuery *CloudWatchAnnotationQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CloudWatchAnnotationQuery` objects.

```go
func (cloudWatchAnnotationQuery *CloudWatchAnnotationQuery) Equals(other CloudWatchAnnotationQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CloudWatchAnnotationQuery` fields for violations and returns them.

```go
func (cloudWatchAnnotationQuery *CloudWatchAnnotationQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [CloudWatchAnnotationQueryBuilder](./builder-CloudWatchAnnotationQueryBuilder.md)
