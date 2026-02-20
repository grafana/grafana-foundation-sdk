// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverageSimpleModelSettings] = (*MovingAverageSimpleModelSettingsBuilder)(nil)

type MovingAverageSimpleModelSettingsBuilder struct {
	internal *MovingAverageSimpleModelSettings
	errors   cog.BuildErrors
}

func NewMovingAverageSimpleModelSettingsBuilder() *MovingAverageSimpleModelSettingsBuilder {
	resource := NewMovingAverageSimpleModelSettings()
	builder := &MovingAverageSimpleModelSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MovingAverageSimpleModelSettingsBuilder) Build() (MovingAverageSimpleModelSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingAverageSimpleModelSettings{}, err
	}

	if len(builder.errors) > 0 {
		return MovingAverageSimpleModelSettings{}, cog.MakeBuildErrors("elasticsearch.movingAverageSimpleModelSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MovingAverageSimpleModelSettingsBuilder) RecordError(path string, err error) *MovingAverageSimpleModelSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *MovingAverageSimpleModelSettingsBuilder) Window(window string) *MovingAverageSimpleModelSettingsBuilder {
	builder.internal.Window = window

	return builder
}

func (builder *MovingAverageSimpleModelSettingsBuilder) Predict(predict string) *MovingAverageSimpleModelSettingsBuilder {
	builder.internal.Predict = predict

	return builder
}
