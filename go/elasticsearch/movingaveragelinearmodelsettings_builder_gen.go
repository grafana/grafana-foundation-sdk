// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverageLinearModelSettings] = (*MovingAverageLinearModelSettingsBuilder)(nil)

type MovingAverageLinearModelSettingsBuilder struct {
	internal *MovingAverageLinearModelSettings
	errors   map[string]cog.BuildErrors
}

func NewMovingAverageLinearModelSettingsBuilder() *MovingAverageLinearModelSettingsBuilder {
	resource := NewMovingAverageLinearModelSettings()
	builder := &MovingAverageLinearModelSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *MovingAverageLinearModelSettingsBuilder) Build() (MovingAverageLinearModelSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingAverageLinearModelSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *MovingAverageLinearModelSettingsBuilder) Window(window string) *MovingAverageLinearModelSettingsBuilder {
	builder.internal.Window = window

	return builder
}

func (builder *MovingAverageLinearModelSettingsBuilder) Predict(predict string) *MovingAverageLinearModelSettingsBuilder {
	builder.internal.Predict = predict

	return builder
}
