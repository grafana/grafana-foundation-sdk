---
title: <span class="badge builder"></span> CloudWatchAnnotationQuery
---
# <span class="badge builder"></span> CloudWatchAnnotationQuery

## Constructor

```python
CloudWatchAnnotationQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> cloudwatch.CloudWatchAnnotationQuery
```

### <span class="badge object-method"></span> account_id

The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.

```python
def account_id(account_id: str) -> typing.Self
```

### <span class="badge object-method"></span> action_prefix

Use this parameter to filter the results of the operation to only those alarms

that use a certain alarm action. For example, you could specify the ARN of

an SNS topic to find all alarms that send notifications to that topic.

e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`

but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`

```python
def action_prefix(action_prefix: str) -> typing.Self
```

### <span class="badge object-method"></span> alarm_name_prefix

An alarm name prefix. If you specify this parameter, you receive information

about all alarms that have names that start with this prefix.

e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`

```python
def alarm_name_prefix(alarm_name_prefix: str) -> typing.Self
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> dimensions

The dimensions of the metric

```python
def dimensions(dimensions: cloudwatch.Dimensions) -> typing.Self
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> match_exact

Only show metrics that exactly match all defined dimension names.

```python
def match_exact(match_exact: bool) -> typing.Self
```

### <span class="badge object-method"></span> metric_name

Name of the metric

```python
def metric_name(metric_name: str) -> typing.Self
```

### <span class="badge object-method"></span> namespace

A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.

```python
def namespace(namespace: str) -> typing.Self
```

### <span class="badge object-method"></span> period

The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes

```python
def period(period: str) -> typing.Self
```

### <span class="badge object-method"></span> prefix_matching

Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix

```python
def prefix_matching(prefix_matching: bool) -> typing.Self
```

### <span class="badge object-method"></span> query_mode

Whether a query is a Metrics, Logs, or Annotations query

```python
def query_mode(query_mode: cloudwatch.CloudWatchQueryMode) -> typing.Self
```

### <span class="badge object-method"></span> query_type

Specify the query flavor

TODO make this required and give it a default

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> region

AWS region to query for the metric

```python
def region(region: str) -> typing.Self
```

### <span class="badge object-method"></span> statistic

Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.

```python
def statistic(statistic: str) -> typing.Self
```

### <span class="badge object-method"></span> statistics

@deprecated use statistic

```python
def statistics(statistics: list[str]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [CloudWatchAnnotationQuery](./object-CloudWatchAnnotationQuery.md)
