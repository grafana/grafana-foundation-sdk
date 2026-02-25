// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AutoGridLayoutKind] = (*AutoGridBuilder)(nil)

type AutoGridBuilder struct {
	internal *AutoGridLayoutKind
	errors   cog.BuildErrors
}

func NewAutoGridBuilder() *AutoGridBuilder {
	resource := NewAutoGridLayoutKind()
	builder := &AutoGridBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "AutoGridLayout"

	return builder
}

func AutoGrid() *AutoGridBuilder {
	builder := NewAutoGridBuilder()

	return builder
}

func (builder *AutoGridBuilder) Build() (AutoGridLayoutKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return AutoGridLayoutKind{}, err
	}

	if len(builder.errors) > 0 {
		return AutoGridLayoutKind{}, cog.MakeBuildErrors("dashboardv2beta1.autoGrid", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AutoGridBuilder) RecordError(path string, err error) *AutoGridBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AutoGridBuilder) MaxColumnCount(maxColumnCount float64) *AutoGridBuilder {
	builder.internal.Spec.MaxColumnCount = &maxColumnCount

	return builder
}

func (builder *AutoGridBuilder) ColumnWidthMode(columnWidthMode AutoGridLayoutSpecColumnWidthMode) *AutoGridBuilder {
	builder.internal.Spec.ColumnWidthMode = columnWidthMode

	return builder
}

func (builder *AutoGridBuilder) ColumnWidth(columnWidth float64) *AutoGridBuilder {
	builder.internal.Spec.ColumnWidth = &columnWidth

	return builder
}

func (builder *AutoGridBuilder) RowHeightMode(rowHeightMode AutoGridLayoutSpecRowHeightMode) *AutoGridBuilder {
	builder.internal.Spec.RowHeightMode = rowHeightMode

	return builder
}

func (builder *AutoGridBuilder) RowHeight(rowHeight float64) *AutoGridBuilder {
	builder.internal.Spec.RowHeight = &rowHeight

	return builder
}

func (builder *AutoGridBuilder) FillScreen(fillScreen bool) *AutoGridBuilder {
	builder.internal.Spec.FillScreen = &fillScreen

	return builder
}

func (builder *AutoGridBuilder) Items(items []cog.Builder[AutoGridLayoutItemKind]) *AutoGridBuilder {
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

func (builder *AutoGridBuilder) Item(item cog.Builder[AutoGridLayoutItemKind]) *AutoGridBuilder {
	itemResource, err := item.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Items = append(builder.internal.Spec.Items, itemResource)

	return builder
}

func (builder *AutoGridBuilder) WithItem(name string) *AutoGridBuilder {
	builder.internal.Spec.Items = append(builder.internal.Spec.Items, AutoGridLayoutItemKind{
		Kind: "AutoGridLayoutItem",
		Spec: AutoGridLayoutItemSpec{
			Element: ElementReference{
				Kind: "ElementReference",
				Name: name,
			},
		},
	})

	return builder
}
