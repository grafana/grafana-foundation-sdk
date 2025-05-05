---
title: <span class="badge builder"></span> CloudWatchMetricsQueryBuilder
---
# <span class="badge builder"></span> CloudWatchMetricsQueryBuilder

## Constructor

```java
new CloudWatchMetricsQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public CloudWatchMetricsQuery build()
```

### <span class="badge object-method"></span> accountId

The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.

```java
public CloudWatchMetricsQueryBuilder accountId(String accountId)
```

### <span class="badge object-method"></span> alias

Deprecated: use label

@deprecated use label

```java
public CloudWatchMetricsQueryBuilder alias(String alias)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public CloudWatchMetricsQueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> dimensions

The dimensions of the metric

```java
public CloudWatchMetricsQueryBuilder dimensions(Map<String, StringOrArrayOfString> dimensions)
```

### <span class="badge object-method"></span> expression

Math expression query

```java
public CloudWatchMetricsQueryBuilder expression(String expression)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public CloudWatchMetricsQueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> id

ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.

```java
public CloudWatchMetricsQueryBuilder id(String id)
```

### <span class="badge object-method"></span> label

Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.

```java
public CloudWatchMetricsQueryBuilder label(String label)
```

### <span class="badge object-method"></span> matchExact

Only show metrics that exactly match all defined dimension names.

```java
public CloudWatchMetricsQueryBuilder matchExact(Boolean matchExact)
```

### <span class="badge object-method"></span> metricEditorMode

Whether to use the query builder or code editor to create the query

```java
public CloudWatchMetricsQueryBuilder metricEditorMode(MetricEditorMode metricEditorMode)
```

### <span class="badge object-method"></span> metricName

Name of the metric

```java
public CloudWatchMetricsQueryBuilder metricName(String metricName)
```

### <span class="badge object-method"></span> metricQueryType

Whether to use a metric search or metric insights query

```java
public CloudWatchMetricsQueryBuilder metricQueryType(MetricQueryType metricQueryType)
```

### <span class="badge object-method"></span> namespace

A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.

```java
public CloudWatchMetricsQueryBuilder namespace(String namespace)
```

### <span class="badge object-method"></span> period

The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes

```java
public CloudWatchMetricsQueryBuilder period(String period)
```

### <span class="badge object-method"></span> queryMode

Whether a query is a Metrics, Logs, or Annotations query

```java
public CloudWatchMetricsQueryBuilder queryMode(CloudWatchQueryMode queryMode)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public CloudWatchMetricsQueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public CloudWatchMetricsQueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> region

AWS region to query for the metric

```java
public CloudWatchMetricsQueryBuilder region(String region)
```

### <span class="badge object-method"></span> sql

When the metric query type is set to `Insights` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.

```java
public CloudWatchMetricsQueryBuilder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql)
```

### <span class="badge object-method"></span> sqlExpression

When the metric query type is set to `Insights`, this field is used to specify the query string.

```java
public CloudWatchMetricsQueryBuilder sqlExpression(String sqlExpression)
```

### <span class="badge object-method"></span> statistic

Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.

```java
public CloudWatchMetricsQueryBuilder statistic(String statistic)
```

### <span class="badge object-method"></span> statistics

@deprecated use statistic

```java
public CloudWatchMetricsQueryBuilder statistics(List<String> statistics)
```

## See also

 * <span class="badge object-type-class"></span> [CloudWatchMetricsQuery](./object-CloudWatchMetricsQuery.md)
