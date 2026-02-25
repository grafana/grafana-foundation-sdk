// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RowsLayoutRowKind] = (*RowBuilder)(nil)

type RowBuilder struct {
	internal *RowsLayoutRowKind
	errors   cog.BuildErrors
}

func NewRowBuilder() *RowBuilder {
	resource := NewRowsLayoutRowKind()
	builder := &RowBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "RowsLayoutRow"

	return builder
}

func Row(title string) *RowBuilder {
	builder := NewRowBuilder()

	builder.Title(title)

	return builder
}

func (builder *RowBuilder) Build() (RowsLayoutRowKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return RowsLayoutRowKind{}, err
	}

	if len(builder.errors) > 0 {
		return RowsLayoutRowKind{}, cog.MakeBuildErrors("dashboardv2beta1.row", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RowBuilder) RecordError(path string, err error) *RowBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *RowBuilder) Title(title string) *RowBuilder {
	builder.internal.Spec.Title = &title

	return builder
}

func (builder *RowBuilder) Collapse(collapse bool) *RowBuilder {
	builder.internal.Spec.Collapse = &collapse

	return builder
}

func (builder *RowBuilder) HideHeader(hideHeader bool) *RowBuilder {
	builder.internal.Spec.HideHeader = &hideHeader

	return builder
}

func (builder *RowBuilder) FillScreen(fillScreen bool) *RowBuilder {
	builder.internal.Spec.FillScreen = &fillScreen

	return builder
}

func (builder *RowBuilder) ConditionalRendering(conditionalRendering cog.Builder[ConditionalRenderingGroupKind]) *RowBuilder {
	conditionalRenderingResource, err := conditionalRendering.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.ConditionalRendering = &conditionalRenderingResource

	return builder
}

func (builder *RowBuilder) Repeat(repeat cog.Builder[RowRepeatOptions]) *RowBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Repeat = &repeatResource

	return builder
}

func (builder *RowBuilder) GridLayout(gridLayoutKind cog.Builder[GridLayoutKind]) *RowBuilder {
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

func (builder *RowBuilder) AutoGridLayout(autoGridLayoutKind cog.Builder[AutoGridLayoutKind]) *RowBuilder {
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

func (builder *RowBuilder) TabsLayout(tabsLayoutKind cog.Builder[TabsLayoutKind]) *RowBuilder {
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

func (builder *RowBuilder) RowsLayout(rowsLayoutKind cog.Builder[RowsLayoutKind]) *RowBuilder {
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
