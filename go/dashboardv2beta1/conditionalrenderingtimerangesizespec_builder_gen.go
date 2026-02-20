// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConditionalRenderingTimeRangeSizeSpec] = (*ConditionalRenderingTimeRangeSizeSpecBuilder)(nil)

type ConditionalRenderingTimeRangeSizeSpecBuilder struct {
	internal *ConditionalRenderingTimeRangeSizeSpec
	errors   cog.BuildErrors
}

func NewConditionalRenderingTimeRangeSizeSpecBuilder() *ConditionalRenderingTimeRangeSizeSpecBuilder {
	resource := NewConditionalRenderingTimeRangeSizeSpec()
	builder := &ConditionalRenderingTimeRangeSizeSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ConditionalRenderingTimeRangeSizeSpecBuilder) Build() (ConditionalRenderingTimeRangeSizeSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConditionalRenderingTimeRangeSizeSpec{}, err
	}

	if len(builder.errors) > 0 {
		return ConditionalRenderingTimeRangeSizeSpec{}, cog.MakeBuildErrors("dashboardv2beta1.conditionalRenderingTimeRangeSizeSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConditionalRenderingTimeRangeSizeSpecBuilder) RecordError(path string, err error) *ConditionalRenderingTimeRangeSizeSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ConditionalRenderingTimeRangeSizeSpecBuilder) Value(value string) *ConditionalRenderingTimeRangeSizeSpecBuilder {
	builder.internal.Value = value

	return builder
}
