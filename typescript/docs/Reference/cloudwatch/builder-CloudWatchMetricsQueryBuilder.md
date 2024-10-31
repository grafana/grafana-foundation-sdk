---
title: <span class="badge builder"></span> CloudWatchMetricsQueryBuilder
---
# <span class="badge builder"></span> CloudWatchMetricsQueryBuilder

## Constructor

```typescript
new CloudWatchMetricsQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> accountId

The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.

```typescript
accountId(accountId: string)
```

### <span class="badge object-method"></span> alias

Deprecated: use label

@deprecated use label

```typescript
alias(alias: string)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```typescript
datasource(datasource: dashboard.DataSourceRef)
```

### <span class="badge object-method"></span> dimensions

The dimensions of the metric

```typescript
dimensions(dimensions: cloudwatch.Dimensions)
```

### <span class="badge object-method"></span> expression

Math expression query

```typescript
expression(expression: string)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> id

ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.

```typescript
id(id: string)
```

### <span class="badge object-method"></span> label

Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.

```typescript
label(label: string)
```

### <span class="badge object-method"></span> matchExact

Only show metrics that exactly match all defined dimension names.

```typescript
matchExact(matchExact: boolean)
```

### <span class="badge object-method"></span> metricEditorMode

Whether to use the query builder or code editor to create the query

```typescript
metricEditorMode(metricEditorMode: cloudwatch.MetricEditorMode)
```

### <span class="badge object-method"></span> metricName

Name of the metric

```typescript
metricName(metricName: string)
```

### <span class="badge object-method"></span> metricQueryType

Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.

```typescript
metricQueryType(metricQueryType: cloudwatch.MetricQueryType)
```

### <span class="badge object-method"></span> namespace

A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.

```typescript
namespace(namespace: string)
```

### <span class="badge object-method"></span> period

The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes

```typescript
period(period: string)
```

### <span class="badge object-method"></span> queryMode

Whether a query is a Metrics, Logs, or Annotations query

```typescript
queryMode(queryMode: cloudwatch.CloudWatchQueryMode)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```typescript
refId(refId: string)
```

### <span class="badge object-method"></span> region

AWS region to query for the metric

```typescript
region(region: string)
```

### <span class="badge object-method"></span> sql

When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.

```typescript
sql(sql: cog.Builder<cloudwatch.SQLExpression>)
```

### <span class="badge object-method"></span> sqlExpression

When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.

```typescript
sqlExpression(sqlExpression: string)
```

### <span class="badge object-method"></span> statistic

Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.

```typescript
statistic(statistic: string)
```

### <span class="badge object-method"></span> statistics

@deprecated use statistic

```typescript
statistics(statistics: string[])
```

## See also

 * <span class="badge object-type-interface"></span> [CloudWatchMetricsQuery](./object-CloudWatchMetricsQuery.md)
