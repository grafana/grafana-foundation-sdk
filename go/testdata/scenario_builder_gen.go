// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Scenario] = (*ScenarioBuilder)(nil)

// TODO: Should this live here given it's not used in the dataquery?
type ScenarioBuilder struct {
	internal *Scenario
	errors   map[string]cog.BuildErrors
}

func NewScenarioBuilder() *ScenarioBuilder {
	resource := &Scenario{}
	builder := &ScenarioBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ScenarioBuilder) Build() (Scenario, error) {
	if err := builder.internal.Validate(); err != nil {
		return Scenario{}, err
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

func (builder *ScenarioBuilder) applyDefaults() {
}
