// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

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
		return Dataquery{}, cog.MakeBuildErrors("prometheus.dataquery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DataqueryBuilder) RecordError(path string, err error) *DataqueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Additional Ad-hoc filters that take precedence over Scope on conflict.
func (builder *DataqueryBuilder) AdhocFilters(adhocFilters []cog.Builder[AdhocFilters]) *DataqueryBuilder {
	adhocFiltersResources := make([]AdhocFilters, 0, len(adhocFilters))
	for _, r1 := range adhocFilters {
		adhocFiltersDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		adhocFiltersResources = append(adhocFiltersResources, adhocFiltersDepth1)
	}
	builder.internal.AdhocFilters = adhocFiltersResources

	return builder
}

// The datasource
func (builder *DataqueryBuilder) Datasource(datasource common.DataSourceRef) *DataqueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// what we should show in the editor
// Possible enum values:
//   - `"builder"`
//   - `"code"`
func (builder *DataqueryBuilder) EditorMode(editorMode QueryEditorMode) *DataqueryBuilder {
	builder.internal.EditorMode = &editorMode

	return builder
}

// Execute an additional query to identify interesting raw samples relevant for the given expr
func (builder *DataqueryBuilder) Exemplar(exemplar bool) *DataqueryBuilder {
	builder.internal.Exemplar = &exemplar

	return builder
}

// The actual expression/query that will be evaluated by Prometheus
func (builder *DataqueryBuilder) Expr(expr string) *DataqueryBuilder {
	builder.internal.Expr = expr

	return builder
}

// The response format
// Possible enum values:
//   - `"time_series"`
//   - `"table"`
//   - `"heatmap"`
func (builder *DataqueryBuilder) Format(format PromQueryFormat) *DataqueryBuilder {
	builder.internal.Format = &format

	return builder
}

// Group By parameters to apply to aggregate expressions in the query
func (builder *DataqueryBuilder) GroupByKeys(groupByKeys []string) *DataqueryBuilder {
	builder.internal.GroupByKeys = groupByKeys

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// NOTE: this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Returns only the latest value that Prometheus has scraped for the requested time series
func (builder *DataqueryBuilder) Instant() *DataqueryBuilder {
	valInstant := true
	builder.internal.Instant = &valInstant
	valRange := false
	builder.internal.Range = &valRange

	return builder
}

// Used to specify how many times to divide max data points by. We use max data points under query options
// See https://github.com/grafana/grafana/issues/48081
// Deprecated: use interval
func (builder *DataqueryBuilder) IntervalFactor(intervalFactor int64) *DataqueryBuilder {
	builder.internal.IntervalFactor = &intervalFactor

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *DataqueryBuilder) IntervalMs(intervalMs float64) *DataqueryBuilder {
	builder.internal.IntervalMs = &intervalMs

	return builder
}

// Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
func (builder *DataqueryBuilder) LegendFormat(legendFormat string) *DataqueryBuilder {
	builder.internal.LegendFormat = &legendFormat

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *DataqueryBuilder) MaxDataPoints(maxDataPoints int64) *DataqueryBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
func (builder *DataqueryBuilder) Range() *DataqueryBuilder {
	valRange := true
	builder.internal.Range = &valRange
	valInstant := false
	builder.internal.Instant = &valInstant

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder {
	builder.internal.RefId = &refId

	return builder
}

// Optionally define expected query result behavior
func (builder *DataqueryBuilder) ResultAssertions(resultAssertions cog.Builder[ResultAssertions]) *DataqueryBuilder {
	resultAssertionsResource, err := resultAssertions.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.ResultAssertions = &resultAssertionsResource

	return builder
}

// A set of filters applied to apply to the query
func (builder *DataqueryBuilder) Scopes(scopes []cog.Builder[Scopes]) *DataqueryBuilder {
	scopesResources := make([]Scopes, 0, len(scopes))
	for _, r1 := range scopes {
		scopesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		scopesResources = append(scopesResources, scopesDepth1)
	}
	builder.internal.Scopes = scopesResources

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *DataqueryBuilder) TimeRange(timeRange cog.Builder[TimeRange]) *DataqueryBuilder {
	timeRangeResource, err := timeRange.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TimeRange = &timeRangeResource

	return builder
}

// An additional lower limit for the step parameter of the Prometheus query and for the
// `$__interval` and `$__rate_interval` variables.
func (builder *DataqueryBuilder) Interval(interval string) *DataqueryBuilder {
	builder.internal.Interval = &interval

	return builder
}

func (builder *DataqueryBuilder) RangeAndInstant() *DataqueryBuilder {
	valRange := true
	builder.internal.Range = &valRange
	valInstant := true
	builder.internal.Instant = &valInstant

	return builder
}
