// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConditionalRenderingGroupKind] = (*ConditionalRenderingGroupBuilder)(nil)

type ConditionalRenderingGroupBuilder struct {
	internal *ConditionalRenderingGroupKind
	errors   cog.BuildErrors
}

func NewConditionalRenderingGroupBuilder() *ConditionalRenderingGroupBuilder {
	resource := NewConditionalRenderingGroupKind()
	builder := &ConditionalRenderingGroupBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "ConditionalRenderingGroup"

	return builder
}

func (builder *ConditionalRenderingGroupBuilder) Build() (ConditionalRenderingGroupKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConditionalRenderingGroupKind{}, err
	}

	if len(builder.errors) > 0 {
		return ConditionalRenderingGroupKind{}, cog.MakeBuildErrors("dashboardv2beta1.conditionalRenderingGroup", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConditionalRenderingGroupBuilder) RecordError(path string, err error) *ConditionalRenderingGroupBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ConditionalRenderingGroupBuilder) Visibility(visibility ConditionalRenderingGroupSpecVisibility) *ConditionalRenderingGroupBuilder {
	builder.internal.Spec.Visibility = visibility

	return builder
}

func (builder *ConditionalRenderingGroupBuilder) Condition(condition ConditionalRenderingGroupSpecCondition) *ConditionalRenderingGroupBuilder {
	builder.internal.Spec.Condition = condition

	return builder
}

func (builder *ConditionalRenderingGroupBuilder) Items(items []ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) *ConditionalRenderingGroupBuilder {
	builder.internal.Spec.Items = items

	return builder
}

func (builder *ConditionalRenderingGroupBuilder) Item(item ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind) *ConditionalRenderingGroupBuilder {
	builder.internal.Spec.Items = append(builder.internal.Spec.Items, item)

	return builder
}
