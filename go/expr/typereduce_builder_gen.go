// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*TypeReduceBuilder)(nil)

type TypeReduceBuilder struct {
	internal *TypeReduce
	errors   map[string]cog.BuildErrors
}

func NewTypeReduceBuilder() *TypeReduceBuilder {
	resource := &TypeReduce{}
	builder := &TypeReduceBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "reduce"

	return builder
}

func (builder *TypeReduceBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return TypeReduce{}, err
	}

	return *builder.internal, nil
}

// The datasource
func (builder *TypeReduceBuilder) Datasource(datasource dashboard.DataSourceRef) *TypeReduceBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Reference to single query result
func (builder *TypeReduceBuilder) Expression(expression string) *TypeReduceBuilder {
	builder.internal.Expression = expression

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// NOTE: this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *TypeReduceBuilder) Hide(hide bool) *TypeReduceBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *TypeReduceBuilder) IntervalMs(intervalMs float64) *TypeReduceBuilder {
	builder.internal.IntervalMs = &intervalMs

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *TypeReduceBuilder) MaxDataPoints(maxDataPoints int64) *TypeReduceBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *TypeReduceBuilder) QueryType(queryType string) *TypeReduceBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// The reducer
// Possible enum values:
//   - `"sum"`
//   - `"mean"`
//   - `"min"`
//   - `"max"`
//   - `"count"`
//   - `"last"`
//   - `"median"`
func (builder *TypeReduceBuilder) Reducer(reducer TypeReduceReducer) *TypeReduceBuilder {
	builder.internal.Reducer = reducer

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *TypeReduceBuilder) RefId(refId string) *TypeReduceBuilder {
	builder.internal.RefId = refId

	return builder
}

// Optionally define expected query result behavior
func (builder *TypeReduceBuilder) ResultAssertions(resultAssertions cog.Builder[ExprTypeReduceResultAssertions]) *TypeReduceBuilder {
	resultAssertionsResource, err := resultAssertions.Build()
	if err != nil {
		builder.errors["resultAssertions"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.ResultAssertions = &resultAssertionsResource

	return builder
}

// Reducer Options
func (builder *TypeReduceBuilder) Settings(settings cog.Builder[ExprTypeReduceSettings]) *TypeReduceBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *TypeReduceBuilder) TimeRange(timeRange cog.Builder[ExprTypeReduceTimeRange]) *TypeReduceBuilder {
	timeRangeResource, err := timeRange.Build()
	if err != nil {
		builder.errors["timeRange"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.TimeRange = &timeRangeResource

	return builder
}

func (builder *TypeReduceBuilder) applyDefaults() {
}
