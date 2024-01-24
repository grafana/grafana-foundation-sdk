// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PipelineVariable] = (*PipelineVariableBuilder)(nil)

type PipelineVariableBuilder struct {
	internal *PipelineVariable
	errors   map[string]cog.BuildErrors
}

func NewPipelineVariableBuilder() *PipelineVariableBuilder {
	resource := &PipelineVariable{}
	builder := &PipelineVariableBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *PipelineVariableBuilder) Build() (PipelineVariable, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("PipelineVariable", err)...)
	}

	if len(errs) != 0 {
		return PipelineVariable{}, errs
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

func (builder *PipelineVariableBuilder) applyDefaults() {
}
