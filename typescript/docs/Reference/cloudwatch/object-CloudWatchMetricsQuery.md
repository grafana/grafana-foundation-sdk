---
title: <span class="badge object-type-interface"></span> CloudWatchMetricsQuery
---
# <span class="badge object-type-interface"></span> CloudWatchMetricsQuery

Shape of a CloudWatch Metrics query

## Definition

```typescript
export interface CloudWatchMetricsQuery {
	// Whether a query is a Metrics, Logs, or Annotations query
	queryMode: cloudwatch.CloudWatchQueryMode;
	// Whether to use a metric search or metric insights query
	metricQueryType?: cloudwatch.MetricQueryType;
	// Whether to use the query builder or code editor to create the query
	metricEditorMode?: cloudwatch.MetricEditorMode;
	// ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
	id: string;
	// Deprecated: use label
	// @deprecated use label
	alias?: string;
	// Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
	label?: string;
	// Math expression query
	expression?: string;
	// When the metric query type is set to `Insights`, this field is used to specify the query string.
	sqlExpression?: string;
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
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
	// When the metric query type is set to `Insights` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
	sql?: cloudwatch.SQLExpression;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	// @deprecated use statistic
	statistics?: string[];
	_implementsDataqueryVariant(): void;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [CloudWatchMetricsQueryBuilder](./builder-CloudWatchMetricsQueryBuilder.md)
