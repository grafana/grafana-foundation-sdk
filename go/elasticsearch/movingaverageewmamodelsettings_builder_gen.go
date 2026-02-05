// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverageEWMAModelSettings] = (*MovingAverageEWMAModelSettingsBuilder)(nil)

type MovingAverageEWMAModelSettingsBuilder struct {
	internal *MovingAverageEWMAModelSettings
	errors   cog.BuildErrors
}

func NewMovingAverageEWMAModelSettingsBuilder() *MovingAverageEWMAModelSettingsBuilder {
	resource := NewMovingAverageEWMAModelSettings()
	builder := &MovingAverageEWMAModelSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MovingAverageEWMAModelSettingsBuilder) Build() (MovingAverageEWMAModelSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingAverageEWMAModelSettings{}, err
	}

	if len(builder.errors) > 0 {
		return MovingAverageEWMAModelSettings{}, cog.MakeBuildErrors("elasticsearch.movingAverageEWMAModelSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MovingAverageEWMAModelSettingsBuilder) Settings(settings cog.Builder[ElasticsearchMovingAverageEWMAModelSettingsSettings]) *MovingAverageEWMAModelSettingsBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *MovingAverageEWMAModelSettingsBuilder) Window(window string) *MovingAverageEWMAModelSettingsBuilder {
	builder.internal.Window = window

	return builder
}

func (builder *MovingAverageEWMAModelSettingsBuilder) Minimize(minimize bool) *MovingAverageEWMAModelSettingsBuilder {
	builder.internal.Minimize = minimize

	return builder
}

func (builder *MovingAverageEWMAModelSettingsBuilder) Predict(predict string) *MovingAverageEWMAModelSettingsBuilder {
	builder.internal.Predict = predict

	return builder
}
