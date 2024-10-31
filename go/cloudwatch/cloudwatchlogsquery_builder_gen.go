// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*CloudWatchLogsQueryBuilder)(nil)

// Shape of a CloudWatch Logs query
type CloudWatchLogsQueryBuilder struct {
	internal *CloudWatchLogsQuery
	errors   map[string]cog.BuildErrors
}

func NewCloudWatchLogsQueryBuilder() *CloudWatchLogsQueryBuilder {
	resource := &CloudWatchLogsQuery{}
	builder := &CloudWatchLogsQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CloudWatchLogsQueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return CloudWatchLogsQuery{}, err
	}

	return *builder.internal, nil
}

// Whether a query is a Metrics, Logs, or Annotations query
func (builder *CloudWatchLogsQueryBuilder) QueryMode(queryMode CloudWatchQueryMode) *CloudWatchLogsQueryBuilder {
	builder.internal.QueryMode = queryMode

	return builder
}

func (builder *CloudWatchLogsQueryBuilder) Id(id string) *CloudWatchLogsQueryBuilder {
	builder.internal.Id = id

	return builder
}

// AWS region to query for the logs
func (builder *CloudWatchLogsQueryBuilder) Region(region string) *CloudWatchLogsQueryBuilder {
	builder.internal.Region = region

	return builder
}

// The CloudWatch Logs Insights query to execute
func (builder *CloudWatchLogsQueryBuilder) Expression(expression string) *CloudWatchLogsQueryBuilder {
	builder.internal.Expression = &expression

	return builder
}

// Fields to group the results by, this field is automatically populated whenever the query is updated
func (builder *CloudWatchLogsQueryBuilder) StatsGroups(statsGroups []string) *CloudWatchLogsQueryBuilder {
	builder.internal.StatsGroups = statsGroups

	return builder
}

// Log groups to query
func (builder *CloudWatchLogsQueryBuilder) LogGroups(logGroups []cog.Builder[LogGroup]) *CloudWatchLogsQueryBuilder {
	logGroupsResources := make([]LogGroup, 0, len(logGroups))
	for _, r1 := range logGroups {
		logGroupsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["logGroups"] = err.(cog.BuildErrors)
			return builder
		}
		logGroupsResources = append(logGroupsResources, logGroupsDepth1)
	}
	builder.internal.LogGroups = logGroupsResources

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *CloudWatchLogsQueryBuilder) RefId(refId string) *CloudWatchLogsQueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *CloudWatchLogsQueryBuilder) Hide(hide bool) *CloudWatchLogsQueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *CloudWatchLogsQueryBuilder) QueryType(queryType string) *CloudWatchLogsQueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// @deprecated use logGroups
func (builder *CloudWatchLogsQueryBuilder) LogGroupNames(logGroupNames []string) *CloudWatchLogsQueryBuilder {
	builder.internal.LogGroupNames = logGroupNames

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *CloudWatchLogsQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *CloudWatchLogsQueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

func (builder *CloudWatchLogsQueryBuilder) applyDefaults() {
	builder.QueryMode("Logs")
}
