// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TracesQuery] = (*TracesQueryBuilder)(nil)

// Application Insights Traces sub-query properties
type TracesQueryBuilder struct {
	internal *TracesQuery
	errors   cog.BuildErrors
}

func NewTracesQueryBuilder() *TracesQueryBuilder {
	resource := NewTracesQuery()
	builder := &TracesQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TracesQueryBuilder) Build() (TracesQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return TracesQuery{}, err
	}

	if len(builder.errors) > 0 {
		return TracesQuery{}, cog.MakeBuildErrors("azuremonitor.tracesQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TracesQueryBuilder) RecordError(path string, err error) *TracesQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Specifies the format results should be returned as.
func (builder *TracesQueryBuilder) ResultFormat(resultFormat ResultFormat) *TracesQueryBuilder {
	builder.internal.ResultFormat = &resultFormat

	return builder
}

// Array of resource URIs to be queried.
func (builder *TracesQueryBuilder) Resources(resources []string) *TracesQueryBuilder {
	builder.internal.Resources = resources

	return builder
}

// Operation ID. Used only for Traces queries.
func (builder *TracesQueryBuilder) OperationId(operationId string) *TracesQueryBuilder {
	builder.internal.OperationId = &operationId

	return builder
}

// Types of events to filter by.
func (builder *TracesQueryBuilder) TraceTypes(traceTypes []string) *TracesQueryBuilder {
	builder.internal.TraceTypes = traceTypes

	return builder
}

// Filters for property values.
func (builder *TracesQueryBuilder) Filters(filters []cog.Builder[TracesFilter]) *TracesQueryBuilder {
	filtersResources := make([]TracesFilter, 0, len(filters))
	for _, r1 := range filters {
		filtersDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		filtersResources = append(filtersResources, filtersDepth1)
	}
	builder.internal.Filters = filtersResources

	return builder
}

// KQL query to be executed.
func (builder *TracesQueryBuilder) Query(query string) *TracesQueryBuilder {
	builder.internal.Query = &query

	return builder
}
