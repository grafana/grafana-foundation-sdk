// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package tempo

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

var _ cog.Builder[dashboardv2beta1.DataQueryKind] = (*QueryBuilder)(nil)

type QueryBuilder struct {
	internal *dashboardv2beta1.DataQueryKind
	errors   cog.BuildErrors
}

func NewQueryBuilder() *QueryBuilder {
	resource := dashboardv2beta1.NewDataQueryKind()
	builder := &QueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"
	builder.internal.Group = "tempo"

	return builder
}

func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2beta1.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2beta1.DataQueryKind{}, cog.MakeBuildErrors("tempo.query", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryBuilder) RecordError(path string, err error) *QueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryBuilder) Version(version string) *QueryBuilder {
	builder.internal.Version = version

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).RefId = &refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).QueryType = &queryType

	return builder
}

// TraceQL query or trace ID
func (builder *QueryBuilder) Query(query string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).Query = &query

	return builder
}

// @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
func (builder *QueryBuilder) Search(search string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).Search = &search

	return builder
}

// @deprecated Query traces by service name
func (builder *QueryBuilder) ServiceName(serviceName string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).ServiceName = &serviceName

	return builder
}

// @deprecated Query traces by span name
func (builder *QueryBuilder) SpanName(spanName string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).SpanName = &spanName

	return builder
}

// @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
func (builder *QueryBuilder) MinDuration(minDuration string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).MinDuration = &minDuration

	return builder
}

// @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
func (builder *QueryBuilder) MaxDuration(maxDuration string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).MaxDuration = &maxDuration

	return builder
}

// Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
func (builder *QueryBuilder) ServiceMapQuery(serviceMapQuery StringOrArrayOfString) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).ServiceMapQuery = &serviceMapQuery

	return builder
}

// Use service.namespace in addition to service.name to uniquely identify a service.
func (builder *QueryBuilder) ServiceMapIncludeNamespace(serviceMapIncludeNamespace bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).ServiceMapIncludeNamespace = &serviceMapIncludeNamespace

	return builder
}

// Defines the maximum number of traces that are returned from Tempo
func (builder *QueryBuilder) Limit(limit int64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).Limit = &limit

	return builder
}

// Defines the maximum number of spans per spanset that are returned from Tempo
func (builder *QueryBuilder) Spss(spss int64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).Spss = &spss

	return builder
}

func (builder *QueryBuilder) Filters(filters []cog.Builder[TraceqlFilter]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	filtersResources := make([]TraceqlFilter, 0, len(filters))
	for _, r1 := range filters {
		filtersDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		filtersResources = append(filtersResources, filtersDepth1)
	}
	builder.internal.Spec.(*TempoQuery).Filters = filtersResources

	return builder
}

// Filters that are used to query the metrics summary
func (builder *QueryBuilder) GroupBy(groupBy []cog.Builder[TraceqlFilter]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	groupByResources := make([]TraceqlFilter, 0, len(groupBy))
	for _, r1 := range groupBy {
		groupByDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		groupByResources = append(groupByResources, groupByDepth1)
	}
	builder.internal.Spec.(*TempoQuery).GroupBy = groupByResources

	return builder
}

// The type of the table that is used to display the search results
func (builder *QueryBuilder) TableType(tableType SearchTableType) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).TableType = &tableType

	return builder
}

// For metric queries, the step size to use
func (builder *QueryBuilder) Step(step string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).Step = &step

	return builder
}

// For metric queries, how many exemplars to request, 0 means no exemplars
func (builder *QueryBuilder) Exemplars(exemplars int64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).Exemplars = &exemplars

	return builder
}

// For metric queries, whether to run instant or range queries
func (builder *QueryBuilder) MetricsQueryType(metricsQueryType MetricsQueryType) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewTempoQuery()
	}
	builder.internal.Spec.(*TempoQuery).MetricsQueryType = &metricsQueryType

	return builder
}
