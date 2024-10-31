// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SingleStatBaseOptions] = (*SingleStatBaseOptionsBuilder)(nil)

// TODO docs
type SingleStatBaseOptionsBuilder struct {
	internal *SingleStatBaseOptions
	errors   map[string]cog.BuildErrors
}

func NewSingleStatBaseOptionsBuilder() *SingleStatBaseOptionsBuilder {
	resource := &SingleStatBaseOptions{}
	builder := &SingleStatBaseOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *SingleStatBaseOptionsBuilder) Build() (SingleStatBaseOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return SingleStatBaseOptions{}, err
	}

	return *builder.internal, nil
}

func (builder *SingleStatBaseOptionsBuilder) ReduceOptions(reduceOptions cog.Builder[ReduceDataOptions]) *SingleStatBaseOptionsBuilder {
	reduceOptionsResource, err := reduceOptions.Build()
	if err != nil {
		builder.errors["reduceOptions"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.ReduceOptions = reduceOptionsResource

	return builder
}

func (builder *SingleStatBaseOptionsBuilder) Text(text cog.Builder[VizTextDisplayOptions]) *SingleStatBaseOptionsBuilder {
	textResource, err := text.Build()
	if err != nil {
		builder.errors["text"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Text = &textResource

	return builder
}

func (builder *SingleStatBaseOptionsBuilder) Orientation(orientation VizOrientation) *SingleStatBaseOptionsBuilder {
	builder.internal.Orientation = orientation

	return builder
}

func (builder *SingleStatBaseOptionsBuilder) applyDefaults() {
}
