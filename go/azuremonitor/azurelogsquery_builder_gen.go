// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AzureLogsQuery] = (*AzureLogsQueryBuilder)(nil)

// Azure Monitor Logs sub-query properties
type AzureLogsQueryBuilder struct {
	internal *AzureLogsQuery
	errors   cog.BuildErrors
}

func NewAzureLogsQueryBuilder() *AzureLogsQueryBuilder {
	resource := NewAzureLogsQuery()
	builder := &AzureLogsQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AzureLogsQueryBuilder) Build() (AzureLogsQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureLogsQuery{}, err
	}

	if len(builder.errors) > 0 {
		return AzureLogsQuery{}, cog.MakeBuildErrors("azuremonitor.azureLogsQuery", builder.errors)
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

// If set to true the intersection of time ranges specified in the query and Grafana will be used. Otherwise the query time ranges will be used. Defaults to false
func (builder *AzureLogsQueryBuilder) IntersectTime(intersectTime bool) *AzureLogsQueryBuilder {
	builder.internal.IntersectTime = &intersectTime

	return builder
}

// Workspace ID. This was removed in Grafana 8, but remains for backwards compat
func (builder *AzureLogsQueryBuilder) Workspace(workspace string) *AzureLogsQueryBuilder {
	builder.internal.Workspace = &workspace

	return builder
}

// @deprecated Use resources instead
func (builder *AzureLogsQueryBuilder) Resource(resource string) *AzureLogsQueryBuilder {
	builder.internal.Resource = &resource

	return builder
}
