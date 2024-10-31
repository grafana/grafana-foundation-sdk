---
title: <span class="badge builder"></span> MetricStatBuilder
---
# <span class="badge builder"></span> MetricStatBuilder

## Constructor

```php
new MetricStatBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> accountId

The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.

```php
accountId(string $accountId)
```

### <span class="badge object-method"></span> dimensions

The dimensions of the metric

@param array<string, string|array<string>> $dimensions

```php
dimensions(array $dimensions)
```

### <span class="badge object-method"></span> matchExact

Only show metrics that exactly match all defined dimension names.

```php
matchExact(bool $matchExact)
```

### <span class="badge object-method"></span> metricName

Name of the metric

```php
metricName(string $metricName)
```

### <span class="badge object-method"></span> namespace

A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.

```php
namespace(string $namespace)
```

### <span class="badge object-method"></span> period

The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes

```php
period(string $period)
```

### <span class="badge object-method"></span> region

AWS region to query for the metric

```php
region(string $region)
```

### <span class="badge object-method"></span> statistic

Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.

```php
statistic(string $statistic)
```

### <span class="badge object-method"></span> statistics

@deprecated use statistic

@param array<string> $statistics

```php
statistics(array $statistics)
```

## See also

 * <span class="badge object-type-class"></span> [MetricStat](./object-MetricStat.md)
