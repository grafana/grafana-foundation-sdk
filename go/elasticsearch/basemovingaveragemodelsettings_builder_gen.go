// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BaseMovingAverageModelSettings] = (*BaseMovingAverageModelSettingsBuilder)(nil)

type BaseMovingAverageModelSettingsBuilder struct {
	internal *BaseMovingAverageModelSettings
	errors   map[string]cog.BuildErrors
}

func NewBaseMovingAverageModelSettingsBuilder() *BaseMovingAverageModelSettingsBuilder {
	resource := &BaseMovingAverageModelSettings{}
	builder := &BaseMovingAverageModelSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *BaseMovingAverageModelSettingsBuilder) Build() (BaseMovingAverageModelSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("BaseMovingAverageModelSettings", err)...)
	}

	if len(errs) != 0 {
		return BaseMovingAverageModelSettings{}, errs
	}

	return *builder.internal, nil
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

func (builder *BaseMovingAverageModelSettingsBuilder) applyDefaults() {
}
