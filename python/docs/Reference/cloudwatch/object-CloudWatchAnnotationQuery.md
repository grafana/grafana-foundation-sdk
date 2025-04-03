---
title: <span class="badge object-type-class"></span> CloudWatchAnnotationQuery
---
# <span class="badge object-type-class"></span> CloudWatchAnnotationQuery

Shape of a CloudWatch Annotation query

TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer

#CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")

## Definition

```python
class CloudWatchAnnotationQuery(cogvariants.Dataquery):
    """
    Shape of a CloudWatch Annotation query
    TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
    #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
    """

    # Whether a query is a Metrics, Logs, or Annotations query
    query_mode: cloudwatch.CloudWatchQueryMode
    # Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
    prefix_matching: typing.Optional[bool]
    # Use this parameter to filter the results of the operation to only those alarms
    # that use a certain alarm action. For example, you could specify the ARN of
    # an SNS topic to find all alarms that send notifications to that topic.
    # e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
    # but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
    action_prefix: typing.Optional[str]
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # AWS region to query for the metric
    region: str
    # A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
    namespace: str
    # Name of the metric
    metric_name: typing.Optional[str]
    # The dimensions of the metric
    dimensions: typing.Optional[cloudwatch.Dimensions]
    # Only show metrics that exactly match all defined dimension names.
    match_exact: typing.Optional[bool]
    # The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
    period: typing.Optional[str]
    # The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
    account_id: typing.Optional[str]
    # Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
    statistic: typing.Optional[str]
    # An alarm name prefix. If you specify this parameter, you receive information
    # about all alarms that have names that start with this prefix.
    # e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
    alarm_name_prefix: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[dashboard.DataSourceRef]
    # @deprecated use statistic
    statistics: typing.Optional[list[str]]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [CloudWatchAnnotationQuery](./builder-CloudWatchAnnotationQuery.md)
