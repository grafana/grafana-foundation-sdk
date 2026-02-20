// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverageHoltWintersModelSettings] = (*MovingAverageHoltWintersModelSettingsBuilder)(nil)

type MovingAverageHoltWintersModelSettingsBuilder struct {
	internal *MovingAverageHoltWintersModelSettings
	errors   cog.BuildErrors
}

func NewMovingAverageHoltWintersModelSettingsBuilder() *MovingAverageHoltWintersModelSettingsBuilder {
	resource := NewMovingAverageHoltWintersModelSettings()
	builder := &MovingAverageHoltWintersModelSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MovingAverageHoltWintersModelSettingsBuilder) Build() (MovingAverageHoltWintersModelSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingAverageHoltWintersModelSettings{}, err
	}

	if len(builder.errors) > 0 {
		return MovingAverageHoltWintersModelSettings{}, cog.MakeBuildErrors("elasticsearch.movingAverageHoltWintersModelSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MovingAverageHoltWintersModelSettingsBuilder) RecordError(path string, err error) *MovingAverageHoltWintersModelSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *MovingAverageHoltWintersModelSettingsBuilder) Settings(settings cog.Builder[ElasticsearchMovingAverageHoltWintersModelSettingsSettings]) *MovingAverageHoltWintersModelSettingsBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
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
