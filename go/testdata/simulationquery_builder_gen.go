// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SimulationQuery] = (*SimulationQueryBuilder)(nil)

type SimulationQueryBuilder struct {
	internal *SimulationQuery
	errors   cog.BuildErrors
}

func NewSimulationQueryBuilder() *SimulationQueryBuilder {
	resource := NewSimulationQuery()
	builder := &SimulationQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SimulationQueryBuilder) Build() (SimulationQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return SimulationQuery{}, err
	}

	if len(builder.errors) > 0 {
		return SimulationQuery{}, cog.MakeBuildErrors("testdata.simulationQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SimulationQueryBuilder) Key(key cog.Builder[Key]) *SimulationQueryBuilder {
	keyResource, err := key.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Key = keyResource

	return builder
}

func (builder *SimulationQueryBuilder) Config(config map[string]any) *SimulationQueryBuilder {
	builder.internal.Config = config

	return builder
}

func (builder *SimulationQueryBuilder) Stream(stream bool) *SimulationQueryBuilder {
	builder.internal.Stream = &stream

	return builder
}

func (builder *SimulationQueryBuilder) Last(last bool) *SimulationQueryBuilder {
	builder.internal.Last = &last

	return builder
}
