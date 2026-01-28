// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package stat

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
		return Options{}, cog.MakeBuildErrors("stat.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) GraphMode(graphMode common.BigValueGraphMode) *OptionsBuilder {
	builder.internal.GraphMode = graphMode

	return builder
}

func (builder *OptionsBuilder) ColorMode(colorMode common.BigValueColorMode) *OptionsBuilder {
	builder.internal.ColorMode = colorMode

	return builder
}

func (builder *OptionsBuilder) JustifyMode(justifyMode common.BigValueJustifyMode) *OptionsBuilder {
	builder.internal.JustifyMode = justifyMode

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

func (builder *OptionsBuilder) TextMode(textMode common.BigValueTextMode) *OptionsBuilder {
	builder.internal.TextMode = textMode

	return builder
}

func (builder *OptionsBuilder) Orientation(orientation common.VizOrientation) *OptionsBuilder {
	builder.internal.Orientation = orientation

	return builder
}
