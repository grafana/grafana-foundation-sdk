---
title: <span class="badge object-type-struct"></span> MetricStat
---
# <span class="badge object-type-struct"></span> MetricStat

## Definition

```go
type MetricStat struct {
    // AWS region to query for the metric
    Region string `json:"region"`
    // A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
    Namespace string `json:"namespace"`
    // Name of the metric
    MetricName *string `json:"metricName,omitempty"`
    // The dimensions of the metric
    Dimensions *cloudwatch.Dimensions `json:"dimensions,omitempty"`
    // Only show metrics that exactly match all defined dimension names.
    MatchExact *bool `json:"matchExact,omitempty"`
    // The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
    Period *string `json:"period,omitempty"`
    // The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
    AccountId *string `json:"accountId,omitempty"`
    // Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
    Statistic *string `json:"statistic,omitempty"`
    // @deprecated use statistic
    Statistics []string `json:"statistics,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricStat` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metricStat *MetricStat) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MetricStat` objects.

```go
func (metricStat *MetricStat) Equals(other MetricStat) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MetricStat` fields for violations and returns them.

```go
func (metricStat *MetricStat) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetricStatBuilder](./builder-MetricStatBuilder.md)
