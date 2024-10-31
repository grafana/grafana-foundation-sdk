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
	if err := builder.internal.Validate(); err != nil {
		return Derivative{}, err
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

func (builder *DerivativeBuilder) Settings(settings cog.Builder[ElasticsearchDerivativeSettings]) *DerivativeBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *DerivativeBuilder) Hide(hide bool) *DerivativeBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *DerivativeBuilder) applyDefaults() {
}
