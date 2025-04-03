---
title: <span class="badge builder"></span> MetricStat
---
# <span class="badge builder"></span> MetricStat

## Constructor

```python
MetricStat()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> cloudwatch.MetricStat
```

### <span class="badge object-method"></span> account_id

The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.

```python
def account_id(account_id: str) -> typing.Self
```

### <span class="badge object-method"></span> dimensions

The dimensions of the metric

```python
def dimensions(dimensions: cloudwatch.Dimensions) -> typing.Self
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

 * <span class="badge object-type-class"></span> [MetricStat](./object-MetricStat.md)
