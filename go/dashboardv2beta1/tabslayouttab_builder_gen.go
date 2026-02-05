// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TabsLayoutTabKind] = (*TabsLayoutTabBuilder)(nil)

type TabsLayoutTabBuilder struct {
	internal *TabsLayoutTabKind
	errors   cog.BuildErrors
}

func NewTabsLayoutTabBuilder(title string) *TabsLayoutTabBuilder {
	resource := NewTabsLayoutTabKind()
	builder := &TabsLayoutTabBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "TabsLayoutTab"
	builder.internal.Spec.Title = &title

	return builder
}

func (builder *TabsLayoutTabBuilder) Build() (TabsLayoutTabKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return TabsLayoutTabKind{}, err
	}

	if len(builder.errors) > 0 {
		return TabsLayoutTabKind{}, cog.MakeBuildErrors("dashboardv2beta1.tabsLayoutTab", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TabsLayoutTabBuilder) Title(title string) *TabsLayoutTabBuilder {
	builder.internal.Spec.Title = &title

	return builder
}

func (builder *TabsLayoutTabBuilder) GridLayout(gridLayoutKind cog.Builder[GridLayoutKind]) *TabsLayoutTabBuilder {
	gridLayoutKindResource, err := gridLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Layout = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{
		GridLayoutKind: &gridLayoutKindResource,
	}

	return builder
}

func (builder *TabsLayoutTabBuilder) RowsLayout(rowsLayoutKind cog.Builder[RowsLayoutKind]) *TabsLayoutTabBuilder {
	rowsLayoutKindResource, err := rowsLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Layout = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{
		RowsLayoutKind: &rowsLayoutKindResource,
	}

	return builder
}

func (builder *TabsLayoutTabBuilder) AutoGridLayout(autoGridLayoutKind cog.Builder[AutoGridLayoutKind]) *TabsLayoutTabBuilder {
	autoGridLayoutKindResource, err := autoGridLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Layout = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{
		AutoGridLayoutKind: &autoGridLayoutKindResource,
	}

	return builder
}

func (builder *TabsLayoutTabBuilder) TabsLayout(tabsLayoutKind cog.Builder[TabsLayoutKind]) *TabsLayoutTabBuilder {
	tabsLayoutKindResource, err := tabsLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Layout = GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind{
		TabsLayoutKind: &tabsLayoutKindResource,
	}

	return builder
}

func (builder *TabsLayoutTabBuilder) ConditionalRendering(conditionalRendering cog.Builder[ConditionalRenderingGroupKind]) *TabsLayoutTabBuilder {
	conditionalRenderingResource, err := conditionalRendering.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.ConditionalRendering = &conditionalRenderingResource

	return builder
}

func (builder *TabsLayoutTabBuilder) Repeat(repeat cog.Builder[TabRepeatOptions]) *TabsLayoutTabBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Repeat = &repeatResource

	return builder
}
