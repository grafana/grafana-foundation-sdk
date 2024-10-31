---
title: <span class="badge object-type-interface"></span> MetricStat
---
# <span class="badge object-type-interface"></span> MetricStat

## Definition

```typescript
export interface MetricStat {
	// AWS region to query for the metric
	region: string;
	// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
	namespace: string;
	// Name of the metric
	metricName?: string;
	// The dimensions of the metric
	dimensions?: cloudwatch.Dimensions;
	// Only show metrics that exactly match all defined dimension names.
	matchExact?: boolean;
	// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
	period?: string;
	// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
	accountId?: string;
	// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
	statistic?: string;
	// @deprecated use statistic
	statistics?: string[];
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [MetricStatBuilder](./builder-MetricStatBuilder.md)
