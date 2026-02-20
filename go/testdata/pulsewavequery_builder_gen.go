// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PulseWaveQuery] = (*PulseWaveQueryBuilder)(nil)

type PulseWaveQueryBuilder struct {
	internal *PulseWaveQuery
	errors   cog.BuildErrors
}

func NewPulseWaveQueryBuilder() *PulseWaveQueryBuilder {
	resource := NewPulseWaveQuery()
	builder := &PulseWaveQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PulseWaveQueryBuilder) Build() (PulseWaveQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return PulseWaveQuery{}, err
	}

	if len(builder.errors) > 0 {
		return PulseWaveQuery{}, cog.MakeBuildErrors("testdata.pulseWaveQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PulseWaveQueryBuilder) RecordError(path string, err error) *PulseWaveQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
