// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PulseWaveQuery] = (*PulseWaveQueryBuilder)(nil)

type PulseWaveQueryBuilder struct {
	internal *PulseWaveQuery
	errors   map[string]cog.BuildErrors
}

func NewPulseWaveQueryBuilder() *PulseWaveQueryBuilder {
	resource := &PulseWaveQuery{}
	builder := &PulseWaveQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *PulseWaveQueryBuilder) Build() (PulseWaveQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return PulseWaveQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *PulseWaveQueryBuilder) OffCount(offCount int64) *PulseWaveQueryBuilder {
	builder.internal.OffCount = &offCount

	return builder
}

func (builder *PulseWaveQueryBuilder) OffValue(offValue float64) *PulseWaveQueryBuilder {
	builder.internal.OffValue = &offValue

	return builder
}

func (builder *PulseWaveQueryBuilder) OnCount(onCount int64) *PulseWaveQueryBuilder {
	builder.internal.OnCount = &onCount

	return builder
}

func (builder *PulseWaveQueryBuilder) OnValue(onValue float64) *PulseWaveQueryBuilder {
	builder.internal.OnValue = &onValue

	return builder
}

func (builder *PulseWaveQueryBuilder) TimeStep(timeStep int64) *PulseWaveQueryBuilder {
	builder.internal.TimeStep = &timeStep

	return builder
}

func (builder *PulseWaveQueryBuilder) applyDefaults() {
}
