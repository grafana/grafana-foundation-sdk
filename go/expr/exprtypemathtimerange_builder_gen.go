// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeMathTimeRange] = (*ExprTypeMathTimeRangeBuilder)(nil)

type ExprTypeMathTimeRangeBuilder struct {
	internal *ExprTypeMathTimeRange
	errors   map[string]cog.BuildErrors
}

func NewExprTypeMathTimeRangeBuilder() *ExprTypeMathTimeRangeBuilder {
	resource := &ExprTypeMathTimeRange{}
	builder := &ExprTypeMathTimeRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeMathTimeRangeBuilder) Build() (ExprTypeMathTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeMathTimeRange{}, err
	}

	return *builder.internal, nil
}

// From is the start time of the query.
func (builder *ExprTypeMathTimeRangeBuilder) From(from string) *ExprTypeMathTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *ExprTypeMathTimeRangeBuilder) To(to string) *ExprTypeMathTimeRangeBuilder {
	builder.internal.To = to

	return builder
}

func (builder *ExprTypeMathTimeRangeBuilder) applyDefaults() {
	builder.From("now-6h")
	builder.To("now")
}
