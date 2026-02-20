// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XYSeriesConfig] = (*XYSeriesConfigBuilder)(nil)

type XYSeriesConfigBuilder struct {
	internal *XYSeriesConfig
	errors   cog.BuildErrors
}

func NewXYSeriesConfigBuilder() *XYSeriesConfigBuilder {
	resource := NewXYSeriesConfig()
	builder := &XYSeriesConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *XYSeriesConfigBuilder) Build() (XYSeriesConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return XYSeriesConfig{}, err
	}

	if len(builder.errors) > 0 {
		return XYSeriesConfig{}, cog.MakeBuildErrors("xychart.xYSeriesConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *XYSeriesConfigBuilder) RecordError(path string, err error) *XYSeriesConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *XYSeriesConfigBuilder) Name(name cog.Builder[XychartXYSeriesConfigName]) *XYSeriesConfigBuilder {
	nameResource, err := name.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Name = &nameResource

	return builder
}

func (builder *XYSeriesConfigBuilder) Frame(frame cog.Builder[XychartXYSeriesConfigFrame]) *XYSeriesConfigBuilder {
	frameResource, err := frame.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Frame = &frameResource

	return builder
}

func (builder *XYSeriesConfigBuilder) X(x cog.Builder[XychartXYSeriesConfigX]) *XYSeriesConfigBuilder {
	xResource, err := x.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.X = &xResource

	return builder
}

func (builder *XYSeriesConfigBuilder) Y(y cog.Builder[XychartXYSeriesConfigY]) *XYSeriesConfigBuilder {
	yResource, err := y.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Y = &yResource

	return builder
}

func (builder *XYSeriesConfigBuilder) Color(color cog.Builder[XychartXYSeriesConfigColor]) *XYSeriesConfigBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Color = &colorResource

	return builder
}

func (builder *XYSeriesConfigBuilder) Size(size cog.Builder[XychartXYSeriesConfigSize]) *XYSeriesConfigBuilder {
	sizeResource, err := size.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Size = &sizeResource

	return builder
}
