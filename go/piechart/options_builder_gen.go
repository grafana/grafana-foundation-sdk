// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package piechart

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
		return Options{}, cog.MakeBuildErrors("piechart.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) PieType(pieType PieChartType) *OptionsBuilder {
	builder.internal.PieType = pieType

	return builder
}

func (builder *OptionsBuilder) DisplayLabels(displayLabels []PieChartLabels) *OptionsBuilder {
	builder.internal.DisplayLabels = displayLabels

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

func (builder *OptionsBuilder) ReduceOptions(reduceOptions cog.Builder[common.ReduceDataOptions]) *OptionsBuilder {
	reduceOptionsResource, err := reduceOptions.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.ReduceOptions = reduceOptionsResource

	return builder
}

func (builder *OptionsBuilder) Text(text cog.Builder[common.VizTextDisplayOptions]) *OptionsBuilder {
	textResource, err := text.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Text = &textResource

	return builder
}

func (builder *OptionsBuilder) Legend(legend cog.Builder[PieChartLegendOptions]) *OptionsBuilder {
	legendResource, err := legend.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Legend = legendResource

	return builder
}

func (builder *OptionsBuilder) Orientation(orientation common.VizOrientation) *OptionsBuilder {
	builder.internal.Orientation = orientation

	return builder
}
