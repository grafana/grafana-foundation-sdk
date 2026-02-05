// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AutoGridLayoutSpec] = (*AutoGridLayoutSpecBuilder)(nil)

type AutoGridLayoutSpecBuilder struct {
	internal *AutoGridLayoutSpec
	errors   cog.BuildErrors
}

func NewAutoGridLayoutSpecBuilder() *AutoGridLayoutSpecBuilder {
	resource := NewAutoGridLayoutSpec()
	builder := &AutoGridLayoutSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AutoGridLayoutSpecBuilder) Build() (AutoGridLayoutSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return AutoGridLayoutSpec{}, err
	}

	if len(builder.errors) > 0 {
		return AutoGridLayoutSpec{}, cog.MakeBuildErrors("dashboardv2beta1.autoGridLayoutSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AutoGridLayoutSpecBuilder) MaxColumnCount(maxColumnCount float64) *AutoGridLayoutSpecBuilder {
	builder.internal.MaxColumnCount = &maxColumnCount

	return builder
}

func (builder *AutoGridLayoutSpecBuilder) ColumnWidthMode(columnWidthMode AutoGridLayoutSpecColumnWidthMode) *AutoGridLayoutSpecBuilder {
	builder.internal.ColumnWidthMode = columnWidthMode

	return builder
}

func (builder *AutoGridLayoutSpecBuilder) ColumnWidth(columnWidth float64) *AutoGridLayoutSpecBuilder {
	builder.internal.ColumnWidth = &columnWidth

	return builder
}

func (builder *AutoGridLayoutSpecBuilder) RowHeightMode(rowHeightMode AutoGridLayoutSpecRowHeightMode) *AutoGridLayoutSpecBuilder {
	builder.internal.RowHeightMode = rowHeightMode

	return builder
}

func (builder *AutoGridLayoutSpecBuilder) RowHeight(rowHeight float64) *AutoGridLayoutSpecBuilder {
	builder.internal.RowHeight = &rowHeight

	return builder
}

func (builder *AutoGridLayoutSpecBuilder) FillScreen(fillScreen bool) *AutoGridLayoutSpecBuilder {
	builder.internal.FillScreen = &fillScreen

	return builder
}

func (builder *AutoGridLayoutSpecBuilder) Items(items []cog.Builder[AutoGridLayoutItemKind]) *AutoGridLayoutSpecBuilder {
	itemsResources := make([]AutoGridLayoutItemKind, 0, len(items))
	for _, r1 := range items {
		itemsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		itemsResources = append(itemsResources, itemsDepth1)
	}
	builder.internal.Items = itemsResources

	return builder
}
