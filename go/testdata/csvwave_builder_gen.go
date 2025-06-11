// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CSVWave] = (*CSVWaveBuilder)(nil)

type CSVWaveBuilder struct {
	internal *CSVWave
	errors   cog.BuildErrors
}

func NewCSVWaveBuilder() *CSVWaveBuilder {
	resource := NewCSVWave()
	builder := &CSVWaveBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CSVWaveBuilder) Build() (CSVWave, error) {
	if err := builder.internal.Validate(); err != nil {
		return CSVWave{}, err
	}

	if len(builder.errors) > 0 {
		return CSVWave{}, cog.MakeBuildErrors("testdata.cSVWave", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CSVWaveBuilder) Labels(labels string) *CSVWaveBuilder {
	builder.internal.Labels = &labels

	return builder
}

func (builder *CSVWaveBuilder) Name(name string) *CSVWaveBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *CSVWaveBuilder) TimeStep(timeStep int64) *CSVWaveBuilder {
	builder.internal.TimeStep = &timeStep

	return builder
}

func (builder *CSVWaveBuilder) ValuesCSV(valuesCSV string) *CSVWaveBuilder {
	builder.internal.ValuesCSV = &valuesCSV

	return builder
}
