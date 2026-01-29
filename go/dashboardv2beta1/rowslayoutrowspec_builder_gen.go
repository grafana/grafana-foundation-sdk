// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RowsLayoutRowSpec] = (*RowsLayoutRowSpecBuilder)(nil)

type RowsLayoutRowSpecBuilder struct {
	internal *RowsLayoutRowSpec
	errors   cog.BuildErrors
}

func NewRowsLayoutRowSpecBuilder() *RowsLayoutRowSpecBuilder {
	resource := NewRowsLayoutRowSpec()
	builder := &RowsLayoutRowSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RowsLayoutRowSpecBuilder) Build() (RowsLayoutRowSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return RowsLayoutRowSpec{}, err
	}

	if len(builder.errors) > 0 {
		return RowsLayoutRowSpec{}, cog.MakeBuildErrors("dashboardv2beta1.rowsLayoutRowSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RowsLayoutRowSpecBuilder) Title(title string) *RowsLayoutRowSpecBuilder {
	builder.internal.Title = &title

	return builder
}

func (builder *RowsLayoutRowSpecBuilder) Collapse(collapse bool) *RowsLayoutRowSpecBuilder {
	builder.internal.Collapse = &collapse

	return builder
}

func (builder *RowsLayoutRowSpecBuilder) HideHeader(hideHeader bool) *RowsLayoutRowSpecBuilder {
	builder.internal.HideHeader = &hideHeader

	return builder
}

func (builder *RowsLayoutRowSpecBuilder) FillScreen(fillScreen bool) *RowsLayoutRowSpecBuilder {
	builder.internal.FillScreen = &fillScreen

	return builder
}

func (builder *RowsLayoutRowSpecBuilder) ConditionalRendering(conditionalRendering cog.Builder[ConditionalRenderingGroupKind]) *RowsLayoutRowSpecBuilder {
	conditionalRenderingResource, err := conditionalRendering.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.ConditionalRendering = &conditionalRenderingResource

	return builder
}

func (builder *RowsLayoutRowSpecBuilder) Repeat(repeat cog.Builder[RowRepeatOptions]) *RowsLayoutRowSpecBuilder {
	repeatResource, err := repeat.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Repeat = &repeatResource

	return builder
}

func (builder *RowsLayoutRowSpecBuilder) Layout(layout GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) *RowsLayoutRowSpecBuilder {
	builder.internal.Layout = layout

	return builder
}
