// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Scenario] = (*ScenarioBuilder)(nil)

// TODO: Should this live here given it's not used in the dataquery?
type ScenarioBuilder struct {
	internal *Scenario
	errors   cog.BuildErrors
}

func NewScenarioBuilder() *ScenarioBuilder {
	resource := NewScenario()
	builder := &ScenarioBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ScenarioBuilder) Build() (Scenario, error) {
	if err := builder.internal.Validate(); err != nil {
		return Scenario{}, err
	}

	if len(builder.errors) > 0 {
		return Scenario{}, cog.MakeBuildErrors("testdata.scenario", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ScenarioBuilder) Id(id string) *ScenarioBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *ScenarioBuilder) Name(name string) *ScenarioBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *ScenarioBuilder) StringInput(stringInput string) *ScenarioBuilder {
	builder.internal.StringInput = stringInput

	return builder
}

func (builder *ScenarioBuilder) Description(description string) *ScenarioBuilder {
	builder.internal.Description = &description

	return builder
}

func (builder *ScenarioBuilder) HideAliasField(hideAliasField bool) *ScenarioBuilder {
	builder.internal.HideAliasField = &hideAliasField

	return builder
}
