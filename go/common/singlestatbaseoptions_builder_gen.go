// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SingleStatBaseOptions] = (*SingleStatBaseOptionsBuilder)(nil)

// TODO docs
type SingleStatBaseOptionsBuilder struct {
	internal *SingleStatBaseOptions
	errors   cog.BuildErrors
}

func NewSingleStatBaseOptionsBuilder() *SingleStatBaseOptionsBuilder {
	resource := NewSingleStatBaseOptions()
	builder := &SingleStatBaseOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SingleStatBaseOptionsBuilder) Build() (SingleStatBaseOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return SingleStatBaseOptions{}, err
	}

	if len(builder.errors) > 0 {
		return SingleStatBaseOptions{}, cog.MakeBuildErrors("common.singleStatBaseOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SingleStatBaseOptionsBuilder) RecordError(path string, err error) *SingleStatBaseOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *SingleStatBaseOptionsBuilder) ReduceOptions(reduceOptions cog.Builder[ReduceDataOptions]) *SingleStatBaseOptionsBuilder {
	reduceOptionsResource, err := reduceOptions.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.ReduceOptions = reduceOptionsResource

	return builder
}

func (builder *SingleStatBaseOptionsBuilder) Text(text cog.Builder[VizTextDisplayOptions]) *SingleStatBaseOptionsBuilder {
	textResource, err := text.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Text = &textResource

	return builder
}

func (builder *SingleStatBaseOptionsBuilder) Orientation(orientation VizOrientation) *SingleStatBaseOptionsBuilder {
	builder.internal.Orientation = orientation

	return builder
}
