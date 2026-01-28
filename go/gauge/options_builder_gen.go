// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package gauge

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
		return Options{}, cog.MakeBuildErrors("gauge.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) ShowThresholdLabels(showThresholdLabels bool) *OptionsBuilder {
	builder.internal.ShowThresholdLabels = showThresholdLabels

	return builder
}

func (builder *OptionsBuilder) ShowThresholdMarkers(showThresholdMarkers bool) *OptionsBuilder {
	builder.internal.ShowThresholdMarkers = showThresholdMarkers

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

func (builder *OptionsBuilder) MinVizHeight(minVizHeight uint32) *OptionsBuilder {
	builder.internal.MinVizHeight = minVizHeight

	return builder
}

func (builder *OptionsBuilder) Orientation(orientation common.VizOrientation) *OptionsBuilder {
	builder.internal.Orientation = orientation

	return builder
}
