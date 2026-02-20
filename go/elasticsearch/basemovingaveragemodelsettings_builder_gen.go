// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BaseMovingAverageModelSettings] = (*BaseMovingAverageModelSettingsBuilder)(nil)

type BaseMovingAverageModelSettingsBuilder struct {
	internal *BaseMovingAverageModelSettings
	errors   cog.BuildErrors
}

func NewBaseMovingAverageModelSettingsBuilder() *BaseMovingAverageModelSettingsBuilder {
	resource := NewBaseMovingAverageModelSettings()
	builder := &BaseMovingAverageModelSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BaseMovingAverageModelSettingsBuilder) Build() (BaseMovingAverageModelSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return BaseMovingAverageModelSettings{}, err
	}

	if len(builder.errors) > 0 {
		return BaseMovingAverageModelSettings{}, cog.MakeBuildErrors("elasticsearch.baseMovingAverageModelSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BaseMovingAverageModelSettingsBuilder) RecordError(path string, err error) *BaseMovingAverageModelSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *BaseMovingAverageModelSettingsBuilder) Model(model MovingAverageModel) *BaseMovingAverageModelSettingsBuilder {
	builder.internal.Model = model

	return builder
}

func (builder *BaseMovingAverageModelSettingsBuilder) Window(window string) *BaseMovingAverageModelSettingsBuilder {
	builder.internal.Window = window

	return builder
}

func (builder *BaseMovingAverageModelSettingsBuilder) Predict(predict string) *BaseMovingAverageModelSettingsBuilder {
	builder.internal.Predict = predict

	return builder
}
