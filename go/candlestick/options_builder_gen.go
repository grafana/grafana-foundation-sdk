// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package candlestick

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
		return Options{}, cog.MakeBuildErrors("candlestick.options", builder.errors)
	}

	return *builder.internal, nil
}

// Sets which dimensions are used for the visualization
func (builder *OptionsBuilder) Mode(mode VizDisplayMode) *OptionsBuilder {
	builder.internal.Mode = mode

	return builder
}

// Sets the style of the candlesticks
func (builder *OptionsBuilder) CandleStyle(candleStyle CandleStyle) *OptionsBuilder {
	builder.internal.CandleStyle = candleStyle

	return builder
}

// Sets the color strategy for the candlesticks
func (builder *OptionsBuilder) ColorStrategy(colorStrategy ColorStrategy) *OptionsBuilder {
	builder.internal.ColorStrategy = colorStrategy

	return builder
}

// Map fields to appropriate dimension
func (builder *OptionsBuilder) Fields(fields cog.Builder[CandlestickFieldMap]) *OptionsBuilder {
	fieldsResource, err := fields.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Fields = fieldsResource

	return builder
}

// Set which colors are used when the price movement is up or down
func (builder *OptionsBuilder) Colors(colors cog.Builder[CandlestickColors]) *OptionsBuilder {
	colorsResource, err := colors.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Colors = colorsResource

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

// When enabled, all fields will be sent to the graph
func (builder *OptionsBuilder) IncludeAllFields(includeAllFields bool) *OptionsBuilder {
	builder.internal.IncludeAllFields = &includeAllFields

	return builder
}
