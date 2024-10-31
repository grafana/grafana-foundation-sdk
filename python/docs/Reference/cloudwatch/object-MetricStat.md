---
title: <span class="badge object-type-class"></span> MetricStat
---
# <span class="badge object-type-class"></span> MetricStat

## Definition

```python
class MetricStat:
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

 * <span class="badge builder"></span> [MetricStat](./builder-MetricStat.md)
