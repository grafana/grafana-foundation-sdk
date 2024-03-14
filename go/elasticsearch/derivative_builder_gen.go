// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Derivative] = (*DerivativeBuilder)(nil)

type DerivativeBuilder struct {
	internal *Derivative
	errors   map[string]cog.BuildErrors
}

func NewDerivativeBuilder() *DerivativeBuilder {
	resource := &Derivative{}
	builder := &DerivativeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "derivative"

	return builder
}

func (builder *DerivativeBuilder) Build() (Derivative, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Derivative", err)...)
	}

	if len(errs) != 0 {
		return Derivative{}, errs
	}

	return *builder.internal, nil
}

func (builder *DerivativeBuilder) PipelineAgg(pipelineAgg string) *DerivativeBuilder {
	builder.internal.PipelineAgg = &pipelineAgg

	return builder
}

func (builder *DerivativeBuilder) Field(field string) *DerivativeBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *DerivativeBuilder) Id(id string) *DerivativeBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *DerivativeBuilder) Settings(settings struct {
	Unit *string `json:"unit,omitempty"`
}) *DerivativeBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *DerivativeBuilder) Hide(hide bool) *DerivativeBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *DerivativeBuilder) applyDefaults() {
}
