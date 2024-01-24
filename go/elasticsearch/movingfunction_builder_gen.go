// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingFunction] = (*MovingFunctionBuilder)(nil)

type MovingFunctionBuilder struct {
	internal *MovingFunction
	errors   map[string]cog.BuildErrors
}

func NewMovingFunctionBuilder() *MovingFunctionBuilder {
	resource := &MovingFunction{}
	builder := &MovingFunctionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "moving_fn"

	return builder
}

func (builder *MovingFunctionBuilder) Build() (MovingFunction, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("MovingFunction", err)...)
	}

	if len(errs) != 0 {
		return MovingFunction{}, errs
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

func (builder *MovingFunctionBuilder) Settings(settings struct {
	Window *string       `json:"window,omitempty"`
	Script *InlineScript `json:"script,omitempty"`
	Shift  *string       `json:"shift,omitempty"`
}) *MovingFunctionBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *MovingFunctionBuilder) Hide(hide bool) *MovingFunctionBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *MovingFunctionBuilder) applyDefaults() {
}
