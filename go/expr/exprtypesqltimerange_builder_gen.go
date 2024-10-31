// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeSqlTimeRange] = (*ExprTypeSqlTimeRangeBuilder)(nil)

type ExprTypeSqlTimeRangeBuilder struct {
	internal *ExprTypeSqlTimeRange
	errors   map[string]cog.BuildErrors
}

func NewExprTypeSqlTimeRangeBuilder() *ExprTypeSqlTimeRangeBuilder {
	resource := &ExprTypeSqlTimeRange{}
	builder := &ExprTypeSqlTimeRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeSqlTimeRangeBuilder) Build() (ExprTypeSqlTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeSqlTimeRange{}, err
	}

	return *builder.internal, nil
}

// From is the start time of the query.
func (builder *ExprTypeSqlTimeRangeBuilder) From(from string) *ExprTypeSqlTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *ExprTypeSqlTimeRangeBuilder) To(to string) *ExprTypeSqlTimeRangeBuilder {
	builder.internal.To = to

	return builder
}

func (builder *ExprTypeSqlTimeRangeBuilder) applyDefaults() {
	builder.From("now-6h")
	builder.To("now")
}
