// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverageHoltWintersModelSettings] = (*MovingAverageHoltWintersModelSettingsBuilder)(nil)

type MovingAverageHoltWintersModelSettingsBuilder struct {
	internal *MovingAverageHoltWintersModelSettings
	errors   map[string]cog.BuildErrors
}

func NewMovingAverageHoltWintersModelSettingsBuilder() *MovingAverageHoltWintersModelSettingsBuilder {
	resource := &MovingAverageHoltWintersModelSettings{}
	builder := &MovingAverageHoltWintersModelSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Model = "holt_winters"

	return builder
}

func (builder *MovingAverageHoltWintersModelSettingsBuilder) Build() (MovingAverageHoltWintersModelSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingAverageHoltWintersModelSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *MovingAverageHoltWintersModelSettingsBuilder) Settings(settings cog.Builder[ElasticsearchMovingAverageHoltWintersModelSettingsSettings]) *MovingAverageHoltWintersModelSettingsBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = settingsResource

	return builder
}

func (builder *MovingAverageHoltWintersModelSettingsBuilder) Window(window string) *MovingAverageHoltWintersModelSettingsBuilder {
	builder.internal.Window = window

	return builder
}

func (builder *MovingAverageHoltWintersModelSettingsBuilder) Minimize(minimize bool) *MovingAverageHoltWintersModelSettingsBuilder {
	builder.internal.Minimize = minimize

	return builder
}

func (builder *MovingAverageHoltWintersModelSettingsBuilder) Predict(predict string) *MovingAverageHoltWintersModelSettingsBuilder {
	builder.internal.Predict = predict

	return builder
}

func (builder *MovingAverageHoltWintersModelSettingsBuilder) applyDefaults() {
}
