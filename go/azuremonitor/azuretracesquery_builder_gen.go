// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AzureTracesQuery] = (*AzureTracesQueryBuilder)(nil)

// Application Insights Traces sub-query properties
type AzureTracesQueryBuilder struct {
	internal *AzureTracesQuery
	errors   map[string]cog.BuildErrors
}

func NewAzureTracesQueryBuilder() *AzureTracesQueryBuilder {
	resource := &AzureTracesQuery{}
	builder := &AzureTracesQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AzureTracesQueryBuilder) Build() (AzureTracesQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureTracesQuery{}, err
	}

	return *builder.internal, nil
}

// Specifies the format results should be returned as.
func (builder *AzureTracesQueryBuilder) ResultFormat(resultFormat ResultFormat) *AzureTracesQueryBuilder {
	builder.internal.ResultFormat = &resultFormat

	return builder
}

// Array of resource URIs to be queried.
func (builder *AzureTracesQueryBuilder) Resources(resources []string) *AzureTracesQueryBuilder {
	builder.internal.Resources = resources

	return builder
}

// Operation ID. Used only for Traces queries.
func (builder *AzureTracesQueryBuilder) OperationId(operationId string) *AzureTracesQueryBuilder {
	builder.internal.OperationId = &operationId

	return builder
}

// Types of events to filter by.
func (builder *AzureTracesQueryBuilder) TraceTypes(traceTypes []string) *AzureTracesQueryBuilder {
	builder.internal.TraceTypes = traceTypes

	return builder
}

// Filters for property values.
func (builder *AzureTracesQueryBuilder) Filters(filters []cog.Builder[AzureTracesFilter]) *AzureTracesQueryBuilder {
	filtersResources := make([]AzureTracesFilter, 0, len(filters))
	for _, r1 := range filters {
		filtersDepth1, err := r1.Build()
		if err != nil {
			builder.errors["filters"] = err.(cog.BuildErrors)
			return builder
		}
		filtersResources = append(filtersResources, filtersDepth1)
	}
	builder.internal.Filters = filtersResources

	return builder
}

// KQL query to be executed.
func (builder *AzureTracesQueryBuilder) Query(query string) *AzureTracesQueryBuilder {
	builder.internal.Query = &query

	return builder
}

func (builder *AzureTracesQueryBuilder) applyDefaults() {
}
