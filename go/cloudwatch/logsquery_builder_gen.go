// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[LogsQuery] = (*LogsQueryBuilder)(nil)

// Shape of a CloudWatch Logs query
type LogsQueryBuilder struct {
	internal *LogsQuery
	errors   cog.BuildErrors
}

func NewLogsQueryBuilder() *LogsQueryBuilder {
	resource := NewLogsQuery()
	builder := &LogsQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *LogsQueryBuilder) Build() (LogsQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return LogsQuery{}, err
	}

	if len(builder.errors) > 0 {
		return LogsQuery{}, cog.MakeBuildErrors("cloudwatch.logsQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *LogsQueryBuilder) RecordError(path string, err error) *LogsQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Whether a query is a Metrics, Logs, or Annotations query
func (builder *LogsQueryBuilder) QueryMode(queryMode QueryMode) *LogsQueryBuilder {
	builder.internal.QueryMode = queryMode

	return builder
}

func (builder *LogsQueryBuilder) Id(id string) *LogsQueryBuilder {
	builder.internal.Id = id

	return builder
}

// AWS region to query for the logs
func (builder *LogsQueryBuilder) Region(region string) *LogsQueryBuilder {
	builder.internal.Region = region

	return builder
}

// The CloudWatch Logs Insights query to execute
func (builder *LogsQueryBuilder) Expression(expression string) *LogsQueryBuilder {
	builder.internal.Expression = &expression

	return builder
}

// Fields to group the results by, this field is automatically populated whenever the query is updated
func (builder *LogsQueryBuilder) StatsGroups(statsGroups []string) *LogsQueryBuilder {
	builder.internal.StatsGroups = statsGroups

	return builder
}

// Log groups to query
func (builder *LogsQueryBuilder) LogGroups(logGroups []cog.Builder[LogGroup]) *LogsQueryBuilder {
	logGroupsResources := make([]LogGroup, 0, len(logGroups))
	for _, r1 := range logGroups {
		logGroupsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		logGroupsResources = append(logGroupsResources, logGroupsDepth1)
	}
	builder.internal.LogGroups = logGroupsResources

	return builder
}

// @deprecated use logGroups
func (builder *LogsQueryBuilder) LogGroupNames(logGroupNames []string) *LogsQueryBuilder {
	builder.internal.LogGroupNames = logGroupNames

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *LogsQueryBuilder) RefId(refId string) *LogsQueryBuilder {
	builder.internal.RefId = &refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *LogsQueryBuilder) Hide(hide bool) *LogsQueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *LogsQueryBuilder) QueryType(queryType string) *LogsQueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// Language used for querying logs, can be CWLI, SQL, or PPL. If empty, the default language is CWLI.
func (builder *LogsQueryBuilder) QueryLanguage(queryLanguage LogsQueryLanguage) *LogsQueryBuilder {
	builder.internal.QueryLanguage = &queryLanguage

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *LogsQueryBuilder) Datasource(datasource common.DataSourceRef) *LogsQueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}
