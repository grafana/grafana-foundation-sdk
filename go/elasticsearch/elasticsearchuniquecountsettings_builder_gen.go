// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchUniqueCountSettings] = (*ElasticsearchUniqueCountSettingsBuilder)(nil)

type ElasticsearchUniqueCountSettingsBuilder struct {
	internal *ElasticsearchUniqueCountSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchUniqueCountSettingsBuilder() *ElasticsearchUniqueCountSettingsBuilder {
	resource := NewElasticsearchUniqueCountSettings()
	builder := &ElasticsearchUniqueCountSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ElasticsearchUniqueCountSettingsBuilder) Build() (ElasticsearchUniqueCountSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchUniqueCountSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchUniqueCountSettingsBuilder) PrecisionThreshold(precisionThreshold string) *ElasticsearchUniqueCountSettingsBuilder {
	builder.internal.PrecisionThreshold = &precisionThreshold

	return builder
}

func (builder *ElasticsearchUniqueCountSettingsBuilder) Missing(missing string) *ElasticsearchUniqueCountSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}
