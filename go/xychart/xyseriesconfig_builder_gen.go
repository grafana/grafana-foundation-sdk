// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XYSeriesConfig] = (*XYSeriesConfigBuilder)(nil)

type XYSeriesConfigBuilder struct {
	internal *XYSeriesConfig
	errors   map[string]cog.BuildErrors
}

func NewXYSeriesConfigBuilder() *XYSeriesConfigBuilder {
	resource := NewXYSeriesConfig()
	builder := &XYSeriesConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *XYSeriesConfigBuilder) Build() (XYSeriesConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return XYSeriesConfig{}, err
	}

	return *builder.internal, nil
}

func (builder *XYSeriesConfigBuilder) Name(name cog.Builder[XychartXYSeriesConfigName]) *XYSeriesConfigBuilder {
	nameResource, err := name.Build()
	if err != nil {
		builder.errors["name"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Name = &nameResource

	return builder
}

func (builder *XYSeriesConfigBuilder) Frame(frame cog.Builder[XychartXYSeriesConfigFrame]) *XYSeriesConfigBuilder {
	frameResource, err := frame.Build()
	if err != nil {
		builder.errors["frame"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Frame = &frameResource

	return builder
}

func (builder *XYSeriesConfigBuilder) X(x cog.Builder[XychartXYSeriesConfigX]) *XYSeriesConfigBuilder {
	xResource, err := x.Build()
	if err != nil {
		builder.errors["x"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.X = &xResource

	return builder
}

func (builder *XYSeriesConfigBuilder) Y(y cog.Builder[XychartXYSeriesConfigY]) *XYSeriesConfigBuilder {
	yResource, err := y.Build()
	if err != nil {
		builder.errors["y"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Y = &yResource

	return builder
}

func (builder *XYSeriesConfigBuilder) Color(color cog.Builder[XychartXYSeriesConfigColor]) *XYSeriesConfigBuilder {
	colorResource, err := color.Build()
	if err != nil {
		builder.errors["color"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Color = &colorResource

	return builder
}

func (builder *XYSeriesConfigBuilder) Size(size cog.Builder[XychartXYSeriesConfigSize]) *XYSeriesConfigBuilder {
	sizeResource, err := size.Build()
	if err != nil {
		builder.errors["size"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Size = &sizeResource

	return builder
}
