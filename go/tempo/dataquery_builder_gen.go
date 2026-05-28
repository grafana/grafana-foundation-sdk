// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package tempo

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[variants.Dataquery] = (*DataqueryBuilder)(nil)

type DataqueryBuilder struct {
	internal *Dataquery
	errors   cog.BuildErrors
}

func NewDataqueryBuilder() *DataqueryBuilder {
	resource := NewDataquery()
	builder := &DataqueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DataqueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dataquery{}, err
	}

	if len(builder.errors) > 0 {
		return Dataquery{}, cog.MakeBuildErrors("tempo.dataquery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DataqueryBuilder) RecordError(path string, err error) *DataqueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// TraceQL query or trace ID
func (builder *DataqueryBuilder) Query(query string) *DataqueryBuilder {
	builder.internal.Query = &query

	return builder
}

// @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
func (builder *DataqueryBuilder) Search(search string) *DataqueryBuilder {
	builder.internal.Search = &search

	return builder
}

// @deprecated Query traces by service name
func (builder *DataqueryBuilder) ServiceName(serviceName string) *DataqueryBuilder {
	builder.internal.ServiceName = &serviceName

	return builder
}

// @deprecated Query traces by span name
func (builder *DataqueryBuilder) SpanName(spanName string) *DataqueryBuilder {
	builder.internal.SpanName = &spanName

	return builder
}

// @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
func (builder *DataqueryBuilder) MinDuration(minDuration string) *DataqueryBuilder {
	builder.internal.MinDuration = &minDuration

	return builder
}

// @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
func (builder *DataqueryBuilder) MaxDuration(maxDuration string) *DataqueryBuilder {
	builder.internal.MaxDuration = &maxDuration

	return builder
}

// Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
func (builder *DataqueryBuilder) ServiceMapQuery(serviceMapQuery StringOrArrayOfString) *DataqueryBuilder {
	builder.internal.ServiceMapQuery = &serviceMapQuery

	return builder
}

// Use service.namespace in addition to service.name to uniquely identify a service.
func (builder *DataqueryBuilder) ServiceMapIncludeNamespace(serviceMapIncludeNamespace bool) *DataqueryBuilder {
	builder.internal.ServiceMapIncludeNamespace = &serviceMapIncludeNamespace

	return builder
}

// Defines the maximum number of traces that are returned from Tempo
func (builder *DataqueryBuilder) Limit(limit int64) *DataqueryBuilder {
	builder.internal.Limit = &limit

	return builder
}

// Defines the maximum number of spans per spanset that are returned from Tempo
func (builder *DataqueryBuilder) Spss(spss int64) *DataqueryBuilder {
	builder.internal.Spss = &spss

	return builder
}

func (builder *DataqueryBuilder) Filters(filters []cog.Builder[TraceqlFilter]) *DataqueryBuilder {
	filtersResources := make([]TraceqlFilter, 0, len(filters))
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

// Filters that are used to query the metrics summary
func (builder *DataqueryBuilder) GroupBy(groupBy []cog.Builder[TraceqlFilter]) *DataqueryBuilder {
	groupByResources := make([]TraceqlFilter, 0, len(groupBy))
	for _, r1 := range groupBy {
		groupByDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		groupByResources = append(groupByResources, groupByDepth1)
	}
	builder.internal.GroupBy = groupByResources

	return builder
}

// The type of the table that is used to display the search results
func (builder *DataqueryBuilder) TableType(tableType SearchTableType) *DataqueryBuilder {
	builder.internal.TableType = &tableType

	return builder
}

// For metric queries, the step size to use
func (builder *DataqueryBuilder) Step(step string) *DataqueryBuilder {
	builder.internal.Step = &step

	return builder
}

// For metric queries, how many exemplars to request, 0 means no exemplars
func (builder *DataqueryBuilder) Exemplars(exemplars int64) *DataqueryBuilder {
	builder.internal.Exemplars = &exemplars

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *DataqueryBuilder) Datasource(datasource common.DataSourceRef) *DataqueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// For metric queries, whether to run instant or range queries
func (builder *DataqueryBuilder) MetricsQueryType(metricsQueryType MetricsQueryType) *DataqueryBuilder {
	builder.internal.MetricsQueryType = &metricsQueryType

	return builder
}
