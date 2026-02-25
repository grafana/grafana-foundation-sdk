// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConditionalRenderingTimeRangeSizeKind] = (*ConditionalRenderingTimeRangeSizeBuilder)(nil)

type ConditionalRenderingTimeRangeSizeBuilder struct {
	internal *ConditionalRenderingTimeRangeSizeKind
	errors   cog.BuildErrors
}

func NewConditionalRenderingTimeRangeSizeBuilder() *ConditionalRenderingTimeRangeSizeBuilder {
	resource := NewConditionalRenderingTimeRangeSizeKind()
	builder := &ConditionalRenderingTimeRangeSizeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "ConditionalRenderingTimeRangeSize"

	return builder
}

func (builder *ConditionalRenderingTimeRangeSizeBuilder) Build() (ConditionalRenderingTimeRangeSizeKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConditionalRenderingTimeRangeSizeKind{}, err
	}

	if len(builder.errors) > 0 {
		return ConditionalRenderingTimeRangeSizeKind{}, cog.MakeBuildErrors("dashboardv2beta1.conditionalRenderingTimeRangeSize", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConditionalRenderingTimeRangeSizeBuilder) RecordError(path string, err error) *ConditionalRenderingTimeRangeSizeBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ConditionalRenderingTimeRangeSizeBuilder) Value(value string) *ConditionalRenderingTimeRangeSizeBuilder {
	builder.internal.Spec.Value = value

	return builder
}
