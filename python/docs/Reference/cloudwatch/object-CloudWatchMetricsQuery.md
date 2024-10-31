---
title: <span class="badge object-type-class"></span> CloudWatchMetricsQuery
---
# <span class="badge object-type-class"></span> CloudWatchMetricsQuery

Shape of a CloudWatch Metrics query

## Definition

```python
class CloudWatchMetricsQuery(cogvariants.Dataquery):
    """
    Shape of a CloudWatch Metrics query
    """

    # Whether a query is a Metrics, Logs, or Annotations query
    query_mode: cloudwatch.CloudWatchQueryMode
    # Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.
    metric_query_type: typing.Optional[cloudwatch.MetricQueryType]
    # Whether to use the query builder or code editor to create the query
    metric_editor_mode: typing.Optional[cloudwatch.MetricEditorMode]
    # ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
    id_val: str
    # Deprecated: use label
    # @deprecated use label
    alias: typing.Optional[str]
    # Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
    label: typing.Optional[str]
    # Math expression query
    expression: typing.Optional[str]
    # When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
    sql_expression: typing.Optional[str]
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # Note this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
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
    # When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
    sql: typing.Optional[cloudwatch.SQLExpression]
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

 * <span class="badge builder"></span> [CloudWatchMetricsQuery](./builder-CloudWatchMetricsQuery.md)
