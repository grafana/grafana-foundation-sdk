// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AutoGridLayoutItemKind] = (*AutoGridItemBuilder)(nil)

type AutoGridItemBuilder struct {
	internal *AutoGridLayoutItemKind
	errors   cog.BuildErrors
}

func NewAutoGridItemBuilder() *AutoGridItemBuilder {
	resource := NewAutoGridLayoutItemKind()
	builder := &AutoGridItemBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "AutoGridLayoutItem"
	builder.internal.Spec.Element.Kind = "ElementReference"

	return builder
}

func AutoGridItem(name string) *AutoGridItemBuilder {
	builder := NewAutoGridItemBuilder()

	builder.Name(name)

	return builder
}

func (builder *AutoGridItemBuilder) Build() (AutoGridLayoutItemKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return AutoGridLayoutItemKind{}, err
	}

	if len(builder.errors) > 0 {
		return AutoGridLayoutItemKind{}, cog.MakeBuildErrors("dashboardv2beta1.autoGridItem", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AutoGridItemBuilder) RecordError(path string, err error) *AutoGridItemBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AutoGridItemBuilder) Repeat(repeat cog.Builder[AutoGridRepeatOptions]) *AutoGridItemBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Repeat = &repeatResource

	return builder
}

func (builder *AutoGridItemBuilder) ConditionalRendering(conditionalRendering cog.Builder[ConditionalRenderingGroupKind]) *AutoGridItemBuilder {
	conditionalRenderingResource, err := conditionalRendering.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.ConditionalRendering = &conditionalRenderingResource

	return builder
}

func (builder *AutoGridItemBuilder) Name(name string) *AutoGridItemBuilder {
	builder.internal.Spec.Element.Name = name

	return builder
}
