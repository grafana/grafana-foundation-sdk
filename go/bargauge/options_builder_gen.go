// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bargauge

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
		return Options{}, cog.MakeBuildErrors("bargauge.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) DisplayMode(displayMode common.BarGaugeDisplayMode) *OptionsBuilder {
	builder.internal.DisplayMode = displayMode

	return builder
}

func (builder *OptionsBuilder) ValueMode(valueMode common.BarGaugeValueMode) *OptionsBuilder {
	builder.internal.ValueMode = valueMode

	return builder
}

func (builder *OptionsBuilder) NamePlacement(namePlacement common.BarGaugeNamePlacement) *OptionsBuilder {
	builder.internal.NamePlacement = namePlacement

	return builder
}

func (builder *OptionsBuilder) ShowUnfilled(showUnfilled bool) *OptionsBuilder {
	builder.internal.ShowUnfilled = showUnfilled

	return builder
}

func (builder *OptionsBuilder) Sizing(sizing common.BarGaugeSizing) *OptionsBuilder {
	builder.internal.Sizing = sizing

	return builder
}

func (builder *OptionsBuilder) MinVizWidth(minVizWidth uint32) *OptionsBuilder {
	builder.internal.MinVizWidth = minVizWidth

	return builder
}

func (builder *OptionsBuilder) MinVizHeight(minVizHeight uint32) *OptionsBuilder {
	builder.internal.MinVizHeight = minVizHeight

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

func (builder *OptionsBuilder) MaxVizHeight(maxVizHeight uint32) *OptionsBuilder {
	builder.internal.MaxVizHeight = maxVizHeight

	return builder
}

func (builder *OptionsBuilder) Orientation(orientation common.VizOrientation) *OptionsBuilder {
	builder.internal.Orientation = orientation

	return builder
}
