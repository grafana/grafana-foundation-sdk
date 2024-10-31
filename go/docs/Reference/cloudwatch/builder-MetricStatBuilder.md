---
title: <span class="badge builder"></span> MetricStatBuilder
---
# <span class="badge builder"></span> MetricStatBuilder

## Constructor

```go
func NewMetricStatBuilder() *MetricStatBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricStatBuilder) Build() (MetricStat, error)
```

### <span class="badge object-method"></span> AccountId

The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.

```go
func (builder *MetricStatBuilder) AccountId(accountId string) *MetricStatBuilder
```

### <span class="badge object-method"></span> Dimensions

The dimensions of the metric

```go
func (builder *MetricStatBuilder) Dimensions(dimensions cloudwatch.Dimensions) *MetricStatBuilder
```

### <span class="badge object-method"></span> MatchExact

Only show metrics that exactly match all defined dimension names.

```go
func (builder *MetricStatBuilder) MatchExact(matchExact bool) *MetricStatBuilder
```

### <span class="badge object-method"></span> MetricName

Name of the metric

```go
func (builder *MetricStatBuilder) MetricName(metricName string) *MetricStatBuilder
```

### <span class="badge object-method"></span> Namespace

A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.

```go
func (builder *MetricStatBuilder) Namespace(namespace string) *MetricStatBuilder
```

### <span class="badge object-method"></span> Period

The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes

```go
func (builder *MetricStatBuilder) Period(period string) *MetricStatBuilder
```

### <span class="badge object-method"></span> Region

AWS region to query for the metric

```go
func (builder *MetricStatBuilder) Region(region string) *MetricStatBuilder
```

### <span class="badge object-method"></span> Statistic

Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.

```go
func (builder *MetricStatBuilder) Statistic(statistic string) *MetricStatBuilder
```

### <span class="badge object-method"></span> Statistics

@deprecated use statistic

```go
func (builder *MetricStatBuilder) Statistics(statistics []string) *MetricStatBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricStat](./object-MetricStat.md)
