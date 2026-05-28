// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package tempo

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

var _ cog.Builder[dashboardv2.DataQueryKind] = (*QueryV2Builder)(nil)

type QueryV2Builder struct {
	internal *dashboardv2.DataQueryKind
	errors   cog.BuildErrors
}

func NewQueryV2Builder() *QueryV2Builder {
	resource := dashboardv2.NewDataQueryKind()
	builder := &QueryV2Builder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"
	builder.internal.Group = "tempo"

	return builder
}

func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2.DataQueryKind{}, cog.MakeBuildErrors("tempo.queryV2", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryV2Builder) RecordError(path string, err error) *QueryV2Builder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryV2Builder) Version(version string) *QueryV2Builder {
	builder.internal.Version = version

	return builder
}

func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder {
	builder.internal.Labels = labels

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder {
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
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RefId = refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).QueryType = &queryType

	return builder
}

// TraceQL query or trace ID
func (builder *QueryV2Builder) Query(query string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Query = &query

	return builder
}

// @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
func (builder *QueryV2Builder) Search(search string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Search = &search

	return builder
}

// @deprecated Query traces by service name
func (builder *QueryV2Builder) ServiceName(serviceName string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).ServiceName = &serviceName

	return builder
}

// @deprecated Query traces by span name
func (builder *QueryV2Builder) SpanName(spanName string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).SpanName = &spanName

	return builder
}

// @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
func (builder *QueryV2Builder) MinDuration(minDuration string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).MinDuration = &minDuration

	return builder
}

// @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
func (builder *QueryV2Builder) MaxDuration(maxDuration string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).MaxDuration = &maxDuration

	return builder
}

// Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
func (builder *QueryV2Builder) ServiceMapQuery(serviceMapQuery StringOrArrayOfString) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).ServiceMapQuery = &serviceMapQuery

	return builder
}

// Use service.namespace in addition to service.name to uniquely identify a service.
func (builder *QueryV2Builder) ServiceMapIncludeNamespace(serviceMapIncludeNamespace bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).ServiceMapIncludeNamespace = &serviceMapIncludeNamespace

	return builder
}

// Defines the maximum number of traces that are returned from Tempo
func (builder *QueryV2Builder) Limit(limit int64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Limit = &limit

	return builder
}

// Defines the maximum number of spans per spanset that are returned from Tempo
func (builder *QueryV2Builder) Spss(spss int64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Spss = &spss

	return builder
}

func (builder *QueryV2Builder) Filters(filters []cog.Builder[TraceqlFilter]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
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
	builder.internal.Spec.(*Dataquery).Filters = filtersResources

	return builder
}

// Filters that are used to query the metrics summary
func (builder *QueryV2Builder) GroupBy(groupBy []cog.Builder[TraceqlFilter]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
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
	builder.internal.Spec.(*Dataquery).GroupBy = groupByResources

	return builder
}

// The type of the table that is used to display the search results
func (builder *QueryV2Builder) TableType(tableType SearchTableType) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).TableType = &tableType

	return builder
}

// For metric queries, the step size to use
func (builder *QueryV2Builder) Step(step string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Step = &step

	return builder
}

// For metric queries, how many exemplars to request, 0 means no exemplars
func (builder *QueryV2Builder) Exemplars(exemplars int64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Exemplars = &exemplars

	return builder
}

// For metric queries, whether to run instant or range queries
func (builder *QueryV2Builder) MetricsQueryType(metricsQueryType MetricsQueryType) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).MetricsQueryType = &metricsQueryType

	return builder
}
