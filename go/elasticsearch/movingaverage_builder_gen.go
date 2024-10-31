// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverage] = (*MovingAverageBuilder)(nil)

// #MovingAverage's settings are overridden in types.ts
type MovingAverageBuilder struct {
	internal *MovingAverage
	errors   map[string]cog.BuildErrors
}

func NewMovingAverageBuilder() *MovingAverageBuilder {
	resource := &MovingAverage{}
	builder := &MovingAverageBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "moving_avg"

	return builder
}

func (builder *MovingAverageBuilder) Build() (MovingAverage, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingAverage{}, err
	}

	return *builder.internal, nil
}

func (builder *MovingAverageBuilder) PipelineAgg(pipelineAgg string) *MovingAverageBuilder {
	builder.internal.PipelineAgg = &pipelineAgg

	return builder
}

func (builder *MovingAverageBuilder) Field(field string) *MovingAverageBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *MovingAverageBuilder) Id(id string) *MovingAverageBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *MovingAverageBuilder) Settings(settings map[string]any) *MovingAverageBuilder {
	builder.internal.Settings = settings

	return builder
}

func (builder *MovingAverageBuilder) Hide(hide bool) *MovingAverageBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *MovingAverageBuilder) applyDefaults() {
}
