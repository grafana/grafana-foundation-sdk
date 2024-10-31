---
title: <span class="badge object-type-class"></span> MetricStat
---
# <span class="badge object-type-class"></span> MetricStat

## Definition

```php
class MetricStat implements \JsonSerializable
{
    /**
     * AWS region to query for the metric
     */
    public string $region;

    /**
     * A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
     */
    public string $namespace;

    /**
     * Name of the metric
     */
    public ?string $metricName;

    /**
     * The dimensions of the metric
     * @var array<string, string|array<string>>
     */
    public array $dimensions;

    /**
     * Only show metrics that exactly match all defined dimension names.
     */
    public ?bool $matchExact;

    /**
     * The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
     */
    public ?string $period;

    /**
     * The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
     */
    public ?string $accountId;

    /**
     * Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
     */
    public ?string $statistic;

    /**
     * @deprecated use statistic
     * @var array<string>|null
     */
    public ?array $statistics;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [MetricStatBuilder](./builder-MetricStatBuilder.md)
