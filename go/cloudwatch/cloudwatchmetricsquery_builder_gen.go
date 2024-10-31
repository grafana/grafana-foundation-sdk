// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*CloudWatchMetricsQueryBuilder)(nil)

// Shape of a CloudWatch Metrics query
type CloudWatchMetricsQueryBuilder struct {
	internal *CloudWatchMetricsQuery
	errors   map[string]cog.BuildErrors
}

func NewCloudWatchMetricsQueryBuilder() *CloudWatchMetricsQueryBuilder {
	resource := &CloudWatchMetricsQuery{}
	builder := &CloudWatchMetricsQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CloudWatchMetricsQueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return CloudWatchMetricsQuery{}, err
	}

	return *builder.internal, nil
}

// Whether a query is a Metrics, Logs, or Annotations query
func (builder *CloudWatchMetricsQueryBuilder) QueryMode(queryMode CloudWatchQueryMode) *CloudWatchMetricsQueryBuilder {
	builder.internal.QueryMode = queryMode

	return builder
}

// Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.
func (builder *CloudWatchMetricsQueryBuilder) MetricQueryType(metricQueryType MetricQueryType) *CloudWatchMetricsQueryBuilder {
	builder.internal.MetricQueryType = &metricQueryType

	return builder
}

// Whether to use the query builder or code editor to create the query
func (builder *CloudWatchMetricsQueryBuilder) MetricEditorMode(metricEditorMode MetricEditorMode) *CloudWatchMetricsQueryBuilder {
	builder.internal.MetricEditorMode = &metricEditorMode

	return builder
}

// ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
func (builder *CloudWatchMetricsQueryBuilder) Id(id string) *CloudWatchMetricsQueryBuilder {
	builder.internal.Id = id

	return builder
}

// Deprecated: use label
// @deprecated use label
func (builder *CloudWatchMetricsQueryBuilder) Alias(alias string) *CloudWatchMetricsQueryBuilder {
	builder.internal.Alias = &alias

	return builder
}

// Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
func (builder *CloudWatchMetricsQueryBuilder) Label(label string) *CloudWatchMetricsQueryBuilder {
	builder.internal.Label = &label

	return builder
}

// Math expression query
func (builder *CloudWatchMetricsQueryBuilder) Expression(expression string) *CloudWatchMetricsQueryBuilder {
	builder.internal.Expression = &expression

	return builder
}

// When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
func (builder *CloudWatchMetricsQueryBuilder) SqlExpression(sqlExpression string) *CloudWatchMetricsQueryBuilder {
	builder.internal.SqlExpression = &sqlExpression

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *CloudWatchMetricsQueryBuilder) RefId(refId string) *CloudWatchMetricsQueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *CloudWatchMetricsQueryBuilder) Hide(hide bool) *CloudWatchMetricsQueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *CloudWatchMetricsQueryBuilder) QueryType(queryType string) *CloudWatchMetricsQueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// AWS region to query for the metric
func (builder *CloudWatchMetricsQueryBuilder) Region(region string) *CloudWatchMetricsQueryBuilder {
	builder.internal.Region = region

	return builder
}

// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
func (builder *CloudWatchMetricsQueryBuilder) Namespace(namespace string) *CloudWatchMetricsQueryBuilder {
	builder.internal.Namespace = namespace

	return builder
}

// Name of the metric
func (builder *CloudWatchMetricsQueryBuilder) MetricName(metricName string) *CloudWatchMetricsQueryBuilder {
	builder.internal.MetricName = &metricName

	return builder
}

// The dimensions of the metric
func (builder *CloudWatchMetricsQueryBuilder) Dimensions(dimensions Dimensions) *CloudWatchMetricsQueryBuilder {
	builder.internal.Dimensions = &dimensions

	return builder
}

// Only show metrics that exactly match all defined dimension names.
func (builder *CloudWatchMetricsQueryBuilder) MatchExact(matchExact bool) *CloudWatchMetricsQueryBuilder {
	builder.internal.MatchExact = &matchExact

	return builder
}

// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
func (builder *CloudWatchMetricsQueryBuilder) Period(period string) *CloudWatchMetricsQueryBuilder {
	builder.internal.Period = &period

	return builder
}

// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
func (builder *CloudWatchMetricsQueryBuilder) AccountId(accountId string) *CloudWatchMetricsQueryBuilder {
	builder.internal.AccountId = &accountId

	return builder
}

// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
func (builder *CloudWatchMetricsQueryBuilder) Statistic(statistic string) *CloudWatchMetricsQueryBuilder {
	builder.internal.Statistic = &statistic

	return builder
}

// When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
func (builder *CloudWatchMetricsQueryBuilder) Sql(sql cog.Builder[SQLExpression]) *CloudWatchMetricsQueryBuilder {
	sqlResource, err := sql.Build()
	if err != nil {
		builder.errors["sql"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Sql = &sqlResource

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *CloudWatchMetricsQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *CloudWatchMetricsQueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// @deprecated use statistic
func (builder *CloudWatchMetricsQueryBuilder) Statistics(statistics []string) *CloudWatchMetricsQueryBuilder {
	builder.internal.Statistics = statistics

	return builder
}

func (builder *CloudWatchMetricsQueryBuilder) applyDefaults() {
	builder.QueryMode("Metrics")
}
