// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RowsLayoutRowKind] = (*RowsLayoutRowBuilder)(nil)

type RowsLayoutRowBuilder struct {
	internal *RowsLayoutRowKind
	errors   cog.BuildErrors
}

func NewRowsLayoutRowBuilder(title string) *RowsLayoutRowBuilder {
	resource := NewRowsLayoutRowKind()
	builder := &RowsLayoutRowBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "RowsLayoutRow"
	builder.internal.Spec.Title = &title

	return builder
}

func (builder *RowsLayoutRowBuilder) Build() (RowsLayoutRowKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return RowsLayoutRowKind{}, err
	}

	if len(builder.errors) > 0 {
		return RowsLayoutRowKind{}, cog.MakeBuildErrors("dashboardv2beta1.rowsLayoutRow", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RowsLayoutRowBuilder) Title(title string) *RowsLayoutRowBuilder {
	builder.internal.Spec.Title = &title

	return builder
}

func (builder *RowsLayoutRowBuilder) Collapse(collapse bool) *RowsLayoutRowBuilder {
	builder.internal.Spec.Collapse = &collapse

	return builder
}

func (builder *RowsLayoutRowBuilder) HideHeader(hideHeader bool) *RowsLayoutRowBuilder {
	builder.internal.Spec.HideHeader = &hideHeader

	return builder
}

func (builder *RowsLayoutRowBuilder) FillScreen(fillScreen bool) *RowsLayoutRowBuilder {
	builder.internal.Spec.FillScreen = &fillScreen

	return builder
}

func (builder *RowsLayoutRowBuilder) ConditionalRendering(conditionalRendering cog.Builder[ConditionalRenderingGroupKind]) *RowsLayoutRowBuilder {
	conditionalRenderingResource, err := conditionalRendering.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.ConditionalRendering = &conditionalRenderingResource

	return builder
}

func (builder *RowsLayoutRowBuilder) Repeat(repeat cog.Builder[RowRepeatOptions]) *RowsLayoutRowBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Repeat = &repeatResource

	return builder
}

func (builder *RowsLayoutRowBuilder) GridLayout(gridLayoutKind cog.Builder[GridLayoutKind]) *RowsLayoutRowBuilder {
	gridLayoutKindResource, err := gridLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Layout = GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind{
		GridLayoutKind: &gridLayoutKindResource,
	}

	return builder
}

func (builder *RowsLayoutRowBuilder) AutoGridLayout(autoGridLayoutKind cog.Builder[AutoGridLayoutKind]) *RowsLayoutRowBuilder {
	autoGridLayoutKindResource, err := autoGridLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Layout = GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind{
		AutoGridLayoutKind: &autoGridLayoutKindResource,
	}

	return builder
}

func (builder *RowsLayoutRowBuilder) TabsLayout(tabsLayoutKind cog.Builder[TabsLayoutKind]) *RowsLayoutRowBuilder {
	tabsLayoutKindResource, err := tabsLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Layout = GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind{
		TabsLayoutKind: &tabsLayoutKindResource,
	}

	return builder
}

func (builder *RowsLayoutRowBuilder) RowsLayout(rowsLayoutKind cog.Builder[RowsLayoutKind]) *RowsLayoutRowBuilder {
	rowsLayoutKindResource, err := rowsLayoutKind.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Layout = GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind{
		RowsLayoutKind: &rowsLayoutKindResource,
	}

	return builder
}
