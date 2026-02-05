// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PipelineVariable] = (*PipelineVariableBuilder)(nil)

type PipelineVariableBuilder struct {
	internal *PipelineVariable
	errors   cog.BuildErrors
}

func NewPipelineVariableBuilder() *PipelineVariableBuilder {
	resource := NewPipelineVariable()
	builder := &PipelineVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PipelineVariableBuilder) Build() (PipelineVariable, error) {
	if err := builder.internal.Validate(); err != nil {
		return PipelineVariable{}, err
	}

	if len(builder.errors) > 0 {
		return PipelineVariable{}, cog.MakeBuildErrors("elasticsearch.pipelineVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PipelineVariableBuilder) Name(name string) *PipelineVariableBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *PipelineVariableBuilder) PipelineAgg(pipelineAgg string) *PipelineVariableBuilder {
	builder.internal.PipelineAgg = pipelineAgg

	return builder
}
