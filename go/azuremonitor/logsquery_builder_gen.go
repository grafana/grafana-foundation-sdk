// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LogsQuery] = (*LogsQueryBuilder)(nil)

// Azure Monitor Logs sub-query properties
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
		return LogsQuery{}, cog.MakeBuildErrors("azuremonitor.logsQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *LogsQueryBuilder) RecordError(path string, err error) *LogsQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// KQL query to be executed.
func (builder *LogsQueryBuilder) Query(query string) *LogsQueryBuilder {
	builder.internal.Query = &query

	return builder
}

// Specifies the format results should be returned as.
func (builder *LogsQueryBuilder) ResultFormat(resultFormat ResultFormat) *LogsQueryBuilder {
	builder.internal.ResultFormat = &resultFormat

	return builder
}

// Array of resource URIs to be queried.
func (builder *LogsQueryBuilder) Resources(resources []string) *LogsQueryBuilder {
	builder.internal.Resources = resources

	return builder
}

// If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
func (builder *LogsQueryBuilder) DashboardTime(dashboardTime bool) *LogsQueryBuilder {
	builder.internal.DashboardTime = &dashboardTime

	return builder
}

// If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
func (builder *LogsQueryBuilder) TimeColumn(timeColumn string) *LogsQueryBuilder {
	builder.internal.TimeColumn = &timeColumn

	return builder
}

// If set to true the query will be run as a basic logs query
func (builder *LogsQueryBuilder) BasicLogsQuery(basicLogsQuery bool) *LogsQueryBuilder {
	builder.internal.BasicLogsQuery = &basicLogsQuery

	return builder
}

// Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
func (builder *LogsQueryBuilder) Workspace(workspace string) *LogsQueryBuilder {
	builder.internal.Workspace = &workspace

	return builder
}

// @deprecated Use resources instead
func (builder *LogsQueryBuilder) Resource(resource string) *LogsQueryBuilder {
	builder.internal.Resource = &resource

	return builder
}

// @deprecated Use dashboardTime instead
func (builder *LogsQueryBuilder) IntersectTime(intersectTime bool) *LogsQueryBuilder {
	builder.internal.IntersectTime = &intersectTime

	return builder
}
