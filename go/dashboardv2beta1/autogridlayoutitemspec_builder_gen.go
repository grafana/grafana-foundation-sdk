// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AutoGridLayoutItemSpec] = (*AutoGridLayoutItemSpecBuilder)(nil)

type AutoGridLayoutItemSpecBuilder struct {
	internal *AutoGridLayoutItemSpec
	errors   cog.BuildErrors
}

func NewAutoGridLayoutItemSpecBuilder() *AutoGridLayoutItemSpecBuilder {
	resource := NewAutoGridLayoutItemSpec()
	builder := &AutoGridLayoutItemSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AutoGridLayoutItemSpecBuilder) Build() (AutoGridLayoutItemSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return AutoGridLayoutItemSpec{}, err
	}

	if len(builder.errors) > 0 {
		return AutoGridLayoutItemSpec{}, cog.MakeBuildErrors("dashboardv2beta1.autoGridLayoutItemSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AutoGridLayoutItemSpecBuilder) Element(element cog.Builder[ElementReference]) *AutoGridLayoutItemSpecBuilder {
	elementResource, err := element.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Element = elementResource

	return builder
}

func (builder *AutoGridLayoutItemSpecBuilder) Repeat(repeat cog.Builder[AutoGridRepeatOptions]) *AutoGridLayoutItemSpecBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Repeat = &repeatResource

	return builder
}

func (builder *AutoGridLayoutItemSpecBuilder) ConditionalRendering(conditionalRendering cog.Builder[ConditionalRenderingGroupKind]) *AutoGridLayoutItemSpecBuilder {
	conditionalRenderingResource, err := conditionalRendering.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.ConditionalRendering = &conditionalRenderingResource

	return builder
}
