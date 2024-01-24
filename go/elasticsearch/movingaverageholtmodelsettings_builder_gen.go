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
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("MovingAverageHoltModelSettings", err)...)
	}

	if len(errs) != 0 {
		return MovingAverageHoltModelSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *MovingAverageHoltModelSettingsBuilder) Settings(settings struct {
	Alpha *string `json:"alpha,omitempty"`
	Beta  *string `json:"beta,omitempty"`
}) *MovingAverageHoltModelSettingsBuilder {
	builder.internal.Settings = settings

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
