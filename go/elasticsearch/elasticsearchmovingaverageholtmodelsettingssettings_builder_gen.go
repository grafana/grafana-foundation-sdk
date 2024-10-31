// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMovingAverageHoltModelSettingsSettings] = (*ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder)(nil)

type ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder struct {
	internal *ElasticsearchMovingAverageHoltModelSettingsSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchMovingAverageHoltModelSettingsSettingsBuilder() *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder {
	resource := &ElasticsearchMovingAverageHoltModelSettingsSettings{}
	builder := &ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder) Build() (ElasticsearchMovingAverageHoltModelSettingsSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMovingAverageHoltModelSettingsSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder) Alpha(alpha string) *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder {
	builder.internal.Alpha = &alpha

	return builder
}

func (builder *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder) Beta(beta string) *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder {
	builder.internal.Beta = &beta

	return builder
}

func (builder *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder) applyDefaults() {
}
