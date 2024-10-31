// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*CloudWatchAnnotationQueryBuilder)(nil)

// Shape of a CloudWatch Annotation query
// TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
// #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
type CloudWatchAnnotationQueryBuilder struct {
	internal *CloudWatchAnnotationQuery
	errors   map[string]cog.BuildErrors
}

func NewCloudWatchAnnotationQueryBuilder() *CloudWatchAnnotationQueryBuilder {
	resource := &CloudWatchAnnotationQuery{}
	builder := &CloudWatchAnnotationQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CloudWatchAnnotationQueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return CloudWatchAnnotationQuery{}, err
	}

	return *builder.internal, nil
}

// Whether a query is a Metrics, Logs, or Annotations query
func (builder *CloudWatchAnnotationQueryBuilder) QueryMode(queryMode CloudWatchQueryMode) *CloudWatchAnnotationQueryBuilder {
	builder.internal.QueryMode = queryMode

	return builder
}

// Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
func (builder *CloudWatchAnnotationQueryBuilder) PrefixMatching(prefixMatching bool) *CloudWatchAnnotationQueryBuilder {
	builder.internal.PrefixMatching = &prefixMatching

	return builder
}

// Use this parameter to filter the results of the operation to only those alarms
// that use a certain alarm action. For example, you could specify the ARN of
// an SNS topic to find all alarms that send notifications to that topic.
// e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
// but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
func (builder *CloudWatchAnnotationQueryBuilder) ActionPrefix(actionPrefix string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.ActionPrefix = &actionPrefix

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *CloudWatchAnnotationQueryBuilder) RefId(refId string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *CloudWatchAnnotationQueryBuilder) Hide(hide bool) *CloudWatchAnnotationQueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *CloudWatchAnnotationQueryBuilder) QueryType(queryType string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// AWS region to query for the metric
func (builder *CloudWatchAnnotationQueryBuilder) Region(region string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.Region = region

	return builder
}

// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
func (builder *CloudWatchAnnotationQueryBuilder) Namespace(namespace string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.Namespace = namespace

	return builder
}

// Name of the metric
func (builder *CloudWatchAnnotationQueryBuilder) MetricName(metricName string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.MetricName = &metricName

	return builder
}

// The dimensions of the metric
func (builder *CloudWatchAnnotationQueryBuilder) Dimensions(dimensions Dimensions) *CloudWatchAnnotationQueryBuilder {
	builder.internal.Dimensions = &dimensions

	return builder
}

// Only show metrics that exactly match all defined dimension names.
func (builder *CloudWatchAnnotationQueryBuilder) MatchExact(matchExact bool) *CloudWatchAnnotationQueryBuilder {
	builder.internal.MatchExact = &matchExact

	return builder
}

// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
func (builder *CloudWatchAnnotationQueryBuilder) Period(period string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.Period = &period

	return builder
}

// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
func (builder *CloudWatchAnnotationQueryBuilder) AccountId(accountId string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.AccountId = &accountId

	return builder
}

// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
func (builder *CloudWatchAnnotationQueryBuilder) Statistic(statistic string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.Statistic = &statistic

	return builder
}

// An alarm name prefix. If you specify this parameter, you receive information
// about all alarms that have names that start with this prefix.
// e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
func (builder *CloudWatchAnnotationQueryBuilder) AlarmNamePrefix(alarmNamePrefix string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.AlarmNamePrefix = &alarmNamePrefix

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *CloudWatchAnnotationQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *CloudWatchAnnotationQueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// @deprecated use statistic
func (builder *CloudWatchAnnotationQueryBuilder) Statistics(statistics []string) *CloudWatchAnnotationQueryBuilder {
	builder.internal.Statistics = statistics

	return builder
}

func (builder *CloudWatchAnnotationQueryBuilder) applyDefaults() {
	builder.QueryMode("Annotations")
}
