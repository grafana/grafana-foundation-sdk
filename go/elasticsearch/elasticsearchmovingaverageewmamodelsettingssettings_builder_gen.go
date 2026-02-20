// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMovingAverageEWMAModelSettingsSettings] = (*ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder)(nil)

type ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder struct {
	internal *ElasticsearchMovingAverageEWMAModelSettingsSettings
	errors   cog.BuildErrors
}

func NewElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder() *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder {
	resource := NewElasticsearchMovingAverageEWMAModelSettingsSettings()
	builder := &ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder) Build() (ElasticsearchMovingAverageEWMAModelSettingsSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMovingAverageEWMAModelSettingsSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchMovingAverageEWMAModelSettingsSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchMovingAverageEWMAModelSettingsSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder) RecordError(path string, err error) *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder) Alpha(alpha string) *ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder {
	builder.internal.Alpha = &alpha

	return builder
}
