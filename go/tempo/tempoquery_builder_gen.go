// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package tempo

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*TempoQueryBuilder)(nil)

type TempoQueryBuilder struct {
	internal *TempoQuery
	errors   map[string]cog.BuildErrors
}

func NewTempoQueryBuilder() *TempoQueryBuilder {
	resource := &TempoQuery{}
	builder := &TempoQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TempoQueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return TempoQuery{}, err
	}

	return *builder.internal, nil
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *TempoQueryBuilder) RefId(refId string) *TempoQueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *TempoQueryBuilder) Hide(hide bool) *TempoQueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *TempoQueryBuilder) QueryType(queryType string) *TempoQueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// TraceQL query or trace ID
func (builder *TempoQueryBuilder) Query(query string) *TempoQueryBuilder {
	builder.internal.Query = &query

	return builder
}

// @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
func (builder *TempoQueryBuilder) Search(search string) *TempoQueryBuilder {
	builder.internal.Search = &search

	return builder
}

// @deprecated Query traces by service name
func (builder *TempoQueryBuilder) ServiceName(serviceName string) *TempoQueryBuilder {
	builder.internal.ServiceName = &serviceName

	return builder
}

// @deprecated Query traces by span name
func (builder *TempoQueryBuilder) SpanName(spanName string) *TempoQueryBuilder {
	builder.internal.SpanName = &spanName

	return builder
}

// @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
func (builder *TempoQueryBuilder) MinDuration(minDuration string) *TempoQueryBuilder {
	builder.internal.MinDuration = &minDuration

	return builder
}

// @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
func (builder *TempoQueryBuilder) MaxDuration(maxDuration string) *TempoQueryBuilder {
	builder.internal.MaxDuration = &maxDuration

	return builder
}

// Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
func (builder *TempoQueryBuilder) ServiceMapQuery(serviceMapQuery StringOrArrayOfString) *TempoQueryBuilder {
	builder.internal.ServiceMapQuery = &serviceMapQuery

	return builder
}

// Use service.namespace in addition to service.name to uniquely identify a service.
func (builder *TempoQueryBuilder) ServiceMapIncludeNamespace(serviceMapIncludeNamespace bool) *TempoQueryBuilder {
	builder.internal.ServiceMapIncludeNamespace = &serviceMapIncludeNamespace

	return builder
}

// Defines the maximum number of traces that are returned from Tempo
func (builder *TempoQueryBuilder) Limit(limit int64) *TempoQueryBuilder {
	builder.internal.Limit = &limit

	return builder
}

// Defines the maximum number of spans per spanset that are returned from Tempo
func (builder *TempoQueryBuilder) Spss(spss int64) *TempoQueryBuilder {
	builder.internal.Spss = &spss

	return builder
}

func (builder *TempoQueryBuilder) Filters(filters []cog.Builder[TraceqlFilter]) *TempoQueryBuilder {
	filtersResources := make([]TraceqlFilter, 0, len(filters))
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

// Filters that are used to query the metrics summary
func (builder *TempoQueryBuilder) GroupBy(groupBy []cog.Builder[TraceqlFilter]) *TempoQueryBuilder {
	groupByResources := make([]TraceqlFilter, 0, len(groupBy))
	for _, r1 := range groupBy {
		groupByDepth1, err := r1.Build()
		if err != nil {
			builder.errors["groupBy"] = err.(cog.BuildErrors)
			return builder
		}
		groupByResources = append(groupByResources, groupByDepth1)
	}
	builder.internal.GroupBy = groupByResources

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *TempoQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *TempoQueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// The type of the table that is used to display the search results
func (builder *TempoQueryBuilder) TableType(tableType SearchTableType) *TempoQueryBuilder {
	builder.internal.TableType = &tableType

	return builder
}

func (builder *TempoQueryBuilder) applyDefaults() {
}
