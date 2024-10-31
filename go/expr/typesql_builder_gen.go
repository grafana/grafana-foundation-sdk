// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*TypeSqlBuilder)(nil)

type TypeSqlBuilder struct {
	internal *TypeSql
	errors   map[string]cog.BuildErrors
}

func NewTypeSqlBuilder() *TypeSqlBuilder {
	resource := &TypeSql{}
	builder := &TypeSqlBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "sql"

	return builder
}

func (builder *TypeSqlBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return TypeSql{}, err
	}

	return *builder.internal, nil
}

// The datasource
func (builder *TypeSqlBuilder) Datasource(datasource dashboard.DataSourceRef) *TypeSqlBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

func (builder *TypeSqlBuilder) Expression(expression string) *TypeSqlBuilder {
	builder.internal.Expression = expression

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// NOTE: this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *TypeSqlBuilder) Hide(hide bool) *TypeSqlBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *TypeSqlBuilder) IntervalMs(intervalMs float64) *TypeSqlBuilder {
	builder.internal.IntervalMs = &intervalMs

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *TypeSqlBuilder) MaxDataPoints(maxDataPoints int64) *TypeSqlBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *TypeSqlBuilder) QueryType(queryType string) *TypeSqlBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *TypeSqlBuilder) RefId(refId string) *TypeSqlBuilder {
	builder.internal.RefId = refId

	return builder
}

// Optionally define expected query result behavior
func (builder *TypeSqlBuilder) ResultAssertions(resultAssertions cog.Builder[ExprTypeSqlResultAssertions]) *TypeSqlBuilder {
	resultAssertionsResource, err := resultAssertions.Build()
	if err != nil {
		builder.errors["resultAssertions"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.ResultAssertions = &resultAssertionsResource

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *TypeSqlBuilder) TimeRange(timeRange cog.Builder[ExprTypeSqlTimeRange]) *TypeSqlBuilder {
	timeRangeResource, err := timeRange.Build()
	if err != nil {
		builder.errors["timeRange"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.TimeRange = &timeRangeResource

	return builder
}

func (builder *TypeSqlBuilder) applyDefaults() {
}
