// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

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
	builder.internal.Group = "prometheus"

	return builder
}

func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2.DataQueryKind{}, cog.MakeBuildErrors("prometheus.queryV2", builder.errors)
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

// Additional Ad-hoc filters that take precedence over Scope on conflict.
func (builder *QueryV2Builder) AdhocFilters(adhocFilters []cog.Builder[AdhocFilters]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	adhocFiltersResources := make([]AdhocFilters, 0, len(adhocFilters))
	for _, r1 := range adhocFilters {
		adhocFiltersDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		adhocFiltersResources = append(adhocFiltersResources, adhocFiltersDepth1)
	}
	builder.internal.Spec.(*Dataquery).AdhocFilters = adhocFiltersResources

	return builder
}

// what we should show in the editor
// Possible enum values:
//   - `"builder"`
//   - `"code"`
func (builder *QueryV2Builder) EditorMode(editorMode QueryEditorMode) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).EditorMode = &editorMode

	return builder
}

// Execute an additional query to identify interesting raw samples relevant for the given expr
func (builder *QueryV2Builder) Exemplar(exemplar bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Exemplar = &exemplar

	return builder
}

// The actual expression/query that will be evaluated by Prometheus
func (builder *QueryV2Builder) Expr(expr string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Expr = expr

	return builder
}

// The response format
// Possible enum values:
//   - `"time_series"`
//   - `"table"`
//   - `"heatmap"`
func (builder *QueryV2Builder) Format(format PromQueryFormat) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Format = &format

	return builder
}

// Group By parameters to apply to aggregate expressions in the query
func (builder *QueryV2Builder) GroupByKeys(groupByKeys []string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).GroupByKeys = groupByKeys

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// NOTE: this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Hide = &hide

	return builder
}

// Returns only the latest value that Prometheus has scraped for the requested time series
func (builder *QueryV2Builder) Instant(instant bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Instant = &instant

	return builder
}

// Used to specify how many times to divide max data points by. We use max data points under query options
// See https://github.com/grafana/grafana/issues/48081
// Deprecated: use interval
func (builder *QueryV2Builder) IntervalFactor(intervalFactor int64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).IntervalFactor = &intervalFactor

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *QueryV2Builder) IntervalMs(intervalMs float64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).IntervalMs = &intervalMs

	return builder
}

// Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
func (builder *QueryV2Builder) LegendFormat(legendFormat string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).LegendFormat = &legendFormat

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *QueryV2Builder) MaxDataPoints(maxDataPoints int64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).MaxDataPoints = &maxDataPoints

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).QueryType = &queryType

	return builder
}

// Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
func (builder *QueryV2Builder) Range(rangeArg bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Range = &rangeArg

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RefId = &refId

	return builder
}

// Optionally define expected query result behavior
func (builder *QueryV2Builder) ResultAssertions(resultAssertions cog.Builder[ResultAssertions]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	resultAssertionsResource, err := resultAssertions.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Dataquery).ResultAssertions = &resultAssertionsResource

	return builder
}

// A set of filters applied to apply to the query
func (builder *QueryV2Builder) Scopes(scopes []cog.Builder[Scopes]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	scopesResources := make([]Scopes, 0, len(scopes))
	for _, r1 := range scopes {
		scopesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		scopesResources = append(scopesResources, scopesDepth1)
	}
	builder.internal.Spec.(*Dataquery).Scopes = scopesResources

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *QueryV2Builder) TimeRange(timeRange cog.Builder[TimeRange]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	timeRangeResource, err := timeRange.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Dataquery).TimeRange = &timeRangeResource

	return builder
}

// An additional lower limit for the step parameter of the Prometheus query and for the
// `$__interval` and `$__rate_interval` variables.
func (builder *QueryV2Builder) Interval(interval string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Interval = &interval

	return builder
}
