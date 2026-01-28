// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statushistory

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
		return Options{}, cog.MakeBuildErrors("statushistory.options", builder.errors)
	}

	return *builder.internal, nil
}

// Set the height of the rows
func (builder *OptionsBuilder) RowHeight(rowHeight float32) *OptionsBuilder {
	builder.internal.RowHeight = rowHeight

	return builder
}

// Show values on the columns
func (builder *OptionsBuilder) ShowValue(showValue common.VisibilityMode) *OptionsBuilder {
	builder.internal.ShowValue = showValue

	return builder
}

func (builder *OptionsBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *OptionsBuilder {
	legendResource, err := legend.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Legend = legendResource

	return builder
}

func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *OptionsBuilder {
	tooltipResource, err := tooltip.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Tooltip = tooltipResource

	return builder
}

func (builder *OptionsBuilder) Timezone(timezone []common.TimeZone) *OptionsBuilder {
	builder.internal.Timezone = timezone

	return builder
}

// Controls the column width
func (builder *OptionsBuilder) ColWidth(colWidth float64) *OptionsBuilder {
	builder.internal.ColWidth = &colWidth

	return builder
}
