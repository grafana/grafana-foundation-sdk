// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMovingAverageEWMAModelSettingsSettings] = (*ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder)(nil)

type ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder struct {
	internal *ElasticsearchMovingAverageEWMAModelSettingsSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder() *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder {
	resource := &ElasticsearchMovingAverageEWMAModelSettingsSettings{}
	builder := &ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder) Build() (ElasticsearchMovingAverageEWMAModelSettingsSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchMovingAverageEWMAModelSettingsSettings", err)...)
	}

	if len(errs) != 0 {
		return ElasticsearchMovingAverageEWMAModelSettingsSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder) Alpha(alpha string) *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder {
	builder.internal.Alpha = &alpha

	return builder
}

func (builder *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder) applyDefaults() {
}
