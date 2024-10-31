// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CSVWave] = (*CSVWaveBuilder)(nil)

type CSVWaveBuilder struct {
	internal *CSVWave
	errors   map[string]cog.BuildErrors
}

func NewCSVWaveBuilder() *CSVWaveBuilder {
	resource := &CSVWave{}
	builder := &CSVWaveBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CSVWaveBuilder) Build() (CSVWave, error) {
	if err := builder.internal.Validate(); err != nil {
		return CSVWave{}, err
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

func (builder *CSVWaveBuilder) applyDefaults() {
}
