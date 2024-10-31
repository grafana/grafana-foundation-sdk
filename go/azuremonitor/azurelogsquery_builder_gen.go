// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AzureLogsQuery] = (*AzureLogsQueryBuilder)(nil)

// Azure Monitor Logs sub-query properties
type AzureLogsQueryBuilder struct {
	internal *AzureLogsQuery
	errors   map[string]cog.BuildErrors
}

func NewAzureLogsQueryBuilder() *AzureLogsQueryBuilder {
	resource := &AzureLogsQuery{}
	builder := &AzureLogsQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AzureLogsQueryBuilder) Build() (AzureLogsQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureLogsQuery{}, err
	}

	return *builder.internal, nil
}

// KQL query to be executed.
func (builder *AzureLogsQueryBuilder) Query(query string) *AzureLogsQueryBuilder {
	builder.internal.Query = &query

	return builder
}

// Specifies the format results should be returned as.
func (builder *AzureLogsQueryBuilder) ResultFormat(resultFormat ResultFormat) *AzureLogsQueryBuilder {
	builder.internal.ResultFormat = &resultFormat

	return builder
}

// Array of resource URIs to be queried.
func (builder *AzureLogsQueryBuilder) Resources(resources []string) *AzureLogsQueryBuilder {
	builder.internal.Resources = resources

	return builder
}

// If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
func (builder *AzureLogsQueryBuilder) DashboardTime(dashboardTime bool) *AzureLogsQueryBuilder {
	builder.internal.DashboardTime = &dashboardTime

	return builder
}

// If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
func (builder *AzureLogsQueryBuilder) TimeColumn(timeColumn string) *AzureLogsQueryBuilder {
	builder.internal.TimeColumn = &timeColumn

	return builder
}

// Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
func (builder *AzureLogsQueryBuilder) Workspace(workspace string) *AzureLogsQueryBuilder {
	builder.internal.Workspace = &workspace

	return builder
}

// @deprecated Use resources instead
func (builder *AzureLogsQueryBuilder) Resource(resource string) *AzureLogsQueryBuilder {
	builder.internal.Resource = &resource

	return builder
}

// @deprecated Use dashboardTime instead
func (builder *AzureLogsQueryBuilder) IntersectTime(intersectTime bool) *AzureLogsQueryBuilder {
	builder.internal.IntersectTime = &intersectTime

	return builder
}

func (builder *AzureLogsQueryBuilder) applyDefaults() {
}
