// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeResampleTimeRange] = (*ExprTypeResampleTimeRangeBuilder)(nil)

type ExprTypeResampleTimeRangeBuilder struct {
	internal *ExprTypeResampleTimeRange
	errors   map[string]cog.BuildErrors
}

func NewExprTypeResampleTimeRangeBuilder() *ExprTypeResampleTimeRangeBuilder {
	resource := NewExprTypeResampleTimeRange()
	builder := &ExprTypeResampleTimeRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ExprTypeResampleTimeRangeBuilder) Build() (ExprTypeResampleTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeResampleTimeRange{}, err
	}

	return *builder.internal, nil
}

// From is the start time of the query.
func (builder *ExprTypeResampleTimeRangeBuilder) From(from string) *ExprTypeResampleTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *ExprTypeResampleTimeRangeBuilder) To(to string) *ExprTypeResampleTimeRangeBuilder {
	builder.internal.To = to

	return builder
}
