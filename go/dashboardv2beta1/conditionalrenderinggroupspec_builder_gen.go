// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConditionalRenderingGroupSpec] = (*ConditionalRenderingGroupSpecBuilder)(nil)

type ConditionalRenderingGroupSpecBuilder struct {
	internal *ConditionalRenderingGroupSpec
	errors   cog.BuildErrors
}

func NewConditionalRenderingGroupSpecBuilder() *ConditionalRenderingGroupSpecBuilder {
	resource := NewConditionalRenderingGroupSpec()
	builder := &ConditionalRenderingGroupSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ConditionalRenderingGroupSpecBuilder) Build() (ConditionalRenderingGroupSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConditionalRenderingGroupSpec{}, err
	}

	if len(builder.errors) > 0 {
		return ConditionalRenderingGroupSpec{}, cog.MakeBuildErrors("dashboardv2beta1.conditionalRenderingGroupSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConditionalRenderingGroupSpecBuilder) RecordError(path string, err error) *ConditionalRenderingGroupSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ConditionalRenderingGroupSpecBuilder) Visibility(visibility ConditionalRenderingGroupSpecVisibility) *ConditionalRenderingGroupSpecBuilder {
	builder.internal.Visibility = visibility

	return builder
}

func (builder *ConditionalRenderingGroupSpecBuilder) Condition(condition ConditionalRenderingGroupSpecCondition) *ConditionalRenderingGroupSpecBuilder {
	builder.internal.Condition = condition

	return builder
}

func (builder *ConditionalRenderingGroupSpecBuilder) Items(items []ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) *ConditionalRenderingGroupSpecBuilder {
	builder.internal.Items = items

	return builder
}
