// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverageHoltModelSettings] = (*MovingAverageHoltModelSettingsBuilder)(nil)

type MovingAverageHoltModelSettingsBuilder struct {
	internal *MovingAverageHoltModelSettings
	errors   map[string]cog.BuildErrors
}

func NewMovingAverageHoltModelSettingsBuilder() *MovingAverageHoltModelSettingsBuilder {
	resource := &MovingAverageHoltModelSettings{}
	builder := &MovingAverageHoltModelSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Model = "holt"

	return builder
}

func (builder *MovingAverageHoltModelSettingsBuilder) Build() (MovingAverageHoltModelSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingAverageHoltModelSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *MovingAverageHoltModelSettingsBuilder) Settings(settings cog.Builder[ElasticsearchMovingAverageHoltModelSettingsSettings]) *MovingAverageHoltModelSettingsBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = settingsResource

	return builder
}

func (builder *MovingAverageHoltModelSettingsBuilder) Window(window string) *MovingAverageHoltModelSettingsBuilder {
	builder.internal.Window = window

	return builder
}

func (builder *MovingAverageHoltModelSettingsBuilder) Minimize(minimize bool) *MovingAverageHoltModelSettingsBuilder {
	builder.internal.Minimize = minimize

	return builder
}

func (builder *MovingAverageHoltModelSettingsBuilder) Predict(predict string) *MovingAverageHoltModelSettingsBuilder {
	builder.internal.Predict = predict

	return builder
}

func (builder *MovingAverageHoltModelSettingsBuilder) applyDefaults() {
}
