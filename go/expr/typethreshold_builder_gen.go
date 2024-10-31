// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*TypeThresholdBuilder)(nil)

type TypeThresholdBuilder struct {
	internal *TypeThreshold
	errors   map[string]cog.BuildErrors
}

func NewTypeThresholdBuilder() *TypeThresholdBuilder {
	resource := &TypeThreshold{}
	builder := &TypeThresholdBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "threshold"

	return builder
}

func (builder *TypeThresholdBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return TypeThreshold{}, err
	}

	return *builder.internal, nil
}

// Threshold Conditions
func (builder *TypeThresholdBuilder) Conditions(conditions []cog.Builder[ExprTypeThresholdConditions]) *TypeThresholdBuilder {
	conditionsResources := make([]ExprTypeThresholdConditions, 0, len(conditions))
	for _, r1 := range conditions {
		conditionsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["conditions"] = err.(cog.BuildErrors)
			return builder
		}
		conditionsResources = append(conditionsResources, conditionsDepth1)
	}
	builder.internal.Conditions = conditionsResources

	return builder
}

// The datasource
func (builder *TypeThresholdBuilder) Datasource(datasource dashboard.DataSourceRef) *TypeThresholdBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Reference to single query result
func (builder *TypeThresholdBuilder) Expression(expression string) *TypeThresholdBuilder {
	builder.internal.Expression = expression

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// NOTE: this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *TypeThresholdBuilder) Hide(hide bool) *TypeThresholdBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *TypeThresholdBuilder) IntervalMs(intervalMs float64) *TypeThresholdBuilder {
	builder.internal.IntervalMs = &intervalMs

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *TypeThresholdBuilder) MaxDataPoints(maxDataPoints int64) *TypeThresholdBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *TypeThresholdBuilder) QueryType(queryType string) *TypeThresholdBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *TypeThresholdBuilder) RefId(refId string) *TypeThresholdBuilder {
	builder.internal.RefId = refId

	return builder
}

// Optionally define expected query result behavior
func (builder *TypeThresholdBuilder) ResultAssertions(resultAssertions cog.Builder[ExprTypeThresholdResultAssertions]) *TypeThresholdBuilder {
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
func (builder *TypeThresholdBuilder) TimeRange(timeRange cog.Builder[ExprTypeThresholdTimeRange]) *TypeThresholdBuilder {
	timeRangeResource, err := timeRange.Build()
	if err != nil {
		builder.errors["timeRange"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.TimeRange = &timeRangeResource

	return builder
}

func (builder *TypeThresholdBuilder) applyDefaults() {
}
