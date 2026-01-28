// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AutoGridLayoutKind] = (*AutoGridLayoutBuilder)(nil)

type AutoGridLayoutBuilder struct {
	internal *AutoGridLayoutKind
	errors   cog.BuildErrors
}

func NewAutoGridLayoutBuilder() *AutoGridLayoutBuilder {
	resource := NewAutoGridLayoutKind()
	builder := &AutoGridLayoutBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "AutoGridLayout"

	return builder
}

func (builder *AutoGridLayoutBuilder) Build() (AutoGridLayoutKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return AutoGridLayoutKind{}, err
	}

	if len(builder.errors) > 0 {
		return AutoGridLayoutKind{}, cog.MakeBuildErrors("dashboardv2beta1.autoGridLayout", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AutoGridLayoutBuilder) MaxColumnCount(maxColumnCount float64) *AutoGridLayoutBuilder {
	builder.internal.Spec.MaxColumnCount = &maxColumnCount

	return builder
}

func (builder *AutoGridLayoutBuilder) ColumnWidthMode(columnWidthMode AutoGridLayoutSpecColumnWidthMode) *AutoGridLayoutBuilder {
	builder.internal.Spec.ColumnWidthMode = columnWidthMode

	return builder
}

func (builder *AutoGridLayoutBuilder) ColumnWidth(columnWidth float64) *AutoGridLayoutBuilder {
	builder.internal.Spec.ColumnWidth = &columnWidth

	return builder
}

func (builder *AutoGridLayoutBuilder) RowHeightMode(rowHeightMode AutoGridLayoutSpecRowHeightMode) *AutoGridLayoutBuilder {
	builder.internal.Spec.RowHeightMode = rowHeightMode

	return builder
}

func (builder *AutoGridLayoutBuilder) RowHeight(rowHeight float64) *AutoGridLayoutBuilder {
	builder.internal.Spec.RowHeight = &rowHeight

	return builder
}

func (builder *AutoGridLayoutBuilder) FillScreen(fillScreen bool) *AutoGridLayoutBuilder {
	builder.internal.Spec.FillScreen = &fillScreen

	return builder
}

func (builder *AutoGridLayoutBuilder) Items(items []cog.Builder[AutoGridLayoutItemKind]) *AutoGridLayoutBuilder {
	itemsResources := make([]AutoGridLayoutItemKind, 0, len(items))
	for _, r1 := range items {
		itemsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		itemsResources = append(itemsResources, itemsDepth1)
	}
	builder.internal.Spec.Items = itemsResources

	return builder
}

func (builder *AutoGridLayoutBuilder) Item(item cog.Builder[AutoGridLayoutItemKind]) *AutoGridLayoutBuilder {
	itemResource, err := item.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Items = append(builder.internal.Spec.Items, itemResource)

	return builder
}
