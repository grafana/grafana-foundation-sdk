// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsTimeRange] = (*ExprTypeClassicConditionsTimeRangeBuilder)(nil)

type ExprTypeClassicConditionsTimeRangeBuilder struct {
	internal *ExprTypeClassicConditionsTimeRange
	errors   map[string]cog.BuildErrors
}

func NewExprTypeClassicConditionsTimeRangeBuilder() *ExprTypeClassicConditionsTimeRangeBuilder {
	resource := &ExprTypeClassicConditionsTimeRange{}
	builder := &ExprTypeClassicConditionsTimeRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeClassicConditionsTimeRangeBuilder) Build() (ExprTypeClassicConditionsTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeClassicConditionsTimeRange{}, err
	}

	return *builder.internal, nil
}

// From is the start time of the query.
func (builder *ExprTypeClassicConditionsTimeRangeBuilder) From(from string) *ExprTypeClassicConditionsTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *ExprTypeClassicConditionsTimeRangeBuilder) To(to string) *ExprTypeClassicConditionsTimeRangeBuilder {
	builder.internal.To = to

	return builder
}

func (builder *ExprTypeClassicConditionsTimeRangeBuilder) applyDefaults() {
	builder.From("now-6h")
	builder.To("now")
}
