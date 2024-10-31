// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricStat] = (*MetricStatBuilder)(nil)

type MetricStatBuilder struct {
	internal *MetricStat
	errors   map[string]cog.BuildErrors
}

func NewMetricStatBuilder() *MetricStatBuilder {
	resource := &MetricStat{}
	builder := &MetricStatBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *MetricStatBuilder) Build() (MetricStat, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricStat{}, err
	}

	return *builder.internal, nil
}

// AWS region to query for the metric
func (builder *MetricStatBuilder) Region(region string) *MetricStatBuilder {
	builder.internal.Region = region

	return builder
}

// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
func (builder *MetricStatBuilder) Namespace(namespace string) *MetricStatBuilder {
	builder.internal.Namespace = namespace

	return builder
}

// Name of the metric
func (builder *MetricStatBuilder) MetricName(metricName string) *MetricStatBuilder {
	builder.internal.MetricName = &metricName

	return builder
}

// The dimensions of the metric
func (builder *MetricStatBuilder) Dimensions(dimensions Dimensions) *MetricStatBuilder {
	builder.internal.Dimensions = &dimensions

	return builder
}

// Only show metrics that exactly match all defined dimension names.
func (builder *MetricStatBuilder) MatchExact(matchExact bool) *MetricStatBuilder {
	builder.internal.MatchExact = &matchExact

	return builder
}

// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
func (builder *MetricStatBuilder) Period(period string) *MetricStatBuilder {
	builder.internal.Period = &period

	return builder
}

// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
func (builder *MetricStatBuilder) AccountId(accountId string) *MetricStatBuilder {
	builder.internal.AccountId = &accountId

	return builder
}

// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
func (builder *MetricStatBuilder) Statistic(statistic string) *MetricStatBuilder {
	builder.internal.Statistic = &statistic

	return builder
}

// @deprecated use statistic
func (builder *MetricStatBuilder) Statistics(statistics []string) *MetricStatBuilder {
	builder.internal.Statistics = statistics

	return builder
}

func (builder *MetricStatBuilder) applyDefaults() {
}
