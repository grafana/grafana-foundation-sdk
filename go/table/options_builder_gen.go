// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package table

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[Options] = (*OptionsBuilder)(nil)

type OptionsBuilder struct {
	internal *Options
	errors   cog.BuildErrors
}

func NewOptionsBuilder() *OptionsBuilder {
	resource := NewOptions()
	builder := &OptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsBuilder) Build() (Options, error) {
	if err := builder.internal.Validate(); err != nil {
		return Options{}, err
	}

	if len(builder.errors) > 0 {
		return Options{}, cog.MakeBuildErrors("table.options", builder.errors)
	}

	return *builder.internal, nil
}

// Represents the index of the selected frame
func (builder *OptionsBuilder) FrameIndex(frameIndex float64) *OptionsBuilder {
	builder.internal.FrameIndex = frameIndex

	return builder
}

// Controls whether the panel should show the header
func (builder *OptionsBuilder) ShowHeader(showHeader bool) *OptionsBuilder {
	builder.internal.ShowHeader = showHeader

	return builder
}

// Controls whether the header should show icons for the column types
func (builder *OptionsBuilder) ShowTypeIcons(showTypeIcons bool) *OptionsBuilder {
	builder.internal.ShowTypeIcons = &showTypeIcons

	return builder
}

// Used to control row sorting
func (builder *OptionsBuilder) SortBy(sortBy []cog.Builder[common.TableSortByFieldState]) *OptionsBuilder {
	sortByResources := make([]common.TableSortByFieldState, 0, len(sortBy))
	for _, r1 := range sortBy {
		sortByDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		sortByResources = append(sortByResources, sortByDepth1)
	}
	builder.internal.SortBy = sortByResources

	return builder
}

// Controls footer options
func (builder *OptionsBuilder) Footer(footer cog.Builder[common.TableFooterOptions]) *OptionsBuilder {
	footerResource, err := footer.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Footer = &footerResource

	return builder
}

// Controls the height of the rows
func (builder *OptionsBuilder) CellHeight(cellHeight common.TableCellHeight) *OptionsBuilder {
	builder.internal.CellHeight = &cellHeight

	return builder
}
