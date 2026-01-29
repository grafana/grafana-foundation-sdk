// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TabsLayoutTabSpec] = (*TabsLayoutTabSpecBuilder)(nil)

type TabsLayoutTabSpecBuilder struct {
	internal *TabsLayoutTabSpec
	errors   cog.BuildErrors
}

func NewTabsLayoutTabSpecBuilder() *TabsLayoutTabSpecBuilder {
	resource := NewTabsLayoutTabSpec()
	builder := &TabsLayoutTabSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TabsLayoutTabSpecBuilder) Build() (TabsLayoutTabSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return TabsLayoutTabSpec{}, err
	}

	if len(builder.errors) > 0 {
		return TabsLayoutTabSpec{}, cog.MakeBuildErrors("dashboardv2beta1.tabsLayoutTabSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TabsLayoutTabSpecBuilder) Title(title string) *TabsLayoutTabSpecBuilder {
	builder.internal.Title = &title

	return builder
}

func (builder *TabsLayoutTabSpecBuilder) Layout(layout GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) *TabsLayoutTabSpecBuilder {
	builder.internal.Layout = layout

	return builder
}

func (builder *TabsLayoutTabSpecBuilder) ConditionalRendering(conditionalRendering cog.Builder[ConditionalRenderingGroupKind]) *TabsLayoutTabSpecBuilder {
	conditionalRenderingResource, err := conditionalRendering.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.ConditionalRendering = &conditionalRenderingResource

	return builder
}

func (builder *TabsLayoutTabSpecBuilder) Repeat(repeat cog.Builder[TabRepeatOptions]) *TabsLayoutTabSpecBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Repeat = &repeatResource

	return builder
}
