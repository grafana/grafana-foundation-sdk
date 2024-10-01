// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeThresholdTimeRange] = (*ExprTypeThresholdTimeRangeBuilder)(nil)

type ExprTypeThresholdTimeRangeBuilder struct {
	internal *ExprTypeThresholdTimeRange
	errors   map[string]cog.BuildErrors
}

func NewExprTypeThresholdTimeRangeBuilder() *ExprTypeThresholdTimeRangeBuilder {
	resource := &ExprTypeThresholdTimeRange{}
	builder := &ExprTypeThresholdTimeRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeThresholdTimeRangeBuilder) Build() (ExprTypeThresholdTimeRange, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeThresholdTimeRange", err)...)
	}

	if len(errs) != 0 {
		return ExprTypeThresholdTimeRange{}, errs
	}

	return *builder.internal, nil
}

// From is the start time of the query.
func (builder *ExprTypeThresholdTimeRangeBuilder) From(from string) *ExprTypeThresholdTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

// To is the end time of the query.
func (builder *ExprTypeThresholdTimeRangeBuilder) To(to string) *ExprTypeThresholdTimeRangeBuilder {
	builder.internal.To = to

	return builder
}

func (builder *ExprTypeThresholdTimeRangeBuilder) applyDefaults() {
	builder.From("now-6h")
	builder.To("now")
}
