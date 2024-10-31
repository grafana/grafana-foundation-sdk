// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CumulativeSum] = (*CumulativeSumBuilder)(nil)

type CumulativeSumBuilder struct {
	internal *CumulativeSum
	errors   map[string]cog.BuildErrors
}

func NewCumulativeSumBuilder() *CumulativeSumBuilder {
	resource := &CumulativeSum{}
	builder := &CumulativeSumBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "cumulative_sum"

	return builder
}

func (builder *CumulativeSumBuilder) Build() (CumulativeSum, error) {
	if err := builder.internal.Validate(); err != nil {
		return CumulativeSum{}, err
	}

	return *builder.internal, nil
}

func (builder *CumulativeSumBuilder) PipelineAgg(pipelineAgg string) *CumulativeSumBuilder {
	builder.internal.PipelineAgg = &pipelineAgg

	return builder
}

func (builder *CumulativeSumBuilder) Field(field string) *CumulativeSumBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *CumulativeSumBuilder) Id(id string) *CumulativeSumBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *CumulativeSumBuilder) Settings(settings cog.Builder[ElasticsearchCumulativeSumSettings]) *CumulativeSumBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *CumulativeSumBuilder) Hide(hide bool) *CumulativeSumBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *CumulativeSumBuilder) applyDefaults() {
}
