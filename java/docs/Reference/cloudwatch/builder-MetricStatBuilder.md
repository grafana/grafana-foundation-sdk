---
title: <span class="badge builder"></span> MetricStatBuilder
---
# <span class="badge builder"></span> MetricStatBuilder

## Constructor

```java
new MetricStatBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public MetricStat build()
```

### <span class="badge object-method"></span> accountId

The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.

```java
public MetricStatBuilder accountId(String accountId)
```

### <span class="badge object-method"></span> dimensions

The dimensions of the metric

```java
public MetricStatBuilder dimensions(Map<String, StringOrArrayOfString> dimensions)
```

### <span class="badge object-method"></span> matchExact

Only show metrics that exactly match all defined dimension names.

```java
public MetricStatBuilder matchExact(Boolean matchExact)
```

### <span class="badge object-method"></span> metricName

Name of the metric

```java
public MetricStatBuilder metricName(String metricName)
```

### <span class="badge object-method"></span> namespace

A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.

```java
public MetricStatBuilder namespace(String namespace)
```

### <span class="badge object-method"></span> period

The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes

```java
public MetricStatBuilder period(String period)
```

### <span class="badge object-method"></span> region

AWS region to query for the metric

```java
public MetricStatBuilder region(String region)
```

### <span class="badge object-method"></span> statistic

Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.

```java
public MetricStatBuilder statistic(String statistic)
```

### <span class="badge object-method"></span> statistics

@deprecated use statistic

```java
public MetricStatBuilder statistics(List<String> statistics)
```

## See also

 * <span class="badge object-type-class"></span> [MetricStat](./object-MetricStat.md)
