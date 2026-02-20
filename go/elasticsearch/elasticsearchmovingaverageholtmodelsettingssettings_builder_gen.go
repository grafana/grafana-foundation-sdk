// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMovingAverageHoltModelSettingsSettings] = (*ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder)(nil)

type ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder struct {
	internal *ElasticsearchMovingAverageHoltModelSettingsSettings
	errors   cog.BuildErrors
}

func NewElasticsearchMovingAverageHoltModelSettingsSettingsBuilder() *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder {
	resource := NewElasticsearchMovingAverageHoltModelSettingsSettings()
	builder := &ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder) Build() (ElasticsearchMovingAverageHoltModelSettingsSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMovingAverageHoltModelSettingsSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchMovingAverageHoltModelSettingsSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchMovingAverageHoltModelSettingsSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder) RecordError(path string, err error) *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder) Alpha(alpha string) *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder {
	builder.internal.Alpha = &alpha

	return builder
}

func (builder *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder) Beta(beta string) *ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder {
	builder.internal.Beta = &beta

	return builder
}
