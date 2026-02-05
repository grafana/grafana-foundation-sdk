// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingFunction] = (*MovingFunctionBuilder)(nil)

type MovingFunctionBuilder struct {
	internal *MovingFunction
	errors   cog.BuildErrors
}

func NewMovingFunctionBuilder() *MovingFunctionBuilder {
	resource := NewMovingFunction()
	builder := &MovingFunctionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MovingFunctionBuilder) Build() (MovingFunction, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingFunction{}, err
	}

	if len(builder.errors) > 0 {
		return MovingFunction{}, cog.MakeBuildErrors("elasticsearch.movingFunction", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MovingFunctionBuilder) PipelineAgg(pipelineAgg string) *MovingFunctionBuilder {
	builder.internal.PipelineAgg = &pipelineAgg

	return builder
}

func (builder *MovingFunctionBuilder) Field(field string) *MovingFunctionBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *MovingFunctionBuilder) Id(id string) *MovingFunctionBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *MovingFunctionBuilder) Settings(settings cog.Builder[ElasticsearchMovingFunctionSettings]) *MovingFunctionBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *MovingFunctionBuilder) Hide(hide bool) *MovingFunctionBuilder {
	builder.internal.Hide = &hide

	return builder
}
