// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverageHoltModelSettings] = (*MovingAverageHoltModelSettingsBuilder)(nil)

type MovingAverageHoltModelSettingsBuilder struct {
	internal *MovingAverageHoltModelSettings
	errors   cog.BuildErrors
}

func NewMovingAverageHoltModelSettingsBuilder() *MovingAverageHoltModelSettingsBuilder {
	resource := NewMovingAverageHoltModelSettings()
	builder := &MovingAverageHoltModelSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MovingAverageHoltModelSettingsBuilder) Build() (MovingAverageHoltModelSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingAverageHoltModelSettings{}, err
	}

	if len(builder.errors) > 0 {
		return MovingAverageHoltModelSettings{}, cog.MakeBuildErrors("elasticsearch.movingAverageHoltModelSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MovingAverageHoltModelSettingsBuilder) RecordError(path string, err error) *MovingAverageHoltModelSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *MovingAverageHoltModelSettingsBuilder) Settings(settings cog.Builder[ElasticsearchMovingAverageHoltModelSettingsSettings]) *MovingAverageHoltModelSettingsBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
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
