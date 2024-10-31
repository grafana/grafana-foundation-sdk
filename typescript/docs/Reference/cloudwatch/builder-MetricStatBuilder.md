---
title: <span class="badge builder"></span> MetricStatBuilder
---
# <span class="badge builder"></span> MetricStatBuilder

## Constructor

```typescript
new MetricStatBuilder()
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

### <span class="badge object-method"></span> dimensions

The dimensions of the metric

```typescript
dimensions(dimensions: cloudwatch.Dimensions)
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

 * <span class="badge object-type-interface"></span> [MetricStat](./object-MetricStat.md)
