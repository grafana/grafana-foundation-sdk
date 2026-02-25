// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TabsLayoutTabKind] = (*TabBuilder)(nil)

type TabBuilder struct {
	internal *TabsLayoutTabKind
	errors   cog.BuildErrors
}

func NewTabBuilder() *TabBuilder {
	resource := NewTabsLayoutTabKind()
	builder := &TabBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "TabsLayoutTab"

	return builder
}

func Tab(title string) *TabBuilder {
	builder := NewTabBuilder()

	builder.Title(title)

	return builder
}

func (builder *TabBuilder) Build() (TabsLayoutTabKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return TabsLayoutTabKind{}, err
	}

	if len(builder.errors) > 0 {
		return TabsLayoutTabKind{}, cog.MakeBuildErrors("dashboardv2beta1.tab", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TabBuilder) RecordError(path string, err error) *TabBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *TabBuilder) Title(title string) *TabBuilder {
	builder.internal.Spec.Title = &title

	return builder
}

func (builder *TabBuilder) GridLayout(gridLayoutKind cog.Builder[GridLayoutKind]) *TabBuilder {
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

func (builder *TabBuilder) RowsLayout(rowsLayoutKind cog.Builder[RowsLayoutKind]) *TabBuilder {
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

func (builder *TabBuilder) AutoGridLayout(autoGridLayoutKind cog.Builder[AutoGridLayoutKind]) *TabBuilder {
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

func (builder *TabBuilder) TabsLayout(tabsLayoutKind cog.Builder[TabsLayoutKind]) *TabBuilder {
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

func (builder *TabBuilder) ConditionalRendering(conditionalRendering cog.Builder[ConditionalRenderingGroupKind]) *TabBuilder {
	conditionalRenderingResource, err := conditionalRendering.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.ConditionalRendering = &conditionalRenderingResource

	return builder
}

func (builder *TabBuilder) Repeat(repeat cog.Builder[TabRepeatOptions]) *TabBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Repeat = &repeatResource

	return builder
}
