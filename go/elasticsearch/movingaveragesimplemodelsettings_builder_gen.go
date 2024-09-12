// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverageSimpleModelSettings] = (*MovingAverageSimpleModelSettingsBuilder)(nil)

type MovingAverageSimpleModelSettingsBuilder struct {
	internal *MovingAverageSimpleModelSettings
	errors   map[string]cog.BuildErrors
}

func NewMovingAverageSimpleModelSettingsBuilder() *MovingAverageSimpleModelSettingsBuilder {
	resource := &MovingAverageSimpleModelSettings{}
	builder := &MovingAverageSimpleModelSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Model = "simple"

	return builder
}

func (builder *MovingAverageSimpleModelSettingsBuilder) Build() (MovingAverageSimpleModelSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("MovingAverageSimpleModelSettings", err)...)
	}

	if len(errs) != 0 {
		return MovingAverageSimpleModelSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *MovingAverageSimpleModelSettingsBuilder) Window(window string) *MovingAverageSimpleModelSettingsBuilder {
	builder.internal.Window = window

	return builder
}

func (builder *MovingAverageSimpleModelSettingsBuilder) Predict(predict string) *MovingAverageSimpleModelSettingsBuilder {
	builder.internal.Predict = predict

	return builder
}

func (builder *MovingAverageSimpleModelSettingsBuilder) applyDefaults() {
}
