// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeReduceTimeRange] = (*ExprTypeReduceTimeRangeBuilder)(nil)

type ExprTypeReduceTimeRangeBuilder struct {
	internal *ExprTypeReduceTimeRange
	errors   map[string]cog.BuildErrors
}

func NewExprTypeReduceTimeRangeBuilder() *ExprTypeReduceTimeRangeBuilder {
	resource := &ExprTypeReduceTimeRange{}
	builder := &ExprTypeReduceTimeRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeReduceTimeRangeBuilder) Build() (ExprTypeReduceTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeReduceTimeRange{}, err
	}

	return *builder.internal, nil
}

// From is the start time of the query.
func (builder *ExprTypeReduceTimeRangeBuilder) From(from string) *ExprTypeReduceTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *ExprTypeReduceTimeRangeBuilder) To(to string) *ExprTypeReduceTimeRangeBuilder {
	builder.internal.To = to

	return builder
}

func (builder *ExprTypeReduceTimeRangeBuilder) applyDefaults() {
	builder.From("now-6h")
	builder.To("now")
}
