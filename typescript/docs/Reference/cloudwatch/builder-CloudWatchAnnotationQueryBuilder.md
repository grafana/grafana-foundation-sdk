---
title: <span class="badge builder"></span> CloudWatchAnnotationQueryBuilder
---
# <span class="badge builder"></span> CloudWatchAnnotationQueryBuilder

## Constructor

```typescript
new CloudWatchAnnotationQueryBuilder()
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

### <span class="badge object-method"></span> actionPrefix

Use this parameter to filter the results of the operation to only those alarms

that use a certain alarm action. For example, you could specify the ARN of

an SNS topic to find all alarms that send notifications to that topic.

e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`

but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`

```typescript
actionPrefix(actionPrefix: string)
```

### <span class="badge object-method"></span> alarmNamePrefix

An alarm name prefix. If you specify this parameter, you receive information

about all alarms that have names that start with this prefix.

e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`

```typescript
alarmNamePrefix(alarmNamePrefix: string)
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

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> matchExact

Only show metrics that exactly match all defined dimension names.

```typescript
matchExact(matchExact: boolean)
```

### <span class="badge object-method"></span> metricName

Name of the metric

```typescript
metricName(metricName: string)
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

### <span class="badge object-method"></span> prefixMatching

Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix

```typescript
prefixMatching(prefixMatching: boolean)
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

 * <span class="badge object-type-interface"></span> [CloudWatchAnnotationQuery](./object-CloudWatchAnnotationQuery.md)