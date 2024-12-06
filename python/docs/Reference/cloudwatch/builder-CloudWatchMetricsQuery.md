---
title: <span class="badge builder"></span> CloudWatchMetricsQuery
---
# <span class="badge builder"></span> CloudWatchMetricsQuery

## Constructor

```python
CloudWatchMetricsQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> cloudwatch.CloudWatchMetricsQuery
```

### <span class="badge object-method"></span> account_id

The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.

```python
def account_id(account_id: str) -> typing.Self
```

### <span class="badge object-method"></span> alias

Deprecated: use label

@deprecated use label

```python
def alias(alias: str) -> typing.Self
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

### <span class="badge object-method"></span> expression

Math expression query

```python
def expression(expression: str) -> typing.Self
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> id_val

ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.

```python
def id_val(id_val: str) -> typing.Self
```

### <span class="badge object-method"></span> label

Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.

```python
def label(label: str) -> typing.Self
```

### <span class="badge object-method"></span> match_exact

Only show metrics that exactly match all defined dimension names.

```python
def match_exact(match_exact: bool) -> typing.Self
```

### <span class="badge object-method"></span> metric_editor_mode

Whether to use the query builder or code editor to create the query

```python
def metric_editor_mode(metric_editor_mode: cloudwatch.MetricEditorMode) -> typing.Self
```

### <span class="badge object-method"></span> metric_name

Name of the metric

```python
def metric_name(metric_name: str) -> typing.Self
```

### <span class="badge object-method"></span> metric_query_type

Whether to use a metric search or metric insights query

```python
def metric_query_type(metric_query_type: cloudwatch.MetricQueryType) -> typing.Self
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

### <span class="badge object-method"></span> sql

When the metric query type is set to `Insights` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.

```python
def sql(sql: cogbuilder.Builder[cloudwatch.SQLExpression]) -> typing.Self
```

### <span class="badge object-method"></span> sql_expression

When the metric query type is set to `Insights`, this field is used to specify the query string.

```python
def sql_expression(sql_expression: str) -> typing.Self
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

 * <span class="badge object-type-class"></span> [CloudWatchMetricsQuery](./object-CloudWatchMetricsQuery.md)
