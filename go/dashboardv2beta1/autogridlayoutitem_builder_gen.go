// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AutoGridLayoutItemKind] = (*AutoGridLayoutItemBuilder)(nil)

type AutoGridLayoutItemBuilder struct {
	internal *AutoGridLayoutItemKind
	errors   cog.BuildErrors
}

func NewAutoGridLayoutItemBuilder(name string) *AutoGridLayoutItemBuilder {
	resource := NewAutoGridLayoutItemKind()
	builder := &AutoGridLayoutItemBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "AutoGridLayoutItem"
	builder.internal.Spec.Element.Kind = "ElementReference"
	builder.internal.Spec.Element.Name = name

	return builder
}

func (builder *AutoGridLayoutItemBuilder) Build() (AutoGridLayoutItemKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return AutoGridLayoutItemKind{}, err
	}

	if len(builder.errors) > 0 {
		return AutoGridLayoutItemKind{}, cog.MakeBuildErrors("dashboardv2beta1.autoGridLayoutItem", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AutoGridLayoutItemBuilder) RecordError(path string, err error) *AutoGridLayoutItemBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AutoGridLayoutItemBuilder) Element(element cog.Builder[ElementReference]) *AutoGridLayoutItemBuilder {
	elementResource, err := element.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Element = elementResource

	return builder
}

func (builder *AutoGridLayoutItemBuilder) Repeat(repeat cog.Builder[AutoGridRepeatOptions]) *AutoGridLayoutItemBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Repeat = &repeatResource

	return builder
}

func (builder *AutoGridLayoutItemBuilder) ConditionalRendering(conditionalRendering cog.Builder[ConditionalRenderingGroupKind]) *AutoGridLayoutItemBuilder {
	conditionalRenderingResource, err := conditionalRendering.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.ConditionalRendering = &conditionalRenderingResource

	return builder
}

func (builder *AutoGridLayoutItemBuilder) Name(name string) *AutoGridLayoutItemBuilder {
	builder.internal.Spec.Element.Name = name

	return builder
}
